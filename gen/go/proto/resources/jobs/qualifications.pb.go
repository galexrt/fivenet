// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: resources/jobs/qualifications.proto

package jobs

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

type QualificationModule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" sql:"primary_key" alias:"id"` // @gotags: sql:"primary_key" alias:"id"
	CreatedAt           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	Job                 string               `protobuf:"bytes,4,opt,name=job,proto3" json:"job,omitempty"`
	Open                bool                 `protobuf:"varint,5,opt,name=open,proto3" json:"open,omitempty"`
	RequireRequirements bool                 `protobuf:"varint,6,opt,name=require_requirements,json=requireRequirements,proto3" json:"require_requirements,omitempty"`
	MinimumGrade        int32                `protobuf:"varint,7,opt,name=minimum_grade,json=minimumGrade,proto3" json:"minimum_grade,omitempty"`
	// @sanitize
	Title string `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
	// @sanitize
	Description string `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *QualificationModule) Reset() {
	*x = QualificationModule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_jobs_qualifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QualificationModule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QualificationModule) ProtoMessage() {}

func (x *QualificationModule) ProtoReflect() protoreflect.Message {
	mi := &file_resources_jobs_qualifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QualificationModule.ProtoReflect.Descriptor instead.
func (*QualificationModule) Descriptor() ([]byte, []int) {
	return file_resources_jobs_qualifications_proto_rawDescGZIP(), []int{0}
}

func (x *QualificationModule) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QualificationModule) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *QualificationModule) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *QualificationModule) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *QualificationModule) GetOpen() bool {
	if x != nil {
		return x.Open
	}
	return false
}

func (x *QualificationModule) GetRequireRequirements() bool {
	if x != nil {
		return x.RequireRequirements
	}
	return false
}

func (x *QualificationModule) GetMinimumGrade() int32 {
	if x != nil {
		return x.MinimumGrade
	}
	return 0
}

func (x *QualificationModule) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *QualificationModule) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type QualificationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	TrainingId  uint64               `protobuf:"varint,3,opt,name=training_id,json=trainingId,proto3" json:"training_id,omitempty"`
	UserId      int32                `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	User        *users.UserShort     `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Rating      uint32               `protobuf:"varint,6,opt,name=rating,proto3" json:"rating,omitempty"`
	Description string               `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	CreatorId   int32                `protobuf:"varint,8,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Creator     *users.UserShort     `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *QualificationResult) Reset() {
	*x = QualificationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_jobs_qualifications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QualificationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QualificationResult) ProtoMessage() {}

func (x *QualificationResult) ProtoReflect() protoreflect.Message {
	mi := &file_resources_jobs_qualifications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QualificationResult.ProtoReflect.Descriptor instead.
func (*QualificationResult) Descriptor() ([]byte, []int) {
	return file_resources_jobs_qualifications_proto_rawDescGZIP(), []int{1}
}

func (x *QualificationResult) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QualificationResult) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *QualificationResult) GetTrainingId() uint64 {
	if x != nil {
		return x.TrainingId
	}
	return 0
}

func (x *QualificationResult) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *QualificationResult) GetUser() *users.UserShort {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *QualificationResult) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *QualificationResult) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *QualificationResult) GetCreatorId() int32 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *QualificationResult) GetCreator() *users.UserShort {
	if x != nil {
		return x.Creator
	}
	return nil
}

var File_resources_jobs_qualifications_proto protoreflect.FileDescriptor

var file_resources_jobs_qualifications_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x73,
	0x2f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x1a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x93, 0x03, 0x0a, 0x13, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x19, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x18, 0x14, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6f,
	0x70, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12,
	0x31, 0x0a, 0x14, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x2c, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02,
	0x20, 0x00, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x47, 0x72, 0x61, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0xfd, 0x02, 0x0a, 0x13, 0x51, 0x75, 0x61, 0x6c, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05,
	0x10, 0x03, 0x18, 0x80, 0x04, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6c, 0x65, 0x78, 0x72, 0x74, 0x2f, 0x66, 0x69, 0x76,
	0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x73,
	0x3b, 0x6a, 0x6f, 0x62, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_jobs_qualifications_proto_rawDescOnce sync.Once
	file_resources_jobs_qualifications_proto_rawDescData = file_resources_jobs_qualifications_proto_rawDesc
)

func file_resources_jobs_qualifications_proto_rawDescGZIP() []byte {
	file_resources_jobs_qualifications_proto_rawDescOnce.Do(func() {
		file_resources_jobs_qualifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_jobs_qualifications_proto_rawDescData)
	})
	return file_resources_jobs_qualifications_proto_rawDescData
}

var file_resources_jobs_qualifications_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_resources_jobs_qualifications_proto_goTypes = []interface{}{
	(*QualificationModule)(nil), // 0: resources.jobs.QualificationModule
	(*QualificationResult)(nil), // 1: resources.jobs.QualificationResult
	(*timestamp.Timestamp)(nil), // 2: resources.timestamp.Timestamp
	(*users.UserShort)(nil),     // 3: resources.users.UserShort
}
var file_resources_jobs_qualifications_proto_depIdxs = []int32{
	2, // 0: resources.jobs.QualificationModule.created_at:type_name -> resources.timestamp.Timestamp
	2, // 1: resources.jobs.QualificationModule.updated_at:type_name -> resources.timestamp.Timestamp
	2, // 2: resources.jobs.QualificationResult.created_at:type_name -> resources.timestamp.Timestamp
	3, // 3: resources.jobs.QualificationResult.user:type_name -> resources.users.UserShort
	3, // 4: resources.jobs.QualificationResult.creator:type_name -> resources.users.UserShort
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_resources_jobs_qualifications_proto_init() }
func file_resources_jobs_qualifications_proto_init() {
	if File_resources_jobs_qualifications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_jobs_qualifications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QualificationModule); i {
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
		file_resources_jobs_qualifications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QualificationResult); i {
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
	file_resources_jobs_qualifications_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_resources_jobs_qualifications_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_jobs_qualifications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_jobs_qualifications_proto_goTypes,
		DependencyIndexes: file_resources_jobs_qualifications_proto_depIdxs,
		MessageInfos:      file_resources_jobs_qualifications_proto_msgTypes,
	}.Build()
	File_resources_jobs_qualifications_proto = out.File
	file_resources_jobs_qualifications_proto_rawDesc = nil
	file_resources_jobs_qualifications_proto_goTypes = nil
	file_resources_jobs_qualifications_proto_depIdxs = nil
}
