// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/mint/v1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params holds parameters for the mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty" yaml:"mint_denom"`
	// maximum annual change in inflation rate
	InflationRateChange github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=inflation_rate_change,json=inflationRateChange,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_rate_change"`
	// maximum inflation rate
	InflationMax github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=inflation_max,json=inflationMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_max"`
	// minimum inflation rate
	InflationMin github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=inflation_min,json=inflationMin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_min"`
	// goal of percent bonded coins
	GoalBonded github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=goal_bonded,json=goalBonded,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"goal_bonded"`
	// expected blocks per year
	// TODO epochs
	BlocksPerYear uint64 `protobuf:"varint,6,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty" yaml:"blocks_per_year"`
	// distribution_proportions defines the proportion of the minted denom
	DistributionProportions DistributionProportions `protobuf:"bytes,7,opt,name=distribution_proportions,json=distributionProportions,proto3" json:"distribution_proportions" yaml:"distribution_proportions"`
	// list of funded addresses
	FundedAddresses            []WeightedAddress                       `protobuf:"bytes,8,rep,name=funded_addresses,json=fundedAddresses,proto3" json:"funded_addresses" yaml:"funded_addresses"`
	ReduceFactor               *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=reduce_factor,json=reduceFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reduce_factor,omitempty" yaml:"reduce_factor"`
	MintDistributionEpochStart uint64                                  `protobuf:"varint,10,opt,name=mint_distribution_epoch_start,json=mintDistributionEpochStart,proto3" json:"mint_distribution_epoch_start,omitempty" yaml:"mint_distribution_epoch_start"`
	GenesisEpochProvisions     *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=genesis_epoch_provisions,json=genesisEpochProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"genesis_epoch_provisions,omitempty" yaml:"genesis_epoch_provisions"`
	EpochId                    string                                  `protobuf:"bytes,12,opt,name=epoch_id,json=epochId,proto3" json:"epoch_id,omitempty" yaml:"epoch_id"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_385a01cd296b92ee, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetBlocksPerYear() uint64 {
	if m != nil {
		return m.BlocksPerYear
	}
	return 0
}

func (m *Params) GetDistributionProportions() DistributionProportions {
	if m != nil {
		return m.DistributionProportions
	}
	return DistributionProportions{}
}

func (m *Params) GetFundedAddresses() []WeightedAddress {
	if m != nil {
		return m.FundedAddresses
	}
	return nil
}

func (m *Params) GetMintDistributionEpochStart() uint64 {
	if m != nil {
		return m.MintDistributionEpochStart
	}
	return 0
}

func (m *Params) GetEpochId() string {
	if m != nil {
		return m.EpochId
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "ollo.mint.v1.Params")
}

func init() { proto.RegisterFile("ollo/mint/v1/params.proto", fileDescriptor_385a01cd296b92ee) }

