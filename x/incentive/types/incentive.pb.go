// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/incentive/v1/incentive.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Incentive struct {
	//
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	//
	Claimable github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=claimable,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"claimable" yaml:"claimable"`
}

func (m *Incentive) Reset()         { *m = Incentive{} }
func (m *Incentive) String() string { return proto.CompactTextString(m) }
func (*Incentive) ProtoMessage()    {}
func (*Incentive) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc2ca0d558551425, []int{0}
}
func (m *Incentive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Incentive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Incentive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Incentive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Incentive.Merge(m, src)
}
func (m *Incentive) XXX_Size() int {
	return m.Size()
}
func (m *Incentive) XXX_DiscardUnknown() {
	xxx_messageInfo_Incentive.DiscardUnknown(m)
}

var xxx_messageInfo_Incentive proto.InternalMessageInfo

type IncentiveProps struct {
	//
	MinOpenRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=min_open_ratio,json=minOpenRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_open_ratio" yaml:"min_open_ratio"`
	//
	MinOpenDepthRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=min_open_depth_ratio,json=minOpenDepthRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_open_depth_ratio" yaml:"min_open_ratio"`
	//
	MaxDowntime uint32 `protobuf:"varint,3,opt,name=max_downtime,json=maxDowntime,proto3" json:"max_downtime,omitempty" yaml:"max_downtime"`
	//
	MaxTotalDowntime uint32 `protobuf:"varint,4,opt,name=max_total_downtime,json=maxTotalDowntime,proto3" json:"max_total_downtime,omitempty" yaml:"max_total_downtime"`
	//
	MinHours uint32 `protobuf:"varint,5,opt,name=min_hours,json=minHours,proto3" json:"min_hours,omitempty" yaml:"min_hours"`
	//
	MinDays uint32 `protobuf:"varint,6,opt,name=min_days,json=minDays,proto3" json:"min_days,omitempty" yaml:"min_days"`
}

func (m *IncentiveProps) Reset()         { *m = IncentiveProps{} }
func (m *IncentiveProps) String() string { return proto.CompactTextString(m) }
func (*IncentiveProps) ProtoMessage()    {}
func (*IncentiveProps) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc2ca0d558551425, []int{1}
}
func (m *IncentiveProps) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentiveProps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentiveProps.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentiveProps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentiveProps.Merge(m, src)
}
func (m *IncentiveProps) XXX_Size() int {
	return m.Size()
}
func (m *IncentiveProps) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentiveProps.DiscardUnknown(m)
}

var xxx_messageInfo_IncentiveProps proto.InternalMessageInfo

type IncentivePair struct {
	//
	PairId uint64 `protobuf:"varint,1,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty" yaml:"pair_id"`
	//
	UpdatedAt *time.Time `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3,stdtime" json:"updated_at,omitempty" yaml:"updated_at"`
	//
	IncentiveWeight github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=incentive_weight,json=incentiveWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"incentive_weight" yaml:"incentive_weight"`
	//
	MaxSpread github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=max_spread,json=maxSpread,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_spread" yaml:"max_spread"`
	//
	MinSpread github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=min_spread,json=minSpread,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_spread" yaml:"max_spread"`
	//
	MinDepth github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=min_depth,json=minDepth,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_depth" yaml:"max_spread"`
}

func (m *IncentivePair) Reset()         { *m = IncentivePair{} }
func (m *IncentivePair) String() string { return proto.CompactTextString(m) }
func (*IncentivePair) ProtoMessage()    {}
func (*IncentivePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc2ca0d558551425, []int{2}
}
func (m *IncentivePair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentivePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentivePair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentivePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentivePair.Merge(m, src)
}
func (m *IncentivePair) XXX_Size() int {
	return m.Size()
}
func (m *IncentivePair) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentivePair.DiscardUnknown(m)
}

var xxx_messageInfo_IncentivePair proto.InternalMessageInfo

type IncentiveDistribution struct {
	//
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	//
	PairId uint64                                   `protobuf:"varint,2,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty" yaml:"pair_id"`
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount" yaml:"amount"`
}

