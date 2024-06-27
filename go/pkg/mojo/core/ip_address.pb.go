// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: mojo/core/ip_address.proto

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

type IPAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to IpAddress:
	//
	//	*IPAddress_Ipv4
	//	*IPAddress_Ipv6
	IpAddress isIPAddress_IpAddress `protobuf_oneof:"ip_address" json:"ipAddress,omitempty"`
}

func (x *IPAddress) Reset() {
	*x = IPAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_ip_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPAddress) ProtoMessage() {}

func (x *IPAddress) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_ip_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPAddress.ProtoReflect.Descriptor instead.
func (*IPAddress) Descriptor() ([]byte, []int) {
	return file_mojo_core_ip_address_proto_rawDescGZIP(), []int{0}
}

func (m *IPAddress) GetIpAddress() isIPAddress_IpAddress {
	if m != nil {
		return m.IpAddress
	}
	return nil
}

func (x *IPAddress) GetIpv4() *IPv4 {
	if x, ok := x.GetIpAddress().(*IPAddress_Ipv4); ok {
		return x.Ipv4
	}
	return nil
}

func (x *IPAddress) GetIpv6() *IPv6 {
	if x, ok := x.GetIpAddress().(*IPAddress_Ipv6); ok {
		return x.Ipv6
	}
	return nil
}

type isIPAddress_IpAddress interface {
	isIPAddress_IpAddress()
}

type IPAddress_Ipv4 struct {
	Ipv4 *IPv4 `protobuf:"bytes,1,opt,name=ipv4,proto3,oneof" json:"ipv4,omitempty"`
}

type IPAddress_Ipv6 struct {
	Ipv6 *IPv6 `protobuf:"bytes,2,opt,name=ipv6,proto3,oneof" json:"ipv6,omitempty"`
}

func (*IPAddress_Ipv4) isIPAddress_IpAddress() {}

func (*IPAddress_Ipv6) isIPAddress_IpAddress() {}

type IPv4 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val uint32 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *IPv4) Reset() {
	*x = IPv4{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_ip_address_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPv4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPv4) ProtoMessage() {}

func (x *IPv4) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_ip_address_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPv4.ProtoReflect.Descriptor instead.
func (*IPv4) Descriptor() ([]byte, []int) {
	return file_mojo_core_ip_address_proto_rawDescGZIP(), []int{1}
}

func (x *IPv4) GetVal() uint32 {
	if x != nil {
		return x.Val
	}
	return 0
}

type IPv6 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val uint64 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *IPv6) Reset() {
	*x = IPv6{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_ip_address_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPv6) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPv6) ProtoMessage() {}

func (x *IPv6) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_ip_address_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPv6.ProtoReflect.Descriptor instead.
func (*IPv6) Descriptor() ([]byte, []int) {
	return file_mojo_core_ip_address_proto_rawDescGZIP(), []int{2}
}

func (x *IPv6) GetVal() uint64 {
	if x != nil {
		return x.Val
	}
	return 0
}

var File_mojo_core_ip_address_proto protoreflect.FileDescriptor

var file_mojo_core_ip_address_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x70, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f,
	0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x67, 0x0a, 0x09, 0x49, 0x50, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49,
	0x50, 0x76, 0x34, 0x48, 0x00, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x25, 0x0a, 0x04, 0x69,
	0x70, 0x76, 0x36, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x50, 0x76, 0x36, 0x48, 0x00, 0x52, 0x04, 0x69, 0x70,
	0x76, 0x36, 0x42, 0x0c, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x18, 0x0a, 0x04, 0x49, 0x50, 0x76, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x18, 0x0a, 0x04, 0x49, 0x50,
	0x76, 0x36, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x76, 0x61, 0x6c, 0x42, 0x5b, 0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x0e,
	0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6a,
	0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b, 0x63, 0x6f, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_ip_address_proto_rawDescOnce sync.Once
	file_mojo_core_ip_address_proto_rawDescData = file_mojo_core_ip_address_proto_rawDesc
)

func file_mojo_core_ip_address_proto_rawDescGZIP() []byte {
	file_mojo_core_ip_address_proto_rawDescOnce.Do(func() {
		file_mojo_core_ip_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_ip_address_proto_rawDescData)
	})
	return file_mojo_core_ip_address_proto_rawDescData
}

var file_mojo_core_ip_address_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mojo_core_ip_address_proto_goTypes = []interface{}{
	(*IPAddress)(nil), // 0: mojo.core.IPAddress
	(*IPv4)(nil),      // 1: mojo.core.IPv4
	(*IPv6)(nil),      // 2: mojo.core.IPv6
}
var file_mojo_core_ip_address_proto_depIdxs = []int32{
	1, // 0: mojo.core.IPAddress.ipv4:type_name -> mojo.core.IPv4
	2, // 1: mojo.core.IPAddress.ipv6:type_name -> mojo.core.IPv6
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mojo_core_ip_address_proto_init() }
func file_mojo_core_ip_address_proto_init() {
	if File_mojo_core_ip_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mojo_core_ip_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPAddress); i {
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
		file_mojo_core_ip_address_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPv4); i {
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
		file_mojo_core_ip_address_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPv6); i {
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
	file_mojo_core_ip_address_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*IPAddress_Ipv4)(nil),
		(*IPAddress_Ipv6)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mojo_core_ip_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_ip_address_proto_goTypes,
		DependencyIndexes: file_mojo_core_ip_address_proto_depIdxs,
		MessageInfos:      file_mojo_core_ip_address_proto_msgTypes,
	}.Build()
	File_mojo_core_ip_address_proto = out.File
	file_mojo_core_ip_address_proto_rawDesc = nil
	file_mojo_core_ip_address_proto_goTypes = nil
	file_mojo_core_ip_address_proto_depIdxs = nil
}
