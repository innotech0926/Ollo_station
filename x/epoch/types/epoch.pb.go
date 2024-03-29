// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/epoch/v1/epoch.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

type Epoch struct {
	//
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	//
	Start time.Time `protobuf:"bytes,2,opt,name=start,proto3,stdtime" json:"start" yaml:"start"`
	//
	Duration time.Duration `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration" yaml:"duration"`
	//
	CurrentEpochNumber uint64 `protobuf:"varint,4,opt,name=current_epoch_number,json=currentEpochNumber,proto3" json:"current_epoch_number,omitempty" yaml:"current_epoch_number"`
	//
	CurrentEpochStart time.Time `protobuf:"bytes,5,opt,name=current_epoch_start,json=currentEpochStart,proto3,stdtime" json:"current_epoch_start" yaml:"current_epoch_start"`
	//
	EpochStarted bool `protobuf:"varint,6,opt,name=epoch_started,json=epochStarted,proto3" json:"epoch_started,omitempty" yaml:"epoch_started"`
	//
	CurrentEpochStartHeight uint64 `protobuf:"varint,8,opt,name=current_epoch_start_height,json=currentEpochStartHeight,proto3" json:"current_epoch_start_height,omitempty" yaml:"current_epoch_start_height"`
}

func (m *Epoch) Reset()         { *m = Epoch{} }
func (m *Epoch) String() string { return proto.CompactTextString(m) }
func (*Epoch) ProtoMessage()    {}
func (*Epoch) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e3c0aaa7374a1ce, []int{0}
}
func (m *Epoch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Epoch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Epoch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Epoch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Epoch.Merge(m, src)
}
func (m *Epoch) XXX_Size() int {
	return m.Size()
}
func (m *Epoch) XXX_DiscardUnknown() {
	xxx_messageInfo_Epoch.DiscardUnknown(m)
}

var xxx_messageInfo_Epoch proto.InternalMessageInfo

func (m *Epoch) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Epoch) GetStart() time.Time {
	if m != nil {
		return m.Start
	}
	return time.Time{}
}

func (m *Epoch) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Epoch) GetCurrentEpochNumber() uint64 {
	if m != nil {
		return m.CurrentEpochNumber
	}
	return 0
}

func (m *Epoch) GetCurrentEpochStart() time.Time {
	if m != nil {
		return m.CurrentEpochStart
	}
	return time.Time{}
}

func (m *Epoch) GetEpochStarted() bool {
	if m != nil {
		return m.EpochStarted
	}
	return false
}

func (m *Epoch) GetCurrentEpochStartHeight() uint64 {
	if m != nil {
		return m.CurrentEpochStartHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Epoch)(nil), "ollo.epoch.v1.Epoch")
}

func init() { proto.RegisterFile("ollo/epoch/v1/epoch.proto", fileDescriptor_0e3c0aaa7374a1ce) }

