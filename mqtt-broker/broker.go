package mqtt_broker

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"github.com/akzj/block-queue"
	"github.com/akzj/streamIO/client"
	"github.com/akzj/streamIO/meta-server/store"
	"github.com/eclipse/paho.mqtt.golang/packets"
	"github.com/golang/protobuf/proto"
	"github.com/google/btree"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type Broker struct {
	Options
	tlsConfig *tls.Config

	sessionsLocker sync.Mutex
	sessions       map[string]*session

	topicTree unsafe.Pointer
	metaTree  *btree.BTree

	eventReader *EventReader
	eventQueue  *block_queue.Queue

	client       client.Client
	eventOffset  int64
	treeChanges  int64
	isCheckpoint int32

	eventWriter client.StreamWriter

	snapshot *Snapshot

	ctx    context.Context
	cancel context.CancelFunc

	listener []net.Listener

	offsetCommitter *offsetCommitter
}

func New(options Options) *Broker {
	if options.LogFile != "" {
		_ = os.MkdirAll(filepath.Dir(options.LogFile), 0777)
		file, err := os.OpenFile(options.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		log.SetOutput(file)
	}
	log.SetFormatter(&log.TextFormatter{DisableQuote: true})
	log.SetReportCaller(true)
	ctx, cancel := context.WithCancel(context.Background())
	metaServerClient, err := client.NewMetaServiceClient(ctx, options.MetaServerAddr)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	cli := client.NewClient(metaServerClient)
	streamInfoItem, err := cli.GetOrCreateStreamInfoItem(ctx, MQTTEventStream)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	sess, err := cli.NewStreamSession(ctx, options.BrokerId, streamInfoItem)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	eventWriter, err := sess.NewWriter()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	eventQueue := block_queue.NewQueue(128)
	eventWatcher, err := newEventReader(options.BrokerId, cli, func(message EventWithOffset) {
		eventQueue.Push(message)
	})
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	broker := &Broker{
		Options:         options,
		tlsConfig:       nil,
		sessionsLocker:  sync.Mutex{},
		sessions:        map[string]*session{},
		topicTree:       unsafe.Pointer(NewTopicTree()),
		metaTree:        btree.New(10),
		eventReader:     eventWatcher,
		eventQueue:      eventQueue,
		client:          cli,
		eventOffset:     0,
		treeChanges:     0,
		isCheckpoint:    0,
		eventWriter:     eventWriter,
		snapshot:        NewSnapshot(options.SnapshotPath),
		ctx:             ctx,
		cancel:          cancel,
		listener:        nil,
		offsetCommitter: newOffsetCommitter(),
	}

	return broker
}

func (broker *Broker) Start() error {
	if err := broker.snapshot.reloadSnapshot(broker); err != nil {
		return err
	}
	go broker.eventReader.readEventLoop()
	if err := broker.clientListenLoop(); err != nil {
		log.Fatal(err.Error())
	}
	go broker.offsetCommitter.commitLoop(broker.ctx, broker.ReadOffsetCommitInterval)
	broker.processEventLoop()
	return nil
}

func (broker *Broker) getSubscribeTree() *TopicTree {
	return (*TopicTree)(atomic.LoadPointer(&broker.topicTree))
}

func (broker *Broker) setSubscribeTree(tree *TopicTree) {
	atomic.StorePointer(&broker.topicTree, unsafe.Pointer(tree))
}

func (broker *Broker) newListener() ([]net.Listener, error) {
	var listeners []net.Listener
	if broker.BindPort != 0 {
		listener, err := net.Listen("tcp",
			net.JoinHostPort(broker.HOST, strconv.Itoa(broker.BindPort)))
		if err != nil {
			log.WithField("broker.BindPort", broker.BindPort).Error(err)
			return nil, err
		}
		listeners = append(listeners, listener)
	}
	if broker.BindTLSPort != 0 {
		listener, err := tls.Listen("tcp",
			net.JoinHostPort(broker.HOST, strconv.Itoa(broker.BindPort)), broker.tlsConfig)
		if err != nil {
			log.WithField("broker.BindTLSPort", broker.BindPort).Error(err)
			return nil, err
		}
		listeners = append(listeners, listener)
	}
	if listeners == nil {
		return nil, errors.New("no listeners")
	}
	return listeners, nil
}

func (broker *Broker) deleteSession(identifier string) error {
	//1 delete client session from meta-server
	info, err := broker.client.DeleteMQTTClientSession(context.Background(), identifier)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	if info == nil {
		return nil
	}

	//2 delete subscribe,and pub `unsubscribe-event` to mqtt-event-queue
	packet := &packets.UnsubscribePacket{}
	for topic := range info.Topics {
		packet.Topics = append(packet.Topics, topic)
	}
	if err := broker.handleUnSubscribePacket(info, packet); err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (broker *Broker) handleConnection(conn net.Conn) {
	defer func() {
		_ = conn.Close()
	}()
	remoteAddr := conn.RemoteAddr()
	var logEntry = log.WithField("remoteAddr", remoteAddr)
	packet, err := packets.ReadPacket(conn)
	if err != nil {
		logEntry.Error(err)
		return
	}
	connectPacket, ok := packet.(*packets.ConnectPacket)
	if ok == false {
		logEntry.Error("first packet is no ConnectPacket error")
		return
	}

	logEntry.WithField("connectPacket", connectPacket).Info("handle connection")

	connackPacket := packets.NewControlPacket(packets.Connack).(*packets.ConnackPacket)

	//validate connect packet
	if returnCode := connectPacket.Validate(); returnCode != packets.Accepted {
		logEntry.Errorf("Validate failed %s", packets.ConnackReturnCodes[returnCode])
		connackPacket.ReturnCode = returnCode
		_ = connackPacket.Write(conn)
		return
	}
	//check auth
	if status, err := broker.checkConnectAuth(connectPacket.ClientIdentifier,
		connectPacket.Username,
		string(connectPacket.Password)); err != nil {
		connackPacket.ReturnCode = packets.ErrRefusedServerUnavailable
		_ = connackPacket.Write(conn)
		logEntry.Error(err)
		return
	} else if status == false {
		connackPacket.ReturnCode = packets.ErrRefusedNotAuthorised
		_ = connackPacket.Write(conn)
		return
	}
	//process session
	if connectPacket.CleanSession {
		if err := broker.deleteSession(connectPacket.ClientIdentifier); err != nil {
			connackPacket.ReturnCode = packets.ErrRefusedServerUnavailable
			_ = connackPacket.Write(conn)
			logEntry.Error(err)
		}
		//[MQTT-3.1.2-6]。
		defer func() {
			if err := broker.deleteSession(connectPacket.ClientIdentifier); err != nil {
				log.Error(err)
			}
		}()
	}
	if connectPacket.Keepalive == 0 {
		connectPacket.Keepalive = broker.DefaultKeepalive
	} else if connectPacket.Keepalive < broker.MinKeepalive {
		connectPacket.Keepalive = broker.MinKeepalive
	}
	sess, err := newSession(broker, connectPacket.Keepalive, conn, broker.client, connectPacket.ClientIdentifier)
	if err != nil {
		log.Error(err.Error())
		connackPacket.ReturnCode = packets.ErrRefusedServerUnavailable
		_ = connackPacket.Write(conn)
		logEntry.Error(err)
		return
	}
	connackPacket.SessionPresent = sess.create
	connackPacket.ReturnCode = packets.Accepted
	if err := connackPacket.Write(conn); err != nil {
		return
	}
	if connectPacket.WillFlag {
		willMessage := packets.NewControlPacket(packets.Publish).(*packets.PublishPacket)
		willMessage.Payload = connectPacket.WillMessage
		willMessage.TopicName = connectPacket.WillTopic
		sess.setWillMessage(willMessage)
	}
	sess.readConnLoop()
}

func (broker *Broker) serve(listener net.Listener) {
	var tempDelay time.Duration
	for {
		conn, err := listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.Errorf("http: Accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
		}
		go broker.handleConnection(conn)
	}
}

func (broker *Broker) clientListenLoop() error {
	listeners, err := broker.newListener()
	if err != nil {
		return err
	}
	for _, listener := range listeners {
		go broker.serve(listener)
	}
	return nil
}

func (broker *Broker) checkConnectAuth(clientIdentifier string, username string, password string) (bool, error) {
	return true, nil
}

func (broker *Broker) handleEvent(packet packets.ControlPacket) {
	broker.eventQueue.Push(packet)
}

func (broker *Broker) newSubscriber(event *SubscribeEvent) ([]Subscriber, error) {
	item := broker.metaTree.Get(&subscriberStatus{sessionID: event.SessionId})
	if item == nil {
		return nil, fmt.Errorf("no find session status")
	}
	var subscribers []Subscriber
	for topic, qos := range event.Topic {
		streamInfo := event.Qos0StreamInfo
		if qos == 1 {
			streamInfo = event.Qos1StreamInfo
		}
		session, err := broker.client.NewStreamSession(context.Background(), event.SessionId, streamInfo)
		if err != nil {
			return nil, err
		}
		writer, err := session.NewWriter()
		if err != nil {
			return nil, err
		}
		subscribers = append(subscribers, &subscriber{
			streamWriter: writer,
			sessionID:    event.SessionId,
			qos:          qos,
			topic:        topic,
			status:       item.(*subscriberStatus),
			streamInfo:   streamInfo,
		})
	}
	return subscribers, nil
}

func (broker *Broker) handleSubscribeEvent(event *SubscribeEvent) {
	log.WithField("event", event).Info("handleSubscribeEvent")
	tree := broker.getSubscribeTree().Clone()
	broker.insertSubscriber2Tree(tree, event)
	broker.setSubscribeTree(tree)
}

func (broker *Broker) insertSubscriber2Tree(tree *TopicTree, event *SubscribeEvent) {
	subs, err := broker.newSubscriber(event)
	if err != nil {
		log.Errorf("%+v\n", err)
		return
	}
	for _, sub := range subs {
		tree.Insert(sub)
	}
}

func (broker *Broker) handleUnSubscribeEvent(event *UnSubscribeEvent) {
	log.WithField("event", event).Info("handleUnSubscribeEvent")
	tree := broker.getSubscribeTree().Clone()
	for _, topic := range event.Topic {
		tree.Delete(&subscriber{topic: topic, sessionID: event.SessionId})
	}
	broker.setSubscribeTree(tree)
}

func (broker *Broker) handleRetainMessageEvent(event *RetainMessageEvent) {
	log.WithField("event", event).Info("handleRetainMessageEvent")
	tree := broker.getSubscribeTree().Clone()
	_ = broker.insertRetainMessage2Tree(tree, event)
	broker.setSubscribeTree(tree)
}

func (broker *Broker) insertRetainMessage2Tree(tree *TopicTree, event *RetainMessageEvent) error {
	var buffer = bytes.NewReader(event.Data)
	packet, err := packets.ReadPacket(buffer)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	tree.UpdateRetainPacket(packet.(*packets.PublishPacket))
	return nil
}

func (broker *Broker) processEventLoop() {
	for {
		items := broker.eventQueue.PopAll(nil)
		for _, item := range items {
			event := item.(EventWithOffset)
			switch event := event.event.(type) {
			case *SubscribeEvent:
				broker.handleSubscribeEvent(event)
			case *UnSubscribeEvent:
				broker.handleUnSubscribeEvent(event)
			case *RetainMessageEvent:
				broker.handleRetainMessageEvent(event)
			case *ClientStatusChangeEvent:
				broker.handleClientStatusChangeEvent(event)
			default:
				log.Fatalf("unknown event %+v\n", event)
			}
			broker.eventOffset = event.offset
			broker.treeChanges++
		}
		if broker.treeChanges > broker.CheckpointEventSize {
			if atomic.CompareAndSwapInt32(&broker.isCheckpoint, 0, 1) == false {
				continue
			}
			clone := broker.getSubscribeTree().Clone()
			go func() {
				defer func() {
					atomic.StoreInt32(&broker.isCheckpoint, 0)
				}()
				broker.checkpoint(clone, broker.eventOffset)
				log.Infof("checkpoint success")
			}()
		}
	}
}

func (broker *Broker) handleUnSubscribePacket(sessionItem *store.MQTTSessionItem,
	packet *packets.UnsubscribePacket) error {
	event := UnSubscribeEvent{
		SessionId: sessionItem.SessionId,
		Topic:     packet.Topics,
	}
	return broker.sendEvent(&event)
}

func (broker *Broker) sendEvent(message proto.Message) error {
	event := Event{}
	switch typ := message.(type) {
	case *SubscribeEvent:
		event.Type = Event_SubscribeEvent
	case *UnSubscribeEvent:
		event.Type = Event_UnSubscribeEvent
	case *RetainMessageEvent:
		event.Type = Event_RetainMessageEvent
	case *ClientStatusChangeEvent:
		event.Type = Event_ClientStatusChangeEvent
	default:
		panic("unknown message type" + typ.String())
	}
	data, err := proto.Marshal(message)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	event.Data = data
	data, err = proto.Marshal(&event)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	var buffer bytes.Buffer
	_ = binary.Write(&buffer, binary.BigEndian, int32(len(data)))
	buffer.Write(data)
	var wg sync.WaitGroup
	wg.Wait()
	wg.Add(1)
	broker.eventWriter.WriteWithCb(buffer.Bytes(), func(e error) {
		wg.Done()
		err = e
	})
	wg.Wait()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (broker *Broker) handleSubscribePacket(sessionItem *store.MQTTSessionItem,
	packet *packets.SubscribePacket) error {
	event := SubscribeEvent{
		SessionId:      sessionItem.SessionId,
		Qos0StreamInfo: sessionItem.Qos0StreamInfo,
		Qos1StreamInfo: sessionItem.Qos1StreamInfo,
		Topic:          map[string]int32{},
	}
	for index, topic := range packet.Topics {
		qos := packet.Qoss[index]
		event.Topic[topic] = int32(qos)
	}
	return broker.sendEvent(&event)
}

func (broker *Broker) handleRetainPacket(packet *packets.PublishPacket) error {
	var buffer bytes.Buffer
	_ = packet.Write(&buffer)
	event := RetainMessageEvent{
		Data: buffer.Bytes(),
	}
	return broker.sendEvent(&event)
}

func (broker *Broker) checkpoint(clone *TopicTree, offset int64) error {
	err := broker.snapshot.WriteSnapshot(SnapshotHeader{
		TS:     time.Now(),
		Offset: offset,
	}, clone, broker.metaTree)
	if err != nil {
		log.Error(err)
		return err
	}
	if err := broker.eventReader.commitReadOffset(offset); err != nil {
		log.Error(err)
	}
	return nil
}

func (broker *Broker) handleClientStatusChange(sessionID int64, offline ClientStatusChangeEvent_Status) error {
	event := &ClientStatusChangeEvent{
		SessionID: sessionID,
		Status:    offline,
	}
	return broker.sendEvent(event)
}

func (broker *Broker) handleClientStatusChangeEvent(event *ClientStatusChangeEvent) {
	log.WithField("event", event).Info("handleClientStatusChangeEvent")
	item := broker.metaTree.Get(&subscriberStatus{
		sessionID: event.SessionID,
		status:    &event.Status,
	})
	if item == nil {
		//copy on write
		metaTree := broker.metaTree.Clone()
		metaTree.ReplaceOrInsert(&subscriberStatus{
			sessionID: event.SessionID,
			status:    &event.Status,
		})
		broker.metaTree = metaTree
	} else {
		atomic.StoreInt32((*int32)(item.(*subscriberStatus).status), int32(event.Status))
	}
}
