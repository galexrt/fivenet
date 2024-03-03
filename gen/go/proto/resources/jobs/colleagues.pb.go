// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: resources/jobs/colleagues.proto

package jobs

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	filestore "github.com/galexrt/fivenet/gen/go/proto/resources/filestore"
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

type JobsUserActivityType int32

const (
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_UNSPECIFIED  JobsUserActivityType = 0
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_HIRED        JobsUserActivityType = 1
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_FIRED        JobsUserActivityType = 2
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_PROMOTED     JobsUserActivityType = 3
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_DEMOTED      JobsUserActivityType = 4
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_ABSENCE_DATE JobsUserActivityType = 5
)

// Enum value maps for JobsUserActivityType.
var (
	JobsUserActivityType_name = map[int32]string{
		0: "JOBS_USER_ACTIVITY_TYPE_UNSPECIFIED",
		1: "JOBS_USER_ACTIVITY_TYPE_HIRED",
		2: "JOBS_USER_ACTIVITY_TYPE_FIRED",
		3: "JOBS_USER_ACTIVITY_TYPE_PROMOTED",
		4: "JOBS_USER_ACTIVITY_TYPE_DEMOTED",
		5: "JOBS_USER_ACTIVITY_TYPE_ABSENCE_DATE",
	}
	JobsUserActivityType_value = map[string]int32{
		"JOBS_USER_ACTIVITY_TYPE_UNSPECIFIED":  0,
		"JOBS_USER_ACTIVITY_TYPE_HIRED":        1,
		"JOBS_USER_ACTIVITY_TYPE_FIRED":        2,
		"JOBS_USER_ACTIVITY_TYPE_PROMOTED":     3,
		"JOBS_USER_ACTIVITY_TYPE_DEMOTED":      4,
		"JOBS_USER_ACTIVITY_TYPE_ABSENCE_DATE": 5,
	}
)

func (x JobsUserActivityType) Enum() *JobsUserActivityType {
	p := new(JobsUserActivityType)
	*p = x
	return p
}

func (x JobsUserActivityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobsUserActivityType) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_jobs_colleagues_proto_enumTypes[0].Descriptor()
}

func (JobsUserActivityType) Type() protoreflect.EnumType {
	return &file_resources_jobs_colleagues_proto_enumTypes[0]
}

func (x JobsUserActivityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobsUserActivityType.Descriptor instead.
func (JobsUserActivityType) EnumDescriptor() ([]byte, []int) {
	return file_resources_jobs_colleagues_proto_rawDescGZIP(), []int{0}
}

type Colleague struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int32           `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" alias:"id"` // @gotags: alias:"id"
	Identifier    string          `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Job           string          `protobuf:"bytes,3,opt,name=job,proto3" json:"job,omitempty"`
	JobLabel      *string         `protobuf:"bytes,4,opt,name=job_label,json=jobLabel,proto3,oneof" json:"job_label,omitempty"`
	JobGrade      int32           `protobuf:"varint,5,opt,name=job_grade,json=jobGrade,proto3" json:"job_grade,omitempty"`
	JobGradeLabel *string         `protobuf:"bytes,6,opt,name=job_grade_label,json=jobGradeLabel,proto3,oneof" json:"job_grade_label,omitempty"`
	Firstname     string          `protobuf:"bytes,7,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname      string          `protobuf:"bytes,8,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Dateofbirth   string          `protobuf:"bytes,9,opt,name=dateofbirth,proto3" json:"dateofbirth,omitempty"`
	PhoneNumber   *string         `protobuf:"bytes,12,opt,name=phone_number,json=phoneNumber,proto3,oneof" json:"phone_number,omitempty"`
	Avatar        *filestore.File `protobuf:"bytes,17,opt,name=avatar,proto3,oneof" json:"avatar,omitempty"`
	Props         *JobsUserProps  `protobuf:"bytes,18,opt,name=props,proto3" json:"props,omitempty" alias:"fivenet_jobs_user_props"` // @gotags: alias:"fivenet_jobs_user_props"
}

func (x *Colleague) Reset() {
	*x = Colleague{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_jobs_colleagues_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Colleague) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Colleague) ProtoMessage() {}

