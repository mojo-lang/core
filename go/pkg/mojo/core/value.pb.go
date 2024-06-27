// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: mojo/core/value.proto

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

type ValueKind int32

const (
	ValueKind_VALUE_KIND_UNSPECIFIED ValueKind = 0
	ValueKind_VALUE_KIND_NULL        ValueKind = 1
	ValueKind_VALUE_KIND_BOOLEAN     ValueKind = 2
	ValueKind_VALUE_KIND_INTEGER     ValueKind = 3
	ValueKind_VALUE_KIND_NUMBER      ValueKind = 4
	ValueKind_VALUE_KIND_STRING      ValueKind = 5
	ValueKind_VALUE_KIND_BYTES       ValueKind = 6
	ValueKind_VALUE_KIND_ARRAY       ValueKind = 10
	ValueKind_VALUE_KIND_OBJECT      ValueKind = 11
)

// Enum value maps for ValueKind.
var (
	ValueKind_name = map[int32]string{
		0:  "VALUE_KIND_UNSPECIFIED",
		1:  "VALUE_KIND_NULL",
		2:  "VALUE_KIND_BOOLEAN",
		3:  "VALUE_KIND_INTEGER",
		4:  "VALUE_KIND_NUMBER",
		5:  "VALUE_KIND_STRING",
		6:  "VALUE_KIND_BYTES",
		10: "VALUE_KIND_ARRAY",
		11: "VALUE_KIND_OBJECT",
	}
	ValueKind_value = map[string]int32{
		"VALUE_KIND_UNSPECIFIED": 0,
		"VALUE_KIND_NULL":        1,
		"VALUE_KIND_BOOLEAN":     2,
		"VALUE_KIND_INTEGER":     3,
		"VALUE_KIND_NUMBER":      4,
		"VALUE_KIND_STRING":      5,
		"VALUE_KIND_BYTES":       6,
		"VALUE_KIND_ARRAY":       10,
		"VALUE_KIND_OBJECT":      11,
	}
)

func (x ValueKind) Enum() *ValueKind {
	p := new(ValueKind)
	*p = x
	return p
}

func (x ValueKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValueKind) Descriptor() protoreflect.EnumDescriptor {
	return file_mojo_core_value_proto_enumTypes[0].Descriptor()
}

func (ValueKind) Type() protoreflect.EnumType {
	return &file_mojo_core_value_proto_enumTypes[0]
}

