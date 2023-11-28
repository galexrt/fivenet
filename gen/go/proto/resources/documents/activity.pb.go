// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: resources/documents/activity.proto

package documents

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	timestamp "github.com/galexrt/fivenet/gen/go/proto/resources/timestamp"
	users "github.com/galexrt/fivenet/gen/go/proto/resources/users"
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

type DocActivityType int32

const (
	DocActivityType_DOC_ACTIVITY_TYPE_UNSPECIFIED DocActivityType = 0
	// Base
	DocActivityType_DOC_ACTIVITY_TYPE_CREATED            DocActivityType = 1
	DocActivityType_DOC_ACTIVITY_TYPE_STATUS_OPEN        DocActivityType = 2
	DocActivityType_DOC_ACTIVITY_TYPE_STATUS_CLOSED      DocActivityType = 3
	DocActivityType_DOC_ACTIVITY_TYPE_UPDATED            DocActivityType = 4
	DocActivityType_DOC_ACTIVITY_TYPE_RELATIONS_UPDATED  DocActivityType = 5
	DocActivityType_DOC_ACTIVITY_TYPE_REFERENCES_UPDATED DocActivityType = 6
	DocActivityType_DOC_ACTIVITY_TYPE_ACCESS_UPDATED     DocActivityType = 7
	DocActivityType_DOC_ACTIVITY_TYPE_OWNER_CHANGED      DocActivityType = 8
	DocActivityType_DOC_ACTIVITY_TYPE_DELETED            DocActivityType = 9
	// Requests
	DocActivityType_DOC_ACTIVITY_TYPE_REQUESTED_ACCESS   DocActivityType = 10
	DocActivityType_DOC_ACTIVITY_TYPE_REQUESTED_CLOSURE  DocActivityType = 11
	DocActivityType_DOC_ACTIVITY_TYPE_REQUESTED_UPDATE   DocActivityType = 12
	DocActivityType_DOC_ACTIVITY_TYPE_REQUESTED_DELETION DocActivityType = 13
	// Comments
	DocActivityType_DOC_ACTIVITY_TYPE_COMMENT_ADDED   DocActivityType = 14
	DocActivityType_DOC_ACTIVITY_TYPE_COMMENT_UPDATED DocActivityType = 15
	DocActivityType_DOC_ACTIVITY_TYPE_COMMENT_DELETED DocActivityType = 16
)

// Enum value maps for DocActivityType.
var (
	DocActivityType_name = map[int32]string{
		0:  "DOC_ACTIVITY_TYPE_UNSPECIFIED",
		1:  "DOC_ACTIVITY_TYPE_CREATED",
		2:  "DOC_ACTIVITY_TYPE_STATUS_OPEN",
		3:  "DOC_ACTIVITY_TYPE_STATUS_CLOSED",
		4:  "DOC_ACTIVITY_TYPE_UPDATED",
		5:  "DOC_ACTIVITY_TYPE_RELATIONS_UPDATED",
		6:  "DOC_ACTIVITY_TYPE_REFERENCES_UPDATED",
		7:  "DOC_ACTIVITY_TYPE_ACCESS_UPDATED",
		8:  "DOC_ACTIVITY_TYPE_OWNER_CHANGED",
		9:  "DOC_ACTIVITY_TYPE_DELETED",
		10: "DOC_ACTIVITY_TYPE_REQUESTED_ACCESS",
		11: "DOC_ACTIVITY_TYPE_REQUESTED_CLOSURE",
		12: "DOC_ACTIVITY_TYPE_REQUESTED_UPDATE",
		13: "DOC_ACTIVITY_TYPE_REQUESTED_DELETION",
		14: "DOC_ACTIVITY_TYPE_COMMENT_ADDED",
		15: "DOC_ACTIVITY_TYPE_COMMENT_UPDATED",
		16: "DOC_ACTIVITY_TYPE_COMMENT_DELETED",
	}
	DocActivityType_value = map[string]int32{
		"DOC_ACTIVITY_TYPE_UNSPECIFIED":        0,
		"DOC_ACTIVITY_TYPE_CREATED":            1,
		"DOC_ACTIVITY_TYPE_STATUS_OPEN":        2,
		"DOC_ACTIVITY_TYPE_STATUS_CLOSED":      3,
		"DOC_ACTIVITY_TYPE_UPDATED":            4,
		"DOC_ACTIVITY_TYPE_RELATIONS_UPDATED":  5,
		"DOC_ACTIVITY_TYPE_REFERENCES_UPDATED": 6,
		"DOC_ACTIVITY_TYPE_ACCESS_UPDATED":     7,
		"DOC_ACTIVITY_TYPE_OWNER_CHANGED":      8,
		"DOC_ACTIVITY_TYPE_DELETED":            9,
		"DOC_ACTIVITY_TYPE_REQUESTED_ACCESS":   10,
		"DOC_ACTIVITY_TYPE_REQUESTED_CLOSURE":  11,
		"DOC_ACTIVITY_TYPE_REQUESTED_UPDATE":   12,
		"DOC_ACTIVITY_TYPE_REQUESTED_DELETION": 13,
		"DOC_ACTIVITY_TYPE_COMMENT_ADDED":      14,
		"DOC_ACTIVITY_TYPE_COMMENT_UPDATED":    15,
		"DOC_ACTIVITY_TYPE_COMMENT_DELETED":    16,
	}
)

