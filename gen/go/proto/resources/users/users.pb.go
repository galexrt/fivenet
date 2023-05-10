// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: resources/users/users.proto

package users

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	jobs "github.com/galexrt/fivenet/gen/go/proto/resources/jobs"
	timestamp "github.com/galexrt/fivenet/gen/go/proto/resources/timestamp"
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

type USER_ACTIVITY_TYPE int32

const (
	USER_ACTIVITY_TYPE_CHANGED   USER_ACTIVITY_TYPE = 0
	USER_ACTIVITY_TYPE_MENTIONED USER_ACTIVITY_TYPE = 1
	USER_ACTIVITY_TYPE_CREATED   USER_ACTIVITY_TYPE = 2
)

// Enum value maps for USER_ACTIVITY_TYPE.
var (
	USER_ACTIVITY_TYPE_name = map[int32]string{
		0: "CHANGED",
		1: "MENTIONED",
		2: "CREATED",
	}
	USER_ACTIVITY_TYPE_value = map[string]int32{
		"CHANGED":   0,
		"MENTIONED": 1,
		"CREATED":   2,
	}
)

func (x USER_ACTIVITY_TYPE) Enum() *USER_ACTIVITY_TYPE {
	p := new(USER_ACTIVITY_TYPE)
	*p = x
	return p
}

func (x USER_ACTIVITY_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (USER_ACTIVITY_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_users_users_proto_enumTypes[0].Descriptor()
}

func (USER_ACTIVITY_TYPE) Type() protoreflect.EnumType {
	return &file_resources_users_users_proto_enumTypes[0]
}

func (x USER_ACTIVITY_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use USER_ACTIVITY_TYPE.Descriptor instead.
func (USER_ACTIVITY_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_resources_users_users_proto_rawDescGZIP(), []int{0}
}

type UserShort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" alias:"id"`                       // @gotags: alias:"id"
	Identifier    string `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty" alias:"identifier"`                              // @gotags: alias:"identifier"
	Job           string `protobuf:"bytes,3,opt,name=job,proto3" json:"job,omitempty" alias:"job"`                                            // @gotags: alias:"job"
	JobLabel      string `protobuf:"bytes,4,opt,name=job_label,json=jobLabel,proto3" json:"job_label,omitempty" alias:"job_label"`                  // @gotags: alias:"job_label"
	JobGrade      int32  `protobuf:"varint,5,opt,name=job_grade,json=jobGrade,proto3" json:"job_grade,omitempty" alias:"job_grade"`                 // @gotags: alias:"job_grade"
	JobGradeLabel string `protobuf:"bytes,6,opt,name=job_grade_label,json=jobGradeLabel,proto3" json:"job_grade_label,omitempty" alias:"job_grade_label"` // @gotags: alias:"job_grade_label"
	Firstname     string `protobuf:"bytes,7,opt,name=firstname,proto3" json:"firstname,omitempty" alias:"firstname"`                                // @gotags: alias:"firstname"
	Lastname      string `protobuf:"bytes,8,opt,name=lastname,proto3" json:"lastname,omitempty" alias:"lastname"`                                  // @gotags: alias:"lastname"
}

func (x *UserShort) Reset() {
	*x = UserShort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_users_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserShort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserShort) ProtoMessage() {}

func (x *UserShort) ProtoReflect() protoreflect.Message {
	mi := &file_resources_users_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserShort.ProtoReflect.Descriptor instead.
func (*UserShort) Descriptor() ([]byte, []int) {
	return file_resources_users_users_proto_rawDescGZIP(), []int{0}
}

func (x *UserShort) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserShort) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *UserShort) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *UserShort) GetJobLabel() string {
	if x != nil {
		return x.JobLabel
	}
	return ""
}

func (x *UserShort) GetJobGrade() int32 {
	if x != nil {
		return x.JobGrade
	}
	return 0
}

func (x *UserShort) GetJobGradeLabel() string {
	if x != nil {
		return x.JobGradeLabel
	}
	return ""
}

