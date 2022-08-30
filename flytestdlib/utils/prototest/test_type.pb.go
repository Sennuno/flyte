// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: test_type.proto

package prototest

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
}

func (x *TestProto) Reset() {
	*x = TestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestProto) ProtoMessage() {}

func (x *TestProto) ProtoReflect() protoreflect.Message {
	mi := &file_test_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestProto.ProtoReflect.Descriptor instead.
func (*TestProto) Descriptor() ([]byte, []int) {
	return file_test_type_proto_rawDescGZIP(), []int{0}
}

func (x *TestProto) GetStringValue() string {
	if x != nil {
		return x.StringValue
	}
	return ""
}

var File_test_type_proto protoreflect.FileDescriptor

var file_test_type_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x22, 0x2e, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x79, 0x66, 0x74, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x73, 0x74, 0x64, 0x6c, 0x69, 0x62, 0x2f,
	0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x65, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_type_proto_rawDescOnce sync.Once
	file_test_type_proto_rawDescData = file_test_type_proto_rawDesc
)

func file_test_type_proto_rawDescGZIP() []byte {
	file_test_type_proto_rawDescOnce.Do(func() {
		file_test_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_type_proto_rawDescData)
	})
	return file_test_type_proto_rawDescData
}

var file_test_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_test_type_proto_goTypes = []interface{}{
	(*TestProto)(nil), // 0: flyteidl.core.TestProto
}
var file_test_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_type_proto_init() }
func file_test_type_proto_init() {
	if File_test_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestProto); i {
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
			RawDescriptor: file_test_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_type_proto_goTypes,
		DependencyIndexes: file_test_type_proto_depIdxs,
		MessageInfos:      file_test_type_proto_msgTypes,
	}.Build()
	File_test_type_proto = out.File
	file_test_type_proto_rawDesc = nil
	file_test_type_proto_goTypes = nil
	file_test_type_proto_depIdxs = nil
}
