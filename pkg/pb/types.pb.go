// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: types.proto

package pb

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

type SizeUnit int32

const (
	SizeUnit_UNKNOWN  SizeUnit = 0
	SizeUnit_BIT      SizeUnit = 1
	SizeUnit_BYTE     SizeUnit = 2
	SizeUnit_KILOBYTE SizeUnit = 3
	SizeUnit_MEGABYTE SizeUnit = 4
	SizeUnit_GIGABYTE SizeUnit = 5
	SizeUnit_TERABYTE SizeUnit = 6
	SizeUnit_PETABYTE SizeUnit = 7
)

// Enum value maps for SizeUnit.
var (
	SizeUnit_name = map[int32]string{
		0: "UNKNOWN",
		1: "BIT",
		2: "BYTE",
		3: "KILOBYTE",
		4: "MEGABYTE",
		5: "GIGABYTE",
		6: "TERABYTE",
		7: "PETABYTE",
	}
	SizeUnit_value = map[string]int32{
		"UNKNOWN":  0,
		"BIT":      1,
		"BYTE":     2,
		"KILOBYTE": 3,
		"MEGABYTE": 4,
		"GIGABYTE": 5,
		"TERABYTE": 6,
		"PETABYTE": 7,
	}
)

func (x SizeUnit) Enum() *SizeUnit {
	p := new(SizeUnit)
	*p = x
	return p
}

func (x SizeUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SizeUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_types_proto_enumTypes[0].Descriptor()
}

func (SizeUnit) Type() protoreflect.EnumType {
	return &file_types_proto_enumTypes[0]
}

func (x SizeUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SizeUnit.Descriptor instead.
func (SizeUnit) EnumDescriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x2a, 0x70, 0x0a, 0x08, 0x53, 0x69, 0x7a, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x49,
	0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x59, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x4b, 0x49, 0x4c, 0x4f, 0x42, 0x59, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4d,
	0x45, 0x47, 0x41, 0x42, 0x59, 0x54, 0x45, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x49, 0x47,
	0x41, 0x42, 0x59, 0x54, 0x45, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x45, 0x52, 0x41, 0x42,
	0x59, 0x54, 0x45, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x45, 0x54, 0x41, 0x42, 0x59, 0x54,
	0x45, 0x10, 0x07, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x71, 0x75, 0x6f, 0x63, 0x62, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_proto_goTypes = []interface{}{
	(SizeUnit)(0), // 0: pb.SizeUnit
}
var file_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		EnumInfos:         file_types_proto_enumTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}