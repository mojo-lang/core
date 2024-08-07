// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: mojo/core/template_string.proto

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

type TemplateString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Segments []*TemplateString_Segment `protobuf:"bytes,1,rep,name=segments,proto3" json:"segments,omitempty"`
}

func (x *TemplateString) Reset() {
	*x = TemplateString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_template_string_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateString) ProtoMessage() {}

func (x *TemplateString) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_template_string_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateString.ProtoReflect.Descriptor instead.
func (*TemplateString) Descriptor() ([]byte, []int) {
	return file_mojo_core_template_string_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateString) GetSegments() []*TemplateString_Segment {
	if x != nil {
		return x.Segments
	}
	return nil
}

type TemplateString_Segment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content   string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Templated bool   `protobuf:"varint,2,opt,name=templated,proto3" json:"templated,omitempty"`
}

func (x *TemplateString_Segment) Reset() {
	*x = TemplateString_Segment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_template_string_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateString_Segment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateString_Segment) ProtoMessage() {}

func (x *TemplateString_Segment) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_template_string_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateString_Segment.ProtoReflect.Descriptor instead.
func (*TemplateString_Segment) Descriptor() ([]byte, []int) {
	return file_mojo_core_template_string_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TemplateString_Segment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TemplateString_Segment) GetTemplated() bool {
	if x != nil {
		return x.Templated
	}
	return false
}

var File_mojo_core_template_string_proto protoreflect.FileDescriptor

var file_mojo_core_template_string_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x92, 0x01, 0x0a,
	0x0e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x3d, 0x0a, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x41,
	0x0a, 0x07, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x60, 0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x13, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b, 0x63,
	0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_template_string_proto_rawDescOnce sync.Once
	file_mojo_core_template_string_proto_rawDescData = file_mojo_core_template_string_proto_rawDesc
)

func file_mojo_core_template_string_proto_rawDescGZIP() []byte {
	file_mojo_core_template_string_proto_rawDescOnce.Do(func() {
		file_mojo_core_template_string_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_template_string_proto_rawDescData)
	})
	return file_mojo_core_template_string_proto_rawDescData
}

var file_mojo_core_template_string_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mojo_core_template_string_proto_goTypes = []interface{}{
	(*TemplateString)(nil),         // 0: mojo.core.TemplateString
	(*TemplateString_Segment)(nil), // 1: mojo.core.TemplateString.Segment
}
var file_mojo_core_template_string_proto_depIdxs = []int32{
	1, // 0: mojo.core.TemplateString.segments:type_name -> mojo.core.TemplateString.Segment
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mojo_core_template_string_proto_init() }
func file_mojo_core_template_string_proto_init() {
	if File_mojo_core_template_string_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mojo_core_template_string_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateString); i {
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
		file_mojo_core_template_string_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateString_Segment); i {
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
			RawDescriptor: file_mojo_core_template_string_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_template_string_proto_goTypes,
		DependencyIndexes: file_mojo_core_template_string_proto_depIdxs,
		MessageInfos:      file_mojo_core_template_string_proto_msgTypes,
	}.Build()
	File_mojo_core_template_string_proto = out.File
	file_mojo_core_template_string_proto_rawDesc = nil
	file_mojo_core_template_string_proto_goTypes = nil
	file_mojo_core_template_string_proto_depIdxs = nil
}
