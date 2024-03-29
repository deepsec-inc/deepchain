// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/configuration.proto

package orderer // import "deepchain/protos/orderer"

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

// MigrationState defines the possible states of a consensus-type migration.
// NONE represents that migration is not taking place.
type ConsensusType_MigrationState int32

const (
	ConsensusType_MIG_STATE_NONE    ConsensusType_MigrationState = 0
	ConsensusType_MIG_STATE_START   ConsensusType_MigrationState = 1
	ConsensusType_MIG_STATE_COMMIT  ConsensusType_MigrationState = 2
	ConsensusType_MIG_STATE_ABORT   ConsensusType_MigrationState = 3
	ConsensusType_MIG_STATE_CONTEXT ConsensusType_MigrationState = 4
)

var ConsensusType_MigrationState_name = map[int32]string{
	0: "MIG_STATE_NONE",
	1: "MIG_STATE_START",
	2: "MIG_STATE_COMMIT",
	3: "MIG_STATE_ABORT",
	4: "MIG_STATE_CONTEXT",
}
var ConsensusType_MigrationState_value = map[string]int32{
	"MIG_STATE_NONE":    0,
	"MIG_STATE_START":   1,
	"MIG_STATE_COMMIT":  2,
	"MIG_STATE_ABORT":   3,
	"MIG_STATE_CONTEXT": 4,
}

func (x ConsensusType_MigrationState) String() string {
	return proto.EnumName(ConsensusType_MigrationState_name, int32(x))
}
func (ConsensusType_MigrationState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_configuration_e51bcc2c9a7e7914, []int{0, 0}
}

type ConsensusType struct {
	// The consensus type: "solo", "kafka" or "etcdraft".
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Opaque metadata, dependent on the consensus type.
	Metadata []byte `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The state of a consensus-type migration.
	// On the system channel this could be NONE | START | COMMIT | ABORT.
	// On a standard channel this could be NONE | CONTEXT.
	// When migration is not taking place it is set to NONE.
	MigrationState ConsensusType_MigrationState `protobuf:"varint,3,opt,name=migration_state,json=migrationState,proto3,enum=orderer.ConsensusType_MigrationState" json:"migration_state,omitempty"`
	// The context of a consensus-type migration.
	// The migration context is the system channel block height of the config-update-tx that START migration.
	// The context must be present in COMMIT, CONTEXT, and ABORT config-updates, and match the START block height.
	// On NONE and START it is set =0.
	MigrationContext     uint64   `protobuf:"varint,4,opt,name=migration_context,json=migrationContext,proto3" json:"migration_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsensusType) Reset()         { *m = ConsensusType{} }
func (m *ConsensusType) String() string { return proto.CompactTextString(m) }
func (*ConsensusType) ProtoMessage()    {}
func (*ConsensusType) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_e51bcc2c9a7e7914, []int{0}
}
func (m *ConsensusType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusType.Unmarshal(m, b)
}
func (m *ConsensusType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusType.Marshal(b, m, deterministic)
}
func (dst *ConsensusType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusType.Merge(dst, src)
}
func (m *ConsensusType) XXX_Size() int {
	return xxx_messageInfo_ConsensusType.Size(m)
}
func (m *ConsensusType) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusType.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusType proto.InternalMessageInfo

func (m *ConsensusType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ConsensusType) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ConsensusType) GetMigrationState() ConsensusType_MigrationState {
	if m != nil {
		return m.MigrationState
	}
	return ConsensusType_MIG_STATE_NONE
}

func (m *ConsensusType) GetMigrationContext() uint64 {
	if m != nil {
		return m.MigrationContext
	}
	return 0
}

