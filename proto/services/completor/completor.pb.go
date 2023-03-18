// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: services/completor/completor.proto

package completor

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	documents "github.com/galexrt/arpanet/proto/resources/documents"
	jobs "github.com/galexrt/arpanet/proto/resources/jobs"
	users "github.com/galexrt/arpanet/proto/resources/users"
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

type CompleteCharNamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *CompleteCharNamesRequest) Reset() {
	*x = CompleteCharNamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_completor_completor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteCharNamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteCharNamesRequest) ProtoMessage() {}

func (x *CompleteCharNamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_completor_completor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteCharNamesRequest.ProtoReflect.Descriptor instead.
func (*CompleteCharNamesRequest) Descriptor() ([]byte, []int) {
	return file_services_completor_completor_proto_rawDescGZIP(), []int{0}
}

func (x *CompleteCharNamesRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type CompleteCharNamesRespoonse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*users.UserShort `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *CompleteCharNamesRespoonse) Reset() {
	*x = CompleteCharNamesRespoonse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_completor_completor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteCharNamesRespoonse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteCharNamesRespoonse) ProtoMessage() {}

func (x *CompleteCharNamesRespoonse) ProtoReflect() protoreflect.Message {
	mi := &file_services_completor_completor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteCharNamesRespoonse.ProtoReflect.Descriptor instead.
func (*CompleteCharNamesRespoonse) Descriptor() ([]byte, []int) {
	return file_services_completor_completor_proto_rawDescGZIP(), []int{1}
}

func (x *CompleteCharNamesRespoonse) GetUsers() []*users.UserShort {
	if x != nil {
		return x.Users
	}
	return nil
}

type CompleteJobNamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *CompleteJobNamesRequest) Reset() {
	*x = CompleteJobNamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_completor_completor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteJobNamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteJobNamesRequest) ProtoMessage() {}

func (x *CompleteJobNamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_completor_completor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteJobNamesRequest.ProtoReflect.Descriptor instead.
func (*CompleteJobNamesRequest) Descriptor() ([]byte, []int) {
	return file_services_completor_completor_proto_rawDescGZIP(), []int{2}
}

func (x *CompleteJobNamesRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type CompleteJobNamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*jobs.Job `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *CompleteJobNamesResponse) Reset() {
	*x = CompleteJobNamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_completor_completor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteJobNamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteJobNamesResponse) ProtoMessage() {}

func (x *CompleteJobNamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_completor_completor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteJobNamesResponse.ProtoReflect.Descriptor instead.
func (*CompleteJobNamesResponse) Descriptor() ([]byte, []int) {
	return file_services_completor_completor_proto_rawDescGZIP(), []int{3}
}

func (x *CompleteJobNamesResponse) GetJobs() []*jobs.Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type CompleteDocumentCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *CompleteDocumentCategoryRequest) Reset() {
	*x = CompleteDocumentCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_completor_completor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteDocumentCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteDocumentCategoryRequest) ProtoMessage() {}

func (x *CompleteDocumentCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_completor_completor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteDocumentCategoryRequest.ProtoReflect.Descriptor instead.
func (*CompleteDocumentCategoryRequest) Descriptor() ([]byte, []int) {
	return file_services_completor_completor_proto_rawDescGZIP(), []int{4}
}

func (x *CompleteDocumentCategoryRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type CompleteDocumentCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*documents.DocumentCategory `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *CompleteDocumentCategoryResponse) Reset() {
	*x = CompleteDocumentCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_completor_completor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteDocumentCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteDocumentCategoryResponse) ProtoMessage() {}

func (x *CompleteDocumentCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_completor_completor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteDocumentCategoryResponse.ProtoReflect.Descriptor instead.
func (*CompleteDocumentCategoryResponse) Descriptor() ([]byte, []int) {
	return file_services_completor_completor_proto_rawDescGZIP(), []int{5}
}

func (x *CompleteDocumentCategoryResponse) GetCategories() []*documents.DocumentCategory {
	if x != nil {
		return x.Categories
	}
	return nil
}

var File_services_completor_completor_proto protoreflect.FileDescriptor

var file_services_completor_completor_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x6f, 0x72, 0x1a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2f, 0x6a, 0x6f,
	0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b,
	0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x18, 0x32, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x4e, 0x0a, 0x1a, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x0a, 0x17, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x43, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f,
	0x62, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0x43, 0x0a, 0x1f,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x01, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x22, 0x69, 0x0a, 0x20, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x32, 0xfc, 0x02, 0x0a,
	0x10, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x71, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6c, 0x65, 0x78, 0x72,
	0x74, 0x2f, 0x61, 0x72, 0x70, 0x61, 0x6e, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x6f, 0x72, 0x3b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_completor_completor_proto_rawDescOnce sync.Once
	file_services_completor_completor_proto_rawDescData = file_services_completor_completor_proto_rawDesc
)

func file_services_completor_completor_proto_rawDescGZIP() []byte {
	file_services_completor_completor_proto_rawDescOnce.Do(func() {
		file_services_completor_completor_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_completor_completor_proto_rawDescData)
	})
	return file_services_completor_completor_proto_rawDescData
}

var file_services_completor_completor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_completor_completor_proto_goTypes = []interface{}{
	(*CompleteCharNamesRequest)(nil),         // 0: services.completor.CompleteCharNamesRequest
	(*CompleteCharNamesRespoonse)(nil),       // 1: services.completor.CompleteCharNamesRespoonse
	(*CompleteJobNamesRequest)(nil),          // 2: services.completor.CompleteJobNamesRequest
	(*CompleteJobNamesResponse)(nil),         // 3: services.completor.CompleteJobNamesResponse
	(*CompleteDocumentCategoryRequest)(nil),  // 4: services.completor.CompleteDocumentCategoryRequest
	(*CompleteDocumentCategoryResponse)(nil), // 5: services.completor.CompleteDocumentCategoryResponse
	(*users.UserShort)(nil),                  // 6: resources.users.UserShort
	(*jobs.Job)(nil),                         // 7: resources.jobs.Job
	(*documents.DocumentCategory)(nil),       // 8: resources.documents.DocumentCategory
}
var file_services_completor_completor_proto_depIdxs = []int32{
	6, // 0: services.completor.CompleteCharNamesRespoonse.users:type_name -> resources.users.UserShort
	7, // 1: services.completor.CompleteJobNamesResponse.jobs:type_name -> resources.jobs.Job
	8, // 2: services.completor.CompleteDocumentCategoryResponse.categories:type_name -> resources.documents.DocumentCategory
	0, // 3: services.completor.CompletorService.CompleteCharNames:input_type -> services.completor.CompleteCharNamesRequest
	2, // 4: services.completor.CompletorService.CompleteJobNames:input_type -> services.completor.CompleteJobNamesRequest
	4, // 5: services.completor.CompletorService.CompleteDocumentCategory:input_type -> services.completor.CompleteDocumentCategoryRequest
	1, // 6: services.completor.CompletorService.CompleteCharNames:output_type -> services.completor.CompleteCharNamesRespoonse
	3, // 7: services.completor.CompletorService.CompleteJobNames:output_type -> services.completor.CompleteJobNamesResponse
	5, // 8: services.completor.CompletorService.CompleteDocumentCategory:output_type -> services.completor.CompleteDocumentCategoryResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_services_completor_completor_proto_init() }
func file_services_completor_completor_proto_init() {
	if File_services_completor_completor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_completor_completor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteCharNamesRequest); i {
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
		file_services_completor_completor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteCharNamesRespoonse); i {
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
		file_services_completor_completor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteJobNamesRequest); i {
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
		file_services_completor_completor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteJobNamesResponse); i {
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
		file_services_completor_completor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteDocumentCategoryRequest); i {
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
		file_services_completor_completor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteDocumentCategoryResponse); i {
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
			RawDescriptor: file_services_completor_completor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_completor_completor_proto_goTypes,
		DependencyIndexes: file_services_completor_completor_proto_depIdxs,
		MessageInfos:      file_services_completor_completor_proto_msgTypes,
	}.Build()
	File_services_completor_completor_proto = out.File
	file_services_completor_completor_proto_rawDesc = nil
	file_services_completor_completor_proto_goTypes = nil
	file_services_completor_completor_proto_depIdxs = nil
}
