package mqtt_broker

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/akzj/streamIO/client"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
)

const MQTTEventStream = "$streamIO-mqtt-broker-event"
const MaxEventSize = 1024 * 1024 * 128

type EventReader struct {
	ctx       context.Context
	cancel    context.CancelFunc
	sessionID int64 //serverID
	session   client.StreamSession
	client    client.Client
	callback  eventCallback
	reader    client.StreamReader
}

type EventWithOffset struct {
	event  proto.Message
	offset int64
}

type eventCallback func(message EventWithOffset)

func newEventReader(sessionID int64, client client.Client, callback eventCallback) (*EventReader, error) {
	ctx, cancel := context.WithCancel(context.Background())
	if _, err := client.GetOrCreateStream(ctx, MQTTEventStream); err != nil {
		return nil, errors.WithStack(err)
	}
	sess, err := client.NewStreamSession(ctx, sessionID, MQTTEventStream)
	if err != nil {
		return nil, err
	}
	reader, err := sess.NewReader()
	if err != nil {
		return nil, err
	}
	offset, err := sess.GetReadOffset()
	if err != nil {
		_ = reader.Close()
		return nil, err
	}
	if offset != 0 {
		if _, err := reader.Seek(offset, io.SeekStart); err != nil {
			_ = reader.Close()
			return nil, err
		}
	}
	return &EventReader{
		ctx:       ctx,
		cancel:    cancel,
		sessionID: sessionID,
		session:   sess,
		client:    client,
		callback:  callback,
		reader:    reader,
	}, nil
}

func (eReader *EventReader) handleEvent(event EventWithOffset) {
	log.WithField("event", event.event).Info("handleEvent")
	eReader.callback(event)
}

func (eReader *EventReader) readEventLoop() {
	for {
		var length int32
		if err := binary.Read(eReader.reader, binary.BigEndian, &length); err != nil {
			log.Errorf("%+v", err)
			return
		}
		if length > MaxEventSize {
			log.WithField("length", length).Fatal("eReader.reader event length error")
		}
		data := make([]byte, length)
		if _, err := io.ReadFull(eReader.reader, data); err != nil {
			log.Errorf("%+v", err)
			return
		}
		var event Event
		if err := proto.Unmarshal(data, &event); err != nil {
			log.Panic(err)
		}
		var message proto.Message
		switch event.Type {
		case Event_SubscribeEvent:
			message = &SubscribeEvent{}
		case Event_UnSubscribeEvent:
			message = &UnSubscribeEvent{}
		case Event_RetainMessageEvent:
			message = &RetainMessageEvent{}
		case Event_ClientStatusChangeEvent:
			message = &ClientStatusChangeEvent{}
		default:
			panic(fmt.Sprintf("unknown event type %d %s", event.Type, event.Data))
		}
		if err := proto.Unmarshal(event.Data, message); err != nil {
			log.Panic(err)
		}
		offset := eReader.reader.Offset()
		eReader.handleEvent(EventWithOffset{
			event:  message,
			offset: offset,
		})
	}
}

func (eReader *EventReader) Close() error {
	eReader.cancel()
	return nil
}

func (eReader *EventReader) commitReadOffset(offset int64) error {
	return eReader.session.SetReadOffset(offset)
}
