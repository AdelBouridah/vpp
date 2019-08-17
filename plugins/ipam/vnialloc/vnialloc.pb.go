// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vnialloc.proto

package vnialloc

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

// CustomVxlanVniAllocation represents vni allocation made for a custom vxlan purpose (like SFC vxlan)
type CustomVxlanVniAllocation struct {
	VxlanName            string   `protobuf:"bytes,1,opt,name=vxlan_name,json=vxlanName,proto3" json:"vxlan_name,omitempty"`
	Vni                  uint32   `protobuf:"varint,2,opt,name=vni,proto3" json:"vni,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomVxlanVniAllocation) Reset()         { *m = CustomVxlanVniAllocation{} }
func (m *CustomVxlanVniAllocation) String() string { return proto.CompactTextString(m) }
func (*CustomVxlanVniAllocation) ProtoMessage()    {}
func (*CustomVxlanVniAllocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_71dc0d993d483727, []int{0}
}

func (m *CustomVxlanVniAllocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomVxlanVniAllocation.Unmarshal(m, b)
}
func (m *CustomVxlanVniAllocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomVxlanVniAllocation.Marshal(b, m, deterministic)
}
func (m *CustomVxlanVniAllocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomVxlanVniAllocation.Merge(m, src)
}
func (m *CustomVxlanVniAllocation) XXX_Size() int {
	return xxx_messageInfo_CustomVxlanVniAllocation.Size(m)
}
func (m *CustomVxlanVniAllocation) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomVxlanVniAllocation.DiscardUnknown(m)
}

var xxx_messageInfo_CustomVxlanVniAllocation proto.InternalMessageInfo

func (m *CustomVxlanVniAllocation) GetVxlanName() string {
	if m != nil {
		return m.VxlanName
	}
	return ""
}

func (m *CustomVxlanVniAllocation) GetVni() uint32 {
	if m != nil {
		return m.Vni
	}
	return 0
}

func init() {
	proto.RegisterType((*CustomVxlanVniAllocation)(nil), "vnialloc.CustomVxlanVniAllocation")
}

func init() { proto.RegisterFile("vnialloc.proto", fileDescriptor_71dc0d993d483727) }

var fileDescriptor_71dc0d993d483727 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xcb, 0xcb, 0x4c,
	0xcc, 0xc9, 0xc9, 0x4f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xbc,
	0xb9, 0x24, 0x9c, 0x4b, 0x8b, 0x4b, 0xf2, 0x73, 0xc3, 0x2a, 0x72, 0x12, 0xf3, 0xc2, 0xf2, 0x32,
	0x1d, 0x41, 0xc2, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x42, 0xb2, 0x5c, 0x5c, 0x65, 0x20, 0xd1, 0xf8,
	0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x4e, 0xb0, 0x88, 0x5f, 0x62,
	0x6e, 0xaa, 0x90, 0x00, 0x17, 0x73, 0x59, 0x5e, 0xa6, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6f, 0x10,
	0x88, 0x99, 0xc4, 0x06, 0x36, 0xdd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xea, 0x37, 0xb4, 0xca,
	0x6f, 0x00, 0x00, 0x00,
}
