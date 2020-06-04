// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protob/signature.proto

package common

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

//
// State object for signatures, either partial (for offline/async "one round" signing) or full (contains the final signature).
type SignatureData struct {
	OneRoundData *SignatureData_OneRoundData `protobuf:"bytes,6,opt,name=one_round_data,json=oneRoundData,proto3" json:"one_round_data,omitempty"`
	Signature    []byte                      `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Ethereum-style recovery byte; only the first byte is relevant
	SignatureRecovery []byte `protobuf:"bytes,2,opt,name=signature_recovery,json=signatureRecovery,proto3" json:"signature_recovery,omitempty"`
	// Signature components R, S
	R []byte `protobuf:"bytes,3,opt,name=r,proto3" json:"r,omitempty"`
	S []byte `protobuf:"bytes,4,opt,name=s,proto3" json:"s,omitempty"`
	// M represents the original message digest that was signed M
	M                    []byte   `protobuf:"bytes,5,opt,name=m,proto3" json:"m,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignatureData) Reset()         { *m = SignatureData{} }
func (m *SignatureData) String() string { return proto.CompactTextString(m) }
func (*SignatureData) ProtoMessage()    {}
func (*SignatureData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73ad5e44f61b3fb, []int{0}
}

func (m *SignatureData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureData.Unmarshal(m, b)
}
func (m *SignatureData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureData.Marshal(b, m, deterministic)
}
func (m *SignatureData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureData.Merge(m, src)
}
func (m *SignatureData) XXX_Size() int {
	return xxx_messageInfo_SignatureData.Size(m)
}
func (m *SignatureData) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureData.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureData proto.InternalMessageInfo

func (m *SignatureData) GetOneRoundData() *SignatureData_OneRoundData {
	if m != nil {
		return m.OneRoundData
	}
	return nil
}

func (m *SignatureData) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignatureData) GetSignatureRecovery() []byte {
	if m != nil {
		return m.SignatureRecovery
	}
	return nil
}

func (m *SignatureData) GetR() []byte {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *SignatureData) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *SignatureData) GetM() []byte {
	if m != nil {
		return m.M
	}
	return nil
}