var fileDescriptor_0e3c0aaa7374a1ce = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6a, 0xd4, 0x40,
	0x18, 0xc7, 0x33, 0x6b, 0xb6, 0xa6, 0x63, 0x17, 0x75, 0x5c, 0x70, 0x9a, 0x62, 0x26, 0x0e, 0x28,
	0x41, 0x30, 0xa1, 0x7a, 0x13, 0xbc, 0x84, 0x0a, 0xd2, 0x83, 0x60, 0xf4, 0xe4, 0x65, 0x49, 0x36,
	0x63, 0x12, 0x48, 0x76, 0x42, 0x32, 0x29, 0xf6, 0x2d, 0x7a, 0xf4, 0x49, 0x7c, 0x86, 0x1e, 0x7b,
	0xf4, 0x14, 0x65, 0xf7, 0x0d, 0xf2, 0x04, 0x92, 0x99, 0xa4, 0x6c, 0xed, 0x8a, 0xb7, 0x99, 0xef,
	0xfb, 0x7f, 0xbf, 0xff, 0xc7, 0x7f, 0x06, 0x1e, 0xf2, 0x3c, 0xe7, 0x1e, 0x2b, 0xf9, 0x32, 0xf5,
	0xce, 0x8e, 0xd5, 0xc1, 0x2d, 0x2b, 0x2e, 0x38, 0x9a, 0xf5, 0x2d, 0x57, 0x55, 0xce, 0x8e, 0xcd,
	0x79, 0xc2, 0x13, 0x2e, 0x3b, 0x5e, 0x7f, 0x52, 0x22, 0x93, 0x24, 0x9c, 0x27, 0x39, 0xf3, 0xe4,
	0x2d, 0x6a, 0xbe, 0x7a, 0x22, 0x2b, 0x58, 0x2d, 0xc2, 0xa2, 0x1c, 0x04, 0xd6, 0xdf, 0x82, 0xb8,
	0xa9, 0x42, 0x91, 0xf1, 0x95, 0xea, 0xd3, 0x1f, 0x3a, 0x9c, 0xbe, 0xeb, 0x3d, 0xd0, 0x13, 0x38,
	0xc9, 0x62, 0x0c, 0x6c, 0xe0, 0xec, 0xfb, 0xb3, 0xae, 0x25, 0xfb, 0xe7, 0x61, 0x91, 0xbf, 0xa1,
	0x59, 0x4c, 0x83, 0x49, 0x16, 0xa3, 0x53, 0x38, 0xad, 0x45, 0x58, 0x09, 0x3c, 0xb1, 0x81, 0x73,
	0xef, 0x95, 0xe9, 0x2a, 0xb0, 0x3b, 0x82, 0xdd, 0xcf, 0xa3, 0xb3, 0x8f, 0x2f, 0x5b, 0xa2, 0x75,
	0x2d, 0x39, 0x50, 0x04, 0x39, 0x46, 0x2f, 0x7e, 0x11, 0x10, 0x28, 0x04, 0x0a, 0xa0, 0x31, 0xae,
	0x81, 0xef, 0x48, 0xdc, 0xe1, 0x2d, 0xdc, 0xc9, 0x20, 0xf0, 0x8f, 0x06, 0xda, 0x7d, 0x45, 0x1b,
	0x07, 0xe9, 0xf7, 0x1e, 0x78, 0xcd, 0x41, 0x1f, 0xe1, 0x7c, 0xd9, 0x54, 0x15, 0x5b, 0x89, 0x85,
	0xcc, 0x6c, 0xb1, 0x6a, 0x8a, 0x88, 0x55, 0x58, 0xb7, 0x81, 0xa3, 0xfb, 0xa4, 0x6b, 0xc9, 0x91,
	0x02, 0xec, 0x52, 0xd1, 0x00, 0x0d, 0x65, 0x99, 0xc5, 0x07, 0x59, 0x44, 0x15, 0x7c, 0x74, 0x53,
	0xac, 0x02, 0x98, 0xfe, 0x37, 0x80, 0xe7, 0xc3, 0xca, 0xe6, 0x2e, 0xc7, 0xad, 0x38, 0x1e, 0x6e,
	0x9b, 0x7e, 0x92, 0xd1, 0xbc, 0x85, 0xb3, 0x2d, 0x19, 0x8b, 0xf1, 0x9e, 0x0d, 0x1c, 0xc3, 0xc7,
	0x5d, 0x4b, 0xe6, 0x8a, 0x76, 0xa3, 0x4d, 0x83, 0x03, 0x76, 0x3d, 0xcc, 0x62, 0x14, 0x41, 0x73,
	0x87, 0xdb, 0x22, 0x65, 0x59, 0x92, 0x0a, 0x6c, 0xc8, 0x2c, 0x9e, 0x75, 0x2d, 0x79, 0xfa, 0xcf,
	0xcd, 0x06, 0x2d, 0x0d, 0x1e, 0xdf, 0x5a, 0xee, 0xbd, 0xec, 0x9c, 0xea, 0xc6, 0xdd, 0x07, 0x86,
	0x7f, 0x72, 0xb9, 0xb6, 0xc0, 0xd5, 0xda, 0x02, 0xbf, 0xd7, 0x16, 0xb8, 0xd8, 0x58, 0xda, 0xd5,
	0xc6, 0xd2, 0x7e, 0x6e, 0x2c, 0xed, 0xcb, 0x8b, 0x24, 0x13, 0x69, 0x13, 0xb9, 0x4b, 0x5e, 0x78,
	0xfd, 0x1f, 0x7e, 0x59, 0x0b, 0xf9, 0x44, 0xf2, 0xe2, 0x7d, 0x1b, 0x7e, 0xbb, 0x38, 0x2f, 0x59,
	0x1d, 0xed, 0xc9, 0xf4, 0x5e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x36, 0xdd, 0x53, 0x67, 0x08,
	0x03, 0x00, 0x00,
}

func (m *Epoch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Epoch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Epoch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrentEpochStartHeight != 0 {
		i = encodeVarintEpoch(dAtA, i, uint64(m.CurrentEpochStartHeight))
		i--
		dAtA[i] = 0x40
	}
	if m.EpochStarted {
		i--
		if m.EpochStarted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CurrentEpochStart, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CurrentEpochStart):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEpoch(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if m.CurrentEpochNumber != 0 {
		i = encodeVarintEpoch(dAtA, i, uint64(m.CurrentEpochNumber))
		i--
		dAtA[i] = 0x20
	}
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintEpoch(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Start, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Start):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintEpoch(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintEpoch(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEpoch(dAtA []byte, offset int, v uint64) int {
	offset -= sovEpoch(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Epoch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovEpoch(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Start)
	n += 1 + l + sovEpoch(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovEpoch(uint64(l))
	if m.CurrentEpochNumber != 0 {
		n += 1 + sovEpoch(uint64(m.CurrentEpochNumber))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CurrentEpochStart)
	n += 1 + l + sovEpoch(uint64(l))
	if m.EpochStarted {
		n += 2
	}
	if m.CurrentEpochStartHeight != 0 {
		n += 1 + sovEpoch(uint64(m.CurrentEpochStartHeight))
	}
	return n
}

func sovEpoch(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEpoch(x uint64) (n int) {
	return sovEpoch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Epoch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpoch
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
			return fmt.Errorf("proto: Epoch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Epoch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
				return ErrInvalidLengthEpoch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEpoch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
				return ErrInvalidLengthEpoch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Start, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
				return ErrInvalidLengthEpoch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochNumber", wireType)
			}
			m.CurrentEpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpochNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochStart", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
				return ErrInvalidLengthEpoch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CurrentEpochStart, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochStarted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
			m.EpochStarted = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochStartHeight", wireType)
			}
			m.CurrentEpochStartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpochStartHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEpoch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEpoch
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
func skipEpoch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEpoch
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
					return 0, ErrIntOverflowEpoch
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
					return 0, ErrIntOverflowEpoch
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
				return 0, ErrInvalidLengthEpoch
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEpoch
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEpoch
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEpoch        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEpoch          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEpoch = fmt.Errorf("proto: unexpected end of group")
)
