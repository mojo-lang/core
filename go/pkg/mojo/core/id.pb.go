// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: mojo/core/id.proto

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

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Id:
	//
	//	*Id_Uint64Val
	//	*Id_StringVal
	//	*Id_UuidVal
	Id isId_Id `protobuf_oneof:"id" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_id_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_id_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_mojo_core_id_proto_rawDescGZIP(), []int{0}
}

func (m *Id) GetId() isId_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (x *Id) GetUint64Val() uint64 {
	if x, ok := x.GetId().(*Id_Uint64Val); ok {
		return x.Uint64Val
	}
	return 0
}

func (x *Id) GetStringVal() string {
	if x, ok := x.GetId().(*Id_StringVal); ok {
		return x.StringVal
	}
	return ""
}

func (x *Id) GetUuidVal() *Uuid {
	if x, ok := x.GetId().(*Id_UuidVal); ok {
		return x.UuidVal
	}
	return nil
}

type isId_Id interface {
	isId_Id()
}

type Id_Uint64Val struct {
	Uint64Val uint64 `protobuf:"varint,1,opt,name=uint64_val,json=uint64Val,proto3,oneof" json:"uint64Val,omitempty"`
}

type Id_StringVal struct {
	StringVal string `protobuf:"bytes,2,opt,name=string_val,json=stringVal,proto3,oneof" json:"stringVal,omitempty"`
}

type Id_UuidVal struct {
	UuidVal *Uuid `protobuf:"bytes,3,opt,name=uuid_val,json=uuidVal,proto3,oneof" json:"uuidVal,omitempty"`
}

func (*Id_Uint64Val) isId_Id() {}

func (*Id_StringVal) isId_Id() {}

func (*Id_UuidVal) isId_Id() {}

var File_mojo_core_id_proto protoreflect.FileDescriptor

var file_mojo_core_id_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a,
	0x14, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0a, 0x75,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48,
	0x00, 0x52, 0x09, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0a,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x12, 0x2c, 0x0a,
	0x08, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x75, 0x69, 0x64,
	0x48, 0x00, 0x52, 0x07, 0x75, 0x75, 0x69, 0x64, 0x56, 0x61, 0x6c, 0x42, 0x04, 0x0a, 0x02, 0x69,
	0x64, 0x42, 0x54, 0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x07, 0x49, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_id_proto_rawDescOnce sync.Once
	file_mojo_core_id_proto_rawDescData = file_mojo_core_id_proto_rawDesc
)

func file_mojo_core_id_proto_rawDescGZIP() []byte {
	file_mojo_core_id_proto_rawDescOnce.Do(func() {
		file_mojo_core_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_id_proto_rawDescData)
	})
	return file_mojo_core_id_proto_rawDescData
}

var file_mojo_core_id_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mojo_core_id_proto_goTypes = []interface{}{
	(*Id)(nil),   // 0: mojo.core.Id
	(*Uuid)(nil), // 1: mojo.core.Uuid
}
var file_mojo_core_id_proto_depIdxs = []int32{
	1, // 0: mojo.core.Id.uuid_val:type_name -> mojo.core.Uuid
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mojo_core_id_proto_init() }
func file_mojo_core_id_proto_init() {
	if File_mojo_core_id_proto != nil {
		return
	}
	file_mojo_core_uuid_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mojo_core_id_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
	file_mojo_core_id_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Id_Uint64Val)(nil),
		(*Id_StringVal)(nil),
		(*Id_UuidVal)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mojo_core_id_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_id_proto_goTypes,
		DependencyIndexes: file_mojo_core_id_proto_depIdxs,
		MessageInfos:      file_mojo_core_id_proto_msgTypes,
	}.Build()
	File_mojo_core_id_proto = out.File
	file_mojo_core_id_proto_rawDesc = nil
	file_mojo_core_id_proto_goTypes = nil
	file_mojo_core_id_proto_depIdxs = nil
}
