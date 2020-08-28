// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: github.com/akzj/streamIO/meta-server/store/items.proto

package store

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StreamInfoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamId       int64  `protobuf:"varint,1,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StreamServerId int64  `protobuf:"varint,3,opt,name=stream_server_id,json=streamServerId,proto3" json:"stream_server_id,omitempty"`
}

func (x *StreamInfoItem) Reset() {
	*x = StreamInfoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamInfoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamInfoItem) ProtoMessage() {}

func (x *StreamInfoItem) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamInfoItem.ProtoReflect.Descriptor instead.
func (*StreamInfoItem) Descriptor() ([]byte, []int) {
	return file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDescGZIP(), []int{0}
}

func (x *StreamInfoItem) GetStreamId() int64 {
	if x != nil {
		return x.StreamId
	}
	return 0
}

func (x *StreamInfoItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StreamInfoItem) GetStreamServerId() int64 {
	if x != nil {
		return x.StreamServerId
	}
	return 0
}

type MetaDataItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextStreamId int64 `protobuf:"varint,1,opt,name=next_stream_id,json=nextStreamId,proto3" json:"next_stream_id,omitempty"`
}

func (x *MetaDataItem) Reset() {
	*x = MetaDataItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaDataItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaDataItem) ProtoMessage() {}

func (x *MetaDataItem) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaDataItem.ProtoReflect.Descriptor instead.
func (*MetaDataItem) Descriptor() ([]byte, []int) {
	return file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDescGZIP(), []int{1}
}

func (x *MetaDataItem) GetNextStreamId() int64 {
	if x != nil {
		return x.NextStreamId
	}
	return 0
}

type SSOffsetItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId int64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	StreamId  int64 `protobuf:"varint,2,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	Offset    int64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *SSOffsetItem) Reset() {
	*x = SSOffsetItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSOffsetItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSOffsetItem) ProtoMessage() {}

func (x *SSOffsetItem) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSOffsetItem.ProtoReflect.Descriptor instead.
func (*SSOffsetItem) Descriptor() ([]byte, []int) {
	return file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDescGZIP(), []int{2}
}

func (x *SSOffsetItem) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *SSOffsetItem) GetStreamId() int64 {
	if x != nil {
		return x.StreamId
	}
	return 0
}

func (x *SSOffsetItem) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ServerInfoBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Leader bool   `protobuf:"varint,2,opt,name=leader,proto3" json:"leader,omitempty"`
	Addr   string `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *ServerInfoBase) Reset() {
	*x = ServerInfoBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfoBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfoBase) ProtoMessage() {}

func (x *ServerInfoBase) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfoBase.ProtoReflect.Descriptor instead.
func (*ServerInfoBase) Descriptor() ([]byte, []int) {
	return file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDescGZIP(), []int{3}
}

func (x *ServerInfoBase) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ServerInfoBase) GetLeader() bool {
	if x != nil {
		return x.Leader
	}
	return false
}

func (x *ServerInfoBase) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type StreamServerInfoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *ServerInfoBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *StreamServerInfoItem) Reset() {
	*x = StreamServerInfoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamServerInfoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamServerInfoItem) ProtoMessage() {}

func (x *StreamServerInfoItem) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamServerInfoItem.ProtoReflect.Descriptor instead.
func (*StreamServerInfoItem) Descriptor() ([]byte, []int) {
	return file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDescGZIP(), []int{4}
}

func (x *StreamServerInfoItem) GetBase() *ServerInfoBase {
	if x != nil {
		return x.Base
	}
	return nil
}

type MetaServerInfoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *ServerInfoBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *MetaServerInfoItem) Reset() {
	*x = MetaServerInfoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaServerInfoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaServerInfoItem) ProtoMessage() {}

func (x *MetaServerInfoItem) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaServerInfoItem.ProtoReflect.Descriptor instead.
func (*MetaServerInfoItem) Descriptor() ([]byte, []int) {
	return file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDescGZIP(), []int{5}
}

func (x *MetaServerInfoItem) GetBase() *ServerInfoBase {
	if x != nil {
		return x.Base
	}
	return nil
}

type StreamServerHeartbeatItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerInfoBase *ServerInfoBase      `protobuf:"bytes,1,opt,name=ServerInfoBase,proto3" json:"ServerInfoBase,omitempty"`
	Timestamp      *timestamp.Timestamp `protobuf:"bytes,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *StreamServerHeartbeatItem) Reset() {
	*x = StreamServerHeartbeatItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamServerHeartbeatItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamServerHeartbeatItem) ProtoMessage() {}

