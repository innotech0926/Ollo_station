// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/exchange/v1beta1/proposal.proto

package types

import (
	fmt "fmt"
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

type MarketParameterChangeProposal struct {
	Title       string                  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Changes     []MarketParameterChange `protobuf:"bytes,3,rep,name=changes,proto3" json:"changes"`
}

func (m *MarketParameterChangeProposal) Reset()      { *m = MarketParameterChangeProposal{} }
func (*MarketParameterChangeProposal) ProtoMessage() {}
func (*MarketParameterChangeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6f44678dc7e78b1, []int{0}
}
func (m *MarketParameterChangeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketParameterChangeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketParameterChangeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketParameterChangeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketParameterChangeProposal.Merge(m, src)
}
func (m *MarketParameterChangeProposal) XXX_Size() int {
	return m.Size()
}
func (m *MarketParameterChangeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketParameterChangeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_MarketParameterChangeProposal proto.InternalMessageInfo

type MarketParameterChange struct {
	MarketId            uint64                                 `protobuf:"varint,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	MakerFeeRate        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=maker_fee_rate,json=makerFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"maker_fee_rate"`
	TakerFeeRate        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=taker_fee_rate,json=takerFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"taker_fee_rate"`
	OrderSourceFeeRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=order_source_fee_ratio,json=orderSourceFeeRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"order_source_fee_ratio"`
}

func (m *MarketParameterChange) Reset()         { *m = MarketParameterChange{} }
func (m *MarketParameterChange) String() string { return proto.CompactTextString(m) }
func (*MarketParameterChange) ProtoMessage()    {}
func (*MarketParameterChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6f44678dc7e78b1, []int{1}
}
func (m *MarketParameterChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketParameterChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketParameterChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketParameterChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketParameterChange.Merge(m, src)
}
func (m *MarketParameterChange) XXX_Size() int {
	return m.Size()
}
func (m *MarketParameterChange) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketParameterChange.DiscardUnknown(m)
}

var xxx_messageInfo_MarketParameterChange proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MarketParameterChangeProposal)(nil), "ollo.exchange.v1beta1.MarketParameterChangeProposal")
	proto.RegisterType((*MarketParameterChange)(nil), "ollo.exchange.v1beta1.MarketParameterChange")
}

func init() {
	proto.RegisterFile("ollo/exchange/v1beta1/proposal.proto", fileDescriptor_f6f44678dc7e78b1)
}

