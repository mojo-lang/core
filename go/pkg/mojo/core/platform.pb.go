// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/core/platform.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Arch int32

const (
	Arch_ARCH_UNSPECIFIED Arch = 0
	Arch_ARCH_X86         Arch = 1
	Arch_ARCH_AMD64       Arch = 2
	Arch_ARCH_ARM         Arch = 5
	Arch_ARCH_ARM64       Arch = 6
	Arch_ARCH_WASM        Arch = 10
)

var Arch_name = map[int32]string{
	0:  "ARCH_UNSPECIFIED",
	1:  "ARCH_X86",
	2:  "ARCH_AMD64",
	5:  "ARCH_ARM",
	6:  "ARCH_ARM64",
	10: "ARCH_WASM",
}

var Arch_value = map[string]int32{
	"ARCH_UNSPECIFIED": 0,
	"ARCH_X86":         1,
	"ARCH_AMD64":       2,
	"ARCH_ARM":         5,
	"ARCH_ARM64":       6,
	"ARCH_WASM":        10,
}

func (x Arch) String() string {
	return proto.EnumName(Arch_name, int32(x))
}

func (Arch) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_85c22e47075a042f, []int{0}
}

type OS int32

const (
	OS_OS_UNSPECIFIED OS = 0
	OS_OS_ANDROID     OS = 1
	OS_OS_DARWIN      OS = 2
	OS_OS_IOS         OS = 6
	OS_OS_LINUX       OS = 10
	OS_OS_WINDOWS     OS = 20
)

var OS_name = map[int32]string{
	0:  "OS_UNSPECIFIED",
	1:  "OS_ANDROID",
	2:  "OS_DARWIN",
	6:  "OS_IOS",
	10: "OS_LINUX",
	20: "OS_WINDOWS",
}

var OS_value = map[string]int32{
	"OS_UNSPECIFIED": 0,
	"OS_ANDROID":     1,
	"OS_DARWIN":      2,
	"OS_IOS":         6,
	"OS_LINUX":       10,
	"OS_WINDOWS":     20,
}

func (x OS) String() string {
	return proto.EnumName(OS_name, int32(x))
}

func (OS) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_85c22e47075a042f, []int{1}
}

type Platform struct {
	Arch                 Arch     `protobuf:"varint,1,opt,name=arch,proto3,enum=mojo.core.Arch" json:"arch,omitempty"`
	Os                   OS       `protobuf:"varint,2,opt,name=os,proto3,enum=mojo.core.OS" json:"os,omitempty"`
	OsVersion            string   `protobuf:"bytes,3,opt,name=os_version,json=osVersion,proto3" json:"osVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" gorm:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" gorm:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" gorm:"-" xml:"-" bson:"-"`
}

func (m *Platform) Reset()         { *m = Platform{} }
func (m *Platform) String() string { return proto.CompactTextString(m) }
func (*Platform) ProtoMessage()    {}
func (*Platform) Descriptor() ([]byte, []int) {
	return fileDescriptor_85c22e47075a042f, []int{0}
}
func (m *Platform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform.Unmarshal(m, b)
}
func (m *Platform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform.Marshal(b, m, deterministic)
}
func (m *Platform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform.Merge(m, src)
}
func (m *Platform) XXX_Size() int {
	return xxx_messageInfo_Platform.Size(m)
}
func (m *Platform) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform.DiscardUnknown(m)
}

var xxx_messageInfo_Platform proto.InternalMessageInfo

func (m *Platform) GetArch() Arch {
	if m != nil {
		return m.Arch
	}
	return Arch_ARCH_UNSPECIFIED
}

func (m *Platform) GetOs() OS {
	if m != nil {
		return m.Os
	}
	return OS_OS_UNSPECIFIED
}

func (m *Platform) GetOsVersion() string {
	if m != nil {
		return m.OsVersion
	}
	return ""
}

func init() {
	proto.RegisterEnum("mojo.core.Arch", Arch_name, Arch_value)
	proto.RegisterEnum("mojo.core.OS", OS_name, OS_value)
	proto.RegisterType((*Platform)(nil), "mojo.core.Platform")
}

func init() { proto.RegisterFile("mojo/core/platform.proto", fileDescriptor_85c22e47075a042f) }

