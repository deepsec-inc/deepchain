// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/signed_cc_dep_spec.proto

package peer // import "deepchain/protos/peer"

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

// SignedChaincodeDeploymentSpec carries the CDS along with endorsements
type SignedChaincodeDeploymentSpec struct {
	// This is the bytes of the ChaincodeDeploymentSpec
	ChaincodeDeploymentSpec []byte `protobuf:"bytes,1,opt,name=chaincode_deployment_spec,json=chaincodeDeploymentSpec,proto3" json:"chaincode_deployment_spec,omitempty"`
	// This is the instantiation policy which is identical in structure
	// to endorsement policy.  This policy is checked by the VSCC at commit
	// time on the instantiation (all peers will get the same policy as it
	// will be part of the LSCC instantation record and will be part of the
	// hash as well)
	InstantiationPolicy []byte `protobuf:"bytes,2,opt,name=instantiation_policy,json=instantiationPolicy,proto3" json:"instantiation_policy,omitempty"`
	// The endorsements of the above deployment spec, the owner's signature over
	// chaincode_deployment_spec and Endorsement.endorser.
	OwnerEndorsements    []*Endorsement `protobuf:"bytes,3,rep,name=owner_endorsements,json=ownerEndorsements,proto3" json:"owner_endorsements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SignedChaincodeDeploymentSpec) Reset()         { *m = SignedChaincodeDeploymentSpec{} }
func (m *SignedChaincodeDeploymentSpec) String() string { return proto.CompactTextString(m) }
func (*SignedChaincodeDeploymentSpec) ProtoMessage()    {}
func (*SignedChaincodeDeploymentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_signed_cc_dep_spec_7acd8064ee280fd9, []int{0}
}
func (m *SignedChaincodeDeploymentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedChaincodeDeploymentSpec.Unmarshal(m, b)
}
func (m *SignedChaincodeDeploymentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedChaincodeDeploymentSpec.Marshal(b, m, deterministic)
}
func (dst *SignedChaincodeDeploymentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedChaincodeDeploymentSpec.Merge(dst, src)
}
func (m *SignedChaincodeDeploymentSpec) XXX_Size() int {
	return xxx_messageInfo_SignedChaincodeDeploymentSpec.Size(m)
}
func (m *SignedChaincodeDeploymentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedChaincodeDeploymentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_SignedChaincodeDeploymentSpec proto.InternalMessageInfo

func (m *SignedChaincodeDeploymentSpec) GetChaincodeDeploymentSpec() []byte {
	if m != nil {
		return m.ChaincodeDeploymentSpec
	}
	return nil
}

func (m *SignedChaincodeDeploymentSpec) GetInstantiationPolicy() []byte {
	if m != nil {
		return m.InstantiationPolicy
	}
	return nil
}

func (m *SignedChaincodeDeploymentSpec) GetOwnerEndorsements() []*Endorsement {
	if m != nil {
		return m.OwnerEndorsements
	}
	return nil
}

func init() {
	proto.RegisterType((*SignedChaincodeDeploymentSpec)(nil), "protos.SignedChaincodeDeploymentSpec")
}

func init() {
	proto.RegisterFile("peer/signed_cc_dep_spec.proto", fileDescriptor_signed_cc_dep_spec_7acd8064ee280fd9)
}

var fileDescriptor_signed_cc_dep_spec_7acd8064ee280fd9 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x05, 0x0f, 0xab, 0x17, 0x53, 0xc5, 0x28, 0x16, 0x4a, 0x4f, 0x3d, 0x25, 0xa8,
	0x37, 0xbd, 0x55, 0xbd, 0x4b, 0x7b, 0xf3, 0xb2, 0xa4, 0xbb, 0xcf, 0x74, 0x21, 0xee, 0x0c, 0x33,
	0x01, 0xc9, 0xdf, 0xf4, 0x17, 0x49, 0x36, 0x16, 0xf5, 0xd0, 0xd3, 0xc0, 0xfb, 0xde, 0x9b, 0x19,
	0x9e, 0x99, 0x31, 0x20, 0x95, 0x86, 0x26, 0xc2, 0x5b, 0xe7, 0xac, 0x07, 0x5b, 0x65, 0xb8, 0x92,
	0x85, 0x3a, 0xca, 0x8f, 0xd3, 0xd0, 0xeb, 0x9b, 0x64, 0x63, 0x21, 0x26, 0xad, 0x5b, 0x2b, 0x50,
	0xa6, 0xa8, 0x18, 0x5d, 0x8b, 0xaf, 0xcc, 0xcc, 0x36, 0x69, 0xc5, 0xd3, 0xae, 0x0e, 0xd1, 0x91,
	0xc7, 0x33, 0xb8, 0xa5, 0xfe, 0x03, 0xb1, 0xdb, 0x30, 0x5c, 0xfe, 0x60, 0xae, 0xdc, 0x1e, 0x0d,
	0x37, 0x7e, 0x58, 0x3a, 0x55, 0x64, 0xf3, 0x6c, 0x79, 0xba, 0xbe, 0x74, 0x07, 0xb2, 0xb7, 0xe6,
	0x3c, 0x44, 0xed, 0xea, 0xd8, 0x85, 0xba, 0x0b, 0x14, 0x2d, 0x53, 0x1b, 0x5c, 0x5f, 0x1c, 0xa5,
	0xd8, 0xf4, 0x1f, 0x7b, 0x4d, 0x28, 0x5f, 0x99, 0x9c, 0x3e, 0x23, 0xc4, 0x22, 0x7a, 0x12, 0xc5,
	0xb0, 0x4b, 0x8b, 0xc9, 0x7c, 0xb2, 0x3c, 0xb9, 0x9b, 0x8e, 0x4f, 0x6b, 0xf9, 0xf2, 0xcb, 0xd6,
	0x67, 0xc9, 0xfe, 0x47, 0xd1, 0xd5, 0xa3, 0x59, 0x90, 0x34, 0xe5, 0xae, 0x67, 0x48, 0x0b, 0xdf,
	0x40, 0xca, 0xf7, 0x7a, 0x2b, 0xc1, 0xed, 0xf3, 0x43, 0x25, 0x6f, 0x17, 0x1e, 0xe0, 0xf4, 0x79,
	0x35, 0xca, 0xd5, 0x20, 0x6f, 0xc7, 0xde, 0xee, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x72,
	0x1b, 0x58, 0x5f, 0x01, 0x00, 0x00,
}
