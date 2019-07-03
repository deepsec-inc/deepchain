// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msgs.proto

package msgs // import "deepchain/core/ledger/kvledger/txmgmt/statedb/statecouchdb/msgs"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VersionFieldProto struct {
	VersionBytes         []byte   `protobuf:"bytes,1,opt,name=version_bytes,json=versionBytes,proto3" json:"version_bytes,omitempty"`
	Metadata             []byte   `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionFieldProto) Reset()         { *m = VersionFieldProto{} }
func (m *VersionFieldProto) String() string { return proto.CompactTextString(m) }
func (*VersionFieldProto) ProtoMessage()    {}
func (*VersionFieldProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_msgs_ca4007467018fecf, []int{0}
}
func (m *VersionFieldProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionFieldProto.Unmarshal(m, b)
}
func (m *VersionFieldProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionFieldProto.Marshal(b, m, deterministic)
}
func (dst *VersionFieldProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionFieldProto.Merge(dst, src)
}
func (m *VersionFieldProto) XXX_Size() int {
	return xxx_messageInfo_VersionFieldProto.Size(m)
}
func (m *VersionFieldProto) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionFieldProto.DiscardUnknown(m)
}

var xxx_messageInfo_VersionFieldProto proto.InternalMessageInfo

func (m *VersionFieldProto) GetVersionBytes() []byte {
	if m != nil {
		return m.VersionBytes
	}
	return nil
}

func (m *VersionFieldProto) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionFieldProto)(nil), "msgs.VersionFieldProto")
}

func init() { proto.RegisterFile("msgs.proto", fileDescriptor_msgs_ca4007467018fecf) }

var fileDescriptor_msgs_ca4007467018fecf = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x2d, 0x4e, 0x2f,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x42, 0xb8, 0x04, 0xc3, 0x52,
	0x8b, 0x8a, 0x33, 0xf3, 0xf3, 0xdc, 0x32, 0x53, 0x73, 0x52, 0x02, 0xc0, 0x52, 0xca, 0x5c, 0xbc,
	0x65, 0x10, 0xc1, 0xf8, 0xa4, 0xca, 0x92, 0xd4, 0x62, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20,
	0x1e, 0xa8, 0xa0, 0x13, 0x48, 0x4c, 0x48, 0x8a, 0x8b, 0x23, 0x37, 0xb5, 0x24, 0x31, 0x25, 0xb1,
	0x24, 0x51, 0x82, 0x09, 0x2c, 0x0f, 0xe7, 0x3b, 0x39, 0x46, 0xd9, 0xa7, 0xa4, 0xa6, 0x16, 0x24,
	0x67, 0x24, 0x66, 0xe6, 0xe9, 0x27, 0xe7, 0x17, 0xa5, 0xea, 0xe7, 0xa4, 0xa6, 0xa4, 0xa7, 0x16,
	0xe9, 0x67, 0x97, 0x41, 0x19, 0x25, 0x15, 0xb9, 0xe9, 0xb9, 0x25, 0xfa, 0xc5, 0x25, 0x89, 0x25,
	0xa9, 0x29, 0x49, 0x10, 0x3a, 0x39, 0xbf, 0x34, 0x39, 0x23, 0x25, 0x49, 0x1f, 0xe4, 0xb0, 0x24,
	0x36, 0xb0, 0x2b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xc3, 0xbf, 0x4d, 0xb3, 0x00,
	0x00, 0x00,
}
