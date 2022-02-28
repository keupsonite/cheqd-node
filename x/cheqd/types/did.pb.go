// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cheqd/v1/did.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Did struct {
	Context              []string              `protobuf:"bytes,1,rep,name=context,proto3" json:"context,omitempty"`
	Id                   string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Controller           []string              `protobuf:"bytes,3,rep,name=controller,proto3" json:"controller,omitempty"`
	VerificationMethod   []*VerificationMethod `protobuf:"bytes,4,rep,name=verification_method,json=verificationMethod,proto3" json:"verification_method,omitempty"`
	Authentication       []string              `protobuf:"bytes,5,rep,name=authentication,proto3" json:"authentication,omitempty"`
	AssertionMethod      []string              `protobuf:"bytes,6,rep,name=assertion_method,json=assertionMethod,proto3" json:"assertion_method,omitempty"`
	CapabilityInvocation []string              `protobuf:"bytes,7,rep,name=capability_invocation,json=capabilityInvocation,proto3" json:"capability_invocation,omitempty"`
	CapabilityDelegation []string              `protobuf:"bytes,8,rep,name=capability_delegation,json=capabilityDelegation,proto3" json:"capability_delegation,omitempty"`
	KeyAgreement         []string              `protobuf:"bytes,9,rep,name=key_agreement,json=keyAgreement,proto3" json:"key_agreement,omitempty"`
	Service              []*Service            `protobuf:"bytes,10,rep,name=service,proto3" json:"service,omitempty"`
	AlsoKnownAs          []string              `protobuf:"bytes,11,rep,name=also_known_as,json=alsoKnownAs,proto3" json:"also_known_as,omitempty"`
}

func (m *Did) Reset()         { *m = Did{} }
func (m *Did) String() string { return proto.CompactTextString(m) }
func (*Did) ProtoMessage()    {}
func (*Did) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb1cddf7c2ece8cb, []int{0}
}
func (m *Did) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Did) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Did.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Did) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Did.Merge(m, src)
}
func (m *Did) XXX_Size() int {
	return m.Size()
}
func (m *Did) XXX_DiscardUnknown() {
	xxx_messageInfo_Did.DiscardUnknown(m)
}

var xxx_messageInfo_Did proto.InternalMessageInfo

func (m *Did) GetContext() []string {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Did) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Did) GetController() []string {
	if m != nil {
		return m.Controller
	}
	return nil
}

func (m *Did) GetVerificationMethod() []*VerificationMethod {
	if m != nil {
		return m.VerificationMethod
	}
	return nil
}

func (m *Did) GetAuthentication() []string {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *Did) GetAssertionMethod() []string {
	if m != nil {
		return m.AssertionMethod
	}
	return nil
}

func (m *Did) GetCapabilityInvocation() []string {
	if m != nil {
		return m.CapabilityInvocation
	}
	return nil
}

func (m *Did) GetCapabilityDelegation() []string {
	if m != nil {
		return m.CapabilityDelegation
	}
	return nil
}

func (m *Did) GetKeyAgreement() []string {
	if m != nil {
		return m.KeyAgreement
	}
	return nil
}

func (m *Did) GetService() []*Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Did) GetAlsoKnownAs() []string {
	if m != nil {
		return m.AlsoKnownAs
	}
	return nil
}

type VerificationMethod struct {
	Id                 string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validate:"required,did-url,did-url-no-path,did-url-no-query,did-url-with-fragment"`
	Type               string          `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Controller         string          `protobuf:"bytes,3,opt,name=controller,proto3" json:"controller,omitempty"`
	PublicKeyJwk       []*KeyValuePair `protobuf:"bytes,4,rep,name=public_key_jwk,json=publicKeyJwk,proto3" json:"public_key_jwk,omitempty" validate:"required_without_all=PublicKeyMultibase"`
	PublicKeyMultibase string          `protobuf:"bytes,5,opt,name=public_key_multibase,json=publicKeyMultibase,proto3" json:"public_key_multibase,omitempty" validate:"required_without_all=PublicKeyJwk"`
}