func (x *Colleague) ProtoReflect() protoreflect.Message {
	mi := &file_resources_jobs_colleagues_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Colleague.ProtoReflect.Descriptor instead.
func (*Colleague) Descriptor() ([]byte, []int) {
	return file_resources_jobs_colleagues_proto_rawDescGZIP(), []int{0}
}

func (x *Colleague) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Colleague) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Colleague) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *Colleague) GetJobLabel() string {
	if x != nil && x.JobLabel != nil {
		return *x.JobLabel
	}
	return ""
}

func (x *Colleague) GetJobGrade() int32 {
	if x != nil {
		return x.JobGrade
	}
	return 0
}

func (x *Colleague) GetJobGradeLabel() string {
	if x != nil && x.JobGradeLabel != nil {
		return *x.JobGradeLabel
	}
	return ""
}

func (x *Colleague) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *Colleague) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *Colleague) GetDateofbirth() string {
	if x != nil {
		return x.Dateofbirth
	}
	return ""
}

func (x *Colleague) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *Colleague) GetAvatar() *filestore.File {
	if x != nil {
		return x.Avatar
	}
	return nil
}

func (x *Colleague) GetProps() *JobsUserProps {
	if x != nil {
		return x.Props
	}
	return nil
}

type JobsUserProps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int32                `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AbsenceBegin *timestamp.Timestamp `protobuf:"bytes,2,opt,name=absence_begin,json=absenceBegin,proto3,oneof" json:"absence_begin,omitempty"`
	AbsenceEnd   *timestamp.Timestamp `protobuf:"bytes,3,opt,name=absence_end,json=absenceEnd,proto3,oneof" json:"absence_end,omitempty"`
}

func (x *JobsUserProps) Reset() {
	*x = JobsUserProps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_jobs_colleagues_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobsUserProps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobsUserProps) ProtoMessage() {}

func (x *JobsUserProps) ProtoReflect() protoreflect.Message {
	mi := &file_resources_jobs_colleagues_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobsUserProps.ProtoReflect.Descriptor instead.
func (*JobsUserProps) Descriptor() ([]byte, []int) {
	return file_resources_jobs_colleagues_proto_rawDescGZIP(), []int{1}
}

func (x *JobsUserProps) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *JobsUserProps) GetAbsenceBegin() *timestamp.Timestamp {
	if x != nil {
		return x.AbsenceBegin
	}
	return nil
}

func (x *JobsUserProps) GetAbsenceEnd() *timestamp.Timestamp {
	if x != nil {
		return x.AbsenceEnd
	}
	return nil
}

type JobsUserActivity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" sql:"primary_key" alias:"id"` // @gotags: sql:"primary_key" alias:"id"
	CreatedAt    *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	Job          string               `protobuf:"bytes,4,opt,name=job,proto3" json:"job,omitempty"`
	SourceUserId int32                `protobuf:"varint,5,opt,name=source_user_id,json=sourceUserId,proto3" json:"source_user_id,omitempty"`
	SourceUser   *users.UserShort     `protobuf:"bytes,6,opt,name=source_user,json=sourceUser,proto3" json:"source_user,omitempty" alias:"source_user"` // @gotags: alias:"source_user"
	TargetUserId int32                `protobuf:"varint,7,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
	TargetUser   *users.UserShort     `protobuf:"bytes,8,opt,name=target_user,json=targetUser,proto3" json:"target_user,omitempty" alias:"target_user"` // @gotags: alias:"target_user"
	ActivityType JobsUserActivityType `protobuf:"varint,9,opt,name=activity_type,json=activityType,proto3,enum=resources.jobs.JobsUserActivityType" json:"activity_type,omitempty"`
	// @sanitize
	Reason string                `protobuf:"bytes,10,opt,name=reason,proto3" json:"reason,omitempty"`
	Data   *JobsUserActivityData `protobuf:"bytes,11,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *JobsUserActivity) Reset() {
	*x = JobsUserActivity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_jobs_colleagues_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobsUserActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobsUserActivity) ProtoMessage() {}

func (x *JobsUserActivity) ProtoReflect() protoreflect.Message {
	mi := &file_resources_jobs_colleagues_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobsUserActivity.ProtoReflect.Descriptor instead.
func (*JobsUserActivity) Descriptor() ([]byte, []int) {
	return file_resources_jobs_colleagues_proto_rawDescGZIP(), []int{2}
}

func (x *JobsUserActivity) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *JobsUserActivity) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *JobsUserActivity) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *JobsUserActivity) GetSourceUserId() int32 {
	if x != nil {
		return x.SourceUserId
	}
	return 0
}