func (x ValueKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValueKind.Descriptor instead.
func (ValueKind) EnumDescriptor() ([]byte, []int) {
	return file_mojo_core_value_proto_rawDescGZIP(), []int{0}
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals map[string]*Value `protobuf:"bytes,1,rep,name=vals,proto3" json:"vals,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_mojo_core_value_proto_rawDescGZIP(), []int{0}
}

func (x *Object) GetVals() map[string]*Value {
	if x != nil {
		return x.Vals
	}
	return nil
}

type Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals []*Value `protobuf:"bytes,1,rep,name=vals,proto3" json:"vals,omitempty"`
}

func (x *Values) Reset() {
	*x = Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_value_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Values) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Values) ProtoMessage() {}

func (x *Values) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_value_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Values.ProtoReflect.Descriptor instead.
func (*Values) Descriptor() ([]byte, []int) {
	return file_mojo_core_value_proto_rawDescGZIP(), []int{1}
}

func (x *Values) GetVals() []*Value {
	if x != nil {
		return x.Vals
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Val:
	//
	//	*Value_ValuesVal
	//	*Value_ObjectVal
	//	*Value_NullVal
	//	*Value_BoolVal
	//	*Value_PositiveVal
	//	*Value_NegativeVal
	//	*Value_DoubleVal
	//	*Value_StringVal
	//	*Value_BytesVal
	Val isValue_Val `protobuf_oneof:"val" json:"val,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_core_value_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_core_value_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_mojo_core_value_proto_rawDescGZIP(), []int{2}
}

func (m *Value) GetVal() isValue_Val {
	if m != nil {
		return m.Val
	}
	return nil
}

func (x *Value) GetValuesVal() *Values {
	if x, ok := x.GetVal().(*Value_ValuesVal); ok {
		return x.ValuesVal
	}
	return nil
}

func (x *Value) GetObjectVal() *Object {
	if x, ok := x.GetVal().(*Value_ObjectVal); ok {
		return x.ObjectVal
	}
	return nil
}

func (x *Value) GetNullVal() *Null {
	if x, ok := x.GetVal().(*Value_NullVal); ok {
		return x.NullVal
	}
	return nil
}

func (x *Value) GetBoolVal() bool {
	if x, ok := x.GetVal().(*Value_BoolVal); ok {
		return x.BoolVal
	}
	return false
}

func (x *Value) GetPositiveVal() uint64 {
	if x, ok := x.GetVal().(*Value_PositiveVal); ok {
		return x.PositiveVal
	}
	return 0
}

func (x *Value) GetNegativeVal() uint64 {
	if x, ok := x.GetVal().(*Value_NegativeVal); ok {
		return x.NegativeVal
	}
	return 0
}

func (x *Value) GetDoubleVal() float64 {
	if x, ok := x.GetVal().(*Value_DoubleVal); ok {
		return x.DoubleVal
	}
	return 0
}

func (x *Value) GetStringVal() string {
	if x, ok := x.GetVal().(*Value_StringVal); ok {
		return x.StringVal
	}
	return ""
}

func (x *Value) GetBytesVal() []byte {
	if x, ok := x.GetVal().(*Value_BytesVal); ok {
		return x.BytesVal
	}
	return nil
}

type isValue_Val interface {
	isValue_Val()
}

type Value_ValuesVal struct {
	ValuesVal *Values `protobuf:"bytes,1,opt,name=values_val,json=valuesVal,proto3,oneof" json:"valuesVal,omitempty"`
}

type Value_ObjectVal struct {
	ObjectVal *Object `protobuf:"bytes,2,opt,name=object_val,json=objectVal,proto3,oneof" json:"objectVal,omitempty"`
}

type Value_NullVal struct {
	NullVal *Null `protobuf:"bytes,3,opt,name=null_val,json=nullVal,proto3,oneof" json:"nullVal,omitempty"`
}

type Value_BoolVal struct {
	BoolVal bool `protobuf:"varint,4,opt,name=bool_val,json=boolVal,proto3,oneof" json:"boolVal,omitempty"`
}

type Value_PositiveVal struct {
	PositiveVal uint64 `protobuf:"varint,5,opt,name=positive_val,json=positiveVal,proto3,oneof" json:"positiveVal,omitempty"`
}

type Value_NegativeVal struct {
	NegativeVal uint64 `protobuf:"varint,6,opt,name=negative_val,json=negativeVal,proto3,oneof" json:"negativeVal,omitempty"`
}

type Value_DoubleVal struct {
	DoubleVal float64 `protobuf:"fixed64,7,opt,name=double_val,json=doubleVal,proto3,oneof" json:"doubleVal,omitempty"`
}

type Value_StringVal struct {
	StringVal string `protobuf:"bytes,8,opt,name=string_val,json=stringVal,proto3,oneof" json:"stringVal,omitempty"`
}

type Value_BytesVal struct {
	BytesVal []byte `protobuf:"bytes,9,opt,name=bytes_val,json=bytesVal,proto3,oneof" json:"bytesVal,omitempty"`
}

func (*Value_ValuesVal) isValue_Val() {}

func (*Value_ObjectVal) isValue_Val() {}

func (*Value_NullVal) isValue_Val() {}

func (*Value_BoolVal) isValue_Val() {}

func (*Value_PositiveVal) isValue_Val() {}

func (*Value_NegativeVal) isValue_Val() {}

func (*Value_DoubleVal) isValue_Val() {}

func (*Value_StringVal) isValue_Val() {}

func (*Value_BytesVal) isValue_Val() {}

var File_mojo_core_value_proto protoreflect.FileDescriptor

var file_mojo_core_value_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x1a, 0x14, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6e, 0x75,
	0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x06, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x76, 0x61, 0x6c, 0x73, 0x1a, 0x49, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x2e, 0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x76, 0x61, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x22,
	0xec, 0x02, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x48, 0x00, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x12, 0x32, 0x0a,
	0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x61,
	0x6c, 0x12, 0x2c, 0x0a, 0x08, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x4e, 0x75, 0x6c, 0x6c, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x12,
	0x1b, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x00, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0c,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x56, 0x61,
	0x6c, 0x12, 0x23, 0x0a, 0x0c, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x76, 0x61,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0b, 0x6e, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0a, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x5f, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x09, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x08, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x42, 0x05, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x2a, 0xdd,
	0x01, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x16,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x16, 0x0a,
	0x12, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x42, 0x4f, 0x4f, 0x4c,
	0x45, 0x41, 0x4e, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4b,
	0x49, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x10, 0x03, 0x12, 0x15, 0x0a,
	0x11, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x4e, 0x55, 0x4d, 0x42,
	0x45, 0x52, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4b, 0x49,
	0x4e, 0x44, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10,
	0x06, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f,
	0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x0a, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x0b, 0x42, 0x57,
	0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6d,
	0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_core_value_proto_rawDescOnce sync.Once
	file_mojo_core_value_proto_rawDescData = file_mojo_core_value_proto_rawDesc
)

