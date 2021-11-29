// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/core/value.proto

package core

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Object struct {
	Vals                 map[string]*Value `protobuf:"bytes,1,rep,name=vals,proto3" json:"vals,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c44d46343c3cc16, []int{0}
}
func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetVals() map[string]*Value {
	if m != nil {
		return m.Vals
	}
	return nil
}

type Values struct {
	Vals                 []*Value `protobuf:"bytes,1,rep,name=vals,proto3" json:"vals,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Values) Reset()         { *m = Values{} }
func (m *Values) String() string { return proto.CompactTextString(m) }
func (*Values) ProtoMessage()    {}
func (*Values) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c44d46343c3cc16, []int{1}
}
func (m *Values) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Values.Unmarshal(m, b)
}
func (m *Values) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Values.Marshal(b, m, deterministic)
}
func (m *Values) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Values.Merge(m, src)
}
func (m *Values) XXX_Size() int {
	return xxx_messageInfo_Values.Size(m)
}
func (m *Values) XXX_DiscardUnknown() {
	xxx_messageInfo_Values.DiscardUnknown(m)
}

var xxx_messageInfo_Values proto.InternalMessageInfo

func (m *Values) GetVals() []*Value {
	if m != nil {
		return m.Vals
	}
	return nil
}

type Value struct {
	// Types that are valid to be assigned to Value:
	//	*Value_ValuesVal
	//	*Value_ObjectVal
	//	*Value_BoolVal
	//	*Value_Int64Val
	//	*Value_DoubleVal
	//	*Value_StringVal
	Value                isValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c44d46343c3cc16, []int{2}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Value interface {
	isValue_Value()
}

type Value_ValuesVal struct {
	ValuesVal *Values `protobuf:"bytes,1,opt,name=values_val,json=valuesVal,proto3,oneof" json:"valuesVal,omitempty"`
}
type Value_ObjectVal struct {
	ObjectVal *Object `protobuf:"bytes,2,opt,name=object_val,json=objectVal,proto3,oneof" json:"objectVal,omitempty"`
}
type Value_BoolVal struct {
	BoolVal bool `protobuf:"varint,3,opt,name=bool_val,json=boolVal,proto3,oneof" json:"boolVal,omitempty"`
}
type Value_Int64Val struct {
	Int64Val int64 `protobuf:"varint,4,opt,name=int64_val,json=int64Val,proto3,oneof" json:"int64Val,omitempty"`
}
type Value_DoubleVal struct {
	DoubleVal float64 `protobuf:"fixed64,5,opt,name=double_val,json=doubleVal,proto3,oneof" json:"doubleVal,omitempty"`
}
type Value_StringVal struct {
	StringVal string `protobuf:"bytes,7,opt,name=string_val,json=stringVal,proto3,oneof" json:"stringVal,omitempty"`
}

func (*Value_ValuesVal) isValue_Value() {}
func (*Value_ObjectVal) isValue_Value() {}
func (*Value_BoolVal) isValue_Value()   {}
func (*Value_Int64Val) isValue_Value()  {}
func (*Value_DoubleVal) isValue_Value() {}
func (*Value_StringVal) isValue_Value() {}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Value) GetValuesVal() *Values {
	if x, ok := m.GetValue().(*Value_ValuesVal); ok {
		return x.ValuesVal
	}
	return nil
}

func (m *Value) GetObjectVal() *Object {
	if x, ok := m.GetValue().(*Value_ObjectVal); ok {
		return x.ObjectVal
	}
	return nil
}

func (m *Value) GetBoolVal() bool {
	if x, ok := m.GetValue().(*Value_BoolVal); ok {
		return x.BoolVal
	}
	return false
}

func (m *Value) GetInt64Val() int64 {
	if x, ok := m.GetValue().(*Value_Int64Val); ok {
		return x.Int64Val
	}
	return 0
}

func (m *Value) GetDoubleVal() float64 {
	if x, ok := m.GetValue().(*Value_DoubleVal); ok {
		return x.DoubleVal
	}
	return 0
}

func (m *Value) GetStringVal() string {
	if x, ok := m.GetValue().(*Value_StringVal); ok {
		return x.StringVal
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Value) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Value_ValuesVal)(nil),
		(*Value_ObjectVal)(nil),
		(*Value_BoolVal)(nil),
		(*Value_Int64Val)(nil),
		(*Value_DoubleVal)(nil),
		(*Value_StringVal)(nil),
	}
}

func init() {
	proto.RegisterType((*Object)(nil), "mojo.core.Object")
	proto.RegisterMapType((map[string]*Value)(nil), "mojo.core.Object.ValsEntry")
	proto.RegisterType((*Values)(nil), "mojo.core.Values")
	proto.RegisterType((*Value)(nil), "mojo.core.Value")
}

func init() { proto.RegisterFile("mojo/core/value.proto", fileDescriptor_5c44d46343c3cc16) }

