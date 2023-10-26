// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: node.proto

package protobufs

import (
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

type GetFramesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter            []byte `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	FromFrameNumber   uint64 `protobuf:"varint,2,opt,name=from_frame_number,json=fromFrameNumber,proto3" json:"from_frame_number,omitempty"`
	ToFrameNumber     uint64 `protobuf:"varint,3,opt,name=to_frame_number,json=toFrameNumber,proto3" json:"to_frame_number,omitempty"`
	IncludeCandidates bool   `protobuf:"varint,4,opt,name=include_candidates,json=includeCandidates,proto3" json:"include_candidates,omitempty"`
}

func (x *GetFramesRequest) Reset() {
	*x = GetFramesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFramesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFramesRequest) ProtoMessage() {}

func (x *GetFramesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFramesRequest.ProtoReflect.Descriptor instead.
func (*GetFramesRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

func (x *GetFramesRequest) GetFilter() []byte {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *GetFramesRequest) GetFromFrameNumber() uint64 {
	if x != nil {
		return x.FromFrameNumber
	}
	return 0
}

func (x *GetFramesRequest) GetToFrameNumber() uint64 {
	if x != nil {
		return x.ToFrameNumber
	}
	return 0
}

func (x *GetFramesRequest) GetIncludeCandidates() bool {
	if x != nil {
		return x.IncludeCandidates
	}
	return false
}

type GetFrameInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter      []byte `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	FrameNumber uint64 `protobuf:"varint,2,opt,name=frame_number,json=frameNumber,proto3" json:"frame_number,omitempty"`
	Selector    []byte `protobuf:"bytes,3,opt,name=selector,proto3" json:"selector,omitempty"`
}

func (x *GetFrameInfoRequest) Reset() {
	*x = GetFrameInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFrameInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFrameInfoRequest) ProtoMessage() {}

func (x *GetFrameInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFrameInfoRequest.ProtoReflect.Descriptor instead.
func (*GetFrameInfoRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{1}
}

func (x *GetFrameInfoRequest) GetFilter() []byte {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *GetFrameInfoRequest) GetFrameNumber() uint64 {
	if x != nil {
		return x.FrameNumber
	}
	return 0
}

func (x *GetFrameInfoRequest) GetSelector() []byte {
	if x != nil {
		return x.Selector
	}
	return nil
}

type GetPeerInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPeerInfoRequest) Reset() {
	*x = GetPeerInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeerInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeerInfoRequest) ProtoMessage() {}

func (x *GetPeerInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeerInfoRequest.ProtoReflect.Descriptor instead.
func (*GetPeerInfoRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{2}
}

type GetNetworkInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNetworkInfoRequest) Reset() {
	*x = GetNetworkInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNetworkInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNetworkInfoRequest) ProtoMessage() {}

func (x *GetNetworkInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNetworkInfoRequest.ProtoReflect.Descriptor instead.
func (*GetNetworkInfoRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{3}
}

type FramesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TruncatedClockFrames []*ClockFrame `protobuf:"bytes,1,rep,name=truncated_clock_frames,json=truncatedClockFrames,proto3" json:"truncated_clock_frames,omitempty"`
}

func (x *FramesResponse) Reset() {
	*x = FramesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FramesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FramesResponse) ProtoMessage() {}

func (x *FramesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FramesResponse.ProtoReflect.Descriptor instead.
func (*FramesResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{4}
}

func (x *FramesResponse) GetTruncatedClockFrames() []*ClockFrame {
	if x != nil {
		return x.TruncatedClockFrames
	}
	return nil
}

type FrameInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClockFrame *ClockFrame `protobuf:"bytes,1,opt,name=clock_frame,json=clockFrame,proto3" json:"clock_frame,omitempty"`
}

func (x *FrameInfoResponse) Reset() {
	*x = FrameInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrameInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrameInfoResponse) ProtoMessage() {}

