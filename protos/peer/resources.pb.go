// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/resources.proto

package peer // import "deepchain/protos/peer"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "deepchain/protos/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ChaincodeIdentifier identifies a piece of chaincode.  For a peer to accept invocations of
// this chaincode, the hash of the installed code must match, as must the version string
// included with the install command.
type ChaincodeIdentifier struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeIdentifier) Reset()         { *m = ChaincodeIdentifier{} }
func (m *ChaincodeIdentifier) String() string { return proto.CompactTextString(m) }
func (*ChaincodeIdentifier) ProtoMessage()    {}
func (*ChaincodeIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_457593128941a568, []int{0}
}
func (m *ChaincodeIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeIdentifier.Unmarshal(m, b)
}
func (m *ChaincodeIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeIdentifier.Marshal(b, m, deterministic)
}
func (dst *ChaincodeIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeIdentifier.Merge(dst, src)
}
func (m *ChaincodeIdentifier) XXX_Size() int {
	return xxx_messageInfo_ChaincodeIdentifier.Size(m)
}
func (m *ChaincodeIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeIdentifier proto.InternalMessageInfo

func (m *ChaincodeIdentifier) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ChaincodeIdentifier) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// ChaincodeValidation instructs the peer how transactions for this chaincode should be
// validated.  The only validation mechanism which ships with fabric today is the standard
// 'vscc' validation mechanism.  This built in validation method utilizes an endorsement policy
// which checks that a sufficient number of signatures have been included.  The 'arguement'
// field encodes any parameters required by the validation implementation.
type ChaincodeValidation struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Argument             []byte   `protobuf:"bytes,2,opt,name=argument,proto3" json:"argument,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeValidation) Reset()         { *m = ChaincodeValidation{} }
func (m *ChaincodeValidation) String() string { return proto.CompactTextString(m) }
func (*ChaincodeValidation) ProtoMessage()    {}
func (*ChaincodeValidation) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_457593128941a568, []int{1}
}
func (m *ChaincodeValidation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeValidation.Unmarshal(m, b)
}
func (m *ChaincodeValidation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeValidation.Marshal(b, m, deterministic)
}
func (dst *ChaincodeValidation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeValidation.Merge(dst, src)
}
func (m *ChaincodeValidation) XXX_Size() int {
	return xxx_messageInfo_ChaincodeValidation.Size(m)
}
func (m *ChaincodeValidation) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeValidation.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeValidation proto.InternalMessageInfo

func (m *ChaincodeValidation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeValidation) GetArgument() []byte {
	if m != nil {
		return m.Argument
	}
	return nil
}

// VSCCArgs is passed (marshaled) as a parameter to the VSCC imlementation via the
// argument field of the ChaincodeValidation message.
type VSCCArgs struct {
	EndorsementPolicyRef string   `protobuf:"bytes,1,opt,name=endorsement_policy_ref,json=endorsementPolicyRef,proto3" json:"endorsement_policy_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VSCCArgs) Reset()         { *m = VSCCArgs{} }
func (m *VSCCArgs) String() string { return proto.CompactTextString(m) }
func (*VSCCArgs) ProtoMessage()    {}
func (*VSCCArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_457593128941a568, []int{2}
}
func (m *VSCCArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VSCCArgs.Unmarshal(m, b)
}
func (m *VSCCArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VSCCArgs.Marshal(b, m, deterministic)
}
func (dst *VSCCArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VSCCArgs.Merge(dst, src)
}
func (m *VSCCArgs) XXX_Size() int {
	return xxx_messageInfo_VSCCArgs.Size(m)
}
func (m *VSCCArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_VSCCArgs.DiscardUnknown(m)
}

var xxx_messageInfo_VSCCArgs proto.InternalMessageInfo

func (m *VSCCArgs) GetEndorsementPolicyRef() string {
	if m != nil {
		return m.EndorsementPolicyRef
	}
	return ""
}

// ChaincodeEndorsement instructs the peer how transactions should be endorsed.  The only
// endorsement mechanism which ships with the fabric today is the standard 'escc' mechanism.
// This code simply simulates the proposal to generate a RW set, then signs the result
// using the peer's local signing identity.
type ChaincodeEndorsement struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeEndorsement) Reset()         { *m = ChaincodeEndorsement{} }
func (m *ChaincodeEndorsement) String() string { return proto.CompactTextString(m) }
func (*ChaincodeEndorsement) ProtoMessage()    {}
func (*ChaincodeEndorsement) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_457593128941a568, []int{3}
}
func (m *ChaincodeEndorsement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeEndorsement.Unmarshal(m, b)
}
func (m *ChaincodeEndorsement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeEndorsement.Marshal(b, m, deterministic)
}
func (dst *ChaincodeEndorsement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeEndorsement.Merge(dst, src)
}
func (m *ChaincodeEndorsement) XXX_Size() int {
	return xxx_messageInfo_ChaincodeEndorsement.Size(m)
}
func (m *ChaincodeEndorsement) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeEndorsement.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeEndorsement proto.InternalMessageInfo

