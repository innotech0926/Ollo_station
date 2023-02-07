// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/loan/v1/loans.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type Loans struct {
	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount     string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee        string `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Collateral string `protobuf:"bytes,4,opt,name=collateral,proto3" json:"collateral,omitempty"`
	Deadline   string `protobuf:"bytes,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
	State      string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	Borrower   string `protobuf:"bytes,7,opt,name=borrower,proto3" json:"borrower,omitempty" yaml:"borrower"`
	Lender     string `protobuf:"bytes,8,opt,name=lender,proto3" json:"lender,omitempty" yaml:"lender"`
}

func (m *Loans) Reset()         { *m = Loans{} }
func (m *Loans) String() string { return proto.CompactTextString(m) }
func (*Loans) ProtoMessage()    {}
func (*Loans) Descriptor() ([]byte, []int) {
	return fileDescriptor_55064bade1929cb0, []int{0}
}
func (m *Loans) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Loans) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Loans.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Loans) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Loans.Merge(m, src)
}
func (m *Loans) XXX_Size() int {
	return m.Size()
}
func (m *Loans) XXX_DiscardUnknown() {
	xxx_messageInfo_Loans.DiscardUnknown(m)
}

var xxx_messageInfo_Loans proto.InternalMessageInfo

func (m *Loans) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Loans) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *Loans) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *Loans) GetCollateral() string {
	if m != nil {
		return m.Collateral
	}
	return ""
}

func (m *Loans) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

func (m *Loans) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Loans) GetBorrower() string {
	if m != nil {
		return m.Borrower
	}
	return ""
}

func (m *Loans) GetLender() string {
	if m != nil {
		return m.Lender
	}
	return ""
}

func init() {
	proto.RegisterType((*Loans)(nil), "ollo.loan.v1.Loans")
}

func init() { proto.RegisterFile("ollo/loan/v1/loans.proto", fileDescriptor_55064bade1929cb0) }

var fileDescriptor_55064bade1929cb0 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4e, 0x02, 0x31,
	0x14, 0xc6, 0x99, 0x01, 0x46, 0x6c, 0xfc, 0xdb, 0x10, 0x53, 0x59, 0x14, 0xc2, 0x0a, 0x43, 0xa4,
	0x99, 0xb8, 0x73, 0x07, 0x1b, 0x37, 0xae, 0x70, 0xe7, 0xc6, 0x74, 0x98, 0x32, 0x99, 0xa4, 0x33,
	0x8f, 0xb4, 0x05, 0xe5, 0x16, 0xde, 0xc0, 0x4b, 0x78, 0x08, 0x97, 0xc4, 0x95, 0x2b, 0x62, 0xe0,
	0x06, 0x9c, 0xc0, 0x4c, 0x3b, 0x12, 0x37, 0xae, 0xfa, 0xbe, 0xef, 0xf7, 0xde, 0x97, 0xb6, 0x0f,
	0x11, 0x90, 0x12, 0x98, 0x04, 0x9e, 0xb3, 0x45, 0x68, 0x4f, 0x3d, 0x98, 0x29, 0x30, 0x80, 0x8f,
	0x0a, 0x32, 0x28, 0x9c, 0xc1, 0x22, 0x6c, 0x35, 0x13, 0x48, 0xc0, 0x02, 0x56, 0x54, 0xae, 0xa7,
	0xd5, 0x4e, 0x00, 0x12, 0x29, 0x98, 0x55, 0xd1, 0x7c, 0xca, 0x4c, 0x9a, 0x09, 0x6d, 0x78, 0x36,
	0x2b, 0x1b, 0x2e, 0x27, 0xa0, 0x33, 0xd0, 0x4f, 0x6e, 0xd2, 0x89, 0x12, 0x51, 0xa7, 0x58, 0xc4,
	0xb5, 0x60, 0x8b, 0x30, 0x12, 0x86, 0x87, 0x6c, 0x02, 0x69, 0xee, 0x78, 0xf7, 0xcd, 0x47, 0xf5,
	0xfb, 0xe2, 0x3e, 0xf8, 0x04, 0xf9, 0x69, 0x4c, 0xbc, 0x8e, 0xd7, 0xab, 0x8d, 0xfd, 0x34, 0xc6,
	0x17, 0x28, 0xe0, 0x19, 0xcc, 0x73, 0x43, 0xfc, 0x8e, 0xd7, 0x3b, 0x1c, 0x97, 0x0a, 0x9f, 0xa1,
	0xea, 0x54, 0x08, 0x52, 0xb5, 0x66, 0x51, 0x62, 0x8a, 0xd0, 0x04, 0xa4, 0xe4, 0x46, 0x28, 0x2e,
	0x49, 0xcd, 0x82, 0x3f, 0x0e, 0x6e, 0xa1, 0x46, 0x2c, 0x78, 0x2c, 0xd3, 0x5c, 0x90, 0xba, 0xa5,
	0x7b, 0x8d, 0x9b, 0xa8, 0xae, 0x0d, 0x37, 0x82, 0x04, 0x16, 0x38, 0x81, 0xef, 0x50, 0x23, 0x02,
	0xa5, 0xe0, 0x59, 0x28, 0x72, 0x50, 0x80, 0x51, 0x7f, 0xb7, 0x6e, 0x9f, 0x2e, 0x79, 0x26, 0x6f,
	0xbb, 0xbf, 0xa4, 0xfb, 0xf9, 0x7e, 0xdd, 0x2c, 0x1f, 0x3b, 0x8c, 0x63, 0x25, 0xb4, 0x7e, 0x30,
	0x2a, 0xcd, 0x93, 0xf1, 0x7e, 0x18, 0x0f, 0x51, 0x20, 0x45, 0x1e, 0x0b, 0x45, 0x1a, 0x36, 0xe6,
	0x6a, 0xb7, 0x6e, 0x1f, 0xbb, 0x18, 0xe7, 0xff, 0x1f, 0x52, 0x0e, 0x8e, 0xfa, 0x1f, 0x1b, 0xea,
	0xad, 0x36, 0xd4, 0xfb, 0xde, 0x50, 0xef, 0x75, 0x4b, 0x2b, 0xab, 0x2d, 0xad, 0x7c, 0x6d, 0x69,
	0xe5, 0xf1, 0xdc, 0x6e, 0xf5, 0xc5, 0xed, 0xd5, 0x2c, 0x67, 0x42, 0x47, 0x81, 0xfd, 0xd5, 0x9b,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x55, 0xe8, 0xcc, 0xf1, 0x01, 0x00, 0x00,
}