func (m *VerificationMethod) Reset()         { *m = VerificationMethod{} }
func (m *VerificationMethod) String() string { return proto.CompactTextString(m) }
func (*VerificationMethod) ProtoMessage()    {}
func (*VerificationMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb1cddf7c2ece8cb, []int{1}
}
func (m *VerificationMethod) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerificationMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerificationMethod.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerificationMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerificationMethod.Merge(m, src)
}
func (m *VerificationMethod) XXX_Size() int {
	return m.Size()
}
func (m *VerificationMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_VerificationMethod.DiscardUnknown(m)
}

var xxx_messageInfo_VerificationMethod proto.InternalMessageInfo

func (m *VerificationMethod) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VerificationMethod) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *VerificationMethod) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *VerificationMethod) GetPublicKeyJwk() []*KeyValuePair {
	if m != nil {
		return m.PublicKeyJwk
	}
	return nil
}

func (m *VerificationMethod) GetPublicKeyMultibase() string {
	if m != nil {
		return m.PublicKeyMultibase
	}
	return ""
}

type Service struct {
	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type            string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ServiceEndpoint string `protobuf:"bytes,3,opt,name=service_endpoint,json=serviceEndpoint,proto3" json:"service_endpoint,omitempty"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb1cddf7c2ece8cb, []int{2}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Service.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return m.Size()
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Service) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Service) GetServiceEndpoint() string {
	if m != nil {
		return m.ServiceEndpoint
	}
	return ""
}

func init() {
	proto.RegisterType((*Did)(nil), "cheqdid.cheqdnode.cheqd.v1.Did")
	proto.RegisterType((*VerificationMethod)(nil), "cheqdid.cheqdnode.cheqd.v1.VerificationMethod")
	proto.RegisterType((*Service)(nil), "cheqdid.cheqdnode.cheqd.v1.Service")
}

func init() { proto.RegisterFile("cheqd/v1/did.proto", fileDescriptor_fb1cddf7c2ece8cb) }

var fileDescriptor_fb1cddf7c2ece8cb = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0xc7, 0x31, 0x01, 0x52, 0x96, 0x4f, 0x6d, 0x41, 0x72, 0x39, 0x18, 0x64, 0xa4, 0x2a, 0x48,
	0xc4, 0x2e, 0x20, 0xf5, 0x50, 0x89, 0x03, 0x94, 0x1e, 0x0a, 0x42, 0x42, 0xae, 0x84, 0xaa, 0x5e,
	0xac, 0xb5, 0x77, 0xe2, 0x6c, 0xe3, 0x78, 0xcd, 0x7a, 0xed, 0xe0, 0x17, 0xe8, 0xb9, 0x2f, 0xd0,
	0xf7, 0xe9, 0x91, 0x63, 0x4f, 0xa8, 0x82, 0x37, 0x88, 0xd4, 0x7b, 0xe5, 0xcd, 0x26, 0x8d, 0x92,
	0x16, 0xf5, 0x62, 0xef, 0xfe, 0x66, 0xc6, 0x33, 0xfb, 0x9f, 0xf1, 0x22, 0x1c, 0xb6, 0xe1, 0x86,
	0xba, 0xc5, 0x81, 0x4b, 0x19, 0x75, 0x52, 0xc1, 0x25, 0xc7, 0x5b, 0x8a, 0x31, 0xea, 0xa8, 0x77,
	0xc2, 0x29, 0x0c, 0x56, 0x4e, 0x71, 0xb0, 0xb5, 0x11, 0xf1, 0x88, 0x2b, 0x37, 0xb7, 0x5a, 0x0d,
	0x22, 0xb6, 0x5e, 0x44, 0x9c, 0x47, 0x31, 0xb8, 0x6a, 0x17, 0xe4, 0x2d, 0x97, 0x24, 0xa5, 0x36,
	0x6d, 0x8e, 0x12, 0x84, 0xbc, 0xdb, 0xe5, 0xc9, 0x00, 0xdb, 0xbf, 0x6a, 0xa8, 0x76, 0xc6, 0x28,
	0x36, 0x51, 0x3d, 0xe4, 0x89, 0x84, 0x5b, 0x69, 0x1a, 0x3b, 0xb5, 0xc6, 0xa2, 0x37, 0xdc, 0xe2,
	0x55, 0x34, 0xcb, 0xa8, 0x39, 0xbb, 0x63, 0x34, 0x16, 0xbd, 0x59, 0x46, 0xb1, 0x85, 0x50, 0x65,
	0x12, 0x3c, 0x8e, 0x41, 0x98, 0x35, 0xe5, 0x3c, 0x46, 0xb0, 0x8f, 0x9e, 0x17, 0x20, 0x58, 0x8b,
	0x85, 0x44, 0x32, 0x9e, 0xf8, 0x5d, 0x90, 0x6d, 0x4e, 0xcd, 0xb9, 0x9d, 0x5a, 0x63, 0xe9, 0xd0,
	0x71, 0xfe, 0x7d, 0x26, 0xe7, 0x7a, 0x2c, 0xec, 0x52, 0x45, 0x79, 0xb8, 0x98, 0x62, 0xf8, 0x25,
	0x5a, 0x25, 0xb9, 0x6c, 0x43, 0x22, 0x35, 0x37, 0xe7, 0x55, 0x11, 0x13, 0x14, 0xef, 0xa1, 0x75,
	0x92, 0x65, 0x20, 0xc6, 0xab, 0x58, 0x50, 0x9e, 0x6b, 0x23, 0xae, 0x3f, 0x79, 0x84, 0x36, 0x43,
	0x92, 0x92, 0x80, 0xc5, 0x4c, 0x96, 0x3e, 0x4b, 0x0a, 0xae, 0xbf, 0x5c, 0x57, 0xfe, 0x1b, 0x7f,
	0x8c, 0xef, 0x47, 0xb6, 0x89, 0x20, 0x0a, 0x31, 0x44, 0x83, 0xa0, 0x67, 0x93, 0x41, 0x67, 0x23,
	0x1b, 0xde, 0x45, 0x2b, 0x1d, 0x28, 0x7d, 0x12, 0x09, 0x80, 0x2e, 0x24, 0xd2, 0x5c, 0x54, 0xce,
	0xcb, 0x1d, 0x28, 0x4f, 0x86, 0x0c, 0x1f, 0xa3, 0x7a, 0x06, 0xa2, 0x60, 0x21, 0x98, 0x48, 0xc9,
	0xb6, 0xfb, 0x94, 0x6c, 0x1f, 0x06, 0xae, 0xde, 0x30, 0x06, 0xdb, 0x68, 0x85, 0xc4, 0x19, 0xf7,
	0x3b, 0x09, 0xef, 0x25, 0x3e, 0xc9, 0xcc, 0x25, 0x95, 0x63, 0xa9, 0x82, 0x17, 0x15, 0x3b, 0xc9,
	0xec, 0x6f, 0x35, 0x84, 0xa7, 0xf5, 0xc6, 0x2d, 0xd5, 0x6c, 0xa3, 0x6a, 0xf6, 0xe9, 0x75, 0xff,
	0x7e, 0xdb, 0x2b, 0x48, 0xcc, 0x28, 0x91, 0xf0, 0xc6, 0x16, 0x70, 0x93, 0x33, 0x01, 0x74, 0x9f,
	0x32, 0xda, 0xcc, 0x45, 0x3c, 0x7c, 0x37, 0x13, 0xde, 0x4c, 0x89, 0x6c, 0x8f, 0xef, 0x6f, 0x72,
	0x10, 0xe5, 0x08, 0xf4, 0x98, 0x6c, 0x37, 0x5b, 0x82, 0x44, 0xd5, 0xe1, 0x6c, 0x35, 0x44, 0x18,
	0xcd, 0xc9, 0x32, 0x05, 0x3d, 0x56, 0x6a, 0x3d, 0x35, 0x58, 0xc6, 0xc4, 0x60, 0x7d, 0x31, 0xd0,
	0x6a, 0x9a, 0x07, 0x31, 0x0b, 0xfd, 0x4a, 0xc2, 0xcf, 0xbd, 0x8e, 0x1e, 0xaa, 0xc6, 0x53, 0xea,
	0x5c, 0x40, 0x79, 0x4d, 0xe2, 0x1c, 0xae, 0x08, 0x13, 0xa7, 0xaf, 0xfb, 0xf7, 0xdb, 0x87, 0xd3,
	0x47, 0xf2, 0xab, 0x0a, 0x79, 0x2e, 0x7d, 0x12, 0xc7, 0xc7, 0x57, 0x2a, 0xc5, 0x05, 0x94, 0x97,
	0x79, 0x2c, 0x59, 0x40, 0x32, 0xb0, 0xbd, 0xe5, 0x74, 0x08, 0xcf, 0x7b, 0x1d, 0x1c, 0xa0, 0x8d,
	0xb1, 0x3a, 0xba, 0x43, 0x37, 0x73, 0x5e, 0xc9, 0xf6, 0xaa, 0x7f, 0xbf, 0xbd, 0xff, 0xbf, 0x39,
	0xce, 0x7b, 0x1d, 0xdb, 0xc3, 0xe9, 0x74, 0xca, 0x8f, 0xa8, 0xae, 0xfb, 0xaa, 0x7f, 0x40, 0x63,
	0xf4, 0x03, 0xfe, 0x4d, 0xbb, 0x3d, 0xb4, 0xae, 0xbb, 0xef, 0x43, 0x42, 0x53, 0xce, 0x12, 0xa9,
	0x15, 0x5c, 0xd3, 0xfc, 0x9d, 0xc6, 0xa7, 0x6f, 0xbf, 0x3f, 0x58, 0xc6, 0xdd, 0x83, 0x65, 0xfc,
	0x7c, 0xb0, 0x8c, 0xaf, 0x8f, 0xd6, 0xcc, 0xdd, 0xa3, 0x35, 0xf3, 0xe3, 0xd1, 0x9a, 0xf9, 0xb4,
	0x17, 0x31, 0xd9, 0xce, 0x03, 0x27, 0xe4, 0x5d, 0x77, 0x70, 0x5b, 0xa8, 0x67, 0xb3, 0x12, 0xd4,
	0xbd, 0xd5, 0xa8, 0x4a, 0x97, 0x05, 0x0b, 0xea, 0xf6, 0x38, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0xa6, 0x99, 0x71, 0x6e, 0xb7, 0x04, 0x00, 0x00,
}

func (m *Did) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Did) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Did) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AlsoKnownAs) > 0 {
		for iNdEx := len(m.AlsoKnownAs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AlsoKnownAs[iNdEx])
			copy(dAtA[i:], m.AlsoKnownAs[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.AlsoKnownAs[iNdEx])))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.Service) > 0 {
		for iNdEx := len(m.Service) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Service[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDid(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.KeyAgreement) > 0 {
		for iNdEx := len(m.KeyAgreement) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.KeyAgreement[iNdEx])
			copy(dAtA[i:], m.KeyAgreement[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.KeyAgreement[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.CapabilityDelegation) > 0 {
		for iNdEx := len(m.CapabilityDelegation) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CapabilityDelegation[iNdEx])
			copy(dAtA[i:], m.CapabilityDelegation[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.CapabilityDelegation[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.CapabilityInvocation) > 0 {
		for iNdEx := len(m.CapabilityInvocation) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CapabilityInvocation[iNdEx])
			copy(dAtA[i:], m.CapabilityInvocation[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.CapabilityInvocation[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AssertionMethod) > 0 {
		for iNdEx := len(m.AssertionMethod) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AssertionMethod[iNdEx])
			copy(dAtA[i:], m.AssertionMethod[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.AssertionMethod[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Authentication) > 0 {
		for iNdEx := len(m.Authentication) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Authentication[iNdEx])
			copy(dAtA[i:], m.Authentication[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.Authentication[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.VerificationMethod) > 0 {
		for iNdEx := len(m.VerificationMethod) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VerificationMethod[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDid(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Controller) > 0 {
		for iNdEx := len(m.Controller) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Controller[iNdEx])
			copy(dAtA[i:], m.Controller[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.Controller[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Context) > 0 {
		for iNdEx := len(m.Context) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Context[iNdEx])
			copy(dAtA[i:], m.Context[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.Context[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *VerificationMethod) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerificationMethod) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerificationMethod) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PublicKeyMultibase) > 0 {
		i -= len(m.PublicKeyMultibase)
		copy(dAtA[i:], m.PublicKeyMultibase)
		i = encodeVarintDid(dAtA, i, uint64(len(m.PublicKeyMultibase)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PublicKeyJwk) > 0 {
		for iNdEx := len(m.PublicKeyJwk) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PublicKeyJwk[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDid(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Service) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Service) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Service) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ServiceEndpoint) > 0 {
		i -= len(m.ServiceEndpoint)
		copy(dAtA[i:], m.ServiceEndpoint)
		i = encodeVarintDid(dAtA, i, uint64(len(m.ServiceEndpoint)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDid(dAtA []byte, offset int, v uint64) int {
	offset -= sovDid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Did) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Context) > 0 {
		for _, s := range m.Context {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	if len(m.Controller) > 0 {
		for _, s := range m.Controller {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.VerificationMethod) > 0 {
		for _, e := range m.VerificationMethod {
			l = e.Size()
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.Authentication) > 0 {
		for _, s := range m.Authentication {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.AssertionMethod) > 0 {
		for _, s := range m.AssertionMethod {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.CapabilityInvocation) > 0 {
		for _, s := range m.CapabilityInvocation {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.CapabilityDelegation) > 0 {
		for _, s := range m.CapabilityDelegation {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.KeyAgreement) > 0 {
		for _, s := range m.KeyAgreement {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.Service) > 0 {
		for _, e := range m.Service {
			l = e.Size()
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.AlsoKnownAs) > 0 {
		for _, s := range m.AlsoKnownAs {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	return n
}

func (m *VerificationMethod) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	if len(m.PublicKeyJwk) > 0 {
		for _, e := range m.PublicKeyJwk {
			l = e.Size()
			n += 1 + l + sovDid(uint64(l))
		}
	}
	l = len(m.PublicKeyMultibase)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	return n
}

func (m *Service) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.ServiceEndpoint)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	return n
}

func sovDid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDid(x uint64) (n int) {
	return sovDid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Did) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Did: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Did: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Context = append(m.Context, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = append(m.Controller, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationMethod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerificationMethod = append(m.VerificationMethod, &VerificationMethod{})
			if err := m.VerificationMethod[len(m.VerificationMethod)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authentication", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authentication = append(m.Authentication, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssertionMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssertionMethod = append(m.AssertionMethod, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapabilityInvocation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CapabilityInvocation = append(m.CapabilityInvocation, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapabilityDelegation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CapabilityDelegation = append(m.CapabilityDelegation, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyAgreement", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyAgreement = append(m.KeyAgreement, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = append(m.Service, &Service{})
			if err := m.Service[len(m.Service)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlsoKnownAs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AlsoKnownAs = append(m.AlsoKnownAs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VerificationMethod) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VerificationMethod: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerificationMethod: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKeyJwk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKeyJwk = append(m.PublicKeyJwk, &KeyValuePair{})
			if err := m.PublicKeyJwk[len(m.PublicKeyJwk)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKeyMultibase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKeyMultibase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Service) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Service: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Service: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceEndpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceEndpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDid
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDid
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDid
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDid = fmt.Errorf("proto: unexpected end of group")
)