var fileDescriptor_385a01cd296b92ee = []byte{
	// 642 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0xcd, 0x7c, 0xed, 0xd7, 0x36, 0x4e, 0xab, 0x82, 0xfb, 0xe7, 0x46, 0x6a, 0x26, 0x1a, 0xf1,
	0x13, 0x16, 0x9d, 0xa8, 0x85, 0x55, 0xc5, 0xa6, 0xd3, 0x80, 0xc4, 0x02, 0x29, 0x1a, 0x16, 0x08,
	0x10, 0x1a, 0x39, 0x63, 0x77, 0x62, 0x35, 0x33, 0x1e, 0xd9, 0x4e, 0xd5, 0x3e, 0x02, 0x3b, 0x36,
	0x48, 0x2c, 0x79, 0x08, 0x9e, 0x80, 0x55, 0x97, 0x15, 0x2b, 0xc4, 0x62, 0x84, 0xda, 0x37, 0xc8,
	0x13, 0x20, 0xdb, 0xd3, 0x36, 0x8d, 0xa8, 0x04, 0x52, 0x57, 0x63, 0xfb, 0x9c, 0x7b, 0xce, 0xbd,
	0x73, 0x7d, 0x0d, 0xd6, 0xf9, 0x60, 0xc0, 0xdb, 0x29, 0xcb, 0x54, 0xfb, 0x70, 0xab, 0x9d, 0x63,
	0x81, 0x53, 0xe9, 0xe7, 0x82, 0x2b, 0x0e, 0xe7, 0x35, 0xe4, 0x6b, 0xc8, 0x3f, 0xdc, 0xaa, 0x2f,
	0x27, 0x3c, 0xe1, 0x06, 0x68, 0xeb, 0x95, 0xe5, 0xd4, 0xd7, 0xae, 0x85, 0x1b, 0xae, 0x05, 0xd6,
	0x63, 0x2e, 0x53, 0x2e, 0x23, 0x1b, 0x61, 0x37, 0x16, 0xf2, 0xbe, 0x55, 0xc1, 0x4c, 0xd7, 0x18,
	0xc1, 0x27, 0x00, 0xe8, 0x98, 0x88, 0xd0, 0x8c, 0xa7, 0xc8, 0x69, 0x3a, 0xad, 0x6a, 0xb0, 0x32,
	0x2a, 0xdc, 0xbb, 0xc7, 0x38, 0x1d, 0xec, 0x78, 0x57, 0x98, 0x17, 0x56, 0xf5, 0xa6, 0xa3, 0xd7,
	0x30, 0x07, 0x2b, 0x2c, 0xdb, 0x1f, 0x60, 0xc5, 0x78, 0x16, 0x09, 0xac, 0x68, 0x14, 0xf7, 0x71,
	0x96, 0x50, 0xf4, 0x9f, 0x11, 0x78, 0x7a, 0x52, 0xb8, 0x95, 0x9f, 0x85, 0xfb, 0x20, 0x61, 0xaa,
	0x3f, 0xec, 0xf9, 0x31, 0x4f, 0xcb, 0x04, 0xca, 0xcf, 0xa6, 0x24, 0x07, 0x6d, 0x75, 0x9c, 0x53,
	0xe9, 0x77, 0x68, 0xfc, 0xfd, 0xeb, 0x26, 0x28, 0xf3, 0xeb, 0xd0, 0x38, 0x5c, 0xba, 0x94, 0x0e,
	0xb1, 0xa2, 0x7b, 0x46, 0x18, 0x62, 0xb0, 0x70, 0xe5, 0x98, 0xe2, 0x23, 0x34, 0x75, 0x0b, 0x4e,
	0xf3, 0x97, 0x92, 0x2f, 0xf1, 0xd1, 0x84, 0x05, 0xcb, 0xd0, 0xf4, 0xed, 0x5a, 0xb0, 0x0c, 0xbe,
	0x07, 0xb5, 0x84, 0xe3, 0x41, 0xd4, 0xe3, 0x19, 0xa1, 0x04, 0xfd, 0x7f, 0x0b, 0x06, 0x40, 0x0b,
	0x06, 0x46, 0x0f, 0x06, 0x60, 0xb1, 0x37, 0xe0, 0xf1, 0x81, 0x8c, 0x72, 0x2a, 0xa2, 0x63, 0x8a,
	0x05, 0x9a, 0x69, 0x3a, 0xad, 0xe9, 0xa0, 0x3e, 0x2a, 0xdc, 0x55, 0xdb, 0xd1, 0x09, 0x82, 0x17,
	0x2e, 0xd8, 0x93, 0x2e, 0x15, 0x6f, 0x28, 0x16, 0xf0, 0x83, 0x03, 0x10, 0x61, 0x52, 0x09, 0xd6,
	0x1b, 0x9a, 0x3f, 0x91, 0x0b, 0x9e, 0x73, 0xa1, 0x97, 0x12, 0xcd, 0x36, 0x9d, 0x56, 0x6d, 0xfb,
	0xbe, 0x3f, 0x7e, 0x2f, 0xfd, 0xce, 0x18, 0xbb, 0x7b, 0x45, 0x0e, 0x1e, 0xea, 0xba, 0x46, 0x85,
	0xeb, 0x5a, 0xe3, 0x9b, 0x44, 0xbd, 0x70, 0x8d, 0xfc, 0x59, 0x01, 0x32, 0x70, 0x67, 0x7f, 0xa8,
	0x2b, 0x8b, 0x30, 0x21, 0x82, 0x4a, 0x49, 0x25, 0x9a, 0x6b, 0x4e, 0xb5, 0x6a, 0xdb, 0x1b, 0xd7,
	0x53, 0x78, 0x4d, 0x59, 0xd2, 0x57, 0x94, 0xec, 0x5a, 0x5a, 0xe0, 0x96, 0xd6, 0x6b, 0xd6, 0x7a,
	0x52, 0xc4, 0x0b, 0x17, 0xed, 0xd1, 0xee, 0xc5, 0x09, 0x1c, 0x82, 0x05, 0x41, 0xc9, 0x30, 0xa6,
	0xd1, 0x3e, 0x8e, 0x15, 0x17, 0xa8, 0x6a, 0x7a, 0xd3, 0xfd, 0xfb, 0xbe, 0x8c, 0x0a, 0x77, 0xd9,
	0xda, 0x5d, 0x13, 0xf2, 0x26, 0x2f, 0x84, 0x45, 0x9f, 0x1b, 0x10, 0x1e, 0x80, 0x0d, 0x3b, 0x62,
	0xe3, 0x3f, 0x87, 0xe6, 0x3c, 0xee, 0x47, 0x52, 0x61, 0xa1, 0x10, 0x30, 0xfd, 0x6b, 0x8d, 0x0a,
	0xf7, 0xde, 0xf8, 0x44, 0xde, 0x40, 0xf7, 0xc2, 0xba, 0x19, 0xd2, 0x31, 0xf8, 0x99, 0x46, 0x5f,
	0x69, 0x10, 0x7e, 0x72, 0x00, 0x4a, 0x68, 0x46, 0x25, 0x93, 0x65, 0x50, 0x2e, 0xf8, 0x21, 0x93,
	0xa6, 0xb5, 0x35, 0x53, 0xef, 0xbb, 0x7f, 0xaa, 0xb7, 0xec, 0xec, 0x4d, 0x9a, 0x93, 0xa5, 0xaf,
	0x96, 0x44, 0x93, 0x52, 0xf7, 0x92, 0x06, 0x7d, 0x30, 0x67, 0x43, 0x19, 0x41, 0xf3, 0x26, 0x8d,
	0xa5, 0x51, 0xe1, 0x2e, 0x5a, 0xf1, 0x0b, 0xc4, 0x0b, 0x67, 0xcd, 0xf2, 0x05, 0xd9, 0x99, 0xfe,
	0xfc, 0xc5, 0xad, 0x04, 0x7b, 0x27, 0x67, 0x0d, 0xe7, 0xf4, 0xac, 0xe1, 0xfc, 0x3a, 0x6b, 0x38,
	0x1f, 0xcf, 0x1b, 0x95, 0xd3, 0xf3, 0x46, 0xe5, 0xc7, 0x79, 0xa3, 0xf2, 0xf6, 0xd1, 0x58, 0x01,
	0xfa, 0x9a, 0x6c, 0x4a, 0x65, 0x26, 0xd0, 0x6c, 0xda, 0x47, 0xf6, 0xb1, 0x34, 0x75, 0xf4, 0x66,
	0xcc, 0x83, 0xf8, 0xf8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x3f, 0x5f, 0xa2, 0x85, 0x05,
	0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EpochId) > 0 {
		i -= len(m.EpochId)
		copy(dAtA[i:], m.EpochId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.EpochId)))
		i--
		dAtA[i] = 0x62
	}
	if m.GenesisEpochProvisions != nil {
		{
			size := m.GenesisEpochProvisions.Size()
			i -= size
			if _, err := m.GenesisEpochProvisions.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if m.MintDistributionEpochStart != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MintDistributionEpochStart))
		i--
		dAtA[i] = 0x50
	}
	if m.ReduceFactor != nil {
		{
			size := m.ReduceFactor.Size()
			i -= size
			if _, err := m.ReduceFactor.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if len(m.FundedAddresses) > 0 {
		for iNdEx := len(m.FundedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FundedAddresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	{
		size, err := m.DistributionProportions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.BlocksPerYear != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BlocksPerYear))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.GoalBonded.Size()
		i -= size
		if _, err := m.GoalBonded.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.InflationMin.Size()
		i -= size
		if _, err := m.InflationMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.InflationMax.Size()
		i -= size
		if _, err := m.InflationMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.InflationRateChange.Size()
		i -= size
		if _, err := m.InflationRateChange.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.InflationRateChange.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.InflationMax.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.InflationMin.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.GoalBonded.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.BlocksPerYear != 0 {
		n += 1 + sovParams(uint64(m.BlocksPerYear))
	}
	l = m.DistributionProportions.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.FundedAddresses) > 0 {
		for _, e := range m.FundedAddresses {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.ReduceFactor != nil {
		l = m.ReduceFactor.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MintDistributionEpochStart != 0 {
		n += 1 + sovParams(uint64(m.MintDistributionEpochStart))
	}
	if m.GenesisEpochProvisions != nil {
		l = m.GenesisEpochProvisions.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.EpochId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRateChange", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationRateChange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoalBonded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GoalBonded.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerYear", wireType)
			}
			m.BlocksPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionProportions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DistributionProportions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FundedAddresses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FundedAddresses = append(m.FundedAddresses, WeightedAddress{})
			if err := m.FundedAddresses[len(m.FundedAddresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReduceFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.ReduceFactor = &v
			if err := m.ReduceFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDistributionEpochStart", wireType)
			}
			m.MintDistributionEpochStart = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MintDistributionEpochStart |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisEpochProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.GenesisEpochProvisions = &v
			if err := m.GenesisEpochProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