var fileDescriptor_f6f44678dc7e78b1 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3f, 0x6f, 0xda, 0x40,
	0x18, 0xc6, 0x6d, 0x70, 0xff, 0x70, 0x54, 0x1d, 0x5c, 0x5a, 0xb9, 0xad, 0x6a, 0x10, 0x43, 0xc5,
	0xc2, 0x5d, 0x69, 0x55, 0xa9, 0xea, 0x48, 0xab, 0x4a, 0x1d, 0x2a, 0x21, 0xb7, 0xea, 0x90, 0xc5,
	0x3a, 0xce, 0x6f, 0x8c, 0x65, 0xec, 0xb3, 0xee, 0x0e, 0x42, 0xbe, 0x45, 0xc6, 0x8c, 0xf9, 0x06,
	0xf9, 0x14, 0x91, 0x18, 0x19, 0xa3, 0x0c, 0x28, 0x81, 0x2f, 0x12, 0x71, 0xb6, 0x13, 0x22, 0x91,
	0x85, 0xe9, 0xee, 0xde, 0xf7, 0x79, 0x7f, 0xcf, 0x23, 0xdd, 0x8b, 0x3a, 0x4c, 0x80, 0x64, 0x90,
	0x2a, 0x02, 0x33, 0x36, 0xa2, 0x69, 0x08, 0x64, 0xda, 0x1b, 0x82, 0xa2, 0x3d, 0x92, 0x09, 0x9e,
	0x71, 0x49, 0xc7, 0x38, 0x13, 0x5c, 0x71, 0xfb, 0x6d, 0xa9, 0xc4, 0xa5, 0x12, 0x17, 0xca, 0x77,
	0x8d, 0x90, 0x87, 0x5c, 0xab, 0xc8, 0xe6, 0x96, 0x0f, 0xb4, 0xcf, 0x4d, 0xf4, 0xe1, 0x0f, 0x15,
	0x31, 0xa8, 0x01, 0x15, 0x34, 0x01, 0x05, 0xe2, 0x87, 0x9e, 0x1b, 0x14, 0x60, 0xbb, 0x81, 0x9e,
	0xa8, 0x48, 0x8d, 0xc1, 0x31, 0x5b, 0x66, 0xa7, 0xe6, 0xe5, 0x0f, 0xbb, 0x85, 0xea, 0x01, 0x48,
	0x26, 0xa2, 0x4c, 0x45, 0x3c, 0x75, 0x2a, 0xba, 0xb7, 0x5d, 0xb2, 0x07, 0xe8, 0x59, 0x9e, 0x40,
	0x3a, 0xd5, 0x56, 0xb5, 0x53, 0xff, 0xfc, 0x09, 0x3f, 0x1a, 0x0e, 0xef, 0x8c, 0xd0, 0xb7, 0xe6,
	0xcb, 0xa6, 0xe1, 0x95, 0x98, 0xef, 0xd6, 0xe9, 0x59, 0xd3, 0x68, 0x5f, 0x54, 0xd0, 0xeb, 0x9d,
	0x72, 0xfb, 0x3d, 0xaa, 0x25, 0xba, 0xe1, 0x47, 0x81, 0x4e, 0x6b, 0x79, 0xcf, 0xf3, 0xc2, 0xef,
	0xc0, 0xfe, 0x87, 0x5e, 0x26, 0x34, 0x06, 0xe1, 0x1f, 0x02, 0xf8, 0x82, 0x2a, 0xc8, 0x33, 0xf7,
	0xf1, 0xc6, 0xe3, 0x6a, 0xd9, 0xfc, 0x18, 0x46, 0x6a, 0x34, 0x19, 0x62, 0xc6, 0x13, 0xc2, 0xb8,
	0x4c, 0xb8, 0x2c, 0x8e, 0xae, 0x0c, 0x62, 0xa2, 0x8e, 0x33, 0x90, 0xf8, 0x27, 0x30, 0xef, 0x85,
	0xa6, 0xfc, 0x02, 0xf0, 0xa8, 0x82, 0x0d, 0x55, 0x3d, 0xa4, 0x56, 0xf7, 0xa3, 0xaa, 0x6d, 0x2a,
	0x43, 0x6f, 0xb8, 0x08, 0x40, 0xf8, 0x92, 0x4f, 0x04, 0x83, 0x12, 0x1e, 0x71, 0xc7, 0xda, 0x8b,
	0xfe, 0x4a, 0xd3, 0xfe, 0x6a, 0x58, 0xee, 0x11, 0xf1, 0xfe, 0xff, 0xf9, 0x8d, 0x6b, 0xcc, 0x57,
	0xae, 0xb9, 0x58, 0xb9, 0xe6, 0xf5, 0xca, 0x35, 0x4f, 0xd6, 0xae, 0xb1, 0x58, 0xbb, 0xc6, 0xe5,
	0xda, 0x35, 0x0e, 0xbe, 0x6d, 0xa3, 0x8b, 0x6f, 0xeb, 0xa6, 0xa0, 0x8e, 0xb8, 0x88, 0xef, 0x0a,
	0x64, 0xfa, 0x95, 0xcc, 0xee, 0x77, 0x52, 0x1b, 0x0e, 0x9f, 0xea, 0xc5, 0xfa, 0x72, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xd9, 0xd6, 0x6e, 0x8a, 0xb5, 0x02, 0x00, 0x00,
}

func (m *MarketParameterChangeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketParameterChangeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketParameterChangeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Changes) > 0 {
		for iNdEx := len(m.Changes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Changes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MarketParameterChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketParameterChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketParameterChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.OrderSourceFeeRatio.Size()
		i -= size
		if _, err := m.OrderSourceFeeRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.TakerFeeRate.Size()
		i -= size
		if _, err := m.TakerFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MakerFeeRate.Size()
		i -= size
		if _, err := m.MakerFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.MarketId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.MarketId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MarketParameterChangeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.Changes) > 0 {
		for _, e := range m.Changes {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *MarketParameterChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MarketId != 0 {
		n += 1 + sovProposal(uint64(m.MarketId))
	}
	l = m.MakerFeeRate.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = m.TakerFeeRate.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = m.OrderSourceFeeRatio.Size()
	n += 1 + l + sovProposal(uint64(l))
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MarketParameterChangeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: MarketParameterChangeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketParameterChangeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Changes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Changes = append(m.Changes, MarketParameterChange{})
			if err := m.Changes[len(m.Changes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *MarketParameterChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: MarketParameterChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketParameterChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
			}
			m.MarketId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MarketId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerFeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MakerFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerFeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TakerFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderSourceFeeRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OrderSourceFeeRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
