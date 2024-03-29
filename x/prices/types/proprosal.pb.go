// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/prices/v1/proprosal.proto

package types

import (
	fmt "fmt"
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

type ProprosalAddDenomFeed struct {
	Title       string    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description string    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	DenomIds    DenomList `protobuf:"bytes,3,rep,name=denom_ids,json=denomIds,proto3,castrepeated=DenomList" json:"denom_ids,omitempty" yaml:"tracking_list"`
}

func (m *ProprosalAddDenomFeed) Reset()      { *m = ProprosalAddDenomFeed{} }
func (*ProprosalAddDenomFeed) ProtoMessage() {}
func (*ProprosalAddDenomFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_67a274ec35829db2, []int{0}
}
func (m *ProprosalAddDenomFeed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProprosalAddDenomFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProprosalAddDenomFeed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProprosalAddDenomFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProprosalAddDenomFeed.Merge(m, src)
}
func (m *ProprosalAddDenomFeed) XXX_Size() int {
	return m.Size()
}
func (m *ProprosalAddDenomFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_ProprosalAddDenomFeed.DiscardUnknown(m)
}

var xxx_messageInfo_ProprosalAddDenomFeed proto.InternalMessageInfo

type ProprosalRemoveDenomFeed struct {
	Title       string    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description string    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	DenomIds    DenomList `protobuf:"bytes,3,rep,name=denom_ids,json=denomIds,proto3,castrepeated=DenomList" json:"denom_ids,omitempty" yaml:"tracking_list"`
}

func (m *ProprosalRemoveDenomFeed) Reset()      { *m = ProprosalRemoveDenomFeed{} }
func (*ProprosalRemoveDenomFeed) ProtoMessage() {}
func (*ProprosalRemoveDenomFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_67a274ec35829db2, []int{1}
}
func (m *ProprosalRemoveDenomFeed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProprosalRemoveDenomFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProprosalRemoveDenomFeed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProprosalRemoveDenomFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProprosalRemoveDenomFeed.Merge(m, src)
}
func (m *ProprosalRemoveDenomFeed) XXX_Size() int {
	return m.Size()
}
func (m *ProprosalRemoveDenomFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_ProprosalRemoveDenomFeed.DiscardUnknown(m)
}

var xxx_messageInfo_ProprosalRemoveDenomFeed proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProprosalAddDenomFeed)(nil), "ollo.prices.v1.ProprosalAddDenomFeed")
	proto.RegisterType((*ProprosalRemoveDenomFeed)(nil), "ollo.prices.v1.ProprosalRemoveDenomFeed")
}

func init() { proto.RegisterFile("ollo/prices/v1/proprosal.proto", fileDescriptor_67a274ec35829db2) }

