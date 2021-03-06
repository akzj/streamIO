// Copyright 2020-2026 The streamIO Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"fmt"
	"github.com/akzj/streamIO/pkg/mmdb"
	. "github.com/akzj/streamIO/proto"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"math"
	"sort"
	"strings"
	"sync"
	"time"
)

type Store struct {
	db mmdb.DB

	//protect for metadata
	metaDataItemLocker sync.Mutex
}

func OpenStore(options mmdb.Options) *Store {
	db, err := mmdb.OpenDB(options.WithUnmarshalBinary(UnmarshalItem))
	if err != nil {
		log.Panic(err)
	}

	if err := db.Update(func(tx mmdb.Transaction) error {
		item := tx.Get(MetaDataItemKey)
		if item == nil {
			tx.ReplaceOrInsert(&MetaDataItem{NextStreamId: 1, Key: 1})
		}
		return nil
	}); err != nil {
		log.Panic(err)
	}
	return &Store{
		db: db,
	}
}

func (store *Store) CreateStream(name string) (item *StreamInfoItem, create bool, err error) {
	store.metaDataItemLocker.Lock()
	defer store.metaDataItemLocker.Unlock()
	var exist = false
	streamInfo := NewStreamInfoItem(0, name)
	heartbeatItems, err := store.ListStreamServerHeartbeat()
	if err != nil {
		return nil, false, err
	}
	if len(heartbeatItems) == 0 {
		return nil, false, errors.New("no find stream server")
	}
	sort.Slice(heartbeatItems, func(i, j int) bool {
		return heartbeatItems[i].Base.Id < heartbeatItems[j].Base.Id
	})
	err = store.db.Update(func(tx mmdb.Transaction) error {
		if item := tx.Get(streamInfo); item != nil {
			streamInfo = item.(*StreamInfoItem)
			exist = true
			return nil
		}
		var metaDataItem *MetaDataItem
		item := tx.Get(MetaDataItemKey)
		if item == nil {
			metaDataItem = new(MetaDataItem)
			metaDataItem.NextStreamId = 1
		} else {
			metaDataItem = item.(*MetaDataItem)
		}
		streamInfo.StreamId = metaDataItem.NextStreamId
		streamInfo.StreamServerId =
			heartbeatItems[streamInfo.StreamId%int64(len(heartbeatItems))].Base.Id
		metaDataItem.NextStreamId++

		tx.ReplaceOrInsert(streamInfo)
		tx.ReplaceOrInsert(metaDataItem)
		return nil
	})
	if err != nil {
		log.Warn(err)
		return nil, false, errors.WithStack(err)
	}
	if exist {
		return streamInfo, false, nil
	}
	return streamInfo, true, nil
}

func (store *Store) GetStream(name string) (*StreamInfoItem, error) {
	var streamInfoItem *StreamInfoItem
	err := store.db.View(func(tx mmdb.Transaction) error {
		if item := tx.Get(NewStreamInfoItem(0, name)); item != nil {
			streamInfoItem = item.(*StreamInfoItem)
		}
		return nil
	})
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	return streamInfoItem, nil
}

func (store *Store) SetOffSet(items []*SSOffsetItem) error {
	err := store.db.Update(func(tx mmdb.Transaction) error {
		for _, item := range items {
			tx.ReplaceOrInsert(item)
		}
		return nil
	})
	if err != nil {
		log.Warn(err)
		return errors.WithStack(err)
	}
	return nil
}

func (store *Store) GetOffset(SessionId int64, StreamId int64) (*SSOffsetItem, error) {
	var ssOffsetItem *SSOffsetItem
	err := store.db.View(func(tx mmdb.Transaction) error {
		item := tx.Get(&SSOffsetItem{SessionId: SessionId, StreamId: StreamId})
		if item != nil {
			ssOffsetItem = item.(*SSOffsetItem)
		}
		return nil
	})
	if err != nil {
		log.Warn(err)
		return nil, errors.WithStack(err)
	}
	return ssOffsetItem, nil
}

func (store *Store) DelOffset(sessionID int64, streamID int64) (*SSOffsetItem, error) {
	var item mmdb.Item
	err := store.db.Update(func(tx mmdb.Transaction) error {
		item = tx.Delete(&SSOffsetItem{SessionId: sessionID, StreamId: streamID})
		return nil
	})
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	if item == nil {
		return nil, nil
	}
	return item.(*SSOffsetItem), nil
}

func (store *Store) GetOffsets() ([]*SSOffsetItem, error) {
	var items []*SSOffsetItem
	err := store.db.View(func(tx mmdb.Transaction) error {
		tx.AscendRange(&SSOffsetItem{SessionId: 0}, &SSOffsetItem{SessionId: math.MaxInt64}, func(item mmdb.Item) bool {
			items = append(items, item.(*SSOffsetItem))
			return true
		})
		return nil
	})
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	return items, nil
}

var streamServerInfoItemBegin = &StreamServerInfoItem{
	Base: &ServerInfoBase{
		Id: 0,
	},
}
var streamServerInfoItemEnd = &StreamServerInfoItem{Base: &ServerInfoBase{
	Id: math.MaxUint32,
}}

func (store *Store) ListStreamServer() ([]*StreamServerInfoItem, error) {
	var streamServerInfoItems []*StreamServerInfoItem
	err := store.db.View(func(tx mmdb.Transaction) error {
		tx.AscendRange(streamServerInfoItemBegin, streamServerInfoItemEnd, func(item mmdb.Item) bool {
			streamServerInfoItems = append(streamServerInfoItems, item.(*StreamServerInfoItem))
			return true
		})
		return nil
	})
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	return streamServerInfoItems, nil
}