func (m *IncentiveDistribution) Reset()         { *m = IncentiveDistribution{} }
func (m *IncentiveDistribution) String() string { return proto.CompactTextString(m) }
func (*IncentiveDistribution) ProtoMessage()    {}
func (*IncentiveDistribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc2ca0d558551425, []int{3}
}
func (m *IncentiveDistribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentiveDistribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentiveDistribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentiveDistribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentiveDistribution.Merge(m, src)
}
func (m *IncentiveDistribution) XXX_Size() int {
	return m.Size()
}
func (m *IncentiveDistribution) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentiveDistribution.DiscardUnknown(m)
}

var xxx_messageInfo_IncentiveDistribution proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Incentive)(nil), "ollo.incentive.v1.Incentive")
	proto.RegisterType((*IncentiveProps)(nil), "ollo.incentive.v1.IncentiveProps")
	proto.RegisterType((*IncentivePair)(nil), "ollo.incentive.v1.IncentivePair")
	proto.RegisterType((*IncentiveDistribution)(nil), "ollo.incentive.v1.IncentiveDistribution")
}

func init() { proto.RegisterFile("ollo/incentive/v1/incentive.proto", fileDescriptor_cc2ca0d558551425) }

var fileDescriptor_cc2ca0d558551425 = []byte{
	// 779 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xbf, 0x6f, 0xeb, 0x44,
	0x1c, 0x8f, 0xd3, 0xbc, 0xf4, 0xf9, 0xfa, 0xda, 0xd7, 0xf8, 0xa5, 0x6a, 0x12, 0x09, 0xbb, 0x78,
	0x40, 0x11, 0xa8, 0xb6, 0x52, 0xb6, 0x6c, 0x75, 0x23, 0x41, 0x05, 0x12, 0xc8, 0xad, 0x84, 0xc4,
	0x62, 0xce, 0xb1, 0x49, 0x4e, 0xe4, 0xee, 0x2c, 0xdf, 0xa5, 0x4d, 0x06, 0x76, 0xc4, 0xd4, 0x3f,
	0xa1, 0x33, 0x33, 0x03, 0x7f, 0x42, 0xc7, 0x8a, 0x09, 0x75, 0x70, 0xa1, 0x5d, 0x90, 0xd8, 0xf2,
	0x17, 0xa0, 0x3b, 0x5f, 0xec, 0x14, 0x86, 0xbe, 0x4a, 0x99, 0xf2, 0xfd, 0xf9, 0xf9, 0x7c, 0x2f,
	0xdf, 0x1f, 0x06, 0x1f, 0xd2, 0xc9, 0x84, 0xba, 0x88, 0x0c, 0x63, 0xc2, 0xd1, 0x45, 0xec, 0x5e,
	0xf4, 0x4a, 0xc5, 0x49, 0x52, 0xca, 0xa9, 0xd1, 0x10, 0x21, 0x4e, 0x69, 0xbd, 0xe8, 0x75, 0xda,
	0x43, 0xca, 0x30, 0x65, 0x81, 0x0c, 0x70, 0x73, 0x25, 0x8f, 0xee, 0x98, 0xb9, 0xe6, 0x86, 0x90,
	0x09, 0xb4, 0x30, 0xe6, 0xb0, 0xe7, 0x0e, 0x29, 0x22, 0xca, 0xdf, 0x1c, 0xd1, 0x11, 0xcd, 0xf3,
	0x84, 0xa4, 0xac, 0xd6, 0x88, 0xd2, 0xd1, 0x24, 0x76, 0xa5, 0x16, 0x4e, 0xbf, 0x77, 0x39, 0xc2,
	0x31, 0xe3, 0x10, 0x27, 0x2a, 0xa0, 0xf3, 0xff, 0x3a, 0x31, 0xce, 0x7d, 0xf6, 0x9d, 0x06, 0xf4,
	0xd3, 0xa5, 0xc7, 0x18, 0x80, 0x4d, 0x18, 0x45, 0x69, 0xcc, 0x58, 0x4b, 0x3b, 0xd0, 0xba, 0xba,
	0xf7, 0xf1, 0x22, 0xb3, 0x76, 0xe6, 0x10, 0x4f, 0xfa, 0xb6, 0x72, 0xd8, 0xbf, 0xff, 0x7a, 0xd8,
	0x54, 0x55, 0x1f, 0xe7, 0xa6, 0x33, 0x9e, 0x22, 0x32, 0xf2, 0x97, 0xa9, 0xc6, 0x8f, 0x40, 0x1f,
	0x4e, 0x20, 0xc2, 0x30, 0x9c, 0xc4, 0xad, 0xea, 0xc1, 0x46, 0x77, 0xeb, 0xa8, 0xed, 0xa8, 0x14,
	0xf1, 0x34, 0x47, 0x3d, 0xcd, 0x39, 0xa1, 0x88, 0x78, 0x83, 0x9b, 0xcc, 0xaa, 0x2c, 0x32, 0x6b,
	0x37, 0xa7, 0x29, 0x32, 0xed, 0x5f, 0xee, 0xad, 0xee, 0x08, 0xf1, 0xf1, 0x34, 0x74, 0x86, 0x14,
	0xab, 0x7f, 0x4a, 0xfd, 0x1c, 0xb2, 0xe8, 0x07, 0x97, 0xcf, 0x93, 0x98, 0x49, 0x10, 0xe6, 0x97,
	0x8c, 0xfd, 0xd7, 0x3f, 0x5d, 0x5b, 0x95, 0xbf, 0xaf, 0xad, 0x8a, 0xfd, 0xcf, 0x06, 0xd8, 0x29,
	0x1e, 0xf7, 0x75, 0x4a, 0x13, 0x66, 0x60, 0xb0, 0x83, 0x11, 0x09, 0x68, 0x12, 0x93, 0x20, 0x85,
	0x1c, 0x51, 0xf5, 0xd0, 0xcf, 0x44, 0x15, 0x77, 0x99, 0xf5, 0xd1, 0x7b, 0x30, 0x0e, 0xe2, 0xe1,
	0x22, 0xb3, 0xf6, 0xf2, 0x7a, 0x9f, 0xa2, 0xd9, 0xfe, 0x1b, 0x8c, 0xc8, 0x57, 0x49, 0x4c, 0x7c,
	0xa1, 0x1a, 0x33, 0xd0, 0x2c, 0x02, 0xa2, 0x38, 0xe1, 0x63, 0x45, 0x5a, 0x5d, 0x2f, 0x69, 0x43,
	0x91, 0x0e, 0x04, 0x45, 0xce, 0xdc, 0x07, 0x6f, 0x30, 0x9c, 0x05, 0x11, 0xbd, 0x24, 0x62, 0x1e,
	0x5a, 0x1b, 0x07, 0x5a, 0x77, 0xdb, 0xdb, 0x5f, 0x64, 0xd6, 0x3b, 0x85, 0xb1, 0xe2, 0xb5, 0xfd,
	0x2d, 0x0c, 0x67, 0x03, 0xa5, 0x19, 0x5f, 0x00, 0x43, 0x78, 0x39, 0xe5, 0x70, 0x52, 0x22, 0xd4,
	0x24, 0xc2, 0x07, 0x8b, 0xcc, 0x6a, 0x97, 0x08, 0x4f, 0x63, 0x6c, 0x7f, 0x17, 0xc3, 0xd9, 0xb9,
	0xb0, 0x15, 0x60, 0x3d, 0xa0, 0x8b, 0x72, 0xc7, 0x74, 0x9a, 0xb2, 0xd6, 0x2b, 0x89, 0xd1, 0x2c,
	0xdb, 0x5d, 0xb8, 0x6c, 0xff, 0x35, 0x46, 0xe4, 0x73, 0x21, 0x1a, 0x0e, 0x10, 0x72, 0x10, 0xc1,
	0x39, 0x6b, 0xd5, 0x65, 0xc6, 0xbb, 0x45, 0x66, 0xbd, 0x2d, 0x33, 0x84, 0xc7, 0xf6, 0x37, 0x31,
	0x22, 0x03, 0x38, 0x67, 0xfd, 0x9a, 0xe8, 0xb8, 0xfd, 0x5b, 0x0d, 0x6c, 0x97, 0xdd, 0x86, 0x28,
	0x35, 0x3e, 0x01, 0x9b, 0x09, 0x44, 0x69, 0x80, 0x22, 0xd9, 0xe5, 0x9a, 0x67, 0x94, 0xe3, 0xac,
	0x1c, 0xb6, 0x5f, 0x17, 0xd2, 0x69, 0x64, 0x9c, 0x03, 0x30, 0x4d, 0x22, 0xc8, 0xe3, 0x28, 0x80,
	0x5c, 0x36, 0x68, 0xeb, 0xa8, 0xe3, 0xe4, 0xbb, 0xe5, 0x2c, 0x77, 0xcb, 0x39, 0x5f, 0xee, 0x96,
	0xd7, 0x5e, 0x64, 0x56, 0x23, 0xc7, 0x2a, 0xf3, 0xec, 0xab, 0x7b, 0x4b, 0xf3, 0x75, 0x65, 0x38,
	0xe6, 0x06, 0x07, 0xbb, 0xc5, 0xe2, 0x05, 0x97, 0x31, 0x1a, 0x8d, 0xb9, 0x6c, 0x85, 0xee, 0x9d,
	0xbe, 0xb8, 0xf9, 0xfb, 0x39, 0xdb, 0x7f, 0xf1, 0x6c, 0xff, 0x6d, 0x61, 0xfa, 0x46, 0x5a, 0x8c,
	0x10, 0x00, 0xd1, 0x1c, 0x96, 0xa4, 0x31, 0x8c, 0x64, 0xe3, 0x74, 0xef, 0xe4, 0xc5, 0x7c, 0x8d,
	0xb2, 0xcd, 0x39, 0x92, 0xed, 0xeb, 0x18, 0xce, 0xce, 0xa4, 0x2c, 0x39, 0x10, 0x59, 0x72, 0xbc,
	0x5a, 0x27, 0x07, 0x22, 0x8a, 0xe3, 0xbb, 0x7c, 0x76, 0xe4, 0xe6, 0xc8, 0x49, 0x58, 0x13, 0x85,
	0x18, 0x2f, 0xb9, 0x2b, 0xfd, 0x9a, 0x3c, 0x14, 0x3f, 0x57, 0xc1, 0x5e, 0x31, 0x3a, 0x03, 0xc4,
	0x78, 0x8a, 0xc2, 0x29, 0x47, 0x94, 0xac, 0xe9, 0x22, 0xae, 0x0c, 0x62, 0xf5, 0xd9, 0x41, 0xe4,
	0xa0, 0x0e, 0x31, 0x9d, 0x12, 0x31, 0x28, 0xcf, 0xdc, 0xce, 0x63, 0x75, 0x3b, 0xb7, 0x55, 0x41,
	0x32, 0xed, 0x65, 0x87, 0x53, 0x71, 0x95, 0x57, 0xd3, 0xfb, 0xf2, 0xe6, 0x2f, 0xb3, 0x72, 0xf3,
	0x60, 0x6a, 0xb7, 0x0f, 0xa6, 0xf6, 0xe7, 0x83, 0xa9, 0x5d, 0x3d, 0x9a, 0x95, 0xdb, 0x47, 0xb3,
	0xf2, 0xc7, 0xa3, 0x59, 0xf9, 0xd6, 0x59, 0x41, 0x16, 0xdf, 0x95, 0x43, 0xc6, 0xc5, 0xbd, 0x21,
	0x52, 0x71, 0x67, 0x2b, 0x9f, 0x19, 0xc9, 0x12, 0xd6, 0xe5, 0xea, 0x7c, 0xfa, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xdb, 0x1e, 0x91, 0xe3, 0x2d, 0x07, 0x00, 0x00,
}