type BatchSize struct {
	// Simply specified as number of messages for now, in the future
	// we may want to allow this to be specified by size in bytes
	MaxMessageCount uint32 `protobuf:"varint,1,opt,name=max_message_count,json=maxMessageCount,proto3" json:"max_message_count,omitempty"`
	// The byte count of the serialized messages in a batch cannot
	// exceed this value.
	AbsoluteMaxBytes uint32 `protobuf:"varint,2,opt,name=absolute_max_bytes,json=absoluteMaxBytes,proto3" json:"absolute_max_bytes,omitempty"`
	// The byte count of the serialized messages in a batch should not
	// exceed this value.
	PreferredMaxBytes    uint32   `protobuf:"varint,3,opt,name=preferred_max_bytes,json=preferredMaxBytes,proto3" json:"preferred_max_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchSize) Reset()         { *m = BatchSize{} }
func (m *BatchSize) String() string { return proto.CompactTextString(m) }
func (*BatchSize) ProtoMessage()    {}
func (*BatchSize) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_e51bcc2c9a7e7914, []int{1}
}
func (m *BatchSize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchSize.Unmarshal(m, b)
}
func (m *BatchSize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchSize.Marshal(b, m, deterministic)
}
func (dst *BatchSize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchSize.Merge(dst, src)
}
func (m *BatchSize) XXX_Size() int {
	return xxx_messageInfo_BatchSize.Size(m)
}
func (m *BatchSize) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchSize.DiscardUnknown(m)
}

var xxx_messageInfo_BatchSize proto.InternalMessageInfo

func (m *BatchSize) GetMaxMessageCount() uint32 {
	if m != nil {
		return m.MaxMessageCount
	}
	return 0
}

func (m *BatchSize) GetAbsoluteMaxBytes() uint32 {
	if m != nil {
		return m.AbsoluteMaxBytes
	}
	return 0
}

func (m *BatchSize) GetPreferredMaxBytes() uint32 {
	if m != nil {
		return m.PreferredMaxBytes
	}
	return 0
}

type BatchTimeout struct {
	// Any duration string parseable by ParseDuration():
	// https://golang.org/pkg/time/#ParseDuration
	Timeout              string   `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchTimeout) Reset()         { *m = BatchTimeout{} }
func (m *BatchTimeout) String() string { return proto.CompactTextString(m) }
func (*BatchTimeout) ProtoMessage()    {}
func (*BatchTimeout) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_e51bcc2c9a7e7914, []int{2}
}
func (m *BatchTimeout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchTimeout.Unmarshal(m, b)
}
func (m *BatchTimeout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchTimeout.Marshal(b, m, deterministic)
}
func (dst *BatchTimeout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchTimeout.Merge(dst, src)
}
func (m *BatchTimeout) XXX_Size() int {
	return xxx_messageInfo_BatchTimeout.Size(m)
}
func (m *BatchTimeout) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchTimeout.DiscardUnknown(m)
}

var xxx_messageInfo_BatchTimeout proto.InternalMessageInfo

func (m *BatchTimeout) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

// Carries a list of bootstrap brokers, i.e. this is not the exclusive set of
// brokers an ordering service
type KafkaBrokers struct {
	// Each broker here should be identified using the (IP|host):port notation,
	// e.g. 127.0.0.1:7050, or localhost:7050 are valid entries
	Brokers              []string `protobuf:"bytes,1,rep,name=brokers,proto3" json:"brokers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaBrokers) Reset()         { *m = KafkaBrokers{} }
func (m *KafkaBrokers) String() string { return proto.CompactTextString(m) }
func (*KafkaBrokers) ProtoMessage()    {}
func (*KafkaBrokers) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_e51bcc2c9a7e7914, []int{3}
}
func (m *KafkaBrokers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaBrokers.Unmarshal(m, b)
}
func (m *KafkaBrokers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaBrokers.Marshal(b, m, deterministic)
}
func (dst *KafkaBrokers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaBrokers.Merge(dst, src)
}
func (m *KafkaBrokers) XXX_Size() int {
	return xxx_messageInfo_KafkaBrokers.Size(m)
}
func (m *KafkaBrokers) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaBrokers.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaBrokers proto.InternalMessageInfo

func (m *KafkaBrokers) GetBrokers() []string {
	if m != nil {
		return m.Brokers
	}
	return nil
}

// ChannelRestrictions is the mssage which conveys restrictions on channel creation for an orderer
type ChannelRestrictions struct {
	MaxCount             uint64   `protobuf:"varint,1,opt,name=max_count,json=maxCount,proto3" json:"max_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelRestrictions) Reset()         { *m = ChannelRestrictions{} }
func (m *ChannelRestrictions) String() string { return proto.CompactTextString(m) }
func (*ChannelRestrictions) ProtoMessage()    {}
func (*ChannelRestrictions) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_e51bcc2c9a7e7914, []int{4}
}
func (m *ChannelRestrictions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelRestrictions.Unmarshal(m, b)
}
func (m *ChannelRestrictions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelRestrictions.Marshal(b, m, deterministic)
}
func (dst *ChannelRestrictions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelRestrictions.Merge(dst, src)
}
func (m *ChannelRestrictions) XXX_Size() int {
	return xxx_messageInfo_ChannelRestrictions.Size(m)
}
func (m *ChannelRestrictions) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelRestrictions.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelRestrictions proto.InternalMessageInfo