type SignatureData_OneRoundData struct {
	// Sanity check in FinalizeGetAndVerifyFinalSig
	T int32 `protobuf:"varint,1,opt,name=t,proto3" json:"t,omitempty"`
	// Components to produce s = sum(s_i)
	KI     []byte   `protobuf:"bytes,2,opt,name=k_i,json=kI,proto3" json:"k_i,omitempty"`
	SigmaI []byte   `protobuf:"bytes,3,opt,name=sigma_i,json=sigmaI,proto3" json:"sigma_i,omitempty"`
	BigR   *ECPoint `protobuf:"bytes,4,opt,name=big_r,json=bigR,proto3" json:"big_r,omitempty"`
	// Components for identifiable aborts during the final phase
	BigRBarJ             map[string]*ECPoint `protobuf:"bytes,5,rep,name=big_r_bar_j,json=bigRBarJ,proto3" json:"big_r_bar_j,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BigSJ                map[string]*ECPoint `protobuf:"bytes,6,rep,name=big_s_j,json=bigSJ,proto3" json:"big_s_j,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SignatureData_OneRoundData) Reset()         { *m = SignatureData_OneRoundData{} }
func (m *SignatureData_OneRoundData) String() string { return proto.CompactTextString(m) }
func (*SignatureData_OneRoundData) ProtoMessage()    {}
func (*SignatureData_OneRoundData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73ad5e44f61b3fb, []int{0, 0}
}

func (m *SignatureData_OneRoundData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureData_OneRoundData.Unmarshal(m, b)
}
func (m *SignatureData_OneRoundData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureData_OneRoundData.Marshal(b, m, deterministic)
}
func (m *SignatureData_OneRoundData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureData_OneRoundData.Merge(m, src)
}
func (m *SignatureData_OneRoundData) XXX_Size() int {
	return xxx_messageInfo_SignatureData_OneRoundData.Size(m)
}
func (m *SignatureData_OneRoundData) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureData_OneRoundData.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureData_OneRoundData proto.InternalMessageInfo

func (m *SignatureData_OneRoundData) GetT() int32 {
	if m != nil {
		return m.T
	}
	return 0
}

func (m *SignatureData_OneRoundData) GetKI() []byte {
	if m != nil {
		return m.KI
	}
	return nil
}

func (m *SignatureData_OneRoundData) GetSigmaI() []byte {
	if m != nil {
		return m.SigmaI
	}
	return nil
}

func (m *SignatureData_OneRoundData) GetBigR() *ECPoint {
	if m != nil {
		return m.BigR
	}
	return nil
}

func (m *SignatureData_OneRoundData) GetBigRBarJ() map[string]*ECPoint {
	if m != nil {
		return m.BigRBarJ
	}
	return nil
}

func (m *SignatureData_OneRoundData) GetBigSJ() map[string]*ECPoint {
	if m != nil {
		return m.BigSJ
	}
	return nil
}

func init() {
	proto.RegisterType((*SignatureData)(nil), "SignatureData")
	proto.RegisterType((*SignatureData_OneRoundData)(nil), "SignatureData.OneRoundData")
	proto.RegisterMapType((map[string]*ECPoint)(nil), "SignatureData.OneRoundData.BigRBarJEntry")
	proto.RegisterMapType((map[string]*ECPoint)(nil), "SignatureData.OneRoundData.BigSJEntry")
}

func init() { proto.RegisterFile("protob/signature.proto", fileDescriptor_a73ad5e44f61b3fb) }

var fileDescriptor_a73ad5e44f61b3fb = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4b, 0x6b, 0xe3, 0x30,
	0x14, 0x85, 0xb1, 0x1d, 0x3b, 0xc9, 0xb5, 0x33, 0x0f, 0x0d, 0xcc, 0x08, 0xcf, 0x83, 0x30, 0x8b,
	0x21, 0xb3, 0x18, 0x07, 0x32, 0x9b, 0x61, 0x16, 0x85, 0xba, 0x0d, 0x34, 0xd9, 0xb4, 0x28, 0xbb,
	0x6e, 0x84, 0x9c, 0x08, 0x57, 0x49, 0x6d, 0x15, 0x59, 0x09, 0xe4, 0x0f, 0xf6, 0x47, 0x75, 0x55,
	0x24, 0x3b, 0xaf, 0x4d, 0x0b, 0xdd, 0xf9, 0x3b, 0xe7, 0x72, 0xce, 0xbd, 0x46, 0xf0, 0xf9, 0x41,
	0x49, 0x2d, 0xb3, 0x61, 0x25, 0xf2, 0x92, 0xe9, 0xb5, 0xe2, 0x89, 0x15, 0xe2, 0x4f, 0x3b, 0xfd,
	0x8e, 0x29, 0xbe, 0xa8, 0xc5, 0x9f, 0x8f, 0x2d, 0xe8, 0xcd, 0x76, 0x83, 0x97, 0x4c, 0x33, 0x74,
	0x0e, 0xef, 0x64, 0xc9, 0xa9, 0x92, 0xeb, 0x72, 0x41, 0x17, 0x4c, 0x33, 0x1c, 0xf4, 0x9d, 0x41,
	0x38, 0xfa, 0x9a, 0x9c, 0xcc, 0x25, 0xd7, 0x25, 0x27, 0x66, 0xc6, 0x00, 0x89, 0xe4, 0x11, 0xa1,
	0x6f, 0xd0, 0xdd, 0x97, 0x63, 0xa7, 0xef, 0x0c, 0x22, 0x72, 0x10, 0xd0, 0x1f, 0x40, 0x7b, 0xa0,
	0x8a, 0xcf, 0xe5, 0x86, 0xab, 0x2d, 0x76, 0xed, 0xd8, 0xc7, 0xbd, 0x43, 0x1a, 0x03, 0x45, 0xe0,
	0x28, 0xec, 0x59, 0xd7, 0x51, 0x86, 0x2a, 0xdc, 0xaa, 0xa9, 0x32, 0x54, 0x60, 0xbf, 0xa6, 0x22,
	0x7e, 0x72, 0x21, 0x3a, 0xde, 0xca, 0xd8, 0xda, 0xf6, 0xfb, 0xc4, 0xd1, 0xe8, 0x3d, 0x78, 0x2b,
	0x2a, 0x9a, 0x22, 0x77, 0x35, 0x41, 0x5f, 0xa0, 0x5d, 0x89, 0xbc, 0x60, 0x54, 0x34, 0xf9, 0x81,
	0xc5, 0x09, 0xfa, 0x0e, 0x7e, 0x26, 0x72, 0xaa, 0x6c, 0x51, 0x38, 0xea, 0x24, 0xe3, 0x8b, 0x1b,
	0x29, 0x4a, 0x4d, 0x5a, 0x99, 0xc8, 0x09, 0xba, 0x82, 0xd0, 0xda, 0x34, 0x63, 0x8a, 0x2e, 0xb1,
	0xdf, 0xf7, 0x06, 0xe1, 0xe8, 0xf7, 0x0b, 0xbf, 0x27, 0x49, 0x45, 0x4e, 0x52, 0xa6, 0xa6, 0xe3,
	0x52, 0xab, 0x2d, 0xe9, 0x64, 0x0d, 0xa2, 0x33, 0x68, 0x9b, 0xa4, 0x8a, 0x2e, 0x71, 0x60, 0x53,
	0x7e, 0xbd, 0x92, 0x32, 0x6b, 0x22, 0xcc, 0x7e, 0xb3, 0x69, 0x3c, 0x86, 0xde, 0x49, 0x34, 0xfa,
	0x00, 0xde, 0x8a, 0x6f, 0xed, 0xcd, 0x5d, 0x62, 0x3e, 0xd1, 0x0f, 0xf0, 0x37, 0xec, 0x7e, 0xcd,
	0xed, 0xdd, 0xc7, 0xb7, 0xd4, 0xf2, 0x7f, 0xf7, 0x9f, 0x13, 0xa7, 0x00, 0x87, 0xec, 0xb7, 0x65,
	0xa4, 0x70, 0xdb, 0x49, 0x86, 0x73, 0x59, 0x14, 0xb2, 0xcc, 0x02, 0xfb, 0xb6, 0xfe, 0x3e, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x99, 0xbf, 0x03, 0x53, 0x8a, 0x02, 0x00, 0x00,
}