var fileDescriptor_5c44d46343c3cc16 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0xce, 0xd2, 0x40,
	0x14, 0x86, 0x1d, 0xfa, 0xf1, 0xd3, 0x43, 0xfc, 0x61, 0x08, 0x48, 0x74, 0xd1, 0x06, 0x37, 0x5d,
	0x68, 0x9b, 0xd4, 0x9f, 0x10, 0xd9, 0x8d, 0x31, 0x31, 0x71, 0x21, 0x61, 0x81, 0x89, 0x1b, 0xd3,
	0x62, 0x53, 0x0b, 0xd3, 0x0e, 0xb6, 0x03, 0x49, 0x2f, 0xc1, 0xbb, 0xf2, 0x4a, 0x7a, 0x01, 0xbd,
	0x0a, 0x33, 0x67, 0xb0, 0x20, 0xcd, 0xb7, 0x61, 0xf1, 0xf0, 0x3e, 0x6f, 0x4f, 0x4f, 0x0f, 0x4c,
	0x52, 0xb1, 0x13, 0xde, 0x56, 0xe4, 0x91, 0x77, 0x0a, 0xf8, 0x31, 0x72, 0x0f, 0xb9, 0x90, 0x82,
	0x9a, 0x0a, 0xbb, 0x0a, 0xcf, 0xff, 0x10, 0xe8, 0x7d, 0x09, 0x77, 0xd1, 0x56, 0xd2, 0x0f, 0x70,
	0x77, 0x0a, 0x78, 0x31, 0x23, 0xb6, 0xe1, 0x0c, 0xfd, 0xe7, 0x6e, 0x13, 0x72, 0x75, 0xc0, 0xdd,
	0x04, 0xbc, 0xf8, 0x98, 0xc9, 0xbc, 0x64, 0xb4, 0xae, 0xac, 0x47, 0x2a, 0xfc, 0x52, 0xa4, 0x89,
	0x8c, 0xd2, 0x83, 0x2c, 0xd7, 0x28, 0x3f, 0xfb, 0x05, 0x66, 0x13, 0xa3, 0x2f, 0xc0, 0xd8, 0x47,
	0xe5, 0x8c, 0xd8, 0xc4, 0x31, 0xd9, 0xa8, 0xae, 0xac, 0x87, 0xfb, 0xa8, 0xbc, 0x52, 0xd4, 0xbf,
	0x74, 0x09, 0x5d, 0x9c, 0x6d, 0xd6, 0xb1, 0x89, 0x33, 0xf4, 0x9f, 0x5c, 0x3d, 0x77, 0xa3, 0x38,
	0x1b, 0xd7, 0x95, 0xf5, 0x18, 0x23, 0x57, 0xaa, 0x76, 0xde, 0x77, 0x16, 0x64, 0xce, 0xa0, 0x87,
	0xc1, 0x82, 0x2e, 0xfe, 0x7b, 0x83, 0x76, 0xd3, 0xbd, 0x63, 0xcf, 0x7f, 0x1b, 0xd0, 0xc5, 0x0c,
	0xfd, 0x0c, 0x80, 0xd5, 0xc5, 0xf7, 0x53, 0xc0, 0x71, 0xf4, 0xa1, 0x3f, 0xba, 0x6d, 0x2a, 0xd8,
	0xd3, 0xba, 0xb2, 0xc6, 0x3a, 0xb8, 0x09, 0xf8, 0xa5, 0xef, 0xd3, 0x83, 0xb5, 0xd9, 0x60, 0x55,
	0x26, 0x70, 0x77, 0x58, 0xd6, 0x69, 0x95, 0xe9, 0xc5, 0xea, 0x32, 0x1d, 0x6c, 0x95, 0x35, 0x98,
	0xfa, 0x30, 0x08, 0x85, 0xe0, 0x58, 0x65, 0xd8, 0xc4, 0x19, 0xb0, 0x49, 0x5d, 0x59, 0x23, 0xc5,
	0x6e, 0xad, 0xfe, 0x19, 0xd2, 0xb7, 0x60, 0x26, 0x99, 0x7c, 0xf7, 0x06, 0xa5, 0x3b, 0x9b, 0x38,
	0x06, 0x9b, 0xd6, 0x95, 0x45, 0x11, 0xde, 0x5a, 0x83, 0x7f, 0x94, 0x2e, 0x00, 0x7e, 0x88, 0x63,
	0xc8, 0x23, 0xf4, 0xba, 0x36, 0x71, 0x88, 0x1e, 0x52, 0xd3, 0xd6, 0x90, 0x0d, 0x56, 0x66, 0x21,
	0xf3, 0x24, 0x8b, 0xd1, 0xec, 0xe3, 0x97, 0x47, 0x53, 0xd3, 0x96, 0xd9, 0x60, 0xd6, 0x3f, 0xdf,
	0x01, 0xfb, 0x0a, 0x53, 0x91, 0xc7, 0xb8, 0x25, 0x1e, 0x64, 0xf1, 0x65, 0x5d, 0x0c, 0x70, 0xf9,
	0x2b, 0x75, 0xc3, 0x2b, 0xf2, 0xcd, 0x8b, 0x13, 0xf9, 0xf3, 0x18, 0xba, 0x5b, 0x91, 0x7a, 0x2a,
	0xf3, 0x4a, 0xa5, 0xf5, 0xb1, 0xc7, 0xc2, 0x3b, 0xec, 0x63, 0xaf, 0xb9, 0xfe, 0xa5, 0xfa, 0x09,
	0x7b, 0x78, 0xfd, 0xaf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x99, 0x6d, 0xa2, 0xa1, 0x16, 0x03,
	0x00, 0x00,
}
