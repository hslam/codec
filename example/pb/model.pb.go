// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package pb

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

// 算术运算请求结构
type Student struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=Age,proto3" json:"Age,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Student) Reset()         { *m = Student{} }
func (m *Student) String() string { return proto.CompactTextString(m) }
func (*Student) ProtoMessage()    {}
func (*Student) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func (m *Student) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Student.Unmarshal(m, b)
}
func (m *Student) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Student.Marshal(b, m, deterministic)
}
func (m *Student) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Student.Merge(m, src)
}
func (m *Student) XXX_Size() int {
	return xxx_messageInfo_Student.Size(m)
}
func (m *Student) XXX_DiscardUnknown() {
	xxx_messageInfo_Student.DiscardUnknown(m)
}

var xxx_messageInfo_Student proto.InternalMessageInfo

func (m *Student) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Student) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Student) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Student)(nil), "pb.Student")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 106 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xf2, 0xe4, 0x62, 0x0f,
	0x2e, 0x29, 0x4d, 0x49, 0xcd, 0x2b, 0x11, 0x12, 0xe2, 0x62, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x1d, 0xd3, 0x53, 0x25, 0x98,
	0x14, 0x18, 0x35, 0x58, 0x83, 0x40, 0x4c, 0x21, 0x09, 0x2e, 0x76, 0xc7, 0x94, 0x94, 0xa2, 0xd4,
	0xe2, 0x62, 0x09, 0x66, 0xb0, 0x42, 0x18, 0x37, 0x89, 0x0d, 0x6c, 0xaa, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x9a, 0x79, 0x38, 0x74, 0x64, 0x00, 0x00, 0x00,
}