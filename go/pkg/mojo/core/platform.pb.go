// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: mojo/core/platform.proto

package core

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

type Architecture int32

const (
	Architecture_ARCHITECTURE_UNSPECIFIED Architecture = 0
	Architecture_ARCHITECTURE_X86         Architecture = 1
	Architecture_ARCHITECTURE_AMD64       Architecture = 2
	Architecture_ARCHITECTURE_ARM         Architecture = 5
	Architecture_ARCHITECTURE_ARM64       Architecture = 6
	Architecture_ARCHITECTURE_WASM        Architecture = 10
)

// Enum value maps for Architecture.
var (
	Architecture_name = map[int32]string{
		0:  "ARCHITECTURE_UNSPECIFIED",
		1:  "ARCHITECTURE_X86",
		2:  "ARCHITECTURE_AMD64",
		5:  "ARCHITECTURE_ARM",
		6:  "ARCHITECTURE_ARM64",
		10: "ARCHITECTURE_WASM",
	}
	Architecture_value = map[string]int32{
		"ARCHITECTURE_UNSPECIFIED": 0,
		"ARCHITECTURE_X86":         1,
		"ARCHITECTURE_AMD64":       2,
		"ARCHITECTURE_ARM":         5,
		"ARCHITECTURE_ARM64":       6,
		"ARCHITECTURE_WASM":        10,
	}
)

func (x Architecture) Enum() *Architecture {
	p := new(Architecture)
	*p = x
	return p
}

func (x Architecture) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Architecture) Descriptor() protoreflect.EnumDescriptor {
	return file_mojo_core_platform_proto_enumTypes[0].Descriptor()
}

func (Architecture) Type() protoreflect.EnumType {
	return &file_mojo_core_platform_proto_enumTypes[0]
}

func (x Architecture) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Architecture.Descriptor instead.
func (Architecture) EnumDescriptor() ([]byte, []int) {
	return file_mojo_core_platform_proto_rawDescGZIP(), []int{0}
}

type OS int32

const (
	OS_OS_UNSPECIFIED OS = 0
	OS_OS_ANDROID     OS = 1
	OS_OS_DARWIN      OS = 2
	OS_OS_IOS         OS = 6
	OS_OS_LINUX       OS = 10
	OS_OS_WINDOWS     OS = 20
	OS_OS_SIMULATION  OS = 30
)

// Enum value maps for OS.
var (
	OS_name = map[int32]string{
		0:  "OS_UNSPECIFIED",
		1:  "OS_ANDROID",
		2:  "OS_DARWIN",
		6:  "OS_IOS",
		10: "OS_LINUX",
		20: "OS_WINDOWS",
		30: "OS_SIMULATION",
	}
	OS_value = map[string]int32{
		"OS_UNSPECIFIED": 0,
		"OS_ANDROID":     1,
		"OS_DARWIN":      2,
		"OS_IOS":         6,
		"OS_LINUX":       10,
		"OS_WINDOWS":     20,
		"OS_SIMULATION":  30,
	}
)

func (x OS) Enum() *OS {
	p := new(OS)
	*p = x
	return p
}

func (x OS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OS) Descriptor() protoreflect.EnumDescriptor {
	return file_mojo_core_platform_proto_enumTypes[1].Descriptor()
}

func (OS) Type() protoreflect.EnumType {
	return &file_mojo_core_platform_proto_enumTypes[1]
}

func (x OS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OS.Descriptor instead.
func (OS) EnumDescriptor() ([]byte, []int) {
	return file_mojo_core_platform_proto_rawDescGZIP(), []int{1}
}

type Platform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Architecture Architecture `protobuf:"varint,1,opt,name=architecture,proto3,enum=mojo.core.Architecture" json:"architecture,omitempty"`
	Variant      string       `protobuf:"bytes,2,opt,name=variant,proto3" json:"variant,omitempty"`
	Os           OS           `protobuf:"varint,10,opt,name=os,proto3,enum=mojo.core.OS" json:"os,omitempty"`
	OsName       string       `protobuf:"bytes,11,opt,name=os_name,json=osName,proto3" json:"osName,omitempty"`
	OsVersion    string       `protobuf:"bytes,12,opt,name=os_version,json=osVersion,proto3" json:"osVersion,omitempty"`
}

func (x *Platform) Reset() {
	*x = Platform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform) ProtoMessage() {}

