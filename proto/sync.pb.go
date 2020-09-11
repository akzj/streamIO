// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: streamIO/proto/sync.proto

package proto

import (
	context "context"
	pb "github.com/akzj/streamIO/pkg/sstore/pb"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SyncSegmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Offset int64  `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *SyncSegmentRequest) Reset() {
	*x = SyncSegmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streamIO_proto_sync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncSegmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncSegmentRequest) ProtoMessage() {}

func (x *SyncSegmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streamIO_proto_sync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncSegmentRequest.ProtoReflect.Descriptor instead.
func (*SyncSegmentRequest) Descriptor() ([]byte, []int) {
	return file_streamIO_proto_sync_proto_rawDescGZIP(), []int{0}
}

func (x *SyncSegmentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SyncSegmentRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index              int64               `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	StreamServerId     int64               `protobuf:"varint,3,opt,name=stream_server_id,json=streamServerId,proto3" json:"stream_server_id,omitempty"`
	SyncSegmentRequest *SyncSegmentRequest `protobuf:"bytes,2,opt,name=sync_segment_request,json=syncSegmentRequest,proto3" json:"sync_segment_request,omitempty"`
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streamIO_proto_sync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streamIO_proto_sync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_streamIO_proto_sync_proto_rawDescGZIP(), []int{1}
}

func (x *SyncRequest) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *SyncRequest) GetStreamServerId() int64 {
	if x != nil {
		return x.StreamServerId
	}
	return 0
}

func (x *SyncRequest) GetSyncSegmentRequest() *SyncSegmentRequest {
	if x != nil {
		return x.SyncSegmentRequest
	}
	return nil
}

type SegmentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *SegmentInfo) Reset() {
	*x = SegmentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streamIO_proto_sync_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegmentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegmentInfo) ProtoMessage() {}

