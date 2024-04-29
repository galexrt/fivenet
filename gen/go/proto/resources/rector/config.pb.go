// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: resources/rector/config.proto

package rector

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AppConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth        *Auth        `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Perms       *Perms       `protobuf:"bytes,2,opt,name=perms,proto3" json:"perms,omitempty"`
	Website     *Website     `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
	JobInfo     *JobInfo     `protobuf:"bytes,4,opt,name=job_info,json=jobInfo,proto3" json:"job_info,omitempty"`
	UserTracker *UserTracker `protobuf:"bytes,5,opt,name=user_tracker,json=userTracker,proto3" json:"user_tracker,omitempty"`
	Discord     *Discord     `protobuf:"bytes,6,opt,name=discord,proto3" json:"discord,omitempty"`
}

func (x *AppConfig) Reset() {
	*x = AppConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_rector_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConfig) ProtoMessage() {}

func (x *AppConfig) ProtoReflect() protoreflect.Message {
	mi := &file_resources_rector_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConfig.ProtoReflect.Descriptor instead.
func (*AppConfig) Descriptor() ([]byte, []int) {
	return file_resources_rector_config_proto_rawDescGZIP(), []int{0}
}

func (x *AppConfig) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *AppConfig) GetPerms() *Perms {
	if x != nil {
		return x.Perms
	}
	return nil
}

func (x *AppConfig) GetWebsite() *Website {
	if x != nil {
		return x.Website
	}
	return nil
}

func (x *AppConfig) GetJobInfo() *JobInfo {
	if x != nil {
		return x.JobInfo
	}
	return nil
}

func (x *AppConfig) GetUserTracker() *UserTracker {
	if x != nil {
		return x.UserTracker
	}
	return nil
}

