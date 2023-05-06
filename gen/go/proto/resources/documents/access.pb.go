// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: resources/documents/access.proto

package documents

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

type ACCESS_LEVEL int32

const (
	ACCESS_LEVEL_BLOCKED ACCESS_LEVEL = 0
	ACCESS_LEVEL_VIEW    ACCESS_LEVEL = 1
	ACCESS_LEVEL_COMMENT ACCESS_LEVEL = 2
	ACCESS_LEVEL_ACCESS  ACCESS_LEVEL = 3
	ACCESS_LEVEL_EDIT    ACCESS_LEVEL = 4
)

// Enum value maps for ACCESS_LEVEL.
var (
	ACCESS_LEVEL_name = map[int32]string{
		0: "BLOCKED",
		1: "VIEW",
		2: "COMMENT",
		3: "ACCESS",
		4: "EDIT",
	}
	ACCESS_LEVEL_value = map[string]int32{
		"BLOCKED": 0,
		"VIEW":    1,
		"COMMENT": 2,
		"ACCESS":  3,
		"EDIT":    4,
	}
)

func (x ACCESS_LEVEL) Enum() *ACCESS_LEVEL {
	p := new(ACCESS_LEVEL)
	*p = x
	return p
}

func (x ACCESS_LEVEL) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ACCESS_LEVEL) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_documents_access_proto_enumTypes[0].Descriptor()
}

func (ACCESS_LEVEL) Type() protoreflect.EnumType {
	return &file_resources_documents_access_proto_enumTypes[0]
}

func (x ACCESS_LEVEL) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ACCESS_LEVEL.Descriptor instead.
func (ACCESS_LEVEL) EnumDescriptor() ([]byte, []int) {
	return file_resources_documents_access_proto_rawDescGZIP(), []int{0}
}

var File_resources_documents_access_proto protoreflect.FileDescriptor

var file_resources_documents_access_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2a, 0x48, 0x0a, 0x0c, 0x41, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4c, 0x4f, 0x43, 0x4b,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x49, 0x45, 0x57, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x41,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x44, 0x49, 0x54, 0x10,
	0x04, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x61, 0x6c, 0x65, 0x78, 0x72, 0x74, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x3b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_resources_documents_access_proto_rawDescOnce sync.Once
	file_resources_documents_access_proto_rawDescData = file_resources_documents_access_proto_rawDesc
)

func file_resources_documents_access_proto_rawDescGZIP() []byte {
	file_resources_documents_access_proto_rawDescOnce.Do(func() {
		file_resources_documents_access_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_documents_access_proto_rawDescData)
	})
	return file_resources_documents_access_proto_rawDescData
}

var file_resources_documents_access_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resources_documents_access_proto_goTypes = []interface{}{
	(ACCESS_LEVEL)(0), // 0: resources.documents.ACCESS_LEVEL
}
var file_resources_documents_access_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resources_documents_access_proto_init() }
func file_resources_documents_access_proto_init() {
	if File_resources_documents_access_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_documents_access_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_documents_access_proto_goTypes,
		DependencyIndexes: file_resources_documents_access_proto_depIdxs,
		EnumInfos:         file_resources_documents_access_proto_enumTypes,
	}.Build()
	File_resources_documents_access_proto = out.File
	file_resources_documents_access_proto_rawDesc = nil
	file_resources_documents_access_proto_goTypes = nil
	file_resources_documents_access_proto_depIdxs = nil
}