func (x *UserShort) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *UserShort) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int32      `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" alias:"id"`                       // @gotags: alias:"id"
	Identifier    string     `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty" alias:"identifier"`                              // @gotags: alias:"identifier"
	Job           string     `protobuf:"bytes,3,opt,name=job,proto3" json:"job,omitempty" alias:"job"`                                            // @gotags: alias:"job"
	JobLabel      string     `protobuf:"bytes,4,opt,name=job_label,json=jobLabel,proto3" json:"job_label,omitempty" alias:"job_label"`                  // @gotags: alias:"job_label"
	JobGrade      int32      `protobuf:"varint,5,opt,name=job_grade,json=jobGrade,proto3" json:"job_grade,omitempty" alias:"job_grade"`                 // @gotags: alias:"job_grade"
	JobGradeLabel string     `protobuf:"bytes,6,opt,name=job_grade_label,json=jobGradeLabel,proto3" json:"job_grade_label,omitempty" alias:"job_grade_label"` // @gotags: alias:"job_grade_label"
	Firstname     string     `protobuf:"bytes,7,opt,name=firstname,proto3" json:"firstname,omitempty" alias:"firstname"`                                // @gotags: alias:"firstname"
	Lastname      string     `protobuf:"bytes,8,opt,name=lastname,proto3" json:"lastname,omitempty" alias:"lastname"`                                  // @gotags: alias:"lastname"
	Dateofbirth   string     `protobuf:"bytes,9,opt,name=dateofbirth,proto3" json:"dateofbirth,omitempty" alias:"dateofbirth"`                            // @gotags: alias:"dateofbirth"
	Sex           string     `protobuf:"bytes,10,opt,name=sex,proto3" json:"sex,omitempty" alias:"sex"`                                           // @gotags: alias:"sex"
	Height        string     `protobuf:"bytes,11,opt,name=height,proto3" json:"height,omitempty" alias:"height"`                                     // @gotags: alias:"height"
	PhoneNumber   string     `protobuf:"bytes,12,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty" alias:"phone_number"`        // @gotags: alias:"phone_number"
	Visum         *int32     `protobuf:"varint,13,opt,name=visum,proto3,oneof" json:"visum,omitempty" alias:"visum"`                                // @gotags: alias:"visum"
	Playtime      *int32     `protobuf:"varint,14,opt,name=playtime,proto3,oneof" json:"playtime,omitempty" alias:"playtime"`                          // @gotags: alias:"playtime"
	Props         *UserProps `protobuf:"bytes,15,opt,name=props,proto3" json:"props,omitempty" alias:"fivenet_user_props"`                                       // @gotags: alias:"fivenet_user_props"
	Licenses      []*License `protobuf:"bytes,16,rep,name=licenses,proto3" json:"licenses,omitempty" alias:"user_licenses"`                                 // @gotags: alias:"user_licenses"
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_users_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_resources_users_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_resources_users_users_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *User) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *User) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *User) GetJobLabel() string {
	if x != nil {
		return x.JobLabel
	}
	return ""
}

func (x *User) GetJobGrade() int32 {
	if x != nil {
		return x.JobGrade
	}
	return 0
}

func (x *User) GetJobGradeLabel() string {
	if x != nil {
		return x.JobGradeLabel
	}
	return ""
}

func (x *User) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *User) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *User) GetDateofbirth() string {
	if x != nil {
		return x.Dateofbirth
	}
	return ""
}

func (x *User) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *User) GetHeight() string {
	if x != nil {
		return x.Height
	}
	return ""
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetVisum() int32 {
	if x != nil && x.Visum != nil {
		return *x.Visum
	}
	return 0
}

func (x *User) GetPlaytime() int32 {
	if x != nil && x.Playtime != nil {
		return *x.Playtime
	}
	return 0
}

func (x *User) GetProps() *UserProps {
	if x != nil {
		return x.Props
	}
	return nil
}

func (x *User) GetLicenses() []*License {
	if x != nil {
		return x.Licenses
	}
	return nil
}

type License struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" alias:"type"`   // @gotags: alias:"type"
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty" alias:"label"` // @gotags: alias:"label"
}

func (x *License) Reset() {
	*x = License{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_users_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *License) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License) ProtoMessage() {}

func (x *License) ProtoReflect() protoreflect.Message {
	mi := &file_resources_users_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License.ProtoReflect.Descriptor instead.
func (*License) Descriptor() ([]byte, []int) {
	return file_resources_users_users_proto_rawDescGZIP(), []int{2}
}

func (x *License) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *License) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type UserProps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32     `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" alias:"user_id"`         // @gotags: alias:"user_id"
	Wanted  *bool     `protobuf:"varint,2,opt,name=wanted,proto3,oneof" json:"wanted,omitempty" alias:"wanted"`                 // @gotags: alias:"wanted"
	JobName *string   `protobuf:"bytes,3,opt,name=job_name,json=jobName,proto3,oneof" json:"job_name,omitempty" alias:"job"` // @gotags: alias:"job"
	Job     *jobs.Job `protobuf:"bytes,4,opt,name=job,proto3,oneof" json:"job,omitempty"`
}