func (x *AppConfig) GetDiscord() *Discord {
	if x != nil {
		return x.Discord
	}
	return nil
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignupEnabled bool `protobuf:"varint,1,opt,name=signup_enabled,json=signupEnabled,proto3" json:"signup_enabled,omitempty"`
	LastCharLock  bool `protobuf:"varint,2,opt,name=last_char_lock,json=lastCharLock,proto3" json:"last_char_lock,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_rector_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_resources_rector_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_resources_rector_config_proto_rawDescGZIP(), []int{1}
}

func (x *Auth) GetSignupEnabled() bool {
	if x != nil {
		return x.SignupEnabled
	}
	return false
}

func (x *Auth) GetLastCharLock() bool {
	if x != nil {
		return x.LastCharLock
	}
	return false
}

type Perms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Default []*Perm `protobuf:"bytes,1,rep,name=default,proto3" json:"default,omitempty"`
}

func (x *Perms) Reset() {
	*x = Perms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_rector_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Perms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Perms) ProtoMessage() {}

func (x *Perms) ProtoReflect() protoreflect.Message {
	mi := &file_resources_rector_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Perms.ProtoReflect.Descriptor instead.
func (*Perms) Descriptor() ([]byte, []int) {
	return file_resources_rector_config_proto_rawDescGZIP(), []int{2}
}

func (x *Perms) GetDefault() []*Perm {
	if x != nil {
		return x.Default
	}
	return nil
}

type Perm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Perm) Reset() {
	*x = Perm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_rector_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Perm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Perm) ProtoMessage() {}

func (x *Perm) ProtoReflect() protoreflect.Message {
	mi := &file_resources_rector_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Perm.ProtoReflect.Descriptor instead.
func (*Perm) Descriptor() ([]byte, []int) {
	return file_resources_rector_config_proto_rawDescGZIP(), []int{3}
}

func (x *Perm) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Perm) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Website struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links *Links `protobuf:"bytes,1,opt,name=links,proto3" json:"links,omitempty"`
}

func (x *Website) Reset() {
	*x = Website{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_rector_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Website) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Website) ProtoMessage() {}

func (x *Website) ProtoReflect() protoreflect.Message {
	mi := &file_resources_rector_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Website.ProtoReflect.Descriptor instead.
func (*Website) Descriptor() ([]byte, []int) {
	return file_resources_rector_config_proto_rawDescGZIP(), []int{4}
}

func (x *Website) GetLinks() *Links {
	if x != nil {
		return x.Links
	}
	return nil
}

type Links struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivacyPolicy *string `protobuf:"bytes,1,opt,name=privacy_policy,json=privacyPolicy,proto3,oneof" json:"privacy_policy,omitempty"`
	Imprint       *string `protobuf:"bytes,2,opt,name=imprint,proto3,oneof" json:"imprint,omitempty"`
}

func (x *Links) Reset() {
	*x = Links{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_rector_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Links) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Links) ProtoMessage() {}

func (x *Links) ProtoReflect() protoreflect.Message {
	mi := &file_resources_rector_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Links.ProtoReflect.Descriptor instead.
func (*Links) Descriptor() ([]byte, []int) {
	return file_resources_rector_config_proto_rawDescGZIP(), []int{5}
}

func (x *Links) GetPrivacyPolicy() string {
	if x != nil && x.PrivacyPolicy != nil {
		return *x.PrivacyPolicy
	}
	return ""
}

func (x *Links) GetImprint() string {
	if x != nil && x.Imprint != nil {
		return *x.Imprint
	}
	return ""
}

type JobInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnemployedJob *UnemployedJob `protobuf:"bytes,1,opt,name=unemployed_job,json=unemployedJob,proto3" json:"unemployed_job,omitempty"`
	PublicJobs    []string       `protobuf:"bytes,2,rep,name=public_jobs,json=publicJobs,proto3" json:"public_jobs,omitempty"`
	HiddenJobs    []string       `protobuf:"bytes,3,rep,name=hidden_jobs,json=hiddenJobs,proto3" json:"hidden_jobs,omitempty"`
}

func (x *JobInfo) Reset() {
	*x = JobInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_rector_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobInfo) ProtoMessage() {}

func (x *JobInfo) ProtoReflect() protoreflect.Message {
	mi := &file_resources_rector_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobInfo.ProtoReflect.Descriptor instead.
func (*JobInfo) Descriptor() ([]byte, []int) {
	return file_resources_rector_config_proto_rawDescGZIP(), []int{6}
}

func (x *JobInfo) GetUnemployedJob() *UnemployedJob {
	if x != nil {
		return x.UnemployedJob
	}
	return nil
}

func (x *JobInfo) GetPublicJobs() []string {
	if x != nil {
		return x.PublicJobs
	}
	return nil
}

func (x *JobInfo) GetHiddenJobs() []string {
	if x != nil {
		return x.HiddenJobs
	}
	return nil
}

type UnemployedJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Grade int32  `protobuf:"varint,2,opt,name=grade,proto3" json:"grade,omitempty"`
}

func (x *UnemployedJob) Reset() {
	*x = UnemployedJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_rector_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnemployedJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnemployedJob) ProtoMessage() {}

func (x *UnemployedJob) ProtoReflect() protoreflect.Message {
	mi := &file_resources_rector_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnemployedJob.ProtoReflect.Descriptor instead.
func (*UnemployedJob) Descriptor() ([]byte, []int) {
	return file_resources_rector_config_proto_rawDescGZIP(), []int{7}
}

func (x *UnemployedJob) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UnemployedJob) GetGrade() int32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

type UserTracker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshTime   *durationpb.Duration `protobuf:"bytes,1,opt,name=refresh_time,json=refreshTime,proto3" json:"refresh_time,omitempty"`
	DbRefreshTime *durationpb.Duration `protobuf:"bytes,2,opt,name=db_refresh_time,json=dbRefreshTime,proto3" json:"db_refresh_time,omitempty"`
	LivemapJobs   []string             `protobuf:"bytes,3,rep,name=livemap_jobs,json=livemapJobs,proto3" json:"livemap_jobs,omitempty"`
}

func (x *UserTracker) Reset() {
	*x = UserTracker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_rector_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTracker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTracker) ProtoMessage() {}

func (x *UserTracker) ProtoReflect() protoreflect.Message {
	mi := &file_resources_rector_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTracker.ProtoReflect.Descriptor instead.
func (*UserTracker) Descriptor() ([]byte, []int) {
	return file_resources_rector_config_proto_rawDescGZIP(), []int{8}
}

func (x *UserTracker) GetRefreshTime() *durationpb.Duration {
	if x != nil {
		return x.RefreshTime
	}
	return nil
}

func (x *UserTracker) GetDbRefreshTime() *durationpb.Duration {
	if x != nil {
		return x.DbRefreshTime
	}
	return nil
}

func (x *UserTracker) GetLivemapJobs() []string {
	if x != nil {
		return x.LivemapJobs
	}
	return nil
}

type Discord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled      bool                 `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	SyncInterval *durationpb.Duration `protobuf:"bytes,2,opt,name=sync_interval,json=syncInterval,proto3" json:"sync_interval,omitempty"`
	InviteUrl    *string              `protobuf:"bytes,3,opt,name=invite_url,json=inviteUrl,proto3,oneof" json:"invite_url,omitempty"`
	IgnoredJobs  []string             `protobuf:"bytes,4,rep,name=ignored_jobs,json=ignoredJobs,proto3" json:"ignored_jobs,omitempty"`
}

