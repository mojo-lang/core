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
	Arch_ARCH_AMD64       Arch = 1
	Arch_ARCH_ARM         Arch = 5
	Arch_ARCH_ARM64       Arch = 6
	Arch_ARCH_WASM        Arch = 10
)

var Arch_name = map[int32]string{
	0:  "ARCH_UNSPECIFIED",
	1:  "ARCH_AMD64",
	5:  "ARCH_ARM",
	6:  "ARCH_ARM64",
	10: "ARCH_WASM",
}

var Arch_value = map[string]int32{
	"ARCH_UNSPECIFIED": 0,
	"ARCH_AMD64":       1,
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
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0xae, 0xd2, 0x40,
	0x14, 0x86, 0x69, 0xc5, 0x86, 0x9e, 0x40, 0x9d, 0x8c, 0x44, 0xbb, 0x12, 0xe2, 0x8a, 0x10, 0x6d,
	0x13, 0x25, 0xb8, 0x70, 0xd5, 0x52, 0x8c, 0x4d, 0x6c, 0xa7, 0xe9, 0x04, 0xab, 0x2e, 0x6c, 0x4a,
	0x83, 0x05, 0xa5, 0x9c, 0xa6, 0xad, 0x26, 0xbe, 0xc9, 0x7d, 0x88, 0xfb, 0x4c, 0x6c, 0xee, 0x8e,
	0xa7, 0xb8, 0x99, 0xc2, 0x05, 0x92, 0xbb, 0x99, 0xe4, 0xfc, 0xff, 0xff, 0xcd, 0x4c, 0xce, 0x0f,
	0x7a, 0x8e, 0xbf, 0xd1, 0x4c, 0xb1, 0x5c, 0x99, 0xc5, 0x36, 0xa9, 0x7f, 0x61, 0x99, 0x1b, 0x45,
	0x89, 0x35, 0x52, 0x55, 0x38, 0x86, 0x70, 0x5e, 0xdf, 0x4a, 0xd0, 0x09, 0x4e, 0x2e, 0xfd, 0x00,
	0xed, 0xa4, 0x4c, 0xd7, 0xba, 0x34, 0x94, 0x46, 0xda, 0xbb, 0x67, 0xc6, 0x39, 0x66, 0x58, 0x65,
	0xba, 0xb6, 0xe9, 0x61, 0x3f, 0xd0, 0x44, 0xe0, 0x0d, 0xe6, 0x9b, 0x7a, 0x95, 0x17, 0xf5, 0xff,
	0xb0, 0x01, 0xa8, 0x09, 0x32, 0x56, 0xba, 0xdc, 0x60, 0xbd, 0x2b, 0x8c, 0x71, 0x9b, 0x1c, 0xf6,
	0x83, 0x2e, 0x56, 0x57, 0x88, 0x8c, 0x15, 0x9d, 0x02, 0x60, 0x15, 0xff, 0x5b, 0x95, 0xd5, 0x06,
	0x77, 0xfa, 0x93, 0xa1, 0x34, 0x52, 0xed, 0x97, 0x87, 0xfd, 0xe0, 0x39, 0x56, 0x5f, 0x8f, 0xe2,
	0x15, 0xa0, 0x9e, 0xc5, 0xf1, 0x77, 0x68, 0x8b, 0xaf, 0xd0, 0x3e, 0x10, 0x2b, 0x9c, 0x7d, 0x8e,
	0x17, 0x3e, 0x0f, 0xe6, 0x33, 0xf7, 0x93, 0x3b, 0x77, 0x48, 0x8b, 0x6a, 0x00, 0x8d, 0x6a, 0x79,
	0xce, 0x74, 0x42, 0x24, 0xda, 0x85, 0xce, 0x71, 0x0e, 0x3d, 0xf2, 0xf4, 0xe2, 0x86, 0xde, 0x74,
	0x42, 0x14, 0xda, 0x03, 0xb5, 0x99, 0x23, 0x8b, 0x7b, 0x04, 0xc6, 0x09, 0xc8, 0x8c, 0x53, 0x0a,
	0x1a, 0xe3, 0x8f, 0xaf, 0x65, 0x3c, 0xb6, 0x7c, 0x27, 0x64, 0xae, 0x43, 0x24, 0x01, 0x32, 0x1e,
	0x3b, 0x56, 0x18, 0xb9, 0x3e, 0x91, 0x29, 0x80, 0xc2, 0x78, 0xec, 0x32, 0x4e, 0x14, 0xf1, 0x22,
	0xe3, 0xf1, 0x17, 0xd7, 0x5f, 0x7c, 0x23, 0x70, 0x02, 0x23, 0xd7, 0x77, 0x58, 0xc4, 0x49, 0xdf,
	0xfe, 0x79, 0x73, 0xf7, 0xaa, 0x05, 0x2f, 0xb0, 0xcc, 0x9a, 0x1d, 0x6d, 0x93, 0x5d, 0x76, 0x59,
	0x96, 0xdd, 0x7b, 0xe8, 0x21, 0x10, 0x25, 0x05, 0xd2, 0x0f, 0x33, 0xdb, 0xd4, 0xeb, 0xbf, 0x4b,
	0x23, 0xc5, 0xdc, 0x14, 0xb1, 0xb7, 0x02, 0x38, 0x16, 0x9a, 0xa1, 0x59, 0xfc, 0xc9, 0xcc, 0x73,
	0xc3, 0x1f, 0xc5, 0xb1, 0x54, 0x9a, 0x7a, 0xdf, 0xdf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x27, 0xf4,
	0xfe, 0x5c, 0xfa, 0x01, 0x00, 0x00,
}