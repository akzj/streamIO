package store

import (
	"encoding/binary"
	"github.com/akzj/mmdb"
	"github.com/golang/protobuf/proto"
	"github.com/google/btree"
	"github.com/pkg/errors"
	"io"
)

const (
	MetaDataItemType = 1 + iota
	StreamInfoItemType
	SSOffsetItemType
	StreamServerInfoItemType
	StreamServerHeartbeatItemType
)

type Item interface {
	mmdb.Item
	GetType() uint16
	UnmarshalBinary(data []byte) error
}

func UnmarshalItem(data []byte) (mmdb.Item, error) {
	if len(data) < 2 {
		return nil, io.ErrUnexpectedEOF
	}
	var item Item
	switch itemType := binary.BigEndian.Uint16(data); itemType {
	case MetaDataItemType:
		item = new(MetaDataItem)
	case StreamInfoItemType:
		item = new(StreamInfoItem)
	case SSOffsetItemType:
		item = new(SSOffsetItem)
	case StreamServerInfoItemType:
		item = new(StreamServerInfoItem)
	case StreamServerHeartbeatItemType:
		item = new(StreamServerHeartbeatItem)
	default:
		return nil, errors.Errorf("unknown type %d", itemType)
	}
	if err := item.UnmarshalBinary(data[2:]); err != nil {
		return nil, err
	}
	return item, nil
}

func NewStreamInfoItem(ID int64, name string) *StreamInfoItem {
	return &StreamInfoItem{
		Name:     name,
		StreamId: ID,
	}
}

func (x *StreamInfoItem) Less(other btree.Item) bool {
	if x.GetType() != other.(Item).GetType() {
		return x.GetType() < other.(Item).GetType()
	}
	return x.Name < other.(*StreamInfoItem).Name
}

func (x *StreamInfoItem) MarshalBinary() (data []byte, err error) {
	var buffers = make([]byte, 2)
	binary.BigEndian.PutUint16(buffers, x.GetType())
	if data, err = proto.Marshal(x); err != nil {
		return nil, err
	}
	return append(buffers, data...), nil
}

func (x *StreamInfoItem) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, x)
}

func (x *StreamInfoItem) GetType() uint16 {
	return StreamInfoItemType
}

//MetaDataItem

var metaDataItemKey = &MetaDataItem{Key: 1}

func (x *MetaDataItem) Less(other btree.Item) bool {
	if x.GetType() != other.(Item).GetType() {
		return x.GetType() < other.(Item).GetType()
	}
	return x.Key < other.(*MetaDataItem).Key
}

func (x *MetaDataItem) MarshalBinary() (data []byte, err error) {
	var buffers = make([]byte, 2)
	binary.BigEndian.PutUint16(buffers, x.GetType())
	if data, err = proto.Marshal(x); err != nil {
		return nil, err
	}
	return append(buffers, data...), nil
}

func (x *MetaDataItem) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, x)
}

func (x *MetaDataItem) GetType() uint16 {
	return MetaDataItemType
}

//SSOffsetItem

func (x *SSOffsetItem) Less(other btree.Item) bool {
	if x.GetType() != other.(Item).GetType() {
		return x.GetType() < other.(Item).GetType()
	}
	if x.SessionId != other.(*SSOffsetItem).SessionId {
		return x.SessionId < other.(*SSOffsetItem).SessionId
	}
	return x.StreamId < other.(*SSOffsetItem).StreamId
}

func (x *SSOffsetItem) MarshalBinary() (data []byte, err error) {
	var buffers = make([]byte, 2)
	binary.BigEndian.PutUint16(buffers, x.GetType())
	if data, err = proto.Marshal(x); err != nil {
		return nil, err
	}
	return append(buffers, data...), nil
}

func (x *SSOffsetItem) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, x)
}

func (x *SSOffsetItem) GetType() uint16 {
	return SSOffsetItemType
}

//StreamServerInfoItem

func (x *StreamServerInfoItem) Less(other btree.Item) bool {
	if x.GetType() != other.(Item).GetType() {
		return x.GetType() < other.(Item).GetType()
	}
	return x.Base.Id < other.(*StreamServerInfoItem).Base.Id
}

func (x *StreamServerInfoItem) MarshalBinary() (data []byte, err error) {
	var buffers = make([]byte, 2)
	binary.BigEndian.PutUint16(buffers, x.GetType())
	if data, err = proto.Marshal(x); err != nil {
		return nil, err
	}
	return append(buffers, data...), nil
}

func (x *StreamServerInfoItem) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, x)
}

func (x *StreamServerInfoItem) GetType() uint16 {
	return StreamServerInfoItemType
}

//StreamServerHeartbeatItem

func (x *StreamServerHeartbeatItem) Less(other btree.Item) bool {
	if x.GetType() != other.(Item).GetType() {
		return x.GetType() < other.(Item).GetType()
	}
	return x.ServerInfoBase.Id < other.(*StreamServerHeartbeatItem).ServerInfoBase.Id
}

func (x *StreamServerHeartbeatItem) MarshalBinary() (data []byte, err error) {
	var buffers = make([]byte, 2)
	binary.BigEndian.PutUint16(buffers, x.GetType())
	if data, err = proto.Marshal(x); err != nil {
		return nil, err
	}
	return append(buffers, data...), nil
}

func (x *StreamServerHeartbeatItem) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, x)
}

func (x *StreamServerHeartbeatItem) GetType() uint16 {
	return StreamServerHeartbeatItemType
}
