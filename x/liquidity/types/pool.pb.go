// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/liquidity/v1/pool.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

type PoolA struct {
	Type                  *PoolType                               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id                    uint64                                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	PairId                uint64                                  `protobuf:"varint,3,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty"`
	Creator               string                                  `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	ReserveAddress        string                                  `protobuf:"bytes,5,opt,name=reserve_address,json=reserveAddress,proto3" json:"reserve_address,omitempty"`
	PoolCoinDenom         string                                  `protobuf:"bytes,6,opt,name=pool_coin_denom,json=poolCoinDenom,proto3" json:"pool_coin_denom,omitempty"`
	MinPrice              *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=min_price,json=minPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_price,omitempty"`
	MaxPrice              *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=max_price,json=maxPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_price,omitempty"`
	LastDepositRequestId  uint64                                  `protobuf:"varint,9,opt,name=last_deposit_request_id,json=lastDepositRequestId,proto3" json:"last_deposit_request_id,omitempty"`
	LastWithdrawRequestId uint64                                  `protobuf:"varint,10,opt,name=last_withdraw_request_id,json=lastWithdrawRequestId,proto3" json:"last_withdraw_request_id,omitempty"`
	Disabled              bool                                    `protobuf:"varint,11,opt,name=disabled,proto3" json:"disabled,omitempty"`
}

func (m *PoolA) Reset()         { *m = PoolA{} }
func (m *PoolA) String() string { return proto.CompactTextString(m) }
func (*PoolA) ProtoMessage()    {}
func (*PoolA) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e204054048950d3, []int{0}
}
func (m *PoolA) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolA.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolA.Merge(m, src)
}
func (m *PoolA) XXX_Size() int {
	return m.Size()
}
func (m *PoolA) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolA.DiscardUnknown(m)
}

var xxx_messageInfo_PoolA proto.InternalMessageInfo

func (m *PoolA) GetType() *PoolType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *PoolA) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PoolA) GetPairId() uint64 {
	if m != nil {
		return m.PairId
	}
	return 0
}

func (m *PoolA) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *PoolA) GetReserveAddress() string {
	if m != nil {
		return m.ReserveAddress
	}
	return ""
}

func (m *PoolA) GetPoolCoinDenom() string {
	if m != nil {
		return m.PoolCoinDenom
	}
	return ""
}

func (m *PoolA) GetLastDepositRequestId() uint64 {
	if m != nil {
		return m.LastDepositRequestId
	}
	return 0
}

func (m *PoolA) GetLastWithdrawRequestId() uint64 {
	if m != nil {
		return m.LastWithdrawRequestId
	}
	return 0
}

func (m *PoolA) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func init() {
	proto.RegisterType((*PoolA)(nil), "ollo.liquidity.v1.PoolA")
}

func init() { proto.RegisterFile("ollo/liquidity/v1/pool.proto", fileDescriptor_5e204054048950d3) }

