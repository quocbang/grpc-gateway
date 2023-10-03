// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: role.proto

package roles

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

type Roles int32

const (
	Roles_UNSPECIFIED_USER Roles = 0
	Roles_USER             Roles = 1
	Roles_LEADER           Roles = 2
	Roles_ADMIN            Roles = 3
)

// Enum value maps for Roles.
var (
	Roles_name = map[int32]string{
		0: "UNSPECIFIED_USER",
		1: "USER",
		2: "LEADER",
		3: "ADMIN",
	}
	Roles_value = map[string]int32{
		"UNSPECIFIED_USER": 0,
		"USER":             1,
		"LEADER":           2,
		"ADMIN":            3,
	}
)

func (x Roles) Enum() *Roles {
	p := new(Roles)
	*p = x
	return p
}

func (x Roles) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Roles) Descriptor() protoreflect.EnumDescriptor {
	return file_role_proto_enumTypes[0].Descriptor()
}

func (Roles) Type() protoreflect.EnumType {
	return &file_role_proto_enumTypes[0]
}

func (x Roles) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Roles.Descriptor instead.
func (Roles) EnumDescriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{0}
}

var File_role_proto protoreflect.FileDescriptor

var file_role_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x2a, 0x3e, 0x0a, 0x05, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x10,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x4c, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49,
	0x4e, 0x10, 0x03, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x71, 0x75, 0x6f, 0x63, 0x62, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_proto_rawDescOnce sync.Once
	file_role_proto_rawDescData = file_role_proto_rawDesc
)

func file_role_proto_rawDescGZIP() []byte {
	file_role_proto_rawDescOnce.Do(func() {
		file_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_proto_rawDescData)
	})
	return file_role_proto_rawDescData
}

var file_role_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_role_proto_goTypes = []interface{}{
	(Roles)(0), // 0: roles.Roles
}
var file_role_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_role_proto_init() }
func file_role_proto_init() {
	if File_role_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_role_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_role_proto_goTypes,
		DependencyIndexes: file_role_proto_depIdxs,
		EnumInfos:         file_role_proto_enumTypes,
	}.Build()
	File_role_proto = out.File
	file_role_proto_rawDesc = nil
	file_role_proto_goTypes = nil
	file_role_proto_depIdxs = nil
}