func (x *JobsUserActivity) GetSourceUser() *users.UserShort {
	if x != nil {
		return x.SourceUser
	}
	return nil
}

func (x *JobsUserActivity) GetTargetUserId() int32 {
	if x != nil {
		return x.TargetUserId
	}
	return 0
}

func (x *JobsUserActivity) GetTargetUser() *users.UserShort {
	if x != nil {
		return x.TargetUser
	}
	return nil
}

func (x *JobsUserActivity) GetActivityType() JobsUserActivityType {
	if x != nil {
		return x.ActivityType
	}
	return JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_UNSPECIFIED
}

func (x *JobsUserActivity) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *JobsUserActivity) GetData() *JobsUserActivityData {
	if x != nil {
		return x.Data
	}
	return nil
}

type JobsUserActivityData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*JobsUserActivityData_AbsenceDate
	//	*JobsUserActivityData_GradeChange
	Data isJobsUserActivityData_Data `protobuf_oneof:"data"`
}

func (x *JobsUserActivityData) Reset() {
	*x = JobsUserActivityData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_jobs_colleagues_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobsUserActivityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobsUserActivityData) ProtoMessage() {}

func (x *JobsUserActivityData) ProtoReflect() protoreflect.Message {
	mi := &file_resources_jobs_colleagues_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobsUserActivityData.ProtoReflect.Descriptor instead.
func (*JobsUserActivityData) Descriptor() ([]byte, []int) {
	return file_resources_jobs_colleagues_proto_rawDescGZIP(), []int{3}
}

func (m *JobsUserActivityData) GetData() isJobsUserActivityData_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *JobsUserActivityData) GetAbsenceDate() *ColleagueAbsenceDate {
	if x, ok := x.GetData().(*JobsUserActivityData_AbsenceDate); ok {
		return x.AbsenceDate
	}
	return nil
}

func (x *JobsUserActivityData) GetGradeChange() *ColleagueGradeChange {
	if x, ok := x.GetData().(*JobsUserActivityData_GradeChange); ok {
		return x.GradeChange
	}
	return nil
}

type isJobsUserActivityData_Data interface {
	isJobsUserActivityData_Data()
}

type JobsUserActivityData_AbsenceDate struct {
	AbsenceDate *ColleagueAbsenceDate `protobuf:"bytes,1,opt,name=absence_date,json=absenceDate,proto3,oneof"`
}

type JobsUserActivityData_GradeChange struct {
	GradeChange *ColleagueGradeChange `protobuf:"bytes,2,opt,name=grade_change,json=gradeChange,proto3,oneof"`
}

func (*JobsUserActivityData_AbsenceDate) isJobsUserActivityData_Data() {}

func (*JobsUserActivityData_GradeChange) isJobsUserActivityData_Data() {}

type ColleagueAbsenceDate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbsenceBegin *timestamp.Timestamp `protobuf:"bytes,1,opt,name=absence_begin,json=absenceBegin,proto3" json:"absence_begin,omitempty"`
	AbsenceEnd   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=absence_end,json=absenceEnd,proto3" json:"absence_end,omitempty"`
}

func (x *ColleagueAbsenceDate) Reset() {
	*x = ColleagueAbsenceDate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_jobs_colleagues_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColleagueAbsenceDate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColleagueAbsenceDate) ProtoMessage() {}

func (x *ColleagueAbsenceDate) ProtoReflect() protoreflect.Message {
	mi := &file_resources_jobs_colleagues_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColleagueAbsenceDate.ProtoReflect.Descriptor instead.
func (*ColleagueAbsenceDate) Descriptor() ([]byte, []int) {
	return file_resources_jobs_colleagues_proto_rawDescGZIP(), []int{4}
}

func (x *ColleagueAbsenceDate) GetAbsenceBegin() *timestamp.Timestamp {
	if x != nil {
		return x.AbsenceBegin
	}
	return nil
}

func (x *ColleagueAbsenceDate) GetAbsenceEnd() *timestamp.Timestamp {
	if x != nil {
		return x.AbsenceEnd
	}
	return nil
}

type ColleagueGradeChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grade      int32  `protobuf:"varint,1,opt,name=grade,proto3" json:"grade,omitempty"`
	GradeLabel string `protobuf:"bytes,2,opt,name=grade_label,json=gradeLabel,proto3" json:"grade_label,omitempty"`
}