var fileDescriptor_5e204054048950d3 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0x34, 0xcd, 0x9f, 0xad, 0x68, 0x85, 0x55, 0xe8, 0x2a, 0x20, 0x13, 0x38, 0x94,
	0x08, 0x09, 0x2f, 0x01, 0x21, 0xce, 0x2d, 0x91, 0x50, 0x6f, 0x95, 0x85, 0x84, 0xc4, 0xc5, 0x5a,
	0x7b, 0x17, 0x77, 0x84, 0xed, 0x71, 0x77, 0xd7, 0x69, 0xf2, 0x16, 0x3c, 0x16, 0xc7, 0x1e, 0x11,
	0x07, 0x84, 0x92, 0x1b, 0x4f, 0x81, 0x76, 0x9d, 0xb4, 0x95, 0x72, 0xeb, 0x29, 0x99, 0xef, 0xf7,
	0x7d, 0xe3, 0x9d, 0xd1, 0x90, 0xa7, 0x98, 0xe7, 0xc8, 0x72, 0xb8, 0xac, 0x41, 0x80, 0x59, 0xb0,
	0xd9, 0x84, 0x55, 0x88, 0x79, 0x58, 0x29, 0x34, 0xe8, 0x3f, 0xb4, 0x34, 0xbc, 0xa1, 0xe1, 0x6c,
	0x32, 0x3c, 0xcc, 0x30, 0x43, 0x47, 0x99, 0xfd, 0xd7, 0x18, 0x87, 0x41, 0x86, 0x98, 0xe5, 0x92,
	0xb9, 0x2a, 0xa9, 0xbf, 0x31, 0x51, 0x2b, 0x6e, 0x00, 0xcb, 0x0d, 0x4f, 0x51, 0x17, 0xa8, 0x59,
	0xc2, 0xb5, 0x64, 0xb3, 0x49, 0x22, 0x0d, 0x9f, 0xb0, 0x14, 0x61, 0xc3, 0x9f, 0x6f, 0x3f, 0xe3,
	0xf6, 0xab, 0xce, 0xf2, 0xe2, 0xdf, 0x0e, 0xd9, 0x3d, 0x47, 0xcc, 0x4f, 0x7c, 0x46, 0x3a, 0x66,
	0x51, 0x49, 0xea, 0x8d, 0xbc, 0xf1, 0xde, 0xdb, 0x27, 0xe1, 0xd6, 0x23, 0x43, 0xeb, 0xfb, 0xbc,
	0xa8, 0x64, 0xe4, 0x8c, 0xfe, 0x3e, 0x69, 0x83, 0xa0, 0xed, 0x91, 0x37, 0xee, 0x44, 0x6d, 0x10,
	0xfe, 0x11, 0xe9, 0x55, 0x1c, 0x54, 0x0c, 0x82, 0xee, 0x38, 0xb1, 0x6b, 0xcb, 0x33, 0xe1, 0x53,
	0xd2, 0x4b, 0x95, 0xe4, 0x06, 0x15, 0xed, 0x8c, 0xbc, 0xf1, 0x20, 0xda, 0x94, 0xfe, 0x4b, 0x72,
	0xa0, 0xa4, 0x96, 0x6a, 0x26, 0x63, 0x2e, 0x84, 0x92, 0x5a, 0xd3, 0x5d, 0xe7, 0xd8, 0x5f, 0xcb,
	0x27, 0x8d, 0xea, 0x1f, 0x93, 0x03, 0xbb, 0xc0, 0xd8, 0x0e, 0x17, 0x0b, 0x59, 0x62, 0x41, 0xbb,
	0xce, 0xf8, 0xc0, 0xca, 0x1f, 0x11, 0xca, 0xa9, 0x15, 0xfd, 0x4f, 0x64, 0x50, 0x40, 0x19, 0x57,
	0x0a, 0x52, 0x49, 0x7b, 0xd6, 0x71, 0xfa, 0xea, 0xf7, 0x9f, 0x67, 0xc7, 0x19, 0x98, 0x8b, 0x3a,
	0x09, 0x53, 0x2c, 0xd8, 0x7a, 0x67, 0xcd, 0xcf, 0x6b, 0x2d, 0xbe, 0x33, 0x3b, 0x8d, 0x0e, 0xa7,
	0x32, 0x8d, 0xfa, 0x05, 0x94, 0xe7, 0x36, 0xeb, 0x1a, 0xf1, 0xf9, 0xba, 0x51, 0xff, 0x1e, 0x8d,
	0xf8, 0xbc, 0x69, 0xf4, 0x9e, 0x1c, 0xe5, 0x5c, 0x9b, 0x58, 0xc8, 0x0a, 0x35, 0x98, 0x58, 0xc9,
	0xcb, 0x5a, 0x6a, 0x63, 0xb7, 0x34, 0x70, 0x5b, 0x3a, 0xb4, 0x78, 0xda, 0xd0, 0xa8, 0x81, 0x67,
	0xc2, 0xff, 0x40, 0xa8, 0x8b, 0x5d, 0x81, 0xb9, 0x10, 0x8a, 0x5f, 0xdd, 0xcd, 0x11, 0x97, 0x7b,
	0x64, 0xf9, 0x97, 0x35, 0xbe, 0x0d, 0x0e, 0x49, 0x5f, 0x80, 0xe6, 0x49, 0x2e, 0x05, 0xdd, 0x1b,
	0x79, 0xe3, 0x7e, 0x74, 0x53, 0x9f, 0xbe, 0xf9, 0xb9, 0x0c, 0xbc, 0xeb, 0x65, 0xe0, 0xfd, 0x5d,
	0x06, 0xde, 0x8f, 0x55, 0xd0, 0xba, 0x5e, 0x05, 0xad, 0x5f, 0xab, 0xa0, 0xf5, 0xf5, 0xb1, 0xbb,
	0x94, 0xf9, 0x9d, 0x5b, 0x71, 0xb3, 0x24, 0x5d, 0x77, 0x25, 0xef, 0xfe, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xac, 0x40, 0xab, 0x66, 0xd1, 0x02, 0x00, 0x00,
}

func (m *PoolA) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolA) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolA) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Disabled {
		i--
		if m.Disabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if m.LastWithdrawRequestId != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.LastWithdrawRequestId))
		i--
		dAtA[i] = 0x50
	}
	if m.LastDepositRequestId != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.LastDepositRequestId))
		i--
		dAtA[i] = 0x48
	}
	if m.MaxPrice != nil {
		{
			size := m.MaxPrice.Size()
			i -= size
			if _, err := m.MaxPrice.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.MinPrice != nil {
		{
			size := m.MinPrice.Size()
			i -= size
			if _, err := m.MinPrice.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.PoolCoinDenom) > 0 {
		i -= len(m.PoolCoinDenom)
		copy(dAtA[i:], m.PoolCoinDenom)
		i = encodeVarintPool(dAtA, i, uint64(len(m.PoolCoinDenom)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ReserveAddress) > 0 {
		i -= len(m.ReserveAddress)
		copy(dAtA[i:], m.ReserveAddress)
		i = encodeVarintPool(dAtA, i, uint64(len(m.ReserveAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if m.PairId != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != nil {
		{
			size, err := m.Type.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolA) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != nil {
		l = m.Type.Size()
		n += 1 + l + sovPool(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovPool(uint64(m.Id))
	}
	if m.PairId != 0 {
		n += 1 + sovPool(uint64(m.PairId))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.ReserveAddress)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.PoolCoinDenom)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	if m.MinPrice != nil {
		l = m.MinPrice.Size()
		n += 1 + l + sovPool(uint64(l))
	}
	if m.MaxPrice != nil {
		l = m.MaxPrice.Size()
		n += 1 + l + sovPool(uint64(l))
	}
	if m.LastDepositRequestId != 0 {
		n += 1 + sovPool(uint64(m.LastDepositRequestId))
	}
	if m.LastWithdrawRequestId != 0 {
		n += 1 + sovPool(uint64(m.LastWithdrawRequestId))
	}
	if m.Disabled {
		n += 2
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolA) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: PoolA: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolA: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Type == nil {
				m.Type = &PoolType{}
			}
			if err := m.Type.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReserveAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCoinDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolCoinDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.MinPrice = &v
			if err := m.MinPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.MaxPrice = &v
			if err := m.MaxPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastDepositRequestId", wireType)
			}
			m.LastDepositRequestId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastDepositRequestId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastWithdrawRequestId", wireType)
			}
			m.LastWithdrawRequestId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastWithdrawRequestId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Disabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