func (x *Discord) Reset() {
	*x = Discord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_rector_config_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Discord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Discord) ProtoMessage() {}

func (x *Discord) ProtoReflect() protoreflect.Message {
	mi := &file_resources_rector_config_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Discord.ProtoReflect.Descriptor instead.
func (*Discord) Descriptor() ([]byte, []int) {
	return file_resources_rector_config_proto_rawDescGZIP(), []int{9}
}

func (x *Discord) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Discord) GetSyncInterval() *durationpb.Duration {
	if x != nil {
		return x.SyncInterval
	}
	return nil
}

func (x *Discord) GetInviteUrl() string {
	if x != nil && x.InviteUrl != nil {
		return *x.InviteUrl
	}
	return ""
}

func (x *Discord) GetIgnoredJobs() []string {
	if x != nil {
		return x.IgnoredJobs
	}
	return nil
}

var File_resources_rector_config_proto protoreflect.FileDescriptor

var file_resources_rector_config_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x02, 0x0a, 0x09, 0x41,
	0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x12, 0x2d, 0x0a, 0x05, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x73, 0x52, 0x05, 0x70, 0x65,
	0x72, 0x6d, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52,
	0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4a, 0x6f,
	0x62, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40,
	0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x12, 0x33, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x53, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x61,
	0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x4c, 0x6f, 0x63, 0x6b, 0x22, 0x43, 0x0a, 0x05, 0x50, 0x65,
	0x72, 0x6d, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x92, 0x01, 0x02, 0x10, 0x64, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22,
	0x4a, 0x0a, 0x04, 0x50, 0x65, 0x72, 0x6d, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03,
	0x18, 0x80, 0x01, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x72, 0x03, 0x18, 0xff, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x07, 0x57,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x05,
	0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x05, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12,
	0x34, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0xff,
	0x01, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x07, 0x69, 0x6d, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0xff, 0x01,
	0x48, 0x01, 0x52, 0x07, 0x69, 0x6d, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x11,
	0x0a, 0x0f, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x69, 0x6d, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x22, 0xb1, 0x01,
	0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x50, 0x0a, 0x0e, 0x75, 0x6e, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x55, 0x6e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x4a,
	0x6f, 0x62, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x75, 0x6e,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x12, 0x29, 0x0a, 0x0b, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x64, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x29, 0x0a, 0x0b, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e,
	0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x10, 0x64, 0x52, 0x0a, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x4a, 0x6f, 0x62,
	0x73, 0x22, 0x4b, 0x0a, 0x0d, 0x55, 0x6e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x4a,
	0x6f, 0x62, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x14, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x22, 0xe7,
	0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x52,
	0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x14, 0xfa, 0x42, 0x11, 0xaa, 0x01, 0x0e, 0x08, 0x01, 0x1a, 0x02, 0x08, 0x3c, 0x32, 0x06, 0x10,
	0x80, 0xca, 0xb5, 0xee, 0x01, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x57, 0x0a, 0x0f, 0x64, 0x62, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x14, 0xfa, 0x42, 0x11, 0xaa, 0x01, 0x0e, 0x08, 0x01,
	0x1a, 0x02, 0x08, 0x3c, 0x32, 0x06, 0x10, 0x80, 0xca, 0xb5, 0xee, 0x01, 0x52, 0x0d, 0x64, 0x62,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x0c, 0x6c,
	0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x64, 0x52, 0x0b, 0x6c, 0x69, 0x76,
	0x65, 0x6d, 0x61, 0x70, 0x4a, 0x6f, 0x62, 0x73, 0x22, 0xe2, 0x01, 0x0a, 0x07, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x53,
	0x0a, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x13, 0xfa, 0x42, 0x10, 0xaa, 0x01, 0x0d, 0x08, 0x01, 0x1a, 0x05, 0x08, 0x80, 0xaa, 0xea,
	0x55, 0x32, 0x02, 0x08, 0x3c, 0x52, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0xff,
	0x01, 0x48, 0x00, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x88, 0x01,
	0x01, 0x12, 0x2b, 0x0a, 0x0c, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x6a, 0x6f, 0x62,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10,
	0x64, 0x52, 0x0b, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x73, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x41, 0x5a,
	0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6c, 0x65,
	0x78, 0x72, 0x74, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x3b, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_rector_config_proto_rawDescOnce sync.Once
	file_resources_rector_config_proto_rawDescData = file_resources_rector_config_proto_rawDesc
)

