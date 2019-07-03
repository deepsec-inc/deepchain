// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/configuration.proto

package common // import "deepchain/protos/common"

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

// HashingAlgorithm is encoded into the configuration transaction as  a configuration item of type Chain
// with a Key of "HashingAlgorithm" and a Value of  HashingAlgorithm as marshaled protobuf bytes
type HashingAlgorithm struct {
	// Currently supported algorithms are: SHAKE256
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashingAlgorithm) Reset()         { *m = HashingAlgorithm{} }
func (m *HashingAlgorithm) String() string { return proto.CompactTextString(m) }
func (*HashingAlgorithm) ProtoMessage()    {}
func (*HashingAlgorithm) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_e4fcb01a1c6c57be, []int{0}
}
func (m *HashingAlgorithm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashingAlgorithm.Unmarshal(m, b)
}
func (m *HashingAlgorithm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashingAlgorithm.Marshal(b, m, deterministic)
}
func (dst *HashingAlgorithm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashingAlgorithm.Merge(dst, src)
}
func (m *HashingAlgorithm) XXX_Size() int {
	return xxx_messageInfo_HashingAlgorithm.Size(m)
}
func (m *HashingAlgorithm) XXX_DiscardUnknown() {
	xxx_messageInfo_HashingAlgorithm.DiscardUnknown(m)
}

var xxx_messageInfo_HashingAlgorithm proto.InternalMessageInfo

func (m *HashingAlgorithm) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// BlockDataHashingStructure is encoded into the configuration transaction as a configuration item of
// type Chain with a Key of "BlockDataHashingStructure" and a Value of HashingAlgorithm as marshaled protobuf bytes
type BlockDataHashingStructure struct {
	// width specifies the width of the Merkle tree to use when computing the BlockDataHash
	// in order to replicate flat hashing, set this width to MAX_UINT32
	Width                uint32   `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockDataHashingStructure) Reset()         { *m = BlockDataHashingStructure{} }
func (m *BlockDataHashingStructure) String() string { return proto.CompactTextString(m) }
func (*BlockDataHashingStructure) ProtoMessage()    {}
func (*BlockDataHashingStructure) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_e4fcb01a1c6c57be, []int{1}
}
func (m *BlockDataHashingStructure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockDataHashingStructure.Unmarshal(m, b)
}
func (m *BlockDataHashingStructure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockDataHashingStructure.Marshal(b, m, deterministic)
}
func (dst *BlockDataHashingStructure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockDataHashingStructure.Merge(dst, src)
}
func (m *BlockDataHashingStructure) XXX_Size() int {
	return xxx_messageInfo_BlockDataHashingStructure.Size(m)
}
func (m *BlockDataHashingStructure) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockDataHashingStructure.DiscardUnknown(m)
}

var xxx_messageInfo_BlockDataHashingStructure proto.InternalMessageInfo

func (m *BlockDataHashingStructure) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

// OrdererAddresses is encoded into the configuration transaction as a configuration item of type Chain
// with a Key of "OrdererAddresses" and a Value of OrdererAddresses as marshaled protobuf bytes
type OrdererAddresses struct {
	Addresses            []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrdererAddresses) Reset()         { *m = OrdererAddresses{} }
func (m *OrdererAddresses) String() string { return proto.CompactTextString(m) }
func (*OrdererAddresses) ProtoMessage()    {}
func (*OrdererAddresses) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_e4fcb01a1c6c57be, []int{2}
}
func (m *OrdererAddresses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrdererAddresses.Unmarshal(m, b)
}
func (m *OrdererAddresses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrdererAddresses.Marshal(b, m, deterministic)
}
func (dst *OrdererAddresses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrdererAddresses.Merge(dst, src)
}
func (m *OrdererAddresses) XXX_Size() int {
	return xxx_messageInfo_OrdererAddresses.Size(m)
}
func (m *OrdererAddresses) XXX_DiscardUnknown() {
	xxx_messageInfo_OrdererAddresses.DiscardUnknown(m)
}

var xxx_messageInfo_OrdererAddresses proto.InternalMessageInfo

func (m *OrdererAddresses) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

// Consortium represents the consortium context in which the channel was created
type Consortium struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Consortium) Reset()         { *m = Consortium{} }
func (m *Consortium) String() string { return proto.CompactTextString(m) }
func (*Consortium) ProtoMessage()    {}
func (*Consortium) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_e4fcb01a1c6c57be, []int{3}
}
func (m *Consortium) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consortium.Unmarshal(m, b)
}
func (m *Consortium) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consortium.Marshal(b, m, deterministic)
}
func (dst *Consortium) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consortium.Merge(dst, src)
}
func (m *Consortium) XXX_Size() int {
	return xxx_messageInfo_Consortium.Size(m)
}
func (m *Consortium) XXX_DiscardUnknown() {
	xxx_messageInfo_Consortium.DiscardUnknown(m)
}

var xxx_messageInfo_Consortium proto.InternalMessageInfo

func (m *Consortium) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Capabilities message defines the capabilities a particular binary must implement
// for that binary to be able to safely participate in the channel.  The capabilities
// message is defined at the /Channel level, the /Channel/Application level, and the
// /Channel/Orderer level.
//
// The /Channel level capabilties define capabilities which both the orderer and peer
// binaries must satisfy.  These capabilties might be things like a new MSP type,
// or a new policy type.
//
// The /Channel/Orderer level capabilties define capabilities which must be supported
// by the orderer, but which have no bearing on the behavior of the peer.  For instance
// if the orderer changes the logic for how it constructs new channels, only all orderers
// must agree on the new logic.  The peers do not need to be aware of this change as
// they only interact with the channel after it has been constructed.
//
// Finally, the /Channel/Application level capabilities define capabilities which the peer
// binary must satisfy, but which have no bearing on the orderer.  For instance, if the
// peer adds a new UTXO transaction type, or changes the chaincode lifecycle requirements,
// all peers must agree on the new logic.  However, orderers never inspect transactions
// this deeply, and therefore have no need to be aware of the change.
//
// The capabilities strings defined in these messages typically correspond to release
// binary versions (e.g. "V1.1"), and are used primarilly as a mechanism for a fully
// upgraded network to switch from one set of logic to a new one.
//
// Although for V1.1, the orderers must be upgraded to V1.1 prior to the rest of the
// network, going forward, because of the split between the /Channel, /Channel/Orderer
// and /Channel/Application capabilities.  It should be possible for the orderer and
// application networks to upgrade themselves independently (with the exception of any
// new capabilities defined at the /Channel level).
type Capabilities struct {
	Capabilities         map[string]*Capability `protobuf:"bytes,1,rep,name=capabilities,proto3" json:"capabilities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Capabilities) Reset()         { *m = Capabilities{} }
func (m *Capabilities) String() string { return proto.CompactTextString(m) }
func (*Capabilities) ProtoMessage()    {}
func (*Capabilities) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_e4fcb01a1c6c57be, []int{4}
}
func (m *Capabilities) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Capabilities.Unmarshal(m, b)
}
func (m *Capabilities) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Capabilities.Marshal(b, m, deterministic)
}
func (dst *Capabilities) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Capabilities.Merge(dst, src)
}
func (m *Capabilities) XXX_Size() int {
	return xxx_messageInfo_Capabilities.Size(m)
}
func (m *Capabilities) XXX_DiscardUnknown() {
	xxx_messageInfo_Capabilities.DiscardUnknown(m)
}

