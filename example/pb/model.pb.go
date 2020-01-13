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

type Object struct {
	A                    uint32   `protobuf:"varint,1,opt,name=A,proto3" json:"A,omitempty"`
	B                    uint64   `protobuf:"varint,2,opt,name=B,proto3" json:"B,omitempty"`
	C                    float32  `protobuf:"fixed32,3,opt,name=C,proto3" json:"C,omitempty"`
	D                    float64  `protobuf:"fixed64,4,opt,name=D,proto3" json:"D,omitempty"`
	E                    string   `protobuf:"bytes,5,opt,name=E,proto3" json:"E,omitempty"`
	F                    bool     `protobuf:"varint,6,opt,name=F,proto3" json:"F,omitempty"`
	G                    []byte   `protobuf:"bytes,7,opt,name=G,proto3" json:"G,omitempty"`
	H                    [][]byte `protobuf:"bytes,8,rep,name=H,proto3" json:"H,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
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

func (m *Object) GetA() uint32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Object) GetB() uint64 {
	if m != nil {
		return m.B
	}
	return 0
}

func (m *Object) GetC() float32 {
	if m != nil {
		return m.C
	}
	return 0
}

func (m *Object) GetD() float64 {
	if m != nil {
		return m.D
	}
	return 0
}

func (m *Object) GetE() string {
	if m != nil {
		return m.E
	}
	return ""
}

func (m *Object) GetF() bool {
	if m != nil {
		return m.F
	}
	return false
}

func (m *Object) GetG() []byte {
	if m != nil {
		return m.G
	}
	return nil
}

func (m *Object) GetH() [][]byte {
	if m != nil {
		return m.H
	}
	return nil
}

func init() {
	proto.RegisterType((*Object)(nil), "pb.Object")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x1c, 0xce, 0x2d, 0x0e, 0xc2, 0x40,
	0x10, 0x06, 0xd0, 0x7c, 0xdb, 0xb2, 0x94, 0xa5, 0x98, 0xaa, 0x91, 0x13, 0xd4, 0x28, 0x0c, 0x27,
	0xe8, 0x7f, 0x1d, 0xc9, 0x1c, 0x61, 0xa1, 0x86, 0x40, 0xb6, 0x21, 0x2b, 0x38, 0x3e, 0x19, 0xe4,
	0x73, 0x2f, 0x1c, 0xdf, 0xe9, 0xb1, 0xbe, 0x2e, 0xdb, 0x27, 0xe5, 0xd4, 0xb8, 0x2d, 0x9e, 0xbf,
	0xc1, 0xdf, 0xe2, 0x73, 0xbd, 0xe7, 0xa6, 0x0e, 0x68, 0x09, 0x0c, 0x39, 0x29, 0x5a, 0x53, 0x47,
	0x8e, 0x21, 0xa5, 0xa2, 0x33, 0xf5, 0x54, 0x30, 0xc4, 0x29, 0x7a, 0xd3, 0x40, 0x25, 0x43, 0xa0,
	0x18, 0x4c, 0x23, 0xed, 0x18, 0x72, 0x50, 0x8c, 0xa6, 0x89, 0x3c, 0x43, 0x2a, 0xc5, 0x64, 0x9a,
	0x69, 0xcf, 0x90, 0x5a, 0x31, 0x9b, 0x16, 0xaa, 0xb8, 0x30, 0x2d, 0xd1, 0xff, 0x13, 0xd7, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0x1a, 0x5b, 0x7f, 0x93, 0x00, 0x00, 0x00,
}