func (m *ChannelRestrictions) GetMaxCount() uint64 {
	if m != nil {
		return m.MaxCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ConsensusType)(nil), "orderer.ConsensusType")
	proto.RegisterType((*BatchSize)(nil), "orderer.BatchSize")
	proto.RegisterType((*BatchTimeout)(nil), "orderer.BatchTimeout")
	proto.RegisterType((*KafkaBrokers)(nil), "orderer.KafkaBrokers")
	proto.RegisterType((*ChannelRestrictions)(nil), "orderer.ChannelRestrictions")
	proto.RegisterEnum("orderer.ConsensusType_MigrationState", ConsensusType_MigrationState_name, ConsensusType_MigrationState_value)
}

func init() {
	proto.RegisterFile("orderer/configuration.proto", fileDescriptor_configuration_e51bcc2c9a7e7914)
}

var fileDescriptor_configuration_e51bcc2c9a7e7914 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xdf, 0x8a, 0xda, 0x40,
	0x14, 0xc6, 0x1b, 0x95, 0xee, 0x7a, 0x58, 0x35, 0x8e, 0x2d, 0x84, 0xee, 0x8d, 0x08, 0x0b, 0xd2,
	0x96, 0x08, 0xdb, 0x27, 0x30, 0x41, 0xca, 0x52, 0xa2, 0x30, 0xce, 0x45, 0xe9, 0x4d, 0x98, 0x24,
	0xc7, 0x18, 0xd6, 0x64, 0xc2, 0xcc, 0x04, 0x62, 0xfb, 0x1c, 0x7d, 0x98, 0xbe, 0x5d, 0xc9, 0x3f,
	0x5d, 0xef, 0xce, 0xf7, 0x7d, 0xbf, 0x9c, 0x99, 0x39, 0x39, 0xf0, 0x28, 0x64, 0x84, 0x12, 0xe5,
	0x2a, 0x14, 0xd9, 0x21, 0x89, 0x0b, 0xc9, 0x75, 0x22, 0x32, 0x3b, 0x97, 0x42, 0x0b, 0x72, 0xd7,
	0x86, 0x8b, 0x7f, 0x3d, 0x18, 0xb9, 0x22, 0x53, 0x98, 0xa9, 0x42, 0xb1, 0x73, 0x8e, 0x84, 0xc0,
	0x40, 0x9f, 0x73, 0xb4, 0x8c, 0xb9, 0xb1, 0x1c, 0xd2, 0xba, 0x26, 0x9f, 0xe0, 0x3e, 0x45, 0xcd,
	0x23, 0xae, 0xb9, 0xd5, 0x9b, 0x1b, 0xcb, 0x07, 0x7a, 0xd1, 0x64, 0x0b, 0x93, 0x34, 0x89, 0x9b,
	0xee, 0xbe, 0xd2, 0x5c, 0xa3, 0xd5, 0x9f, 0x1b, 0xcb, 0xf1, 0xf3, 0x93, 0xdd, 0x1e, 0x62, 0xdf,
	0x1c, 0x60, 0x7b, 0x1d, 0xbd, 0xaf, 0x60, 0x3a, 0x4e, 0x6f, 0x34, 0xf9, 0x02, 0xd3, 0x6b, 0xbf,
	0x50, 0x64, 0x1a, 0x4b, 0x6d, 0x0d, 0xe6, 0xc6, 0x72, 0x40, 0xcd, 0x4b, 0xe0, 0x36, 0xfe, 0xe2,
	0x0f, 0x8c, 0x6f, 0xdb, 0x11, 0x02, 0x63, 0xef, 0xe5, 0xbb, 0xbf, 0x67, 0x6b, 0xb6, 0xf1, 0xb7,
	0xbb, 0xed, 0xc6, 0x7c, 0x47, 0x66, 0x30, 0xb9, 0x7a, 0x7b, 0xb6, 0xa6, 0xcc, 0x34, 0xc8, 0x07,
	0x30, 0xaf, 0xa6, 0xbb, 0xf3, 0xbc, 0x17, 0x66, 0xf6, 0x6e, 0xd1, 0xb5, 0xb3, 0xa3, 0xcc, 0xec,
	0x93, 0x8f, 0x30, 0x7d, 0x8b, 0x6e, 0xd9, 0xe6, 0x27, 0x33, 0x07, 0x8b, 0xbf, 0x06, 0x0c, 0x1d,
	0xae, 0xc3, 0xe3, 0x3e, 0xf9, 0x8d, 0xe4, 0x33, 0x4c, 0x53, 0x5e, 0xfa, 0x29, 0x2a, 0xc5, 0x63,
	0xf4, 0x43, 0x51, 0x64, 0xba, 0x1e, 0xe2, 0x88, 0x4e, 0x52, 0x5e, 0x7a, 0x8d, 0xef, 0x56, 0x36,
	0xf9, 0x0a, 0x84, 0x07, 0x4a, 0x9c, 0x0a, 0x8d, 0x7e, 0xf5, 0x51, 0x70, 0xd6, 0xa8, 0xea, 0xc9,
	0x8e, 0xa8, 0xd9, 0x25, 0x1e, 0x2f, 0x9d, 0xca, 0x27, 0x36, 0xcc, 0x72, 0x89, 0x07, 0x94, 0x12,
	0xa3, 0x37, 0x78, 0xbf, 0xc6, 0xa7, 0x97, 0xa8, 0xe3, 0x17, 0x4b, 0x78, 0xa8, 0xaf, 0xc5, 0x92,
	0x14, 0x45, 0xa1, 0x89, 0x05, 0x77, 0xba, 0x29, 0xdb, 0x9f, 0xda, 0xc9, 0x8a, 0xfc, 0xc1, 0x0f,
	0xaf, 0xdc, 0x91, 0xe2, 0x15, 0xa5, 0xaa, 0xc8, 0xa0, 0x29, 0x2d, 0x63, 0xde, 0xaf, 0xc8, 0x56,
	0x2e, 0x9e, 0x61, 0xe6, 0x1e, 0x79, 0x96, 0xe1, 0x89, 0xa2, 0xd2, 0x32, 0x09, 0xab, 0x89, 0x2b,
	0xf2, 0x08, 0xc3, 0xea, 0x42, 0xd7, 0xc7, 0x0e, 0xe8, 0x7d, 0xca, 0xcb, 0xfa, 0x95, 0xce, 0x1a,
	0x9e, 0x84, 0x8c, 0xed, 0xe3, 0x39, 0x47, 0x79, 0xc2, 0x28, 0x46, 0x69, 0x1f, 0x78, 0x20, 0x93,
	0xb0, 0x59, 0x42, 0xd5, 0xed, 0xc7, 0x2f, 0x2b, 0x42, 0xcc, 0xc3, 0x23, 0x4f, 0xb2, 0x55, 0x93,
	0xac, 0xda, 0x24, 0x78, 0x5f, 0xeb, 0x6f, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x63, 0x22, 0xac,
	0xec, 0xcd, 0x02, 0x00, 0x00,
}
