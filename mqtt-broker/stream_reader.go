package mqtt_broker

import (
	"context"
	"github.com/akzj/streamIO/client"
	"github.com/akzj/streamIO/meta-server/store"
	"github.com/eclipse/paho.mqtt.golang/packets"
	log "github.com/sirupsen/logrus"
	"sync"
	"sync/atomic"
)

type streamPacketReader struct {
	ackMapLocker    sync.Mutex
	ackMap          map[uint16]int64
	Reader          client.StreamReader
	streamSession   client.StreamSession
	AckOffset       int64
	Offset          int64
	committedOffset int64
	Qos             int32
	sess            *session

	ctx    context.Context
	cancel context.CancelFunc
}

func newStreamPacketReader(ctx context.Context, sess *session,
	item *store.StreamInfoItem, qos int32, client client.Client) (*streamPacketReader, error) {
	streamSession, reader, err := client.CreateSessionAndReader(ctx, sess.MQTTSessionInfo.SessionId, item)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	return &streamPacketReader{
		ackMapLocker:    sync.Mutex{},
		ackMap:          map[uint16]int64{},
		Reader:          reader,
		streamSession:   streamSession,
		AckOffset:       reader.Offset(),
		Offset:          reader.Offset(),
		committedOffset: reader.Offset(),
		Qos:             qos,
		sess:            sess,
		ctx:             ctx,
		cancel:          cancel,
	}, nil
}

func (spp *streamPacketReader) commitOffsetLoop() {
	ackOffset := atomic.LoadInt64(&spp.AckOffset)
	if ackOffset == atomic.LoadInt64(&spp.committedOffset) {
		return
	}
	spp.streamSession.SetReadOffsetWithCb(ackOffset, func(err error) {
		if err == nil {
			atomic.StoreInt64(&spp.AckOffset, ackOffset)
		} else {
			log.Errorf("%+v", err)
			_ = spp.sess.Close()
		}
	})
}

func (spp *streamPacketReader) readPacketLoop() error {
	for {
		controlPacket, err := packets.ReadPacket(spp.Reader)
		if err != nil {
			select {
			case <-spp.ctx.Done():
				spp.sess.log.Infof("readPacketLoop qos %d done", spp.Qos)
				return nil
			default:
				return err
			}
		}
		spp.Offset = spp.Reader.Offset()
		switch packet := controlPacket.(type) {
		case *packets.PublishPacket:
			if packet.Qos == 0 {
				atomic.StoreInt64(&spp.AckOffset, spp.Offset)
			} else if packet.Qos == 1 {
				spp.ackMapLocker.Lock()
				spp.ackMap[packet.MessageID] = spp.Offset
				spp.ackMapLocker.Unlock()
			}
			if err := spp.sess.handleOutPublishPacket(packet); err != nil {
				return err
			}
		}
	}
}

func (spp *streamPacketReader) Close() {
	spp.cancel()
}

func (spp *streamPacketReader) handleAck(id uint16) {
	spp.ackMapLocker.Lock()
	if offset, _ := spp.ackMap[id]; offset > atomic.LoadInt64(&spp.AckOffset) {
		atomic.StoreInt64(&spp.AckOffset, offset)
	}
	spp.ackMapLocker.Unlock()
}
