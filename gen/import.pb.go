// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: import.proto

package gen

import (
	validate "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_import_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*validate.Int32Rules)(nil),
		ExtensionType: ([]int32)(nil),
		Field:         1800,
		Name:          "custom.abs_not_in",
		Tag:           "varint,1800,rep,name=abs_not_in",
		Filename:      "import.proto",
	},
}

// Extension fields to validate.Int32Rules.
var (
	// repeated int32 abs_not_in = 1800;
	E_AbsNotIn = &file_import_proto_extTypes[0]
)

var File_import_proto protoreflect.FileDescriptor

var file_import_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3a, 0xa2, 0x01, 0x0a, 0x0a, 0x61, 0x62, 0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x5f,
	0x69, 0x6e, 0x12, 0x18, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x88, 0x0e, 0x20,
	0x03, 0x28, 0x05, 0x42, 0x69, 0xc2, 0x48, 0x66, 0x0a, 0x64, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x2e, 0x61, 0x62, 0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x12, 0x27, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x69, 0x6e, 0x20, 0x61,
	0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x6f, 0x66,
	0x20, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x27, 0x74, 0x68, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x72,
	0x75, 0x6c, 0x65, 0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x72,
	0x75, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x28, 0x6e, 0x2c, 0x20, 0x2d, 0x6e, 0x29, 0x52, 0x08,
	0x61, 0x62, 0x73, 0x4e, 0x6f, 0x74, 0x49, 0x6e, 0x42, 0x61, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x42, 0x0b, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0e, 0x70, 0x76, 0x2f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x2f, 0x67, 0x65, 0x6e, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0xca, 0x02, 0x06, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0xe2, 0x02, 0x12,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x06, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
}

var file_import_proto_goTypes = []any{
	(*validate.Int32Rules)(nil), // 0: buf.validate.Int32Rules
}
var file_import_proto_depIdxs = []int32{
	0, // 0: custom.abs_not_in:extendee -> buf.validate.Int32Rules
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_import_proto_init() }
func file_import_proto_init() {
	if File_import_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_import_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_import_proto_goTypes,
		DependencyIndexes: file_import_proto_depIdxs,
		ExtensionInfos:    file_import_proto_extTypes,
	}.Build()
	File_import_proto = out.File
	file_import_proto_rawDesc = nil
	file_import_proto_goTypes = nil
	file_import_proto_depIdxs = nil
}