func (x *UserProps) Reset() {
	*x = UserProps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_users_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProps) ProtoMessage() {}

func (x *UserProps) ProtoReflect() protoreflect.Message {
	mi := &file_resources_users_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProps.ProtoReflect.Descriptor instead.
func (*UserProps) Descriptor() ([]byte, []int) {
	return file_resources_users_users_proto_rawDescGZIP(), []int{3}
}

func (x *UserProps) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserProps) GetWanted() bool {
	if x != nil && x.Wanted != nil {
		return *x.Wanted
	}
	return false
}

func (x *UserProps) GetJobName() string {
	if x != nil && x.JobName != nil {
		return *x.JobName
	}
	return ""
}

func (x *UserProps) GetJob() *jobs.Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type UserActivity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" alias:"id"`                                             // @gotags: alias:"id"
	Type       USER_ACTIVITY_TYPE   `protobuf:"varint,2,opt,name=type,proto3,enum=resources.users.USER_ACTIVITY_TYPE" json:"type,omitempty" alias:"type"` // @gotags: alias:"type"
	CreatedAt  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" alias:"created_at"`               // @gotags: alias:"created_at"
	SourceUser *UserShort           `protobuf:"bytes,4,opt,name=source_user,json=sourceUser,proto3" json:"source_user,omitempty" alias:"source_user"`            // @gotags: alias:"source_user"
	TargetUser *UserShort           `protobuf:"bytes,5,opt,name=target_user,json=targetUser,proto3" json:"target_user,omitempty" alias:"target_user"`            // @gotags: alias:"target_user"
	Key        string               `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty" alias:"key"`                                            // @gotags: alias:"key"
	OldValue   string               `protobuf:"bytes,7,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty" alias:"old_value"`                  // @gotags: alias:"old_value"
	NewValue   string               `protobuf:"bytes,8,opt,name=new_value,json=newValue,proto3" json:"new_value,omitempty" alias:"new_value"`                  // @gotags: alias:"new_value"
	Reason     string               `protobuf:"bytes,9,opt,name=reason,proto3" json:"reason,omitempty" alias:"reason"`                                      // @gotags: alias:"reason"
}

func (x *UserActivity) Reset() {
	*x = UserActivity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_users_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserActivity) ProtoMessage() {}

func (x *UserActivity) ProtoReflect() protoreflect.Message {
	mi := &file_resources_users_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserActivity.ProtoReflect.Descriptor instead.
func (*UserActivity) Descriptor() ([]byte, []int) {
	return file_resources_users_users_proto_rawDescGZIP(), []int{4}
}

func (x *UserActivity) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserActivity) GetType() USER_ACTIVITY_TYPE {
	if x != nil {
		return x.Type
	}
	return USER_ACTIVITY_TYPE_CHANGED
}

func (x *UserActivity) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserActivity) GetSourceUser() *UserShort {
	if x != nil {
		return x.SourceUser
	}
	return nil
}

func (x *UserActivity) GetTargetUser() *UserShort {
	if x != nil {
		return x.TargetUser
	}
	return nil
}

func (x *UserActivity) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UserActivity) GetOldValue() string {
	if x != nil {
		return x.OldValue
	}
	return ""
}

func (x *UserActivity) GetNewValue() string {
	if x != nil {
		return x.NewValue
	}
	return ""
}

func (x *UserActivity) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_resources_users_users_proto protoreflect.FileDescriptor