func (x *FrameInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrameInfoResponse.ProtoReflect.Descriptor instead.
func (*FrameInfoResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{5}
}

func (x *FrameInfoResponse) GetClockFrame() *ClockFrame {
	if x != nil {
		return x.ClockFrame
	}
	return nil
}

type PeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId     []byte   `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Multiaddrs []string `protobuf:"bytes,2,rep,name=multiaddrs,proto3" json:"multiaddrs,omitempty"`
	MaxFrame   uint64   `protobuf:"varint,3,opt,name=max_frame,json=maxFrame,proto3" json:"max_frame,omitempty"`
	Timestamp  int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PeerInfo) Reset() {
	*x = PeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerInfo) ProtoMessage() {}

func (x *PeerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerInfo.ProtoReflect.Descriptor instead.
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{6}
}

func (x *PeerInfo) GetPeerId() []byte {
	if x != nil {
		return x.PeerId
	}
	return nil
}

func (x *PeerInfo) GetMultiaddrs() []string {
	if x != nil {
		return x.Multiaddrs
	}
	return nil
}

func (x *PeerInfo) GetMaxFrame() uint64 {
	if x != nil {
		return x.MaxFrame
	}
	return 0
}

func (x *PeerInfo) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type PeerInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerInfo              []*PeerInfo `protobuf:"bytes,1,rep,name=peer_info,json=peerInfo,proto3" json:"peer_info,omitempty"`
	UncooperativePeerInfo []*PeerInfo `protobuf:"bytes,2,rep,name=uncooperative_peer_info,json=uncooperativePeerInfo,proto3" json:"uncooperative_peer_info,omitempty"`
}

func (x *PeerInfoResponse) Reset() {
	*x = PeerInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerInfoResponse) ProtoMessage() {}

func (x *PeerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerInfoResponse.ProtoReflect.Descriptor instead.
func (*PeerInfoResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{7}
}

func (x *PeerInfoResponse) GetPeerInfo() []*PeerInfo {
	if x != nil {
		return x.PeerInfo
	}
	return nil
}

func (x *PeerInfoResponse) GetUncooperativePeerInfo() []*PeerInfo {
	if x != nil {
		return x.UncooperativePeerInfo
	}
	return nil
}

type NetworkInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId     []byte   `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Multiaddrs []string `protobuf:"bytes,2,rep,name=multiaddrs,proto3" json:"multiaddrs,omitempty"`
	PeerScore  float64  `protobuf:"fixed64,3,opt,name=peer_score,json=peerScore,proto3" json:"peer_score,omitempty"`
}

func (x *NetworkInfo) Reset() {
	*x = NetworkInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInfo) ProtoMessage() {}

func (x *NetworkInfo) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInfo.ProtoReflect.Descriptor instead.
func (*NetworkInfo) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{8}
}

func (x *NetworkInfo) GetPeerId() []byte {
	if x != nil {
		return x.PeerId
	}
	return nil
}

func (x *NetworkInfo) GetMultiaddrs() []string {
	if x != nil {
		return x.Multiaddrs
	}
	return nil
}

func (x *NetworkInfo) GetPeerScore() float64 {
	if x != nil {
		return x.PeerScore
	}
	return 0
}

type NetworkInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkInfo []*NetworkInfo `protobuf:"bytes,1,rep,name=network_info,json=networkInfo,proto3" json:"network_info,omitempty"`
}

func (x *NetworkInfoResponse) Reset() {
	*x = NetworkInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInfoResponse) ProtoMessage() {}

func (x *NetworkInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInfoResponse.ProtoReflect.Descriptor instead.
func (*NetworkInfoResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{9}
}

func (x *NetworkInfoResponse) GetNetworkInfo() []*NetworkInfo {
	if x != nil {
		return x.NetworkInfo
	}
	return nil
}

var File_node_proto protoreflect.FileDescriptor

var file_node_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x71, 0x75,
	0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x62, 0x1a, 0x0b, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x11, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x66, 0x72, 0x6f, 0x6d,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x74,
	0x6f, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x74, 0x6f, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x63,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x11, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x22, 0x6c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x6c, 0x0a, 0x0e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5a, 0x0a, 0x16, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f,
	0x63, 0x6b, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x14, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x5a, 0x0a,
	0x11, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62,
	0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x0a, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x22, 0x7e, 0x0a, 0x08, 0x50, 0x65, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x61, 0x64, 0x64, 0x72, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xad, 0x01, 0x0a, 0x10, 0x50, 0x65,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x09, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x59,
	0x0a, 0x17, 0x75, 0x6e, 0x63, 0x6f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x15, 0x75, 0x6e, 0x63, 0x6f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x65, 0x0a, 0x0b, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x61, 0x64, 0x64, 0x72,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x70, 0x65, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x22, 0x5e, 0x0a, 0x13, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x32, 0xaf, 0x03, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x29, 0x2e,
	0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69,
	0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x68, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2c, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x2e, 0x71, 0x75, 0x69,
	0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62,
	0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x6e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75,
	0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75,
	0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x71, 0x75, 0x69,
	0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x75, 0x69, 0x6c,
	0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_proto_rawDescOnce sync.Once
	file_node_proto_rawDescData = file_node_proto_rawDesc
)