func (x *ColleagueGradeChange) Reset() {
	*x = ColleagueGradeChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_jobs_colleagues_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColleagueGradeChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColleagueGradeChange) ProtoMessage() {}

func (x *ColleagueGradeChange) ProtoReflect() protoreflect.Message {
	mi := &file_resources_jobs_colleagues_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColleagueGradeChange.ProtoReflect.Descriptor instead.
func (*ColleagueGradeChange) Descriptor() ([]byte, []int) {
	return file_resources_jobs_colleagues_proto_rawDescGZIP(), []int{5}
}

func (x *ColleagueGradeChange) GetGrade() int32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *ColleagueGradeChange) GetGradeLabel() string {
	if x != nil {
		return x.GradeLabel
	}
	return ""
}

var File_resources_jobs_colleagues_proto protoreflect.FileDescriptor

var file_resources_jobs_colleagues_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x73,
	0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62,
	0x73, 0x1a, 0x1e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x04, 0x0a,
	0x09, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x1a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0a,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x2e, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x14, 0x52, 0x03, 0x6a, 0x6f,
	0x62, 0x12, 0x29, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x48, 0x00, 0x52,
	0x08, 0x6a, 0x6f, 0x62, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x09,
	0x6a, 0x6f, 0x62, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x10, 0xfa, 0x42, 0x0d, 0x1a, 0x0b, 0x20, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x0f, 0x6a,
	0x6f, 0x62, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x48, 0x01, 0x52,
	0x0d, 0x6a, 0x6f, 0x62, 0x47, 0x72, 0x61, 0x64, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01,
	0x01, 0x12, 0x27, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x6f, 0x66, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x0a,
	0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x6f, 0x66, 0x62, 0x69, 0x72, 0x74, 0x68, 0x12, 0x2f, 0x0a,
	0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x14, 0x48, 0x02, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x36,
	0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x03, 0x52, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x70, 0x73, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x6a, 0x6f, 0x62, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6a, 0x6f,
	0x62, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x0f, 0x0a,
	0x0d, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0xe3, 0x01, 0x0a, 0x0d, 0x4a, 0x6f,
	0x62, 0x73, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x20, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x48, 0x0a,
	0x0d, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0c, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x42,
	0x65, 0x67, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0b, 0x61, 0x62, 0x73, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x0a,
	0x61, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x22,
	0x93, 0x04, 0x0a, 0x10, 0x4a, 0x6f, 0x62, 0x73, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x03,
	0x6a, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x18, 0x14, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x2d, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x1a, 0x02, 0x20, 0x00, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x49, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0x18, 0xff, 0x01, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x73,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0xb4, 0x01, 0x0a, 0x14, 0x4a, 0x6f, 0x62, 0x73, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x49,
	0x0a, 0x0c, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x41,
	0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x62,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73,
	0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9c, 0x01, 0x0a,
	0x14, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x41, 0x62, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x61, 0x62,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x3f, 0x0a, 0x0b, 0x61, 0x62,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x22, 0x4d, 0x0a, 0x14, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x2a, 0xfa, 0x01, 0x0a, 0x14, 0x4a,
	0x6f, 0x62, 0x73, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x23, 0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d,
	0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49,
	0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x49, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x21, 0x0a, 0x1d, 0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x52, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52,
	0x4f, 0x4d, 0x4f, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x4a, 0x4f, 0x42, 0x53,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x28, 0x0a,
	0x24, 0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x42, 0x53, 0x45, 0x4e, 0x43, 0x45,
	0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x05, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6c, 0x65, 0x78, 0x72, 0x74, 0x2f, 0x66, 0x69,
	0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62,
	0x73, 0x3b, 0x6a, 0x6f, 0x62, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_jobs_colleagues_proto_rawDescOnce sync.Once
	file_resources_jobs_colleagues_proto_rawDescData = file_resources_jobs_colleagues_proto_rawDesc
)

