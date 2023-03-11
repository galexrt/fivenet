// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: resources/users/users.proto

package users

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	timestamp "github.com/galexrt/arpanet/proto/resources/timestamp"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      int32      `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty" sql:"primary_key" alias:"id"`          // @gotags: sql:"primary_key" alias:"id"
	Identifier  string     `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty" alias:"identifier"`   // @gotags: alias:"identifier"
	Job         string     `protobuf:"bytes,3,opt,name=job,proto3" json:"job,omitempty" alias:"job"`                 // @gotags: alias:"job"
	JobGrade    int32      `protobuf:"varint,4,opt,name=jobGrade,proto3" json:"jobGrade,omitempty" alias:"job_grade"`      // @gotags: alias:"job_grade"
	Firstname   string     `protobuf:"bytes,5,opt,name=firstname,proto3" json:"firstname,omitempty" alias:"firstname"`     // @gotags: alias:"firstname"
	Lastname    string     `protobuf:"bytes,6,opt,name=lastname,proto3" json:"lastname,omitempty" alias:"lastname"`       // @gotags: alias:"lastname"
	Dateofbirth string     `protobuf:"bytes,7,opt,name=dateofbirth,proto3" json:"dateofbirth,omitempty" alias:"dateofbirth"` // @gotags: alias:"dateofbirth"
	Sex         string     `protobuf:"bytes,8,opt,name=sex,proto3" json:"sex,omitempty" alias:"sex"`                 // @gotags: alias:"sex"
	Height      string     `protobuf:"bytes,9,opt,name=height,proto3" json:"height,omitempty" alias:"height"`           // @gotags: alias:"height"
	Visum       int32      `protobuf:"varint,10,opt,name=visum,proto3" json:"visum,omitempty" alias:"visum"`           // @gotags: alias:"visum"
	Playtime    int32      `protobuf:"varint,11,opt,name=playtime,proto3" json:"playtime,omitempty" alias:"playtime"`     // @gotags: alias:"playtime"
	Props       *Props     `protobuf:"bytes,12,opt,name=props,proto3" json:"props,omitempty" alias:"arpanet_user_props"`            // @gotags: alias:"arpanet_user_props"
	Licenses    []*License `protobuf:"bytes,13,rep,name=licenses,proto3" json:"licenses,omitempty" alias:"user_licenses"`      // @gotags: alias:"user_licenses"
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_users_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_resources_users_users_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserID() int32 {
	if x != nil {
		return x.UserID
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

func (x *User) GetJobGrade() int32 {
	if x != nil {
		return x.JobGrade
	}
	return 0
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

func (x *User) GetVisum() int32 {
	if x != nil {
		return x.Visum
	}
	return 0
}

func (x *User) GetPlaytime() int32 {
	if x != nil {
		return x.Playtime
	}
	return 0
}

func (x *User) GetProps() *Props {
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

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" sql:"primary_key" alias:"type"` // @gotags: sql:"primary_key" alias:"type"
}

func (x *License) Reset() {
	*x = License{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_users_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *License) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License) ProtoMessage() {}

func (x *License) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use License.ProtoReflect.Descriptor instead.
func (*License) Descriptor() ([]byte, []int) {
	return file_resources_users_users_proto_rawDescGZIP(), []int{1}
}

func (x *License) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Props struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wanted bool `protobuf:"varint,1,opt,name=wanted,proto3" json:"wanted,omitempty" alias:"wanted"` // @gotags: alias:"wanted"
}

func (x *Props) Reset() {
	*x = Props{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_users_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Props) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Props) ProtoMessage() {}

func (x *Props) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Props.ProtoReflect.Descriptor instead.
func (*Props) Descriptor() ([]byte, []int) {
	return file_resources_users_users_proto_rawDescGZIP(), []int{2}
}

func (x *Props) GetWanted() bool {
	if x != nil {
		return x.Wanted
	}
	return false
}

type ShortUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     int32  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty" sql:"primary_key" alias:"id"`        // @gotags: sql:"primary_key" alias:"id"
	Identifier string `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty" alias:"identifier"` // @gotags: alias:"identifier"
	Job        string `protobuf:"bytes,3,opt,name=job,proto3" json:"job,omitempty" alias:"job"`               // @gotags: alias:"job"
	JobGrade   int32  `protobuf:"varint,4,opt,name=jobGrade,proto3" json:"jobGrade,omitempty" alias:"job_grade"`    // @gotags: alias:"job_grade"
	Firstname  string `protobuf:"bytes,5,opt,name=firstname,proto3" json:"firstname,omitempty" alias:"firstname"`   // @gotags: alias:"firstname"
	Lastname   string `protobuf:"bytes,6,opt,name=lastname,proto3" json:"lastname,omitempty" alias:"lastname"`     // @gotags: alias:"lastname"
}

func (x *ShortUser) Reset() {
	*x = ShortUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_users_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortUser) ProtoMessage() {}

func (x *ShortUser) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ShortUser.ProtoReflect.Descriptor instead.
func (*ShortUser) Descriptor() ([]byte, []int) {
	return file_resources_users_users_proto_rawDescGZIP(), []int{3}
}

func (x *ShortUser) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *ShortUser) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *ShortUser) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *ShortUser) GetJobGrade() int32 {
	if x != nil {
		return x.JobGrade
	}
	return 0
}