func (x DocActivityType) Enum() *DocActivityType {
	p := new(DocActivityType)
	*p = x
	return p
}

func (x DocActivityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DocActivityType) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_documents_activity_proto_enumTypes[0].Descriptor()
}

func (DocActivityType) Type() protoreflect.EnumType {
	return &file_resources_documents_activity_proto_enumTypes[0]
}

func (x DocActivityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DocActivityType.Descriptor instead.
func (DocActivityType) EnumDescriptor() ([]byte, []int) {
	return file_resources_documents_activity_proto_rawDescGZIP(), []int{0}
}

type DocActivity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt       *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DocumentId      uint64               `protobuf:"varint,3,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	ActivityType    DocActivityType      `protobuf:"varint,4,opt,name=activity_type,json=activityType,proto3,enum=resources.documents.DocActivityType" json:"activity_type,omitempty"`
	CreatorId       *int32               `protobuf:"varint,5,opt,name=creator_id,json=creatorId,proto3,oneof" json:"creator_id,omitempty"`
	Creator         *users.UserShort     `protobuf:"bytes,6,opt,name=creator,proto3,oneof" json:"creator,omitempty" alias:"creator"` // @gotags: alias:"creator"
	CreatorJob      string               `protobuf:"bytes,7,opt,name=creator_job,json=creatorJob,proto3" json:"creator_job,omitempty"`
	CreatorJobLabel *string              `protobuf:"bytes,8,opt,name=creator_job_label,json=creatorJobLabel,proto3,oneof" json:"creator_job_label,omitempty"`
	Reason          *string              `protobuf:"bytes,9,opt,name=reason,proto3,oneof" json:"reason,omitempty"`
	Data            *DocActivityData     `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DocActivity) Reset() {
	*x = DocActivity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_documents_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocActivity) ProtoMessage() {}

func (x *DocActivity) ProtoReflect() protoreflect.Message {
	mi := &file_resources_documents_activity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocActivity.ProtoReflect.Descriptor instead.
func (*DocActivity) Descriptor() ([]byte, []int) {
	return file_resources_documents_activity_proto_rawDescGZIP(), []int{0}
}

func (x *DocActivity) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DocActivity) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DocActivity) GetDocumentId() uint64 {
	if x != nil {
		return x.DocumentId
	}
	return 0
}

func (x *DocActivity) GetActivityType() DocActivityType {
	if x != nil {
		return x.ActivityType
	}
	return DocActivityType_DOC_ACTIVITY_TYPE_UNSPECIFIED
}

func (x *DocActivity) GetCreatorId() int32 {
	if x != nil && x.CreatorId != nil {
		return *x.CreatorId
	}
	return 0
}

func (x *DocActivity) GetCreator() *users.UserShort {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *DocActivity) GetCreatorJob() string {
	if x != nil {
		return x.CreatorJob
	}
	return ""
}

func (x *DocActivity) GetCreatorJobLabel() string {
	if x != nil && x.CreatorJobLabel != nil {
		return *x.CreatorJobLabel
	}
	return ""
}

func (x *DocActivity) GetReason() string {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return ""
}