func (m *ChaincodeEndorsement) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// ConfigTree encapsulates channel and resources configuration of a channel.
// Both configurations are represented as common.Config
type ConfigTree struct {
	ChannelConfig        *common.Config `protobuf:"bytes,1,opt,name=channel_config,json=channelConfig,proto3" json:"channel_config,omitempty"`
	ResourcesConfig      *common.Config `protobuf:"bytes,2,opt,name=resources_config,json=resourcesConfig,proto3" json:"resources_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ConfigTree) Reset()         { *m = ConfigTree{} }
func (m *ConfigTree) String() string { return proto.CompactTextString(m) }
func (*ConfigTree) ProtoMessage()    {}
func (*ConfigTree) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_457593128941a568, []int{4}
}
func (m *ConfigTree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigTree.Unmarshal(m, b)
}
func (m *ConfigTree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigTree.Marshal(b, m, deterministic)
}
func (dst *ConfigTree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigTree.Merge(dst, src)
}
func (m *ConfigTree) XXX_Size() int {
	return xxx_messageInfo_ConfigTree.Size(m)
}
func (m *ConfigTree) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigTree.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigTree proto.InternalMessageInfo

func (m *ConfigTree) GetChannelConfig() *common.Config {
	if m != nil {
		return m.ChannelConfig
	}
	return nil
}

func (m *ConfigTree) GetResourcesConfig() *common.Config {
	if m != nil {
		return m.ResourcesConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*ChaincodeIdentifier)(nil), "protos.ChaincodeIdentifier")
	proto.RegisterType((*ChaincodeValidation)(nil), "protos.ChaincodeValidation")
	proto.RegisterType((*VSCCArgs)(nil), "protos.VSCCArgs")
	proto.RegisterType((*ChaincodeEndorsement)(nil), "protos.ChaincodeEndorsement")
	proto.RegisterType((*ConfigTree)(nil), "protos.ConfigTree")
}

func init() { proto.RegisterFile("peer/resources.proto", fileDescriptor_resources_457593128941a568) }

var fileDescriptor_resources_457593128941a568 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xbf, 0x4f, 0xf3, 0x30,
	0x10, 0x55, 0xab, 0x4f, 0x1f, 0xed, 0x51, 0x0a, 0x32, 0x2d, 0xaa, 0x3a, 0x55, 0x99, 0x2a, 0x86,
	0x44, 0xe2, 0xc7, 0x80, 0x58, 0x80, 0xa8, 0x03, 0x1b, 0x0a, 0xa8, 0x03, 0x4b, 0xe5, 0xda, 0x97,
	0xc4, 0x52, 0x62, 0x47, 0xe7, 0x14, 0xd1, 0x85, 0xbf, 0x1d, 0xc5, 0x6e, 0x43, 0x87, 0x4e, 0xbe,
	0xbb, 0xf7, 0xee, 0xe9, 0xfc, 0x1e, 0x8c, 0x2a, 0x44, 0x8a, 0x08, 0xad, 0xd9, 0x90, 0x40, 0x1b,
	0x56, 0x64, 0x6a, 0xc3, 0xfe, 0xbb, 0xc7, 0x4e, 0xc7, 0xc2, 0x94, 0xa5, 0xd1, 0x91, 0x30, 0x3a,
	0x55, 0x59, 0xfd, 0xed, 0xe1, 0x20, 0x86, 0xcb, 0x38, 0xe7, 0x4a, 0x0b, 0x23, 0xf1, 0x55, 0xa2,
	0xae, 0x55, 0xaa, 0x90, 0x18, 0x83, 0x7f, 0x39, 0xb7, 0xf9, 0xa4, 0x33, 0xeb, 0xcc, 0x07, 0x89,
	0xab, 0xd9, 0x04, 0x4e, 0xbe, 0x90, 0xac, 0x32, 0x7a, 0xd2, 0x9d, 0x75, 0xe6, 0xfd, 0x64, 0xdf,
	0x06, 0x8b, 0x03, 0x91, 0x25, 0x2f, 0x94, 0xe4, 0xb5, 0x32, 0xba, 0x11, 0xd1, 0xbc, 0x44, 0x27,
	0xd2, 0x4f, 0x5c, 0xcd, 0xa6, 0xd0, 0xe3, 0x94, 0x6d, 0x4a, 0xd4, 0xb5, 0x53, 0x19, 0x24, 0x6d,
	0x1f, 0x3c, 0x41, 0x6f, 0xf9, 0x1e, 0xc7, 0xcf, 0x94, 0x59, 0x76, 0x07, 0x57, 0xa8, 0xa5, 0x21,
	0x8b, 0x0d, 0xb4, 0xaa, 0x4c, 0xa1, 0xc4, 0x76, 0x45, 0x98, 0xee, 0xd4, 0x46, 0x07, 0xe8, 0x9b,
	0x03, 0x13, 0x4c, 0x83, 0x6b, 0x18, 0xb5, 0x87, 0x2c, 0xfe, 0x08, 0xc7, 0x2e, 0x09, 0x7e, 0x00,
	0x62, 0xe7, 0xc5, 0x07, 0x21, 0xb2, 0x7b, 0x18, 0x8a, 0x9c, 0x6b, 0x8d, 0xc5, 0xca, 0x3b, 0xe4,
	0xb8, 0xa7, 0x37, 0xc3, 0xd0, 0xfb, 0x16, 0x7a, 0x6e, 0x72, 0xb6, 0x63, 0xf9, 0x96, 0x3d, 0xc0,
	0x45, 0x6b, 0xf8, 0x7e, 0xb1, 0x7b, 0x74, 0xf1, 0xbc, 0xe5, 0xf9, 0xc1, 0xcb, 0x23, 0x04, 0x86,
	0xb2, 0x30, 0xdf, 0x56, 0x48, 0x05, 0xca, 0x0c, 0x29, 0x4c, 0xf9, 0x9a, 0x94, 0xf0, 0xc9, 0xd8,
	0xb0, 0x89, 0xf3, 0x73, 0x2c, 0x11, 0x2b, 0xd1, 0xfc, 0x29, 0xf2, 0xe3, 0xa8, 0x19, 0xaf, 0x7d,
	0xaa, 0xb7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xef, 0x53, 0xb7, 0x0f, 0xf4, 0x01, 0x00, 0x00,
}