func file_node_proto_rawDescGZIP() []byte {
	file_node_proto_rawDescOnce.Do(func() {
		file_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_proto_rawDescData)
	})
	return file_node_proto_rawDescData
}

var file_node_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_node_proto_goTypes = []interface{}{
	(*GetFramesRequest)(nil),      // 0: quilibrium.node.node.pb.GetFramesRequest
	(*GetFrameInfoRequest)(nil),   // 1: quilibrium.node.node.pb.GetFrameInfoRequest
	(*GetPeerInfoRequest)(nil),    // 2: quilibrium.node.node.pb.GetPeerInfoRequest
	(*GetNetworkInfoRequest)(nil), // 3: quilibrium.node.node.pb.GetNetworkInfoRequest
	(*FramesResponse)(nil),        // 4: quilibrium.node.node.pb.FramesResponse
	(*FrameInfoResponse)(nil),     // 5: quilibrium.node.node.pb.FrameInfoResponse
	(*PeerInfo)(nil),              // 6: quilibrium.node.node.pb.PeerInfo
	(*PeerInfoResponse)(nil),      // 7: quilibrium.node.node.pb.PeerInfoResponse
	(*NetworkInfo)(nil),           // 8: quilibrium.node.node.pb.NetworkInfo
	(*NetworkInfoResponse)(nil),   // 9: quilibrium.node.node.pb.NetworkInfoResponse
	(*ClockFrame)(nil),            // 10: quilibrium.node.clock.pb.ClockFrame
}
var file_node_proto_depIdxs = []int32{
	10, // 0: quilibrium.node.node.pb.FramesResponse.truncated_clock_frames:type_name -> quilibrium.node.clock.pb.ClockFrame
	10, // 1: quilibrium.node.node.pb.FrameInfoResponse.clock_frame:type_name -> quilibrium.node.clock.pb.ClockFrame
	6,  // 2: quilibrium.node.node.pb.PeerInfoResponse.peer_info:type_name -> quilibrium.node.node.pb.PeerInfo
	6,  // 3: quilibrium.node.node.pb.PeerInfoResponse.uncooperative_peer_info:type_name -> quilibrium.node.node.pb.PeerInfo
	8,  // 4: quilibrium.node.node.pb.NetworkInfoResponse.network_info:type_name -> quilibrium.node.node.pb.NetworkInfo
	0,  // 5: quilibrium.node.node.pb.NodeService.GetFrames:input_type -> quilibrium.node.node.pb.GetFramesRequest
	1,  // 6: quilibrium.node.node.pb.NodeService.GetFrameInfo:input_type -> quilibrium.node.node.pb.GetFrameInfoRequest
	2,  // 7: quilibrium.node.node.pb.NodeService.GetPeerInfo:input_type -> quilibrium.node.node.pb.GetPeerInfoRequest
	3,  // 8: quilibrium.node.node.pb.NodeService.GetNetworkInfo:input_type -> quilibrium.node.node.pb.GetNetworkInfoRequest
	4,  // 9: quilibrium.node.node.pb.NodeService.GetFrames:output_type -> quilibrium.node.node.pb.FramesResponse
	5,  // 10: quilibrium.node.node.pb.NodeService.GetFrameInfo:output_type -> quilibrium.node.node.pb.FrameInfoResponse
	7,  // 11: quilibrium.node.node.pb.NodeService.GetPeerInfo:output_type -> quilibrium.node.node.pb.PeerInfoResponse
	9,  // 12: quilibrium.node.node.pb.NodeService.GetNetworkInfo:output_type -> quilibrium.node.node.pb.NetworkInfoResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_node_proto_init() }
func file_node_proto_init() {
	if File_node_proto != nil {
		return
	}
	file_clock_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFramesRequest); i {
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
		file_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFrameInfoRequest); i {
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
		file_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPeerInfoRequest); i {
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
		file_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNetworkInfoRequest); i {
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
		file_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FramesResponse); i {
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
		file_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrameInfoResponse); i {
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
		file_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerInfo); i {
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
		file_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerInfoResponse); i {
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
		file_node_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInfo); i {
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
		file_node_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInfoResponse); i {
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
			RawDescriptor: file_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_proto_goTypes,
		DependencyIndexes: file_node_proto_depIdxs,
		MessageInfos:      file_node_proto_msgTypes,
	}.Build()
	File_node_proto = out.File
	file_node_proto_rawDesc = nil
	file_node_proto_goTypes = nil
	file_node_proto_depIdxs = nil
}