var fileDescriptor_85c22e47075a042f = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0xae, 0x93, 0x40,
	0x14, 0x86, 0x2f, 0x78, 0x25, 0x97, 0x93, 0x16, 0x27, 0x63, 0xa3, 0xac, 0x6c, 0xe3, 0xaa, 0x69,
	0x14, 0x12, 0x6d, 0xd0, 0xc4, 0x15, 0x94, 0x1a, 0x49, 0x84, 0x21, 0x4c, 0x2a, 0x8d, 0x0b, 0x09,
	0x25, 0x95, 0x56, 0x4b, 0x0f, 0x01, 0x34, 0xf1, 0x4d, 0x7c, 0x08, 0x9f, 0xa9, 0x1b, 0x77, 0x7d,
	0x0a, 0x33, 0xb4, 0xd2, 0xe6, 0x6e, 0x26, 0x39, 0xff, 0xff, 0x7f, 0x39, 0x93, 0xf3, 0x83, 0x5e,
	0xe0, 0x37, 0x34, 0x33, 0xac, 0xd6, 0x66, 0xb9, 0x4b, 0x9b, 0xaf, 0x58, 0x15, 0x46, 0x59, 0x61,
	0x83, 0x54, 0x15, 0x8e, 0x21, 0x9c, 0xe7, 0x7f, 0x24, 0xb8, 0x0b, 0xcf, 0x2e, 0x7d, 0x03, 0xb7,
	0x69, 0x95, 0x6d, 0x74, 0x69, 0x24, 0x8d, 0xb5, 0x57, 0x8f, 0x8c, 0x2e, 0x66, 0xd8, 0x55, 0xb6,
	0x71, 0xe8, 0xf1, 0x30, 0xd4, 0x44, 0xe0, 0x05, 0x16, 0xdb, 0x66, 0x5d, 0x94, 0xcd, 0xaf, 0xa8,
	0x05, 0xa8, 0x09, 0x32, 0xd6, 0xba, 0xdc, 0x62, 0xfd, 0x2b, 0x8c, 0x71, 0x87, 0x1c, 0x0f, 0xc3,
	0x1e, 0xd6, 0x57, 0x88, 0x8c, 0x35, 0xb5, 0x00, 0xb0, 0x4e, 0x7e, 0xae, 0xab, 0x7a, 0x8b, 0x7b,
	0xfd, 0xc1, 0x48, 0x1a, 0xab, 0xce, 0xd3, 0xe3, 0x61, 0xf8, 0x18, 0xeb, 0x4f, 0x27, 0xf1, 0x0a,
	0x50, 0x3b, 0x71, 0x92, 0xc3, 0xad, 0xf8, 0x0a, 0x1d, 0x00, 0xb1, 0xa3, 0xd9, 0x87, 0x64, 0x11,
	0xf0, 0x70, 0x3e, 0xf3, 0xde, 0x7b, 0x73, 0x97, 0xdc, 0xd0, 0x1e, 0xdc, 0xb5, 0xea, 0xf2, 0xad,
	0x45, 0x24, 0xaa, 0x01, 0xb4, 0x93, 0xed, 0xbb, 0xd6, 0x94, 0xc8, 0x9d, 0x6b, 0x47, 0x3e, 0x79,
	0x78, 0x71, 0x23, 0xdf, 0x9a, 0x12, 0x85, 0xf6, 0x41, 0x6d, 0xe7, 0xd8, 0xe6, 0x3e, 0x81, 0x49,
	0x0a, 0x32, 0xe3, 0x94, 0x82, 0xc6, 0xf8, 0xbd, 0x25, 0x1a, 0x00, 0xe3, 0x89, 0x1d, 0xb8, 0x11,
	0xf3, 0x5c, 0x22, 0x09, 0x90, 0xf1, 0xc4, 0xb5, 0xa3, 0xd8, 0x0b, 0x88, 0x4c, 0x01, 0x14, 0xc6,
	0x13, 0x8f, 0x71, 0xa2, 0x88, 0x8d, 0x8c, 0x27, 0x1f, 0xbd, 0x60, 0xb1, 0x24, 0x70, 0x06, 0x63,
	0x2f, 0x70, 0x59, 0xcc, 0xc9, 0xc0, 0xf9, 0xf2, 0xfb, 0xef, 0xb3, 0x1b, 0x78, 0x82, 0x55, 0xde,
	0x5e, 0x6c, 0x97, 0xee, 0xf3, 0xcb, 0xe9, 0x9c, 0xfe, 0xff, 0x56, 0x42, 0x51, 0x59, 0x28, 0x7d,
	0x36, 0xf3, 0x6d, 0xb3, 0xf9, 0xb1, 0x32, 0x32, 0x2c, 0x4c, 0x11, 0x7b, 0x29, 0x80, 0x53, 0xbd,
	0x39, 0x9a, 0xe5, 0xf7, 0xdc, 0xec, 0xfa, 0x7e, 0x27, 0x9e, 0x95, 0xd2, 0x96, 0xfd, 0xfa, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0x66, 0x25, 0x4b, 0x08, 0x02, 0x00, 0x00,
}