func (x *Platform) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform.ProtoReflect.Descriptor instead.
func (*Platform) Descriptor() ([]byte, []int) {
	return file_mojo_core_platform_proto_rawDescGZIP(), []int{0}
}

func (x *Platform) GetArchitecture() Architecture {
	if x != nil {
		return x.Architecture
	}
	return Architecture_ARCHITECTURE_UNSPECIFIED
}

func (x *Platform) GetVariant() string {
	if x != nil {
		return x.Variant
	}
	return ""
}

func (x *Platform) GetOs() OS {
	if x != nil {
		return x.Os
	}
	return OS_OS_UNSPECIFIED
}

func (x *Platform) GetOsName() string {
	if x != nil {
		return x.OsName
	}
	return ""
}

func (x *Platform) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

var File_mojo_core_platform_proto protoreflect.FileDescriptor

var file_mojo_core_platform_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x6a, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x3b, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x02, 0x6f, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x4f, 0x53, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x73, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2a, 0x9f, 0x01, 0x0a, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x52, 0x43, 0x48, 0x49, 0x54, 0x45, 0x43, 0x54, 0x55, 0x52,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x41, 0x52, 0x43, 0x48, 0x49, 0x54, 0x45, 0x43, 0x54, 0x55, 0x52, 0x45, 0x5f,
	0x58, 0x38, 0x36, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x52, 0x43, 0x48, 0x49, 0x54, 0x45,
	0x43, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x41, 0x4d, 0x44, 0x36, 0x34, 0x10, 0x02, 0x12, 0x14, 0x0a,
	0x10, 0x41, 0x52, 0x43, 0x48, 0x49, 0x54, 0x45, 0x43, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x41, 0x52,
	0x4d, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x52, 0x43, 0x48, 0x49, 0x54, 0x45, 0x43, 0x54,
	0x55, 0x52, 0x45, 0x5f, 0x41, 0x52, 0x4d, 0x36, 0x34, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x41,
	0x52, 0x43, 0x48, 0x49, 0x54, 0x45, 0x43, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x57, 0x41, 0x53, 0x4d,
	0x10, 0x0a, 0x2a, 0x74, 0x0a, 0x02, 0x4f, 0x53, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x4f, 0x53, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x4f, 0x53, 0x5f, 0x44, 0x41, 0x52, 0x57, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4f,
	0x53, 0x5f, 0x49, 0x4f, 0x53, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x53, 0x5f, 0x4c, 0x49,
	0x4e, 0x55, 0x58, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x53, 0x5f, 0x57, 0x49, 0x4e, 0x44,
	0x4f, 0x57, 0x53, 0x10, 0x14, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x53, 0x5f, 0x53, 0x49, 0x4d, 0x55,
	0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x1e, 0x42, 0x5a, 0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e,
	0x6d, 0x6f, 0x6a, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x42, 0x0d, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b,
	0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_platform_proto_rawDescOnce sync.Once
	file_mojo_core_platform_proto_rawDescData = file_mojo_core_platform_proto_rawDesc
)

func file_mojo_core_platform_proto_rawDescGZIP() []byte {
	file_mojo_core_platform_proto_rawDescOnce.Do(func() {
		file_mojo_core_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_platform_proto_rawDescData)
	})
	return file_mojo_core_platform_proto_rawDescData
}

var file_mojo_core_platform_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mojo_core_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mojo_core_platform_proto_goTypes = []interface{}{
	(Architecture)(0), // 0: mojo.core.Architecture
	(OS)(0),           // 1: mojo.core.OS
	(*Platform)(nil),  // 2: mojo.core.Platform
}
var file_mojo_core_platform_proto_depIdxs = []int32{
	0, // 0: mojo.core.Platform.architecture:type_name -> mojo.core.Architecture
	1, // 1: mojo.core.Platform.os:type_name -> mojo.core.OS
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mojo_core_platform_proto_init() }
func file_mojo_core_platform_proto_init() {
	if File_mojo_core_platform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mojo_core_platform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platform); i {
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
			RawDescriptor: file_mojo_core_platform_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_platform_proto_goTypes,
		DependencyIndexes: file_mojo_core_platform_proto_depIdxs,
		EnumInfos:         file_mojo_core_platform_proto_enumTypes,
		MessageInfos:      file_mojo_core_platform_proto_msgTypes,
	}.Build()
	File_mojo_core_platform_proto = out.File
	file_mojo_core_platform_proto_rawDesc = nil
	file_mojo_core_platform_proto_goTypes = nil
	file_mojo_core_platform_proto_depIdxs = nil
}