func file_mojo_core_value_proto_rawDescGZIP() []byte {
	file_mojo_core_value_proto_rawDescOnce.Do(func() {
		file_mojo_core_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_core_value_proto_rawDescData)
	})
	return file_mojo_core_value_proto_rawDescData
}

var file_mojo_core_value_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mojo_core_value_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mojo_core_value_proto_goTypes = []interface{}{
	(ValueKind)(0), // 0: mojo.core.ValueKind
	(*Object)(nil), // 1: mojo.core.Object
	(*Values)(nil), // 2: mojo.core.Values
	(*Value)(nil),  // 3: mojo.core.Value
	nil,            // 4: mojo.core.Object.ValsEntry
	(*Null)(nil),   // 5: mojo.core.Null
}
var file_mojo_core_value_proto_depIdxs = []int32{
	4, // 0: mojo.core.Object.vals:type_name -> mojo.core.Object.ValsEntry
	3, // 1: mojo.core.Values.vals:type_name -> mojo.core.Value
	2, // 2: mojo.core.Value.values_val:type_name -> mojo.core.Values
	1, // 3: mojo.core.Value.object_val:type_name -> mojo.core.Object
	5, // 4: mojo.core.Value.null_val:type_name -> mojo.core.Null
	3, // 5: mojo.core.Object.ValsEntry.value:type_name -> mojo.core.Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_mojo_core_value_proto_init() }
func file_mojo_core_value_proto_init() {
	if File_mojo_core_value_proto != nil {
		return
	}
	file_mojo_core_null_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mojo_core_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_mojo_core_value_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Values); i {
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
		file_mojo_core_value_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
	file_mojo_core_value_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Value_ValuesVal)(nil),
		(*Value_ObjectVal)(nil),
		(*Value_NullVal)(nil),
		(*Value_BoolVal)(nil),
		(*Value_PositiveVal)(nil),
		(*Value_NegativeVal)(nil),
		(*Value_DoubleVal)(nil),
		(*Value_StringVal)(nil),
		(*Value_BytesVal)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mojo_core_value_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mojo_core_value_proto_goTypes,
		DependencyIndexes: file_mojo_core_value_proto_depIdxs,
		EnumInfos:         file_mojo_core_value_proto_enumTypes,
		MessageInfos:      file_mojo_core_value_proto_msgTypes,
	}.Build()
	File_mojo_core_value_proto = out.File
	file_mojo_core_value_proto_rawDesc = nil
	file_mojo_core_value_proto_goTypes = nil
	file_mojo_core_value_proto_depIdxs = nil
}