func (x *DocActivity) GetData() *DocActivityData {
	if x != nil {
		return x.Data
	}
	return nil
}

type DocActivityData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*DocActivityData_Updated
	//	*DocActivityData_OwnerChanged
	Data isDocActivityData_Data `protobuf_oneof:"data"`
}

func (x *DocActivityData) Reset() {
	*x = DocActivityData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_documents_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocActivityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocActivityData) ProtoMessage() {}

func (x *DocActivityData) ProtoReflect() protoreflect.Message {
	mi := &file_resources_documents_activity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocActivityData.ProtoReflect.Descriptor instead.
func (*DocActivityData) Descriptor() ([]byte, []int) {
	return file_resources_documents_activity_proto_rawDescGZIP(), []int{1}
}

func (m *DocActivityData) GetData() isDocActivityData_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *DocActivityData) GetUpdated() *DocUpdated {
	if x, ok := x.GetData().(*DocActivityData_Updated); ok {
		return x.Updated
	}
	return nil
}

func (x *DocActivityData) GetOwnerChanged() *DocOwnerChanged {
	if x, ok := x.GetData().(*DocActivityData_OwnerChanged); ok {
		return x.OwnerChanged
	}
	return nil
}

type isDocActivityData_Data interface {
	isDocActivityData_Data()
}

type DocActivityData_Updated struct {
	Updated *DocUpdated `protobuf:"bytes,1,opt,name=updated,proto3,oneof"`
}

type DocActivityData_OwnerChanged struct {
	OwnerChanged *DocOwnerChanged `protobuf:"bytes,2,opt,name=owner_changed,json=ownerChanged,proto3,oneof"`
}

func (*DocActivityData_Updated) isDocActivityData_Data() {}

func (*DocActivityData_OwnerChanged) isDocActivityData_Data() {}

type DocUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Diff string `protobuf:"bytes,1,opt,name=diff,proto3" json:"diff,omitempty"`
}

func (x *DocUpdated) Reset() {
	*x = DocUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_documents_activity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocUpdated) ProtoMessage() {}