func file_resources_rector_config_proto_rawDescGZIP() []byte {
	file_resources_rector_config_proto_rawDescOnce.Do(func() {
		file_resources_rector_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_rector_config_proto_rawDescData)
	})
	return file_resources_rector_config_proto_rawDescData
}

var file_resources_rector_config_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_resources_rector_config_proto_goTypes = []interface{}{
	(*AppConfig)(nil),           // 0: resources.rector.AppConfig
	(*Auth)(nil),                // 1: resources.rector.Auth
	(*Perms)(nil),               // 2: resources.rector.Perms
	(*Perm)(nil),                // 3: resources.rector.Perm
	(*Website)(nil),             // 4: resources.rector.Website
	(*Links)(nil),               // 5: resources.rector.Links
	(*JobInfo)(nil),             // 6: resources.rector.JobInfo
	(*UnemployedJob)(nil),       // 7: resources.rector.UnemployedJob
	(*UserTracker)(nil),         // 8: resources.rector.UserTracker
	(*Discord)(nil),             // 9: resources.rector.Discord
	(*durationpb.Duration)(nil), // 10: google.protobuf.Duration
}
var file_resources_rector_config_proto_depIdxs = []int32{
	1,  // 0: resources.rector.AppConfig.auth:type_name -> resources.rector.Auth
	2,  // 1: resources.rector.AppConfig.perms:type_name -> resources.rector.Perms
	4,  // 2: resources.rector.AppConfig.website:type_name -> resources.rector.Website
	6,  // 3: resources.rector.AppConfig.job_info:type_name -> resources.rector.JobInfo
	8,  // 4: resources.rector.AppConfig.user_tracker:type_name -> resources.rector.UserTracker
	9,  // 5: resources.rector.AppConfig.discord:type_name -> resources.rector.Discord
	3,  // 6: resources.rector.Perms.default:type_name -> resources.rector.Perm
	5,  // 7: resources.rector.Website.links:type_name -> resources.rector.Links
	7,  // 8: resources.rector.JobInfo.unemployed_job:type_name -> resources.rector.UnemployedJob
	10, // 9: resources.rector.UserTracker.refresh_time:type_name -> google.protobuf.Duration
	10, // 10: resources.rector.UserTracker.db_refresh_time:type_name -> google.protobuf.Duration
	10, // 11: resources.rector.Discord.sync_interval:type_name -> google.protobuf.Duration
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_resources_rector_config_proto_init() }
func file_resources_rector_config_proto_init() {
	if File_resources_rector_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_rector_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConfig); i {
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
		file_resources_rector_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_resources_rector_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Perms); i {
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
		file_resources_rector_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Perm); i {
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
		file_resources_rector_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Website); i {
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
		file_resources_rector_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Links); i {
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
		file_resources_rector_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobInfo); i {
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
		file_resources_rector_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnemployedJob); i {
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
		file_resources_rector_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTracker); i {
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
		file_resources_rector_config_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Discord); i {
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
	file_resources_rector_config_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_resources_rector_config_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_rector_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_rector_config_proto_goTypes,
		DependencyIndexes: file_resources_rector_config_proto_depIdxs,
		MessageInfos:      file_resources_rector_config_proto_msgTypes,
	}.Build()
	File_resources_rector_config_proto = out.File
	file_resources_rector_config_proto_rawDesc = nil
	file_resources_rector_config_proto_goTypes = nil
	file_resources_rector_config_proto_depIdxs = nil
}