func (x *SegmentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_streamIO_proto_sync_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegmentInfo.ProtoReflect.Descriptor instead.
func (*SegmentInfo) Descriptor() ([]byte, []int) {
	return file_streamIO_proto_sync_proto_rawDescGZIP(), []int{2}
}

func (x *SegmentInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SegmentInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type SegmentData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Data   []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SegmentData) Reset() {
	*x = SegmentData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streamIO_proto_sync_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegmentData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegmentData) ProtoMessage() {}

func (x *SegmentData) ProtoReflect() protoreflect.Message {
	mi := &file_streamIO_proto_sync_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegmentData.ProtoReflect.Descriptor instead.
func (*SegmentData) Descriptor() ([]byte, []int) {
	return file_streamIO_proto_sync_proto_rawDescGZIP(), []int{3}
}

func (x *SegmentData) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SegmentData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SegmentEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SegmentEnd) Reset() {
	*x = SegmentEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streamIO_proto_sync_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegmentEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegmentEnd) ProtoMessage() {}

func (x *SegmentEnd) ProtoReflect() protoreflect.Message {
	mi := &file_streamIO_proto_sync_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegmentEnd.ProtoReflect.Descriptor instead.
func (*SegmentEnd) Descriptor() ([]byte, []int) {
	return file_streamIO_proto_sync_proto_rawDescGZIP(), []int{4}
}

type SyncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SegmentInfo *SegmentInfo `protobuf:"bytes,1,opt,name=segment_info,json=segmentInfo,proto3" json:"segment_info,omitempty"`
	SegmentData *SegmentData `protobuf:"bytes,2,opt,name=segment_data,json=segmentData,proto3" json:"segment_data,omitempty"`
	SegmentEnd  *SegmentEnd  `protobuf:"bytes,3,opt,name=segment_end,json=segmentEnd,proto3" json:"segment_end,omitempty"`
	Entry       *pb.Entry    `protobuf:"bytes,4,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *SyncResponse) Reset() {
	*x = SyncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streamIO_proto_sync_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncResponse) ProtoMessage() {}

func (x *SyncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streamIO_proto_sync_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncResponse.ProtoReflect.Descriptor instead.
func (*SyncResponse) Descriptor() ([]byte, []int) {
	return file_streamIO_proto_sync_proto_rawDescGZIP(), []int{5}
}

func (x *SyncResponse) GetSegmentInfo() *SegmentInfo {
	if x != nil {
		return x.SegmentInfo
	}
	return nil
}

func (x *SyncResponse) GetSegmentData() *SegmentData {
	if x != nil {
		return x.SegmentData
	}
	return nil
}

func (x *SyncResponse) GetSegmentEnd() *SegmentEnd {
	if x != nil {
		return x.SegmentEnd
	}
	return nil
}

func (x *SyncResponse) GetEntry() *pb.Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

var File_streamIO_proto_sync_proto protoreflect.FileDescriptor

var file_streamIO_proto_sync_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x4f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x4f, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x73, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x12, 0x53, 0x79, 0x6e, 0x63, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x0b, 0x53, 0x79,
	0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x28, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x14, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x12, 0x73, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x0b, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x39, 0x0a,
	0x0b, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x22, 0xd1, 0x01, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0b, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x35,
	0x0a, 0x0c, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x0b, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x52, 0x0a, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x32, 0x4a, 0x0a, 0x0b, 0x53, 0x79,
	0x6e, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x7a, 0x6a, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x49, 0x4f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_streamIO_proto_sync_proto_rawDescOnce sync.Once
	file_streamIO_proto_sync_proto_rawDescData = file_streamIO_proto_sync_proto_rawDesc
)

func file_streamIO_proto_sync_proto_rawDescGZIP() []byte {
	file_streamIO_proto_sync_proto_rawDescOnce.Do(func() {
		file_streamIO_proto_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_streamIO_proto_sync_proto_rawDescData)
	})
	return file_streamIO_proto_sync_proto_rawDescData
}

var file_streamIO_proto_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_streamIO_proto_sync_proto_goTypes = []interface{}{
	(*SyncSegmentRequest)(nil), // 0: proto.SyncSegmentRequest
	(*SyncRequest)(nil),        // 1: proto.SyncRequest
	(*SegmentInfo)(nil),        // 2: proto.SegmentInfo
	(*SegmentData)(nil),        // 3: proto.SegmentData
	(*SegmentEnd)(nil),         // 4: proto.SegmentEnd
	(*SyncResponse)(nil),       // 5: proto.SyncResponse
	(*pb.Entry)(nil),           // 6: pb.Entry
}
var file_streamIO_proto_sync_proto_depIdxs = []int32{
	0, // 0: proto.SyncRequest.sync_segment_request:type_name -> proto.SyncSegmentRequest
	2, // 1: proto.SyncResponse.segment_info:type_name -> proto.SegmentInfo
	3, // 2: proto.SyncResponse.segment_data:type_name -> proto.SegmentData
	4, // 3: proto.SyncResponse.segment_end:type_name -> proto.SegmentEnd
	6, // 4: proto.SyncResponse.entry:type_name -> pb.Entry
	1, // 5: proto.SyncService.sync_request:input_type -> proto.SyncRequest
	5, // 6: proto.SyncService.sync_request:output_type -> proto.SyncResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_streamIO_proto_sync_proto_init() }
func file_streamIO_proto_sync_proto_init() {
	if File_streamIO_proto_sync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_streamIO_proto_sync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncSegmentRequest); i {
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
		file_streamIO_proto_sync_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRequest); i {
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
		file_streamIO_proto_sync_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegmentInfo); i {
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
		file_streamIO_proto_sync_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegmentData); i {
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
		file_streamIO_proto_sync_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegmentEnd); i {
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
		file_streamIO_proto_sync_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncResponse); i {
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
			RawDescriptor: file_streamIO_proto_sync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_streamIO_proto_sync_proto_goTypes,
		DependencyIndexes: file_streamIO_proto_sync_proto_depIdxs,
		MessageInfos:      file_streamIO_proto_sync_proto_msgTypes,
	}.Build()
	File_streamIO_proto_sync_proto = out.File
	file_streamIO_proto_sync_proto_rawDesc = nil
	file_streamIO_proto_sync_proto_goTypes = nil
	file_streamIO_proto_sync_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SyncServiceClient is the client API for SyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncServiceClient interface {
	SyncRequest(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (SyncService_SyncRequestClient, error)
}

type syncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncServiceClient(cc grpc.ClientConnInterface) SyncServiceClient {
	return &syncServiceClient{cc}
}

func (c *syncServiceClient) SyncRequest(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (SyncService_SyncRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SyncService_serviceDesc.Streams[0], "/proto.SyncService/sync_request", opts...)
	if err != nil {
		return nil, err
	}
	x := &syncServiceSyncRequestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SyncService_SyncRequestClient interface {
	Recv() (*SyncResponse, error)
	grpc.ClientStream
}

type syncServiceSyncRequestClient struct {
	grpc.ClientStream
}

func (x *syncServiceSyncRequestClient) Recv() (*SyncResponse, error) {
	m := new(SyncResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SyncServiceServer is the server API for SyncService service.
type SyncServiceServer interface {
	SyncRequest(*SyncRequest, SyncService_SyncRequestServer) error
}

// UnimplementedSyncServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSyncServiceServer struct {
}

func (*UnimplementedSyncServiceServer) SyncRequest(*SyncRequest, SyncService_SyncRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncRequest not implemented")
}

func RegisterSyncServiceServer(s *grpc.Server, srv SyncServiceServer) {
	s.RegisterService(&_SyncService_serviceDesc, srv)
}

func _SyncService_SyncRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyncRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SyncServiceServer).SyncRequest(m, &syncServiceSyncRequestServer{stream})
}

type SyncService_SyncRequestServer interface {
	Send(*SyncResponse) error
	grpc.ServerStream
}

type syncServiceSyncRequestServer struct {
	grpc.ServerStream
}

func (x *syncServiceSyncRequestServer) Send(m *SyncResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _SyncService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SyncService",
	HandlerType: (*SyncServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "sync_request",
			Handler:       _SyncService_SyncRequest_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "streamIO/proto/sync.proto",
}