var xxx_messageInfo_Capabilities proto.InternalMessageInfo

func (m *Capabilities) GetCapabilities() map[string]*Capability {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

// Capability is an empty message for the time being.  It is defined as a protobuf
// message rather than a constant, so that we may extend capabilities with other fields
// if the need arises in the future.  For the time being, a capability being in the
// capabilities map requires that that capability be supported.
type Capability struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Capability) Reset()         { *m = Capability{} }
func (m *Capability) String() string { return proto.CompactTextString(m) }
func (*Capability) ProtoMessage()    {}
func (*Capability) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_e4fcb01a1c6c57be, []int{5}
}
func (m *Capability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Capability.Unmarshal(m, b)
}
func (m *Capability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Capability.Marshal(b, m, deterministic)
}
func (dst *Capability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Capability.Merge(dst, src)
}
func (m *Capability) XXX_Size() int {
	return xxx_messageInfo_Capability.Size(m)
}
func (m *Capability) XXX_DiscardUnknown() {
	xxx_messageInfo_Capability.DiscardUnknown(m)
}

var xxx_messageInfo_Capability proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HashingAlgorithm)(nil), "common.HashingAlgorithm")
	proto.RegisterType((*BlockDataHashingStructure)(nil), "common.BlockDataHashingStructure")
	proto.RegisterType((*OrdererAddresses)(nil), "common.OrdererAddresses")
	proto.RegisterType((*Consortium)(nil), "common.Consortium")
	proto.RegisterType((*Capabilities)(nil), "common.Capabilities")
	proto.RegisterMapType((map[string]*Capability)(nil), "common.Capabilities.CapabilitiesEntry")
	proto.RegisterType((*Capability)(nil), "common.Capability")
}