func (x *DocUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_resources_documents_activity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocUpdated.ProtoReflect.Descriptor instead.
func (*DocUpdated) Descriptor() ([]byte, []int) {
	return file_resources_documents_activity_proto_rawDescGZIP(), []int{2}
}

func (x *DocUpdated) GetDiff() string {
	if x != nil {
		return x.Diff
	}
	return ""
}

type DocOwnerChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId int32 `protobuf:"varint,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *DocOwnerChanged) Reset() {
	*x = DocOwnerChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_documents_activity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocOwnerChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocOwnerChanged) ProtoMessage() {}

func (x *DocOwnerChanged) ProtoReflect() protoreflect.Message {
	mi := &file_resources_documents_activity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocOwnerChanged.ProtoReflect.Descriptor instead.
func (*DocOwnerChanged) Descriptor() ([]byte, []int) {
	return file_resources_documents_activity_proto_rawDescGZIP(), []int{3}
}

func (x *DocOwnerChanged) GetOwnerId() int32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

var File_resources_documents_activity_proto protoreflect.FileDescriptor

var file_resources_documents_activity_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x04, 0x0a, 0x0b, 0x44, 0x6f, 0x63, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01,
	0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x0d,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x48, 0x01, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x18, 0x14, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x4a, 0x6f, 0x62,
	0x12, 0x38, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x5f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x18, 0x32, 0x48, 0x02, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x4a,
	0x6f, 0x62, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0x18, 0xff, 0x01, 0x48, 0x03, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0xa3, 0x01, 0x0a, 0x0f, 0x44, 0x6f, 0x63, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x4b, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x20, 0x0a,
	0x0a, 0x44, 0x6f, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x69, 0x66, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x69, 0x66, 0x66, 0x22,
	0x2c, 0x0a, 0x0f, 0x44, 0x6f, 0x63, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x8d, 0x05,
	0x0a, 0x0f, 0x44, 0x6f, 0x63, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x4f, 0x43, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54,
	0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x4f, 0x43, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x4f, 0x43, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x4f, 0x50, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x44, 0x4f, 0x43, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x44,
	0x4f, 0x43, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x27, 0x0a, 0x23, 0x44, 0x4f,
	0x43, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x05, 0x12, 0x28, 0x0a, 0x24, 0x44, 0x4f, 0x43, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e,
	0x43, 0x45, 0x53, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x24, 0x0a,
	0x20, 0x44, 0x4f, 0x43, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x07, 0x12, 0x23, 0x0a, 0x1f, 0x44, 0x4f, 0x43, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x08, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x4f, 0x43, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x09, 0x12, 0x26, 0x0a, 0x22, 0x44, 0x4f, 0x43, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x0a, 0x12,
	0x27, 0x0a, 0x23, 0x44, 0x4f, 0x43, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x43,
	0x4c, 0x4f, 0x53, 0x55, 0x52, 0x45, 0x10, 0x0b, 0x12, 0x26, 0x0a, 0x22, 0x44, 0x4f, 0x43, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x0c,
	0x12, 0x28, 0x0a, 0x24, 0x44, 0x4f, 0x43, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0d, 0x12, 0x23, 0x0a, 0x1f, 0x44, 0x4f,
	0x43, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x45, 0x44, 0x10, 0x0e, 0x12,
	0x25, 0x0a, 0x21, 0x44, 0x4f, 0x43, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x0f, 0x12, 0x25, 0x0a, 0x21, 0x44, 0x4f, 0x43, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x10, 0x42, 0x47, 0x5a,
	0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6c, 0x65,
	0x78, 0x72, 0x74, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3b, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_documents_activity_proto_rawDescOnce sync.Once
	file_resources_documents_activity_proto_rawDescData = file_resources_documents_activity_proto_rawDesc
)

func file_resources_documents_activity_proto_rawDescGZIP() []byte {
	file_resources_documents_activity_proto_rawDescOnce.Do(func() {
		file_resources_documents_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_documents_activity_proto_rawDescData)
	})
	return file_resources_documents_activity_proto_rawDescData
}

var file_resources_documents_activity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resources_documents_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_resources_documents_activity_proto_goTypes = []interface{}{
	(DocActivityType)(0),        // 0: resources.documents.DocActivityType
	(*DocActivity)(nil),         // 1: resources.documents.DocActivity
	(*DocActivityData)(nil),     // 2: resources.documents.DocActivityData
	(*DocUpdated)(nil),          // 3: resources.documents.DocUpdated
	(*DocOwnerChanged)(nil),     // 4: resources.documents.DocOwnerChanged
	(*timestamp.Timestamp)(nil), // 5: resources.timestamp.Timestamp
	(*users.UserShort)(nil),     // 6: resources.users.UserShort
}
var file_resources_documents_activity_proto_depIdxs = []int32{
	5, // 0: resources.documents.DocActivity.created_at:type_name -> resources.timestamp.Timestamp
	0, // 1: resources.documents.DocActivity.activity_type:type_name -> resources.documents.DocActivityType
	6, // 2: resources.documents.DocActivity.creator:type_name -> resources.users.UserShort
	2, // 3: resources.documents.DocActivity.data:type_name -> resources.documents.DocActivityData
	3, // 4: resources.documents.DocActivityData.updated:type_name -> resources.documents.DocUpdated
	4, // 5: resources.documents.DocActivityData.owner_changed:type_name -> resources.documents.DocOwnerChanged
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_resources_documents_activity_proto_init() }
func file_resources_documents_activity_proto_init() {
	if File_resources_documents_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_documents_activity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocActivity); i {
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
		file_resources_documents_activity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocActivityData); i {
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
		file_resources_documents_activity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocUpdated); i {
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
		file_resources_documents_activity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocOwnerChanged); i {
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
	file_resources_documents_activity_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_resources_documents_activity_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*DocActivityData_Updated)(nil),
		(*DocActivityData_OwnerChanged)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_documents_activity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_documents_activity_proto_goTypes,
		DependencyIndexes: file_resources_documents_activity_proto_depIdxs,
		EnumInfos:         file_resources_documents_activity_proto_enumTypes,
		MessageInfos:      file_resources_documents_activity_proto_msgTypes,
	}.Build()
	File_resources_documents_activity_proto = out.File
	file_resources_documents_activity_proto_rawDesc = nil
	file_resources_documents_activity_proto_goTypes = nil
	file_resources_documents_activity_proto_depIdxs = nil
}
