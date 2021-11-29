// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/core/pixel.proto

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

type Pixel struct {
	Val                  uint64   `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pixel) Reset()         { *m = Pixel{} }
func (m *Pixel) String() string { return proto.CompactTextString(m) }
func (*Pixel) ProtoMessage()    {}
func (*Pixel) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8431557c25dd588, []int{0}
}
func (m *Pixel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pixel.Unmarshal(m, b)
}
func (m *Pixel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pixel.Marshal(b, m, deterministic)
}
func (m *Pixel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pixel.Merge(m, src)
}
func (m *Pixel) XXX_Size() int {
	return xxx_messageInfo_Pixel.Size(m)
}
func (m *Pixel) XXX_DiscardUnknown() {
	xxx_messageInfo_Pixel.DiscardUnknown(m)
}

var xxx_messageInfo_Pixel proto.InternalMessageInfo

func (m *Pixel) GetVal() uint64 {
	if m != nil {
		return m.Val
	}
	return 0
}

func init() {
	proto.RegisterType((*Pixel)(nil), "mojo.core.Pixel")
}

func init() { proto.RegisterFile("mojo/core/pixel.proto", fileDescriptor_f8431557c25dd588) }

var fileDescriptor_f8431557c25dd588 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xcd, 0xcf, 0xca,
	0xd7, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0xc8, 0xac, 0x48, 0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x04, 0x09, 0xeb, 0x81, 0x84, 0x95, 0x74, 0xb8, 0x58, 0x03, 0x40, 0x32, 0x42,
	0xca, 0x5c, 0xcc, 0x65, 0x89, 0x39, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x4e, 0x82, 0xaf, 0xee,
	0xc9, 0xf3, 0x96, 0x25, 0xe6, 0xe8, 0xe4, 0xe7, 0x66, 0x96, 0xa4, 0xe6, 0x16, 0x94, 0x54, 0x06,
	0x81, 0x64, 0x9d, 0xc2, 0xb9, 0xc4, 0xf2, 0x8b, 0xd2, 0xf5, 0x40, 0xda, 0x73, 0x12, 0xf3, 0x20,
	0x0c, 0xb0, 0x39, 0x4e, 0x5c, 0x60, 0x53, 0x02, 0x40, 0xc6, 0x07, 0x30, 0x46, 0xe9, 0xa7, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x83, 0xd4, 0xe8, 0x82, 0x54, 0x43, 0xdc,
	0x91, 0x9e, 0xaf, 0x5f, 0x90, 0x9d, 0xae, 0x0f, 0x77, 0x98, 0x35, 0x88, 0x48, 0x62, 0x03, 0x3b,
	0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x65, 0xf5, 0xff, 0xb1, 0x00, 0x00, 0x00,
}