var fileDescriptor_67a274ec35829db2 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x91, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x1b, 0x87, 0x62, 0xab, 0x88, 0x94, 0x29, 0x65, 0x87, 0x74, 0x14, 0x94, 0x81, 0xd8,
	0x32, 0xbc, 0x88, 0x37, 0x8b, 0x08, 0x82, 0x07, 0x29, 0x78, 0xf1, 0x32, 0xba, 0x26, 0xd4, 0x60,
	0xbb, 0x7f, 0x48, 0xe2, 0x70, 0x6f, 0xe1, 0x63, 0x88, 0xef, 0x20, 0x78, 0xdc, 0x71, 0xc7, 0x9d,
	0xe6, 0xd6, 0xbd, 0x41, 0x9f, 0x40, 0x9a, 0x3a, 0x19, 0xbe, 0x81, 0xb7, 0x24, 0xbf, 0x2f, 0xff,
	0xef, 0xfb, 0xf3, 0x59, 0x18, 0xb2, 0x0c, 0x02, 0x2e, 0x58, 0x42, 0x65, 0x30, 0xec, 0x06, 0x5c,
	0x00, 0x17, 0x20, 0xe3, 0xcc, 0xe7, 0x02, 0x14, 0xd8, 0x7b, 0x15, 0xf7, 0x6b, 0xee, 0x0f, 0xbb,
	0xad, 0x66, 0x0a, 0x29, 0x68, 0x14, 0x54, 0xa7, 0x5a, 0xd5, 0x6a, 0xfd, 0x99, 0x22, 0x15, 0x08,
	0x5a, 0x33, 0xef, 0x03, 0x59, 0x07, 0x77, 0xab, 0xa9, 0x97, 0x84, 0x5c, 0xd1, 0x01, 0xe4, 0xd7,
	0x94, 0x12, 0xfb, 0xd8, 0xda, 0x54, 0x4c, 0x65, 0xd4, 0x41, 0x6d, 0xd4, 0x31, 0xc3, 0xfd, 0x72,
	0xe6, 0xee, 0x8e, 0xe2, 0x3c, 0xbb, 0xf0, 0xf4, 0xb3, 0x17, 0xd5, 0xd8, 0x3e, 0xb7, 0x76, 0x08,
	0x95, 0x89, 0x60, 0x5c, 0x31, 0x18, 0x38, 0x1b, 0x5a, 0x7d, 0x58, 0xce, 0x5c, 0xbb, 0x56, 0xaf,
	0x41, 0x2f, 0x5a, 0x97, 0xda, 0xa1, 0x65, 0x92, 0xca, 0xae, 0xc7, 0x88, 0x74, 0x1a, 0xed, 0x46,
	0xc7, 0x0c, 0x8f, 0xca, 0x99, 0xdb, 0xfc, 0x71, 0x11, 0x71, 0xf2, 0xc4, 0x06, 0x69, 0x2f, 0x63,
	0x52, 0x79, 0xef, 0x5f, 0xae, 0xa9, 0xb3, 0xdd, 0x32, 0xa9, 0xa2, 0x6d, 0xfd, 0xef, 0x86, 0x48,
	0xef, 0x13, 0x59, 0xce, 0x6f, 0xfe, 0x88, 0xe6, 0x30, 0xa4, 0xff, 0x6c, 0x85, 0xf0, 0x7e, 0xbc,
	0xc0, 0xc6, 0x74, 0x81, 0x8d, 0xf9, 0x02, 0xa3, 0xb7, 0x02, 0xa3, 0x71, 0x81, 0xd1, 0xa4, 0xc0,
	0x68, 0x5e, 0x60, 0xf4, 0xba, 0xc4, 0xc6, 0x64, 0x89, 0x8d, 0xe9, 0x12, 0x1b, 0x0f, 0x27, 0x29,
	0x53, 0x8f, 0xcf, 0x7d, 0x3f, 0x81, 0x3c, 0xa8, 0x3a, 0x3d, 0x95, 0x2a, 0xae, 0x62, 0xe8, 0x4b,
	0xf0, 0xb2, 0xaa, 0x58, 0x8d, 0x38, 0x95, 0xfd, 0x2d, 0x5d, 0xf0, 0xd9, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa3, 0xeb, 0x88, 0xbd, 0x44, 0x02, 0x00, 0x00,
}