func (x *StreamServerHeartbeatItem) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamServerHeartbeatItem.ProtoReflect.Descriptor instead.
func (*StreamServerHeartbeatItem) Descriptor() ([]byte, []int) {
	return file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDescGZIP(), []int{6}
}

func (x *StreamServerHeartbeatItem) GetServerInfoBase() *ServerInfoBase {
	if x != nil {
		return x.ServerInfoBase
	}
	return nil
}

func (x *StreamServerHeartbeatItem) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_github_com_akzj_streamIO_meta_server_store_items_proto protoreflect.FileDescriptor

var file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x7a,
	0x6a, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x4f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x0e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61,
	0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x6e, 0x65, 0x78, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x0c,
	0x53, 0x53, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x4c, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x61,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x3b,
	0x0a, 0x14, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x12, 0x4d,
	0x65, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x19, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x37, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x52, 0x0e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x7a, 0x6a, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x49, 0x4f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDescOnce sync.Once
	file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDescData = file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDesc
)

func file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDescGZIP() []byte {
	file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDescOnce.Do(func() {
		file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDescData)
	})
	return file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDescData
}

var file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_akzj_streamIO_meta_server_store_items_proto_goTypes = []interface{}{
	(*StreamInfoItem)(nil),            // 0: StreamInfoItem
	(*MetaDataItem)(nil),              // 1: MetaDataItem
	(*SSOffsetItem)(nil),              // 2: SSOffsetItem
	(*ServerInfoBase)(nil),            // 3: ServerInfoBase
	(*StreamServerInfoItem)(nil),      // 4: StreamServerInfoItem
	(*MetaServerInfoItem)(nil),        // 5: MetaServerInfoItem
	(*StreamServerHeartbeatItem)(nil), // 6: StreamServerHeartbeatItem
	(*timestamp.Timestamp)(nil),       // 7: google.protobuf.Timestamp
}
var file_github_com_akzj_streamIO_meta_server_store_items_proto_depIdxs = []int32{
	3, // 0: StreamServerInfoItem.base:type_name -> ServerInfoBase
	3, // 1: MetaServerInfoItem.base:type_name -> ServerInfoBase
	3, // 2: StreamServerHeartbeatItem.ServerInfoBase:type_name -> ServerInfoBase
	7, // 3: StreamServerHeartbeatItem.Timestamp:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_akzj_streamIO_meta_server_store_items_proto_init() }
func file_github_com_akzj_streamIO_meta_server_store_items_proto_init() {
	if File_github_com_akzj_streamIO_meta_server_store_items_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamInfoItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaDataItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSOffsetItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfoBase); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamServerInfoItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaServerInfoItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamServerHeartbeatItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_akzj_streamIO_meta_server_store_items_proto_goTypes,
		DependencyIndexes: file_github_com_akzj_streamIO_meta_server_store_items_proto_depIdxs,
		MessageInfos:      file_github_com_akzj_streamIO_meta_server_store_items_proto_msgTypes,
	}.Build()
	File_github_com_akzj_streamIO_meta_server_store_items_proto = out.File
	file_github_com_akzj_streamIO_meta_server_store_items_proto_rawDesc = nil
	file_github_com_akzj_streamIO_meta_server_store_items_proto_goTypes = nil
	file_github_com_akzj_streamIO_meta_server_store_items_proto_depIdxs = nil
}
