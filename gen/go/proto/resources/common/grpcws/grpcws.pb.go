// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: resources/common/grpcws/grpcws.proto

package grpcws

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

type GrpcFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamId uint32 `protobuf:"varint,1,opt,name=streamId,proto3" json:"streamId,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*GrpcFrame_Header
	//	*GrpcFrame_Body
	//	*GrpcFrame_Complete
	//	*GrpcFrame_Failure
	//	*GrpcFrame_Cancel
	Payload isGrpcFrame_Payload `protobuf_oneof:"payload"`
}

func (x *GrpcFrame) Reset() {
	*x = GrpcFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_common_grpcws_grpcws_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcFrame) ProtoMessage() {}

func (x *GrpcFrame) ProtoReflect() protoreflect.Message {
	mi := &file_resources_common_grpcws_grpcws_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcFrame.ProtoReflect.Descriptor instead.
func (*GrpcFrame) Descriptor() ([]byte, []int) {
	return file_resources_common_grpcws_grpcws_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcFrame) GetStreamId() uint32 {
	if x != nil {
		return x.StreamId
	}
	return 0
}

func (m *GrpcFrame) GetPayload() isGrpcFrame_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *GrpcFrame) GetHeader() *Header {
	if x, ok := x.GetPayload().(*GrpcFrame_Header); ok {
		return x.Header
	}
	return nil
}

func (x *GrpcFrame) GetBody() *Body {
	if x, ok := x.GetPayload().(*GrpcFrame_Body); ok {
		return x.Body
	}
	return nil
}

func (x *GrpcFrame) GetComplete() *Complete {
	if x, ok := x.GetPayload().(*GrpcFrame_Complete); ok {
		return x.Complete
	}
	return nil
}

func (x *GrpcFrame) GetFailure() *Failure {
	if x, ok := x.GetPayload().(*GrpcFrame_Failure); ok {
		return x.Failure
	}
	return nil
}

func (x *GrpcFrame) GetCancel() *Cancel {
	if x, ok := x.GetPayload().(*GrpcFrame_Cancel); ok {
		return x.Cancel
	}
	return nil
}

type isGrpcFrame_Payload interface {
	isGrpcFrame_Payload()
}

type GrpcFrame_Header struct {
	Header *Header `protobuf:"bytes,3,opt,name=header,proto3,oneof"`
}

type GrpcFrame_Body struct {
	Body *Body `protobuf:"bytes,4,opt,name=body,proto3,oneof"`
}

type GrpcFrame_Complete struct {
	Complete *Complete `protobuf:"bytes,5,opt,name=complete,proto3,oneof"`
}

type GrpcFrame_Failure struct {
	Failure *Failure `protobuf:"bytes,6,opt,name=failure,proto3,oneof"`
}

type GrpcFrame_Cancel struct {
	Cancel *Cancel `protobuf:"bytes,7,opt,name=cancel,proto3,oneof"`
}

func (*GrpcFrame_Header) isGrpcFrame_Payload() {}

func (*GrpcFrame_Body) isGrpcFrame_Payload() {}

func (*GrpcFrame_Complete) isGrpcFrame_Payload() {}

func (*GrpcFrame_Failure) isGrpcFrame_Payload() {}

func (*GrpcFrame_Cancel) isGrpcFrame_Payload() {}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation string                  `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Headers   map[string]*HeaderValue `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Status    int32                   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_common_grpcws_grpcws_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_resources_common_grpcws_grpcws_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_resources_common_grpcws_grpcws_proto_rawDescGZIP(), []int{1}
}

func (x *Header) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *Header) GetHeaders() map[string]*HeaderValue {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Header) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type HeaderValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *HeaderValue) Reset() {
	*x = HeaderValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_common_grpcws_grpcws_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValue) ProtoMessage() {}

func (x *HeaderValue) ProtoReflect() protoreflect.Message {
	mi := &file_resources_common_grpcws_grpcws_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValue.ProtoReflect.Descriptor instead.
func (*HeaderValue) Descriptor() ([]byte, []int) {
	return file_resources_common_grpcws_grpcws_proto_rawDescGZIP(), []int{2}
}

func (x *HeaderValue) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Complete bool   `protobuf:"varint,2,opt,name=complete,proto3" json:"complete,omitempty"`
}

func (x *Body) Reset() {
	*x = Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_common_grpcws_grpcws_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Body) ProtoMessage() {}

func (x *Body) ProtoReflect() protoreflect.Message {
	mi := &file_resources_common_grpcws_grpcws_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Body.ProtoReflect.Descriptor instead.
func (*Body) Descriptor() ([]byte, []int) {
	return file_resources_common_grpcws_grpcws_proto_rawDescGZIP(), []int{3}
}

func (x *Body) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Body) GetComplete() bool {
	if x != nil {
		return x.Complete
	}
	return false
}

type Complete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Complete) Reset() {
	*x = Complete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_common_grpcws_grpcws_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Complete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Complete) ProtoMessage() {}

func (x *Complete) ProtoReflect() protoreflect.Message {
	mi := &file_resources_common_grpcws_grpcws_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Complete.ProtoReflect.Descriptor instead.
func (*Complete) Descriptor() ([]byte, []int) {
	return file_resources_common_grpcws_grpcws_proto_rawDescGZIP(), []int{4}
}

type Failure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO use enum errors to properly define the protocol
	ErrorMessage string                  `protobuf:"bytes,1,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	ErrorStatus  string                  `protobuf:"bytes,2,opt,name=errorStatus,proto3" json:"errorStatus,omitempty"`
	Headers      map[string]*HeaderValue `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Failure) Reset() {
	*x = Failure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_common_grpcws_grpcws_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Failure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Failure) ProtoMessage() {}

func (x *Failure) ProtoReflect() protoreflect.Message {
	mi := &file_resources_common_grpcws_grpcws_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Failure.ProtoReflect.Descriptor instead.
func (*Failure) Descriptor() ([]byte, []int) {
	return file_resources_common_grpcws_grpcws_proto_rawDescGZIP(), []int{5}
}

func (x *Failure) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *Failure) GetErrorStatus() string {
	if x != nil {
		return x.ErrorStatus
	}
	return ""
}

func (x *Failure) GetHeaders() map[string]*HeaderValue {
	if x != nil {
		return x.Headers
	}
	return nil
}

type Cancel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Cancel) Reset() {
	*x = Cancel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_common_grpcws_grpcws_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cancel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cancel) ProtoMessage() {}

func (x *Cancel) ProtoReflect() protoreflect.Message {
	mi := &file_resources_common_grpcws_grpcws_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cancel.ProtoReflect.Descriptor instead.
func (*Cancel) Descriptor() ([]byte, []int) {
	return file_resources_common_grpcws_grpcws_proto_rawDescGZIP(), []int{6}
}

var File_resources_common_grpcws_grpcws_proto protoreflect.FileDescriptor

var file_resources_common_grpcws_grpcws_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x77, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x77, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x77, 0x73, 0x22,
	0xdc, 0x02, 0x0a, 0x09, 0x47, 0x72, 0x70, 0x63, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x77, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x77, 0x73, 0x2e, 0x42, 0x6f, 0x64,
	0x79, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x3f, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x77, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x00,
	0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x77, 0x73, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52,
	0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x77, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x06, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xe8,
	0x01, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x77, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x60, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x77, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x23, 0x0a, 0x0b, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x36,
	0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x22, 0xfa, 0x01, 0x0a, 0x07, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x47, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x77, 0x73, 0x2e,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x60, 0x0a,
	0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x77, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x08, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2d,
	0x61, 0x70, 0x70, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x77, 0x73,
	0x3b, 0x67, 0x72, 0x70, 0x63, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_common_grpcws_grpcws_proto_rawDescOnce sync.Once
	file_resources_common_grpcws_grpcws_proto_rawDescData = file_resources_common_grpcws_grpcws_proto_rawDesc
)

func file_resources_common_grpcws_grpcws_proto_rawDescGZIP() []byte {
	file_resources_common_grpcws_grpcws_proto_rawDescOnce.Do(func() {
		file_resources_common_grpcws_grpcws_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_common_grpcws_grpcws_proto_rawDescData)
	})
	return file_resources_common_grpcws_grpcws_proto_rawDescData
}

var file_resources_common_grpcws_grpcws_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_resources_common_grpcws_grpcws_proto_goTypes = []interface{}{
	(*GrpcFrame)(nil),   // 0: resources.common.grpcws.GrpcFrame
	(*Header)(nil),      // 1: resources.common.grpcws.Header
	(*HeaderValue)(nil), // 2: resources.common.grpcws.HeaderValue
	(*Body)(nil),        // 3: resources.common.grpcws.Body
	(*Complete)(nil),    // 4: resources.common.grpcws.Complete
	(*Failure)(nil),     // 5: resources.common.grpcws.Failure
	(*Cancel)(nil),      // 6: resources.common.grpcws.Cancel
	nil,                 // 7: resources.common.grpcws.Header.HeadersEntry
	nil,                 // 8: resources.common.grpcws.Failure.HeadersEntry
}
var file_resources_common_grpcws_grpcws_proto_depIdxs = []int32{
	1, // 0: resources.common.grpcws.GrpcFrame.header:type_name -> resources.common.grpcws.Header
	3, // 1: resources.common.grpcws.GrpcFrame.body:type_name -> resources.common.grpcws.Body
	4, // 2: resources.common.grpcws.GrpcFrame.complete:type_name -> resources.common.grpcws.Complete
	5, // 3: resources.common.grpcws.GrpcFrame.failure:type_name -> resources.common.grpcws.Failure
	6, // 4: resources.common.grpcws.GrpcFrame.cancel:type_name -> resources.common.grpcws.Cancel
	7, // 5: resources.common.grpcws.Header.headers:type_name -> resources.common.grpcws.Header.HeadersEntry
	8, // 6: resources.common.grpcws.Failure.headers:type_name -> resources.common.grpcws.Failure.HeadersEntry
	2, // 7: resources.common.grpcws.Header.HeadersEntry.value:type_name -> resources.common.grpcws.HeaderValue
	2, // 8: resources.common.grpcws.Failure.HeadersEntry.value:type_name -> resources.common.grpcws.HeaderValue
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_resources_common_grpcws_grpcws_proto_init() }
func file_resources_common_grpcws_grpcws_proto_init() {
	if File_resources_common_grpcws_grpcws_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_common_grpcws_grpcws_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcFrame); i {
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
		file_resources_common_grpcws_grpcws_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_resources_common_grpcws_grpcws_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderValue); i {
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
		file_resources_common_grpcws_grpcws_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Body); i {
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
		file_resources_common_grpcws_grpcws_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Complete); i {
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
		file_resources_common_grpcws_grpcws_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Failure); i {
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
		file_resources_common_grpcws_grpcws_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cancel); i {
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
	file_resources_common_grpcws_grpcws_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GrpcFrame_Header)(nil),
		(*GrpcFrame_Body)(nil),
		(*GrpcFrame_Complete)(nil),
		(*GrpcFrame_Failure)(nil),
		(*GrpcFrame_Cancel)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_common_grpcws_grpcws_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_common_grpcws_grpcws_proto_goTypes,
		DependencyIndexes: file_resources_common_grpcws_grpcws_proto_depIdxs,
		MessageInfos:      file_resources_common_grpcws_grpcws_proto_msgTypes,
	}.Build()
	File_resources_common_grpcws_grpcws_proto = out.File
	file_resources_common_grpcws_grpcws_proto_rawDesc = nil
	file_resources_common_grpcws_grpcws_proto_goTypes = nil
	file_resources_common_grpcws_grpcws_proto_depIdxs = nil
}