func (this *ProprosalAddDenomFeed) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ProprosalAddDenomFeed)
	if !ok {
		that2, ok := that.(ProprosalAddDenomFeed)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ProprosalAddDenomFeed")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ProprosalAddDenomFeed but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ProprosalAddDenomFeed but is not nil && this == nil")
	}
	if this.Title != that1.Title {
		return fmt.Errorf("Title this(%v) Not Equal that(%v)", this.Title, that1.Title)
	}
	if this.Description != that1.Description {
		return fmt.Errorf("Description this(%v) Not Equal that(%v)", this.Description, that1.Description)
	}
	if len(this.DenomIds) != len(that1.DenomIds) {
		return fmt.Errorf("DenomIds this(%v) Not Equal that(%v)", len(this.DenomIds), len(that1.DenomIds))
	}
	for i := range this.DenomIds {
		if this.DenomIds[i] != that1.DenomIds[i] {
			return fmt.Errorf("DenomIds this[%v](%v) Not Equal that[%v](%v)", i, this.DenomIds[i], i, that1.DenomIds[i])
		}
	}
	return nil
}
func (this *ProprosalAddDenomFeed) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProprosalAddDenomFeed)
	if !ok {
		that2, ok := that.(ProprosalAddDenomFeed)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.DenomIds) != len(that1.DenomIds) {
		return false
	}
	for i := range this.DenomIds {
		if this.DenomIds[i] != that1.DenomIds[i] {
			return false
		}
	}
	return true
}
func (this *ProprosalRemoveDenomFeed) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ProprosalRemoveDenomFeed)
	if !ok {
		that2, ok := that.(ProprosalRemoveDenomFeed)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ProprosalRemoveDenomFeed")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ProprosalRemoveDenomFeed but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ProprosalRemoveDenomFeed but is not nil && this == nil")
	}
	if this.Title != that1.Title {
		return fmt.Errorf("Title this(%v) Not Equal that(%v)", this.Title, that1.Title)
	}
	if this.Description != that1.Description {
		return fmt.Errorf("Description this(%v) Not Equal that(%v)", this.Description, that1.Description)
	}
	if len(this.DenomIds) != len(that1.DenomIds) {
		return fmt.Errorf("DenomIds this(%v) Not Equal that(%v)", len(this.DenomIds), len(that1.DenomIds))
	}
	for i := range this.DenomIds {
		if this.DenomIds[i] != that1.DenomIds[i] {
			return fmt.Errorf("DenomIds this[%v](%v) Not Equal that[%v](%v)", i, this.DenomIds[i], i, that1.DenomIds[i])
		}
	}
	return nil
}
func (this *ProprosalRemoveDenomFeed) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProprosalRemoveDenomFeed)
	if !ok {
		that2, ok := that.(ProprosalRemoveDenomFeed)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.DenomIds) != len(that1.DenomIds) {
		return false
	}
	for i := range this.DenomIds {
		if this.DenomIds[i] != that1.DenomIds[i] {
			return false
		}
	}
	return true
}
func (m *ProprosalAddDenomFeed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProprosalAddDenomFeed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProprosalAddDenomFeed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DenomIds) > 0 {
		for iNdEx := len(m.DenomIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DenomIds[iNdEx])
			copy(dAtA[i:], m.DenomIds[iNdEx])
			i = encodeVarintProprosal(dAtA, i, uint64(len(m.DenomIds[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProprosal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProprosal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProprosalRemoveDenomFeed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProprosalRemoveDenomFeed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProprosalRemoveDenomFeed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DenomIds) > 0 {
		for iNdEx := len(m.DenomIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DenomIds[iNdEx])
			copy(dAtA[i:], m.DenomIds[iNdEx])
			i = encodeVarintProprosal(dAtA, i, uint64(len(m.DenomIds[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProprosal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProprosal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProprosal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProprosal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProprosalAddDenomFeed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProprosal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProprosal(uint64(l))
	}
	if len(m.DenomIds) > 0 {
		for _, s := range m.DenomIds {
			l = len(s)
			n += 1 + l + sovProprosal(uint64(l))
		}
	}
	return n
}

func (m *ProprosalRemoveDenomFeed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProprosal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProprosal(uint64(l))
	}
	if len(m.DenomIds) > 0 {
		for _, s := range m.DenomIds {
			l = len(s)
			n += 1 + l + sovProprosal(uint64(l))
		}
	}
	return n
}

func sovProprosal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProprosal(x uint64) (n int) {
	return sovProprosal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProprosalAddDenomFeed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProprosal
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
			return fmt.Errorf("proto: ProprosalAddDenomFeed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProprosalAddDenomFeed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProprosal
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
				return ErrInvalidLengthProprosal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProprosal
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
					return ErrIntOverflowProprosal
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
				return ErrInvalidLengthProprosal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProprosal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProprosal
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
				return ErrInvalidLengthProprosal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProprosal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomIds = append(m.DenomIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProprosal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProprosal
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
func (m *ProprosalRemoveDenomFeed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProprosal
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
			return fmt.Errorf("proto: ProprosalRemoveDenomFeed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProprosalRemoveDenomFeed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProprosal
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
				return ErrInvalidLengthProprosal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProprosal
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
					return ErrIntOverflowProprosal
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
				return ErrInvalidLengthProprosal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProprosal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProprosal
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
				return ErrInvalidLengthProprosal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProprosal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomIds = append(m.DenomIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProprosal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProprosal
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
func skipProprosal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProprosal
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
					return 0, ErrIntOverflowProprosal
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
					return 0, ErrIntOverflowProprosal
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
				return 0, ErrInvalidLengthProprosal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProprosal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProprosal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProprosal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProprosal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProprosal = fmt.Errorf("proto: unexpected end of group")
)