func (m *Incentive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Incentive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Incentive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Claimable) > 0 {
		for iNdEx := len(m.Claimable) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Claimable[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentive(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintIncentive(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IncentiveProps) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentiveProps) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentiveProps) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinDays != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.MinDays))
		i--
		dAtA[i] = 0x30
	}
	if m.MinHours != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.MinHours))
		i--
		dAtA[i] = 0x28
	}
	if m.MaxTotalDowntime != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.MaxTotalDowntime))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxDowntime != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.MaxDowntime))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.MinOpenDepthRatio.Size()
		i -= size
		if _, err := m.MinOpenDepthRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinOpenRatio.Size()
		i -= size
		if _, err := m.MinOpenRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IncentivePair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentivePair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentivePair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MinDepth.Size()
		i -= size
		if _, err := m.MinDepth.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.MinSpread.Size()
		i -= size
		if _, err := m.MinSpread.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.MaxSpread.Size()
		i -= size
		if _, err := m.MaxSpread.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.IncentiveWeight.Size()
		i -= size
		if _, err := m.IncentiveWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.UpdatedAt != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.UpdatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.UpdatedAt):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintIncentive(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x12
	}
	if m.PairId != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *IncentiveDistribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentiveDistribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentiveDistribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentive(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PairId != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintIncentive(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIncentive(dAtA []byte, offset int, v uint64) int {
	offset -= sovIncentive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Incentive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovIncentive(uint64(l))
	}
	if len(m.Claimable) > 0 {
		for _, e := range m.Claimable {
			l = e.Size()
			n += 1 + l + sovIncentive(uint64(l))
		}
	}
	return n
}

func (m *IncentiveProps) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinOpenRatio.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.MinOpenDepthRatio.Size()
	n += 1 + l + sovIncentive(uint64(l))
	if m.MaxDowntime != 0 {
		n += 1 + sovIncentive(uint64(m.MaxDowntime))
	}
	if m.MaxTotalDowntime != 0 {
		n += 1 + sovIncentive(uint64(m.MaxTotalDowntime))
	}
	if m.MinHours != 0 {
		n += 1 + sovIncentive(uint64(m.MinHours))
	}
	if m.MinDays != 0 {
		n += 1 + sovIncentive(uint64(m.MinDays))
	}
	return n
}

