// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: clock.proto

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

// Represents a clock frame for a given filter. Clock frames are the primary
// sequencing mechanism upon which the network derives consensus. As the master
// pulse clock, this provides deterministic but random leader election. At the
// data pulse clock level, this provides the same, within a quorum for data
// sequencers.
type ClockFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The filter is used as a domain separator for input, but in the context of
	// verifiable delay functions, is simply prepended to the input field as input
	// for the VDF.
	Filter []byte `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// A strictly monotonically-increasing frame number. Used for culling old
	// frames past a configurable cutoff point.
	FrameNumber uint64 `protobuf:"varint,2,opt,name=frame_number,json=frameNumber,proto3" json:"frame_number,omitempty"`
	// The self-reported timestamp from the proof publisher, encoded as an int64
	// of the Unix epoch in milliseconds. Should be good until
	// 292278994-08-17 07:12:55.807, at which point, this is someone else's
	// problem. Timestamps are imperfect, but smoothed in a rolling window to
	// ensure a network and quorum-stable difficulty adjustment. Anomalies are
	// bounded such that a timestamp beyond ten times the average issuance rate
	// is discarded in preference to the runner up electees, unless there is
	// simply no alternative available (for example, if a network outage occurred
	// from an upgrade or bug).
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The difficulty level used for the frame. Difficulty is calculated based on
	// the previous 60 timestamps correlated with difficulties, such that the
	// interval smooths out to align to the type-defined rate. This is expected to
	// increase subtly with clock speed and future hardware implementations, but
	// due to incentive alignment associated with data proofs, not fastest clock
	// in the west, should be gradual.
	Difficulty uint32 `protobuf:"varint,4,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	// The selector value of the previous frame's output, produced as a Poseidon
	// hash of the output.
	ParentSelector []byte `protobuf:"bytes,5,opt,name=parent_selector,json=parentSelector,proto3" json:"parent_selector,omitempty"`
	// The input data used for the VDF proof. For the master pulse clock, this is
	// the concatenation of the filter, frame number, difficulty, previous frame's
	// output, and the rolled state proof commitment input. For the data pulse
	// clocks, this is the concatenation of the filter, frame number, timestamp,
	// difficulty, issuer address, previous frame's output, along with data
	// mutation and availability proofs. Elements that are also in the fields of
	// the clock frame are not included in this field due to redundancy. For the
	// ceremony phase, this is a singular clock fusing master and data pulses.
	Input []byte `protobuf:"bytes,6,opt,name=input,proto3" json:"input,omitempty"`
	// The output data from the VDF, serialized as bytes. For Wesolowski, this is
	// an encoding of the 258 byte Y value concatenated with the 258 byte proof
	// value.
	Output []byte `protobuf:"bytes,7,opt,name=output,proto3" json:"output,omitempty"`
	// Any aggregate proofs to be rolled into the committed clock frame.
	AggregateProofs []*InclusionAggregateProof `protobuf:"bytes,8,rep,name=aggregate_proofs,json=aggregateProofs,proto3" json:"aggregate_proofs,omitempty"`
	// The signature of the proof issuer.
	//
	// Types that are assignable to PublicKeySignature:
	//
	//	*ClockFrame_PublicKeySignatureEd448
	PublicKeySignature isClockFrame_PublicKeySignature `protobuf_oneof:"public_key_signature"`
}

func (x *ClockFrame) Reset() {
	*x = ClockFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClockFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockFrame) ProtoMessage() {}

func (x *ClockFrame) ProtoReflect() protoreflect.Message {
	mi := &file_clock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockFrame.ProtoReflect.Descriptor instead.
func (*ClockFrame) Descriptor() ([]byte, []int) {
	return file_clock_proto_rawDescGZIP(), []int{0}
}

func (x *ClockFrame) GetFilter() []byte {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ClockFrame) GetFrameNumber() uint64 {
	if x != nil {
		return x.FrameNumber
	}
	return 0
}

func (x *ClockFrame) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ClockFrame) GetDifficulty() uint32 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *ClockFrame) GetParentSelector() []byte {
	if x != nil {
		return x.ParentSelector
	}
	return nil
}