func (m *Loans) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Loans) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Loans) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Lender) > 0 {
		i -= len(m.Lender)
		copy(dAtA[i:], m.Lender)
		i = encodeVarintLoans(dAtA, i, uint64(len(m.Lender)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Borrower) > 0 {
		i -= len(m.Borrower)
		copy(dAtA[i:], m.Borrower)
		i = encodeVarintLoans(dAtA, i, uint64(len(m.Borrower)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintLoans(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Deadline) > 0 {
		i -= len(m.Deadline)
		copy(dAtA[i:], m.Deadline)
		i = encodeVarintLoans(dAtA, i, uint64(len(m.Deadline)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Collateral) > 0 {
		i -= len(m.Collateral)
		copy(dAtA[i:], m.Collateral)
		i = encodeVarintLoans(dAtA, i, uint64(len(m.Collateral)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintLoans(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintLoans(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintLoans(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLoans(dAtA []byte, offset int, v uint64) int {
	offset -= sovLoans(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Loans) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLoans(uint64(m.Id))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovLoans(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovLoans(uint64(l))
	}
	l = len(m.Collateral)
	if l > 0 {
		n += 1 + l + sovLoans(uint64(l))
	}
	l = len(m.Deadline)
	if l > 0 {
		n += 1 + l + sovLoans(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovLoans(uint64(l))
	}
	l = len(m.Borrower)
	if l > 0 {
		n += 1 + l + sovLoans(uint64(l))
	}
	l = len(m.Lender)
	if l > 0 {
		n += 1 + l + sovLoans(uint64(l))
	}
	return n
}

func sovLoans(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLoans(x uint64) (n int) {
	return sovLoans(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Loans) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoans
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
			return fmt.Errorf("proto: Loans: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Loans: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoans
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoans
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
				return ErrInvalidLengthLoans
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoans
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoans
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
				return ErrInvalidLengthLoans
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoans
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoans
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
				return ErrInvalidLengthLoans
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoans
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collateral = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoans
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
				return ErrInvalidLengthLoans
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoans
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deadline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoans
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
				return ErrInvalidLengthLoans
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoans
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Borrower", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoans
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
				return ErrInvalidLengthLoans
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoans
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Borrower = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoans
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
				return ErrInvalidLengthLoans
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoans
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoans(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLoans
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
func skipLoans(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLoans
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
					return 0, ErrIntOverflowLoans
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
					return 0, ErrIntOverflowLoans
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
				return 0, ErrInvalidLengthLoans
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLoans
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLoans
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLoans        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLoans          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLoans = fmt.Errorf("proto: unexpected end of group")
)