func (m *IncentivePair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PairId != 0 {
		n += 1 + sovIncentive(uint64(m.PairId))
	}
	if m.UpdatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.UpdatedAt)
		n += 1 + l + sovIncentive(uint64(l))
	}
	l = m.IncentiveWeight.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.MaxSpread.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.MinSpread.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.MinDepth.Size()
	n += 1 + l + sovIncentive(uint64(l))
	return n
}

func (m *IncentiveDistribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovIncentive(uint64(l))
	}
	if m.PairId != 0 {
		n += 1 + sovIncentive(uint64(m.PairId))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovIncentive(uint64(l))
		}
	}
	return n
}

func sovIncentive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIncentive(x uint64) (n int) {
	return sovIncentive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Incentive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: Incentive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Incentive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claimable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Claimable = append(m.Claimable, types.Coin{})
			if err := m.Claimable[len(m.Claimable)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func (m *IncentiveProps) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: IncentiveProps: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentiveProps: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinOpenRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinOpenRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinOpenDepthRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinOpenDepthRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDowntime", wireType)
			}
			m.MaxDowntime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxDowntime |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTotalDowntime", wireType)
			}
			m.MaxTotalDowntime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTotalDowntime |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinHours", wireType)
			}
			m.MinHours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinHours |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDays", wireType)
			}
			m.MinDays = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinDays |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func (m *IncentivePair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: IncentivePair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentivePair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpdatedAt == nil {
				m.UpdatedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncentiveWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IncentiveWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSpread", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxSpread.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSpread", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinSpread.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDepth", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinDepth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func (m *IncentiveDistribution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: IncentiveDistribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentiveDistribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func skipIncentive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
				return 0, ErrInvalidLengthIncentive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIncentive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIncentive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIncentive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIncentive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIncentive = fmt.Errorf("proto: unexpected end of group")
)
