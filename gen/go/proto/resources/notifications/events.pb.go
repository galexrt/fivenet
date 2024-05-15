// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: resources/notifications/events.proto

package notifications

import (
	users "github.com/fivenet-app/fivenet/gen/go/proto/resources/users"
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

type UserEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*UserEvent_Notification
	//	*UserEvent_RefreshToken
	Data isUserEvent_Data `protobuf_oneof:"data"`
}

func (x *UserEvent) Reset() {
	*x = UserEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_notifications_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEvent) ProtoMessage() {}

func (x *UserEvent) ProtoReflect() protoreflect.Message {
	mi := &file_resources_notifications_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEvent.ProtoReflect.Descriptor instead.
func (*UserEvent) Descriptor() ([]byte, []int) {
	return file_resources_notifications_events_proto_rawDescGZIP(), []int{0}
}

func (m *UserEvent) GetData() isUserEvent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UserEvent) GetNotification() *Notification {
	if x, ok := x.GetData().(*UserEvent_Notification); ok {
		return x.Notification
	}
	return nil
}

func (x *UserEvent) GetRefreshToken() bool {
	if x, ok := x.GetData().(*UserEvent_RefreshToken); ok {
		return x.RefreshToken
	}
	return false
}

type isUserEvent_Data interface {
	isUserEvent_Data()
}

type UserEvent_Notification struct {
	Notification *Notification `protobuf:"bytes,3,opt,name=notification,proto3,oneof"`
}

type UserEvent_RefreshToken struct {
	RefreshToken bool `protobuf:"varint,4,opt,name=refresh_token,json=refreshToken,proto3,oneof"`
}

func (*UserEvent_Notification) isUserEvent_Data() {}

func (*UserEvent_RefreshToken) isUserEvent_Data() {}

type JobEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*JobEvent_JobProps
	Data isJobEvent_Data `protobuf_oneof:"data"`
}

func (x *JobEvent) Reset() {
	*x = JobEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_notifications_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobEvent) ProtoMessage() {}

func (x *JobEvent) ProtoReflect() protoreflect.Message {
	mi := &file_resources_notifications_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobEvent.ProtoReflect.Descriptor instead.
func (*JobEvent) Descriptor() ([]byte, []int) {
	return file_resources_notifications_events_proto_rawDescGZIP(), []int{1}
}

func (m *JobEvent) GetData() isJobEvent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *JobEvent) GetJobProps() *users.JobProps {
	if x, ok := x.GetData().(*JobEvent_JobProps); ok {
		return x.JobProps
	}
	return nil
}

type isJobEvent_Data interface {
	isJobEvent_Data()
}

type JobEvent_JobProps struct {
	JobProps *users.JobProps `protobuf:"bytes,1,opt,name=job_props,json=jobProps,proto3,oneof"`
}

func (*JobEvent_JobProps) isJobEvent_Data() {}

type SystemEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SystemEvent) Reset() {
	*x = SystemEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_notifications_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemEvent) ProtoMessage() {}

func (x *SystemEvent) ProtoReflect() protoreflect.Message {
	mi := &file_resources_notifications_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemEvent.ProtoReflect.Descriptor instead.
func (*SystemEvent) Descriptor() ([]byte, []int) {
	return file_resources_notifications_events_proto_rawDescGZIP(), []int{2}
}

var File_resources_notifications_events_proto protoreflect.FileDescriptor

var file_resources_notifications_events_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x2b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x6a, 0x6f,
	0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x4c, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x38,
	0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x48, 0x00, 0x52, 0x08,
	0x6a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42,
	0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69,
	0x76, 0x65, 0x6e, 0x65, 0x74, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65,
	0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_notifications_events_proto_rawDescOnce sync.Once
	file_resources_notifications_events_proto_rawDescData = file_resources_notifications_events_proto_rawDesc
)

func file_resources_notifications_events_proto_rawDescGZIP() []byte {
	file_resources_notifications_events_proto_rawDescOnce.Do(func() {
		file_resources_notifications_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_notifications_events_proto_rawDescData)
	})
	return file_resources_notifications_events_proto_rawDescData
}

var file_resources_notifications_events_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_resources_notifications_events_proto_goTypes = []interface{}{
	(*UserEvent)(nil),      // 0: resources.notifications.UserEvent
	(*JobEvent)(nil),       // 1: resources.notifications.JobEvent
	(*SystemEvent)(nil),    // 2: resources.notifications.SystemEvent
	(*Notification)(nil),   // 3: resources.notifications.Notification
	(*users.JobProps)(nil), // 4: resources.users.JobProps
}
var file_resources_notifications_events_proto_depIdxs = []int32{
	3, // 0: resources.notifications.UserEvent.notification:type_name -> resources.notifications.Notification
	4, // 1: resources.notifications.JobEvent.job_props:type_name -> resources.users.JobProps
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_resources_notifications_events_proto_init() }
func file_resources_notifications_events_proto_init() {
	if File_resources_notifications_events_proto != nil {
		return
	}
	file_resources_notifications_notifications_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_resources_notifications_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEvent); i {
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
		file_resources_notifications_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobEvent); i {
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
		file_resources_notifications_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemEvent); i {
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
	file_resources_notifications_events_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UserEvent_Notification)(nil),
		(*UserEvent_RefreshToken)(nil),
	}
	file_resources_notifications_events_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*JobEvent_JobProps)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_notifications_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_notifications_events_proto_goTypes,
		DependencyIndexes: file_resources_notifications_events_proto_depIdxs,
		MessageInfos:      file_resources_notifications_events_proto_msgTypes,
	}.Build()
	File_resources_notifications_events_proto = out.File
	file_resources_notifications_events_proto_rawDesc = nil
	file_resources_notifications_events_proto_goTypes = nil
	file_resources_notifications_events_proto_depIdxs = nil
}