var file_resources_users_users_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x19,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2f, 0x6a,
	0x6f, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x02, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x72, 0x03, 0x98, 0x01, 0x2e, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x19, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x24, 0x0a, 0x09,
	0x6a, 0x6f, 0x62, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x2d, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0x1a, 0x0b, 0x20, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x12, 0x2f, 0x0a, 0x0f, 0x6a, 0x6f, 0x62, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x18, 0x32, 0x52, 0x0d, 0x6a, 0x6f, 0x62, 0x47, 0x72, 0x61, 0x64, 0x65, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x27, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x9d, 0x05, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x2e, 0x52, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x52, 0x03, 0x6a,
	0x6f, 0x62, 0x12, 0x24, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x52, 0x08,
	0x6a, 0x6f, 0x62, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x2d, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x10, 0xfa, 0x42, 0x0d,
	0x1a, 0x0b, 0x20, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x08, 0x6a,
	0x6f, 0x62, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x0f, 0x6a, 0x6f, 0x62, 0x5f, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x52, 0x0d, 0x6a, 0x6f, 0x62, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x27, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x6f, 0x66, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x0a, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x6f, 0x66, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x02, 0x52, 0x03, 0x73, 0x65,
	0x78, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2a, 0x0a, 0x0c, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x14, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x69, 0x73, 0x75, 0x6d, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x48, 0x00, 0x52,
	0x05, 0x76, 0x69, 0x73, 0x75, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x1a, 0x02, 0x28, 0x00, 0x48, 0x01, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x74, 0x69, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x05,
	0x70, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x34, 0x0a, 0x08, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x52, 0x08, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x76, 0x69, 0x73, 0x75, 0x6d, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x03, 0x18, 0x3c, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x22, 0xb6, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x73,
	0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x1e, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x2a, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x4a, 0x6f,
	0x62, 0x48, 0x02, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6a, 0x6f, 0x62, 0x22, 0x9d, 0x03, 0x0a, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x41, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x3d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b,
	0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52,
	0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x40,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18,
	0x80, 0x02, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x09,
	0x6e, 0x65, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x02, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2a, 0x3d, 0x0a, 0x12, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x4d, 0x45, 0x4e, 0x54, 0x49, 0x4f, 0x4e, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6c, 0x65, 0x78, 0x72, 0x74,
	0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_resources_users_users_proto_rawDescOnce sync.Once
	file_resources_users_users_proto_rawDescData = file_resources_users_users_proto_rawDesc
)

func file_resources_users_users_proto_rawDescGZIP() []byte {
	file_resources_users_users_proto_rawDescOnce.Do(func() {
		file_resources_users_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_users_users_proto_rawDescData)
	})
	return file_resources_users_users_proto_rawDescData
}

var file_resources_users_users_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resources_users_users_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_resources_users_users_proto_goTypes = []interface{}{
	(USER_ACTIVITY_TYPE)(0),     // 0: resources.users.USER_ACTIVITY_TYPE
	(*UserShort)(nil),           // 1: resources.users.UserShort
	(*User)(nil),                // 2: resources.users.User
	(*License)(nil),             // 3: resources.users.License
	(*UserProps)(nil),           // 4: resources.users.UserProps
	(*UserActivity)(nil),        // 5: resources.users.UserActivity
	(*jobs.Job)(nil),            // 6: resources.jobs.Job
	(*timestamp.Timestamp)(nil), // 7: resources.timestamp.Timestamp
}
var file_resources_users_users_proto_depIdxs = []int32{
	4, // 0: resources.users.User.props:type_name -> resources.users.UserProps
	3, // 1: resources.users.User.licenses:type_name -> resources.users.License
	6, // 2: resources.users.UserProps.job:type_name -> resources.jobs.Job
	0, // 3: resources.users.UserActivity.type:type_name -> resources.users.USER_ACTIVITY_TYPE
	7, // 4: resources.users.UserActivity.created_at:type_name -> resources.timestamp.Timestamp
	1, // 5: resources.users.UserActivity.source_user:type_name -> resources.users.UserShort
	1, // 6: resources.users.UserActivity.target_user:type_name -> resources.users.UserShort
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_resources_users_users_proto_init() }
func file_resources_users_users_proto_init() {
	if File_resources_users_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_users_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserShort); i {
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
		file_resources_users_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_resources_users_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*License); i {
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
		file_resources_users_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProps); i {
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
		file_resources_users_users_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserActivity); i {
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
	file_resources_users_users_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_resources_users_users_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_users_users_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_users_users_proto_goTypes,
		DependencyIndexes: file_resources_users_users_proto_depIdxs,
		EnumInfos:         file_resources_users_users_proto_enumTypes,
		MessageInfos:      file_resources_users_users_proto_msgTypes,
	}.Build()
	File_resources_users_users_proto = out.File
	file_resources_users_users_proto_rawDesc = nil
	file_resources_users_users_proto_goTypes = nil
	file_resources_users_users_proto_depIdxs = nil
}
