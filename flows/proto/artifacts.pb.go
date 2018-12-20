// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artifacts.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	proto1 "www.velocidex.com/golang/velociraptor/actions/proto"
	_ "www.velocidex.com/golang/velociraptor/proto"
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

type Artifacts struct {
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifacts) Reset()         { *m = Artifacts{} }
func (m *Artifacts) String() string { return proto.CompactTextString(m) }
func (*Artifacts) ProtoMessage()    {}
func (*Artifacts) Descriptor() ([]byte, []int) {
	return fileDescriptor_a56fa9fd5f0eb7ff, []int{0}
}

func (m *Artifacts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifacts.Unmarshal(m, b)
}
func (m *Artifacts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifacts.Marshal(b, m, deterministic)
}
func (m *Artifacts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifacts.Merge(m, src)
}
func (m *Artifacts) XXX_Size() int {
	return xxx_messageInfo_Artifacts.Size(m)
}
func (m *Artifacts) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifacts.DiscardUnknown(m)
}

var xxx_messageInfo_Artifacts proto.InternalMessageInfo

func (m *Artifacts) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type ArtifactCollectorArgs struct {
	Artifacts            *Artifacts       `protobuf:"bytes,1,opt,name=artifacts,proto3" json:"artifacts,omitempty"`
	Env                  []*proto1.VQLEnv `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ArtifactCollectorArgs) Reset()         { *m = ArtifactCollectorArgs{} }
func (m *ArtifactCollectorArgs) String() string { return proto.CompactTextString(m) }
func (*ArtifactCollectorArgs) ProtoMessage()    {}
func (*ArtifactCollectorArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_a56fa9fd5f0eb7ff, []int{1}
}

func (m *ArtifactCollectorArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactCollectorArgs.Unmarshal(m, b)
}
func (m *ArtifactCollectorArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactCollectorArgs.Marshal(b, m, deterministic)
}
func (m *ArtifactCollectorArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactCollectorArgs.Merge(m, src)
}
func (m *ArtifactCollectorArgs) XXX_Size() int {
	return xxx_messageInfo_ArtifactCollectorArgs.Size(m)
}
func (m *ArtifactCollectorArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactCollectorArgs.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactCollectorArgs proto.InternalMessageInfo

func (m *ArtifactCollectorArgs) GetArtifacts() *Artifacts {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func (m *ArtifactCollectorArgs) GetEnv() []*proto1.VQLEnv {
	if m != nil {
		return m.Env
	}
	return nil
}

func init() {
	proto.RegisterType((*Artifacts)(nil), "proto.Artifacts")
	proto.RegisterType((*ArtifactCollectorArgs)(nil), "proto.ArtifactCollectorArgs")
}

func init() { proto.RegisterFile("artifacts.proto", fileDescriptor_a56fa9fd5f0eb7ff) }

var fileDescriptor_a56fa9fd5f0eb7ff = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8e, 0xb1, 0x6a, 0x32, 0x41,
	0x14, 0x85, 0x59, 0x16, 0x7f, 0x70, 0xe4, 0x27, 0x61, 0x42, 0x60, 0xb1, 0xba, 0x58, 0x59, 0xcd,
	0x82, 0x62, 0x8a, 0x84, 0x14, 0x2a, 0x76, 0x36, 0x11, 0x49, 0x95, 0x66, 0x1c, 0xaf, 0x3a, 0x30,
	0x3b, 0x57, 0x67, 0xc7, 0xd9, 0xe4, 0xfd, 0xf2, 0x18, 0xa9, 0x92, 0xd7, 0x48, 0x11, 0x32, 0xab,
	0x9b, 0x2a, 0x90, 0xea, 0xc2, 0xe5, 0x9c, 0xef, 0x3b, 0xec, 0x42, 0x3a, 0xaf, 0x37, 0x52, 0xf9,
	0x52, 0xec, 0x1d, 0x79, 0xe2, 0xad, 0x78, 0xba, 0xb7, 0x55, 0x55, 0x89, 0x80, 0x86, 0x94, 0x5e,
	0xe3, 0xb3, 0x50, 0x54, 0xe4, 0x5b, 0x32, 0xd2, 0x6e, 0xf3, 0xfa, 0xe9, 0xe4, 0xde, 0x93, 0xcb,
	0x63, 0x38, 0x2f, 0xb1, 0x90, 0xd6, 0x6b, 0x55, 0x23, 0xba, 0xf7, 0x7f, 0xeb, 0x4a, 0xe5, 0x35,
	0xd9, 0xf2, 0xc4, 0x08, 0x07, 0x53, 0xd7, 0x7b, 0x53, 0xd6, 0x1e, 0x9f, 0x47, 0xf1, 0x1b, 0xd6,
	0xb2, 0xb2, 0xc0, 0x32, 0x4b, 0x20, 0xed, 0xb7, 0x27, 0xf0, 0xfe, 0xf9, 0xf1, 0x9a, 0x74, 0x79,
	0xb6, 0xdc, 0x21, 0x34, 0xd3, 0xc1, 0x13, 0x18, 0x79, 0xb4, 0x6a, 0x27, 0x16, 0x75, 0xbc, 0xf7,
	0x96, 0xb0, 0xeb, 0x33, 0x65, 0x4a, 0xc6, 0xa0, 0xf2, 0xe4, 0xc6, 0x6e, 0x5b, 0x72, 0x64, 0xed,
	0xa6, 0x98, 0x25, 0x90, 0xf4, 0x3b, 0x83, 0xcb, 0xda, 0x2c, 0x1a, 0xed, 0x64, 0x14, 0x3d, 0xf9,
	0xef, 0x9e, 0xde, 0x55, 0x13, 0x86, 0x25, 0xc1, 0x3c, 0x7e, 0x17, 0x3f, 0x64, 0xfe, 0xc4, 0x52,
	0xb4, 0x21, 0x4b, 0x21, 0xed, 0x77, 0x06, 0xff, 0x4f, 0x82, 0xc7, 0x87, 0xf9, 0xcc, 0x86, 0xc9,
	0x5d, 0xa4, 0x8f, 0xf8, 0x70, 0x66, 0x83, 0x76, 0x64, 0x0b, 0xb4, 0x1e, 0x82, 0x74, 0x5a, 0xae,
	0x0c, 0x46, 0xcb, 0x0a, 0x61, 0xef, 0x28, 0xe8, 0x35, 0xae, 0x61, 0x43, 0x0e, 0xfc, 0x0e, 0xe1,
	0x70, 0x44, 0xf7, 0x22, 0x16, 0xdf, 0xd8, 0xd5, 0xbf, 0xc8, 0x1b, 0x7e, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x91, 0x28, 0xaf, 0xba, 0xbf, 0x01, 0x00, 0x00,
}