func (x *ClockFrame) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *ClockFrame) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *ClockFrame) GetAggregateProofs() []*InclusionAggregateProof {
	if x != nil {
		return x.AggregateProofs
	}
	return nil
}

func (m *ClockFrame) GetPublicKeySignature() isClockFrame_PublicKeySignature {
	if m != nil {
		return m.PublicKeySignature
	}
	return nil
}

func (x *ClockFrame) GetPublicKeySignatureEd448() *Ed448Signature {
	if x, ok := x.GetPublicKeySignature().(*ClockFrame_PublicKeySignatureEd448); ok {
		return x.PublicKeySignatureEd448
	}
	return nil
}

type isClockFrame_PublicKeySignature interface {
	isClockFrame_PublicKeySignature()
}

type ClockFrame_PublicKeySignatureEd448 struct {
	PublicKeySignatureEd448 *Ed448Signature `protobuf:"bytes,9,opt,name=public_key_signature_ed448,json=publicKeySignatureEd448,proto3,oneof"`
}

func (*ClockFrame_PublicKeySignatureEd448) isClockFrame_PublicKeySignature() {}

// Represents a request for a range of clock frames. Used to stay synchronized
// to the latest state.
type ClockFramesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The filter is used as a domain separator for input to the frames.
	Filter []byte `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The earliest frame in the range requested.
	FromFrameNumber uint64 `protobuf:"varint,2,opt,name=from_frame_number,json=fromFrameNumber,proto3" json:"from_frame_number,omitempty"`
	// The latest frame in the range requested, if provided. Capped to a maximum
	// size of 128 frames.
	ToFrameNumber uint64 `protobuf:"varint,3,opt,name=to_frame_number,json=toFrameNumber,proto3" json:"to_frame_number,omitempty"`
}

func (x *ClockFramesRequest) Reset() {
	*x = ClockFramesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClockFramesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockFramesRequest) ProtoMessage() {}

func (x *ClockFramesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockFramesRequest.ProtoReflect.Descriptor instead.
func (*ClockFramesRequest) Descriptor() ([]byte, []int) {
	return file_clock_proto_rawDescGZIP(), []int{1}
}

func (x *ClockFramesRequest) GetFilter() []byte {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ClockFramesRequest) GetFromFrameNumber() uint64 {
	if x != nil {
		return x.FromFrameNumber
	}
	return 0
}

func (x *ClockFramesRequest) GetToFrameNumber() uint64 {
	if x != nil {
		return x.ToFrameNumber
	}
	return 0
}

// Represents a response for a range of clock frames. Used to stay synchronized
// to the latest state.
type ClockFramesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The filter is used as a domain separator for input to the frames.
	Filter []byte `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The earliest frame in the range response. Paginated to a maximum size of
	// 128 frames per response.
	FromFrameNumber uint64 `protobuf:"varint,2,opt,name=from_frame_number,json=fromFrameNumber,proto3" json:"from_frame_number,omitempty"`
	// The latest frame in the range response.
	ToFrameNumber uint64 `protobuf:"varint,3,opt,name=to_frame_number,json=toFrameNumber,proto3" json:"to_frame_number,omitempty"`
	// The set of clock frames within the provided range.
	ClockFrames []*ClockFrame `protobuf:"bytes,4,rep,name=clock_frames,json=clockFrames,proto3" json:"clock_frames,omitempty"`
}

func (x *ClockFramesResponse) Reset() {
	*x = ClockFramesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClockFramesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockFramesResponse) ProtoMessage() {}

