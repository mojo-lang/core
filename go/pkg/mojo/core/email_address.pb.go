// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: mojo/core/email_address.proto

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

type EmailAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalPart string  `protobuf:"bytes,1,opt,name=local_part,json=localPart,proto3" json:"localPart,omitempty"`
	Domain    *Domain `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *EmailAddress) Reset() {
	*x = EmailAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_email_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailAddress) ProtoMessage() {}

func (x *EmailAddress) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_email_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailAddress.ProtoReflect.Descriptor instead.
func (*EmailAddress) Descriptor() ([]byte, []int) {
	return file_mojo_core_email_address_proto_rawDescGZIP(), []int{0}
}

func (x *EmailAddress) GetLocalPart() string {
	if x != nil {
		return x.LocalPart
	}
	return ""
}

func (x *EmailAddress) GetDomain() *Domain {
	if x != nil {
		return x.Domain
	}
	return nil
}

var File_mojo_core_email_address_proto protoreflect.FileDescriptor

var file_mojo_core_email_address_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x16, 0x6d, 0x6f, 0x6a, 0x6f,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61, 0x72,
	0x74, 0x12, 0x29, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x5e, 0x0a, 0x16,
	0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x6a,
	0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x11, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f,
	0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_email_address_proto_rawDescOnce sync.Once
	file_mojo_core_email_address_proto_rawDescData = file_mojo_core_email_address_proto_rawDesc
)

func file_mojo_core_email_address_proto_rawDescGZIP() []byte {
	file_mojo_core_email_address_proto_rawDescOnce.Do(func() {
		file_mojo_core_email_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_email_address_proto_rawDescData)
	})
	return file_mojo_core_email_address_proto_rawDescData
}

var file_mojo_core_email_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mojo_core_email_address_proto_goTypes = []interface{}{
	(*EmailAddress)(nil), // 0: mojo.core.EmailAddress
	(*Domain)(nil),       // 1: mojo.core.Domain
}
var file_mojo_core_email_address_proto_depIdxs = []int32{
	1, // 0: mojo.core.EmailAddress.domain:type_name -> mojo.core.Domain
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mojo_core_email_address_proto_init() }
func file_mojo_core_email_address_proto_init() {
	if File_mojo_core_email_address_proto != nil {
		return
	}
	file_mojo_core_domain_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mojo_core_email_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailAddress); i {
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
			RawDescriptor: file_mojo_core_email_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_email_address_proto_goTypes,
		DependencyIndexes: file_mojo_core_email_address_proto_depIdxs,
		MessageInfos:      file_mojo_core_email_address_proto_msgTypes,
	}.Build()
	File_mojo_core_email_address_proto = out.File
	file_mojo_core_email_address_proto_rawDesc = nil
	file_mojo_core_email_address_proto_goTypes = nil
	file_mojo_core_email_address_proto_depIdxs = nil
}