func (store *Store) GetStreamServerInfo(id int64) (*StreamServerInfoItem, error) {
	var item mmdb.Item
	err := store.db.View(func(tx mmdb.Transaction) error {
		item = tx.Get(&StreamServerInfoItem{
			Base: &ServerInfoBase{Id: id},
		})
		return nil
	})
	if err != nil {
		log.Warning(err)
		return nil, err
	}
	if item == nil {
		return nil, nil
	}
	return item.(*StreamServerInfoItem), err
}

func (store *Store) AddStreamServer(item *StreamServerInfoItem) (*StreamServerInfoItem, error) {
	err := store.db.Update(func(tx mmdb.Transaction) error {
		if item.Base.Id == 0 {
			var lastItem mmdb.Item
			tx.AscendRange(streamServerInfoItemBegin, streamServerInfoItemEnd, func(item mmdb.Item) bool {
				lastItem = item
				return true
			})
			if lastItem == nil {
				item.Base.Id = 1
			} else {
				item.Base.Id = 1 + lastItem.(*StreamServerInfoItem).Base.Id
			}
			tx.ReplaceOrInsert(item)
		} else {
			tx.ReplaceOrInsert(item)
		}
		return nil
	})
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	return item, nil
}

func (store *Store) DeleteStreamServer(item *StreamServerInfoItem) error {
	err := store.db.Update(func(tx mmdb.Transaction) error {
		tx.Delete(item)
		return nil
	})
	if err != nil {
		log.Warn(err)
		return errors.WithStack(err)
	}
	return nil
}

func (store *Store) InsertStreamServerHeartbeatItem(item *StreamServerHeartbeatItem) error {
	var find = false
	err := store.db.Update(func(tx mmdb.Transaction) error {
		if tx.Get(&StreamServerInfoItem{Base: item.Base}) == nil {
			return nil
		}
		tx.ReplaceOrInsert(item)
		find = true
		return nil
	})
	if err != nil {
		log.Warn(err)
		return err
	}
	if find == false {
		err = errors.Errorf("no find stream server ID %d", item.Base.Id)
	}
	return err
}

func (store *Store) ListStreamServerHeartbeat() ([]*StreamServerHeartbeatItem, error) {
	var items []*StreamServerHeartbeatItem
	err := store.db.View(func(tx mmdb.Transaction) error {
		tx.AscendRange(StreamServerHeartbeatItemKeyMin, StreamServerHeartbeatItemKeyMax, func(item mmdb.Item) bool {
			items = append(items, item.(*StreamServerHeartbeatItem))
			return true
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (store *Store) GetStreamServerHeartbeatItem(ID int64) (*StreamServerHeartbeatItem, error) {
	var item mmdb.Item
	err := store.db.View(func(tx mmdb.Transaction) error {
		item = tx.Get(&StreamServerHeartbeatItem{Base: &ServerInfoBase{Id: ID}})
		return nil
	})
	if err != nil {
		log.Warning(err)
		return nil, err
	}
	if item == nil {
		return nil, nil
	}
	return item.(*StreamServerHeartbeatItem), nil
}

func (store *Store) GetOrCreateMQTTSession(identifier string) (*MQTTSessionItem, bool, error) {
	var item mmdb.Item
	var create bool
	err := store.db.Update(func(tx mmdb.Transaction) error {
		item = tx.Get(&MQTTSessionItem{ClientIdentifier: identifier})
		if item == nil {
			create = true
			session := &MQTTSessionItem{
				Qos1StreamInfo:   nil,
				Qos0StreamInfo:   nil,
				SessionId:        0,
				ClientIdentifier: identifier,
				CreateTs:         time.Now().Unix(),
				AccessTs:         time.Now().Unix(),
				Topics:           nil,
			}
			for retry := 0; retry < 1000; retry++ {
				streamName := strings.ReplaceAll(uuid.New().String(), "-", "")
				info, create, err := store.CreateStream(streamName)
				if err != nil {
					log.Error(err)
					return err
				}
				if create == false {
					log.Warnf("create stream %s exist ....", streamName)
					continue
				}
				if session.Qos0StreamInfo == nil {
					session.Qos0StreamInfo = info
					continue
				}
				session.Qos1StreamInfo = info
				break
			}
			session.SessionId = session.Qos1StreamInfo.StreamId
			tx.ReplaceOrInsert(session)
			item = session
		}
		return nil
	})
	if err != nil {
		log.Warn(err)
		return nil, false, err
	}
	return item.(*MQTTSessionItem), create, nil
}

func (store *Store) DeleteMQTTClientSession(identifier string) (*MQTTSessionItem, error) {
	var item mmdb.Item
	err := store.db.Update(func(tx mmdb.Transaction) error {
		item = tx.Delete(&MQTTSessionItem{ClientIdentifier: identifier})
		if item != nil {
			session := item.(*MQTTSessionItem)
			tx.Delete(session.Qos0StreamInfo)
			tx.Delete(session.Qos1StreamInfo)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, nil
	}
	return item.(*MQTTSessionItem), nil
}

func (store *Store) UpdateMQTTClientSession(ClientIdentifier string,
	UnSubscribe []string, Subscribe map[string]int32) error {
	return store.db.Update(func(tx mmdb.Transaction) error {
		item := tx.Get(&MQTTSessionItem{ClientIdentifier: ClientIdentifier})
		if item == nil {
			return fmt.Errorf("no find mqtt client session with %s",
				ClientIdentifier)
		}
		session := item.(*MQTTSessionItem).Clone()
		for _, topic := range UnSubscribe {
			delete(session.Topics, topic)
		}
		for topic, qos := range Subscribe {
			session.Topics[topic] = qos
		}
		session.AccessTs = time.Now().Unix()
		tx.ReplaceOrInsert(session)
		return nil
	})
}