func (x *ClockFramesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockFramesResponse.ProtoReflect.Descriptor instead.
func (*ClockFramesResponse) Descriptor() ([]byte, []int) {
	return file_clock_proto_rawDescGZIP(), []int{2}
}

func (x *ClockFramesResponse) GetFilter() []byte {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ClockFramesResponse) GetFromFrameNumber() uint64 {
	if x != nil {
		return x.FromFrameNumber
	}
	return 0
}

func (x *ClockFramesResponse) GetToFrameNumber() uint64 {
	if x != nil {
		return x.ToFrameNumber
	}
	return 0
}

func (x *ClockFramesResponse) GetClockFrames() []*ClockFrame {
	if x != nil {
		return x.ClockFrames
	}
	return nil
}

var File_clock_proto protoreflect.FileDescriptor

var file_clock_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x71,
	0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x62, 0x1a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbc, 0x03, 0x0a, 0x0a, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69,
	0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x5e, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x71, 0x75,
	0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69,
	0x6f, 0x6e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x52, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x73, 0x12, 0x66, 0x0a, 0x1a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x64, 0x34, 0x34, 0x38, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69,
	0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x62, 0x2e,
	0x45, 0x64, 0x34, 0x34, 0x38, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48, 0x00,
	0x52, 0x17, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x45, 0x64, 0x34, 0x34, 0x38, 0x42, 0x16, 0x0a, 0x14, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0x80, 0x01, 0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x2a, 0x0a, 0x11, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x66, 0x72, 0x6f,
	0x6d, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f,
	0x74, 0x6f, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x74, 0x6f, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0xca, 0x01, 0x0a, 0x13, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x74, 0x6f, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x0c, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x52, 0x0b, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x73, 0x42, 0x3a, 0x5a, 0x38, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x71, 0x75, 0x69, 0x6c,
	0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x75, 0x69, 0x6c, 0x69,
	0x62, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clock_proto_rawDescOnce sync.Once
	file_clock_proto_rawDescData = file_clock_proto_rawDesc
)

func file_clock_proto_rawDescGZIP() []byte {
	file_clock_proto_rawDescOnce.Do(func() {
		file_clock_proto_rawDescData = protoimpl.X.CompressGZIP(file_clock_proto_rawDescData)
	})
	return file_clock_proto_rawDescData
}

var file_clock_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_clock_proto_goTypes = []interface{}{
	(*ClockFrame)(nil),              // 0: quilibrium.node.clock.pb.ClockFrame
	(*ClockFramesRequest)(nil),      // 1: quilibrium.node.clock.pb.ClockFramesRequest
	(*ClockFramesResponse)(nil),     // 2: quilibrium.node.clock.pb.ClockFramesResponse
	(*InclusionAggregateProof)(nil), // 3: quilibrium.node.channel.pb.InclusionAggregateProof
	(*Ed448Signature)(nil),          // 4: quilibrium.node.keys.pb.Ed448Signature
}
var file_clock_proto_depIdxs = []int32{
	3, // 0: quilibrium.node.clock.pb.ClockFrame.aggregate_proofs:type_name -> quilibrium.node.channel.pb.InclusionAggregateProof
	4, // 1: quilibrium.node.clock.pb.ClockFrame.public_key_signature_ed448:type_name -> quilibrium.node.keys.pb.Ed448Signature
	0, // 2: quilibrium.node.clock.pb.ClockFramesResponse.clock_frames:type_name -> quilibrium.node.clock.pb.ClockFrame
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_clock_proto_init() }
func file_clock_proto_init() {
	if File_clock_proto != nil {
		return
	}
	file_channel_proto_init()
	file_keys_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_clock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClockFrame); i {
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
		file_clock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClockFramesRequest); i {
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
		file_clock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClockFramesResponse); i {
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
	file_clock_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ClockFrame_PublicKeySignatureEd448)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_clock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_clock_proto_goTypes,
		DependencyIndexes: file_clock_proto_depIdxs,
		MessageInfos:      file_clock_proto_msgTypes,
	}.Build()
	File_clock_proto = out.File
	file_clock_proto_rawDesc = nil
	file_clock_proto_goTypes = nil
	file_clock_proto_depIdxs = nil
}