func file_resources_jobs_colleagues_proto_rawDescGZIP() []byte {
	file_resources_jobs_colleagues_proto_rawDescOnce.Do(func() {
		file_resources_jobs_colleagues_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_jobs_colleagues_proto_rawDescData)
	})
	return file_resources_jobs_colleagues_proto_rawDescData
}

var file_resources_jobs_colleagues_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resources_jobs_colleagues_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_resources_jobs_colleagues_proto_goTypes = []interface{}{
	(JobsUserActivityType)(0),    // 0: resources.jobs.JobsUserActivityType
	(*Colleague)(nil),            // 1: resources.jobs.Colleague
	(*JobsUserProps)(nil),        // 2: resources.jobs.JobsUserProps
	(*JobsUserActivity)(nil),     // 3: resources.jobs.JobsUserActivity
	(*JobsUserActivityData)(nil), // 4: resources.jobs.JobsUserActivityData
	(*ColleagueAbsenceDate)(nil), // 5: resources.jobs.ColleagueAbsenceDate
	(*ColleagueGradeChange)(nil), // 6: resources.jobs.ColleagueGradeChange
	(*filestore.File)(nil),       // 7: resources.filestore.File
	(*timestamp.Timestamp)(nil),  // 8: resources.timestamp.Timestamp
	(*users.UserShort)(nil),      // 9: resources.users.UserShort
}
var file_resources_jobs_colleagues_proto_depIdxs = []int32{
	7,  // 0: resources.jobs.Colleague.avatar:type_name -> resources.filestore.File
	2,  // 1: resources.jobs.Colleague.props:type_name -> resources.jobs.JobsUserProps
	8,  // 2: resources.jobs.JobsUserProps.absence_begin:type_name -> resources.timestamp.Timestamp
	8,  // 3: resources.jobs.JobsUserProps.absence_end:type_name -> resources.timestamp.Timestamp
	8,  // 4: resources.jobs.JobsUserActivity.created_at:type_name -> resources.timestamp.Timestamp
	9,  // 5: resources.jobs.JobsUserActivity.source_user:type_name -> resources.users.UserShort
	9,  // 6: resources.jobs.JobsUserActivity.target_user:type_name -> resources.users.UserShort
	0,  // 7: resources.jobs.JobsUserActivity.activity_type:type_name -> resources.jobs.JobsUserActivityType
	4,  // 8: resources.jobs.JobsUserActivity.data:type_name -> resources.jobs.JobsUserActivityData
	5,  // 9: resources.jobs.JobsUserActivityData.absence_date:type_name -> resources.jobs.ColleagueAbsenceDate
	6,  // 10: resources.jobs.JobsUserActivityData.grade_change:type_name -> resources.jobs.ColleagueGradeChange
	8,  // 11: resources.jobs.ColleagueAbsenceDate.absence_begin:type_name -> resources.timestamp.Timestamp
	8,  // 12: resources.jobs.ColleagueAbsenceDate.absence_end:type_name -> resources.timestamp.Timestamp
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_resources_jobs_colleagues_proto_init() }
func file_resources_jobs_colleagues_proto_init() {
	if File_resources_jobs_colleagues_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_jobs_colleagues_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Colleague); i {
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
		file_resources_jobs_colleagues_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobsUserProps); i {
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
		file_resources_jobs_colleagues_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobsUserActivity); i {
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
		file_resources_jobs_colleagues_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobsUserActivityData); i {
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
		file_resources_jobs_colleagues_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColleagueAbsenceDate); i {
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
		file_resources_jobs_colleagues_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColleagueGradeChange); i {
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
	file_resources_jobs_colleagues_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_resources_jobs_colleagues_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_resources_jobs_colleagues_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_resources_jobs_colleagues_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*JobsUserActivityData_AbsenceDate)(nil),
		(*JobsUserActivityData_GradeChange)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_jobs_colleagues_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_jobs_colleagues_proto_goTypes,
		DependencyIndexes: file_resources_jobs_colleagues_proto_depIdxs,
		EnumInfos:         file_resources_jobs_colleagues_proto_enumTypes,
		MessageInfos:      file_resources_jobs_colleagues_proto_msgTypes,
	}.Build()
	File_resources_jobs_colleagues_proto = out.File
	file_resources_jobs_colleagues_proto_rawDesc = nil
	file_resources_jobs_colleagues_proto_goTypes = nil
	file_resources_jobs_colleagues_proto_depIdxs = nil
}