func init() {
	proto.RegisterFile("common/configuration.proto", fileDescriptor_configuration_e4fcb01a1c6c57be)
}

var fileDescriptor_configuration_e4fcb01a1c6c57be = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x51, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0xe9, 0xf6, 0xdf, 0x60, 0x77, 0xfb, 0xc3, 0x0c, 0x82, 0x73, 0xf8, 0x50, 0x8a, 0x8c,
	0x3e, 0x75, 0x3a, 0x5f, 0xc4, 0x17, 0xd9, 0xa6, 0x20, 0xbe, 0x08, 0xdd, 0x9b, 0x6f, 0x59, 0x7a,
	0xd7, 0x86, 0xb5, 0x49, 0xb9, 0x49, 0x95, 0x7e, 0x2a, 0xbf, 0xa2, 0xd8, 0x4c, 0xb6, 0x31, 0xdf,
	0xee, 0x2f, 0xe7, 0x9c, 0x9b, 0x13, 0x02, 0x63, 0xa1, 0x8b, 0x42, 0xab, 0xa9, 0xd0, 0x6a, 0x23,
	0xd3, 0x8a, 0xb8, 0x95, 0x5a, 0x45, 0x25, 0x69, 0xab, 0x59, 0xd7, 0x69, 0xc1, 0x04, 0x86, 0x2f,
	0xdc, 0x64, 0x52, 0xa5, 0xf3, 0x3c, 0xd5, 0x24, 0x6d, 0x56, 0x30, 0x06, 0xff, 0x14, 0x2f, 0x70,
	0xe4, 0xf9, 0x5e, 0xd8, 0x8b, 0x9b, 0x39, 0xb8, 0x85, 0xcb, 0x45, 0xae, 0xc5, 0xf6, 0x89, 0x5b,
	0xbe, 0x0b, 0xac, 0x2c, 0x55, 0xc2, 0x56, 0x84, 0xec, 0x1c, 0x3a, 0x9f, 0x32, 0xb1, 0x59, 0x93,
	0xf8, 0x1f, 0x3b, 0x08, 0x6e, 0x60, 0xf8, 0x46, 0x09, 0x12, 0xd2, 0x3c, 0x49, 0x08, 0x8d, 0x41,
	0xc3, 0xae, 0xa0, 0xc7, 0x7f, 0x61, 0xe4, 0xf9, 0xed, 0xb0, 0x17, 0xef, 0x0f, 0x02, 0x1f, 0x60,
	0xa9, 0x95, 0xd1, 0x64, 0x65, 0xf5, 0x77, 0x8d, 0x2f, 0x0f, 0x06, 0x4b, 0x5e, 0xf2, 0xb5, 0xcc,
	0xa5, 0x95, 0x68, 0xd8, 0x2b, 0x0c, 0xc4, 0x01, 0x37, 0x3b, 0xfb, 0xb3, 0x49, 0xe4, 0x9e, 0x17,
	0x1d, 0x7a, 0x8f, 0xe0, 0x59, 0x59, 0xaa, 0xe3, 0xa3, 0xec, 0x78, 0x05, 0x67, 0x27, 0x16, 0x36,
	0x84, 0xf6, 0x16, 0xeb, 0x5d, 0x89, 0x9f, 0x91, 0x85, 0xd0, 0xf9, 0xe0, 0x79, 0x85, 0xa3, 0x96,
	0xef, 0x85, 0xfd, 0x19, 0x3b, 0xb9, 0xab, 0x8e, 0x9d, 0xe1, 0xa1, 0x75, 0xef, 0x05, 0x03, 0x80,
	0xbd, 0xb0, 0x78, 0x84, 0x6b, 0x4d, 0x69, 0x94, 0xd5, 0x25, 0x52, 0x8e, 0x49, 0x8a, 0x14, 0x6d,
	0xf8, 0x9a, 0xa4, 0x70, 0xdf, 0x62, 0x76, 0xbb, 0xde, 0x2f, 0x12, 0xc4, 0x52, 0x64, 0x5c, 0xaa,
	0xa9, 0x13, 0xa6, 0x4e, 0x58, 0x77, 0x1b, 0xbc, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x61,
	0x8d, 0x64, 0xdc, 0x01, 0x00, 0x00,
}