func (x *ShortUser) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *ShortUser) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

type UserActivity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" sql:"primary_key" alias:"arpanet_user_activity.id"`                // @gotags: sql:"primary_key" alias:"arpanet_user_activity.id"
	Type       string               `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty" alias:"arpanet_user_activity.type"`             // @gotags: alias:"arpanet_user_activity.type"
	CreatedAt  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty" alias:"arpanet_user_activity.created_at"`   // @gotags: alias:"arpanet_user_activity.created_at"
	TargetUser *ShortUser           `protobuf:"bytes,4,opt,name=targetUser,proto3" json:"targetUser,omitempty" alias:"target_user"` // @gotags: alias:"target_user"
	CauseUser  *ShortUser           `protobuf:"bytes,5,opt,name=causeUser,proto3" json:"causeUser,omitempty" alias:"cause_user"`   // @gotags: alias:"cause_user"
	Key        string               `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty" alias:"arpanet_user_activity.key"`               // @gotags: alias:"arpanet_user_activity.key"
	OldValue   string               `protobuf:"bytes,7,opt,name=oldValue,proto3" json:"oldValue,omitempty" alias:"arpanet_user_activity.old_value"`     // @gotags: alias:"arpanet_user_activity.old_value"
	NewValue   string               `protobuf:"bytes,8,opt,name=newValue,proto3" json:"newValue,omitempty" alias:"arpanet_user_activity.new_value"`     // @gotags: alias:"arpanet_user_activity.new_value"
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

func (x *UserActivity) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UserActivity) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserActivity) GetTargetUser() *ShortUser {
	if x != nil {
		return x.TargetUser
	}
	return nil
}

func (x *UserActivity) GetCauseUser() *ShortUser {
	if x != nil {
		return x.CauseUser
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

var File_resources_users_users_proto protoreflect.FileDescriptor

var file_resources_users_users_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x03, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x6f, 0x66, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x6f, 0x66, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69,
	0x73, 0x75, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x69, 0x73, 0x75, 0x6d,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x05,
	0x70, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x73, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x34, 0x0a, 0x08, 0x6c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73,
	0x22, 0x1d, 0x0a, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x1f, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x61, 0x6e, 0x74,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64,
	0x22, 0xb4, 0x01, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f,
	0x62, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x47, 0x72, 0x61, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb9, 0x02, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x38, 0x0a, 0x09, 0x63, 0x61, 0x75, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09,
	0x63, 0x61, 0x75, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6f,
	0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x61, 0x6c, 0x65, 0x78, 0x72, 0x74, 0x2f, 0x61, 0x72, 0x70, 0x61, 0x6e, 0x65,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_resources_users_users_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_resources_users_users_proto_goTypes = []interface{}{
	(*User)(nil),                // 0: resources.users.User
	(*License)(nil),             // 1: resources.users.License
	(*Props)(nil),               // 2: resources.users.Props
	(*ShortUser)(nil),           // 3: resources.users.ShortUser
	(*UserActivity)(nil),        // 4: resources.users.UserActivity
	(*timestamp.Timestamp)(nil), // 5: resources.timestamp.Timestamp
}
var file_resources_users_users_proto_depIdxs = []int32{
	2, // 0: resources.users.User.props:type_name -> resources.users.Props
	1, // 1: resources.users.User.licenses:type_name -> resources.users.License
	5, // 2: resources.users.UserActivity.createdAt:type_name -> resources.timestamp.Timestamp
	3, // 3: resources.users.UserActivity.targetUser:type_name -> resources.users.ShortUser
	3, // 4: resources.users.UserActivity.causeUser:type_name -> resources.users.ShortUser
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_resources_users_users_proto_init() }
func file_resources_users_users_proto_init() {
	if File_resources_users_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_users_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_resources_users_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_resources_users_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Props); i {
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
			switch v := v.(*ShortUser); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_users_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_users_users_proto_goTypes,
		DependencyIndexes: file_resources_users_users_proto_depIdxs,
		MessageInfos:      file_resources_users_users_proto_msgTypes,
	}.Build()
	File_resources_users_users_proto = out.File
	file_resources_users_users_proto_rawDesc = nil
	file_resources_users_users_proto_goTypes = nil
	file_resources_users_users_proto_depIdxs = nil
}
