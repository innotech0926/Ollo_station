// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/claim/v1/claim.proto

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

type InitialClaim struct {
	Enabled bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	GoalId  uint64 `protobuf:"varint,2,opt,name=goal_id,json=goalId,proto3" json:"goal_id,omitempty"`
}

func (m *InitialClaim) Reset()         { *m = InitialClaim{} }
func (m *InitialClaim) String() string { return proto.CompactTextString(m) }
func (*InitialClaim) ProtoMessage()    {}
func (*InitialClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcf5997285c872eb, []int{0}
}
func (m *InitialClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InitialClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InitialClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InitialClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitialClaim.Merge(m, src)
}
func (m *InitialClaim) XXX_Size() int {
	return m.Size()
}
func (m *InitialClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_InitialClaim.DiscardUnknown(m)
}

var xxx_messageInfo_InitialClaim proto.InternalMessageInfo

func (m *InitialClaim) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *InitialClaim) GetGoalId() uint64 {
	if m != nil {
		return m.GoalId
	}
	return 0
}

type ClaimRecord struct {
	Address        string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Claimable      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=claimable,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"claimable"`
	CompletedGoals []uint64                               `protobuf:"varint,3,rep,packed,name=completed_goals,json=completedGoals,proto3" json:"completed_goals,omitempty"`
	ClaimedGoals   []uint64                               `protobuf:"varint,4,rep,packed,name=claimed_goals,json=claimedGoals,proto3" json:"claimed_goals,omitempty"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcf5997285c872eb, []int{1}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func (m *ClaimRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimRecord) GetCompletedGoals() []uint64 {
	if m != nil {
		return m.CompletedGoals
	}
	return nil
}

func (m *ClaimRecord) GetClaimedGoals() []uint64 {
	if m != nil {
		return m.ClaimedGoals
	}
	return nil
}

func init() {
	proto.RegisterType((*InitialClaim)(nil), "ollo.claim.v1.InitialClaim")
	proto.RegisterType((*ClaimRecord)(nil), "ollo.claim.v1.ClaimRecord")
}

func init() { proto.RegisterFile("ollo/claim/v1/claim.proto", fileDescriptor_fcf5997285c872eb) }

var fileDescriptor_fcf5997285c872eb = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x51, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0xed, 0x0a, 0x01, 0x59, 0x41, 0x93, 0x0d, 0x89, 0x85, 0x43, 0x21, 0x98, 0x28, 0x07, 0x69,
	0x83, 0x5e, 0xbd, 0x80, 0x07, 0xd3, 0x6b, 0xbd, 0x71, 0x21, 0xa5, 0xbb, 0xa9, 0x8d, 0xdb, 0x0e,
	0xe9, 0xae, 0x44, 0xff, 0xc2, 0x8f, 0xe1, 0x23, 0x38, 0x12, 0x4e, 0xc6, 0x03, 0x31, 0x70, 0xf4,
	0x27, 0xcc, 0xee, 0x16, 0x3c, 0x75, 0xe6, 0xbd, 0x37, 0x6f, 0x5e, 0x77, 0x70, 0x0b, 0x38, 0x07,
	0x2f, 0xe2, 0x61, 0x92, 0x7a, 0x8b, 0xa1, 0x29, 0xdc, 0x79, 0x0e, 0x12, 0x48, 0x43, 0x51, 0xae,
	0x41, 0x16, 0xc3, 0x76, 0x33, 0x86, 0x18, 0x34, 0xe3, 0xa9, 0xca, 0x88, 0xda, 0xad, 0x08, 0x44,
	0x0a, 0x62, 0x6a, 0x08, 0xd3, 0x18, 0xaa, 0x37, 0xc2, 0x75, 0x3f, 0x4b, 0x64, 0x12, 0xf2, 0x47,
	0xe5, 0x41, 0x6c, 0x5c, 0x65, 0x59, 0x38, 0xe3, 0x8c, 0xda, 0xa8, 0x8b, 0xfa, 0xa7, 0xc1, 0xa1,
	0x25, 0x97, 0xb8, 0x1a, 0x43, 0xc8, 0xa7, 0x09, 0xb5, 0x4f, 0xba, 0xa8, 0x5f, 0x0e, 0x2a, 0xaa,
	0xf5, 0x69, 0xef, 0x17, 0xe1, 0x33, 0x3d, 0x1c, 0xb0, 0x08, 0x72, 0x4a, 0xee, 0x70, 0x35, 0xa4,
	0x34, 0x67, 0x42, 0x68, 0x8b, 0xda, 0xd8, 0xde, 0x2c, 0x07, 0xcd, 0x62, 0xeb, 0xc8, 0x30, 0xcf,
	0x32, 0x4f, 0xb2, 0x38, 0x38, 0x08, 0xc9, 0x04, 0xd7, 0xf4, 0x3f, 0xa8, 0x55, 0xda, 0xbe, 0x36,
	0x7e, 0x58, 0x6d, 0x3b, 0xd6, 0xf7, 0xb6, 0x73, 0x1d, 0x27, 0xf2, 0xe5, 0x6d, 0xe6, 0x46, 0x90,
	0x16, 0xd1, 0x8b, 0xcf, 0x40, 0xd0, 0x57, 0x4f, 0x7e, 0xcc, 0x99, 0x70, 0xfd, 0x4c, 0x6e, 0x96,
	0x03, 0x5c, 0xec, 0xf0, 0x33, 0x19, 0xfc, 0xdb, 0x91, 0x1b, 0x7c, 0x11, 0x41, 0x3a, 0xe7, 0x4c,
	0x32, 0x3a, 0x55, 0x99, 0x85, 0x5d, 0xea, 0x96, 0xfa, 0xe5, 0xe0, 0xfc, 0x08, 0x3f, 0x29, 0x94,
	0x5c, 0xe1, 0x86, 0x9e, 0x3a, 0xca, 0xca, 0x5a, 0x56, 0x2f, 0x40, 0x2d, 0x1a, 0xdf, 0xae, 0x76,
	0x0e, 0x5a, 0xef, 0x1c, 0xf4, 0xb3, 0x73, 0xd0, 0xe7, 0xde, 0xb1, 0xd6, 0x7b, 0xc7, 0xfa, 0xda,
	0x3b, 0xd6, 0x84, 0xe8, 0x2b, 0xbd, 0x17, 0x77, 0xd2, 0xc1, 0x66, 0x15, 0xfd, 0xca, 0xf7, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x94, 0x1a, 0x4b, 0xc2, 0x01, 0x00, 0x00,
}

func (m *InitialClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InitialClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InitialClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GoalId != 0 {
		i = encodeVarintClaim(dAtA, i, uint64(m.GoalId))
		i--
		dAtA[i] = 0x10
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimedGoals) > 0 {
		dAtA2 := make([]byte, len(m.ClaimedGoals)*10)
		var j1 int
		for _, num := range m.ClaimedGoals {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintClaim(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CompletedGoals) > 0 {
		dAtA4 := make([]byte, len(m.CompletedGoals)*10)
		var j3 int
		for _, num := range m.CompletedGoals {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintClaim(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Claimable.Size()
		i -= size
		if _, err := m.Claimable.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaim(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaim(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InitialClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	if m.GoalId != 0 {
		n += 1 + sovClaim(uint64(m.GoalId))
	}
	return n
}

func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	l = m.Claimable.Size()
	n += 1 + l + sovClaim(uint64(l))
	if len(m.CompletedGoals) > 0 {
		l = 0
		for _, e := range m.CompletedGoals {
			l += sovClaim(uint64(e))
		}
		n += 1 + sovClaim(uint64(l)) + l
	}
	if len(m.ClaimedGoals) > 0 {
		l = 0
		for _, e := range m.ClaimedGoals {
			l += sovClaim(uint64(e))
		}
		n += 1 + sovClaim(uint64(l)) + l
	}
	return n
}

func sovClaim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaim(x uint64) (n int) {
	return sovClaim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InitialClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
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
			return fmt.Errorf("proto: InitialClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InitialClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoalId", wireType)
			}
			m.GoalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GoalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
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
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
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
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
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
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Claimable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.CompletedGoals = append(m.CompletedGoals, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.CompletedGoals) == 0 {
					m.CompletedGoals = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.CompletedGoals = append(m.CompletedGoals, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletedGoals", wireType)
			}
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ClaimedGoals = append(m.ClaimedGoals, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ClaimedGoals) == 0 {
					m.ClaimedGoals = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ClaimedGoals = append(m.ClaimedGoals, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedGoals", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
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
func skipClaim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
				return 0, ErrInvalidLengthClaim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaim = fmt.Errorf("proto: unexpected end of group")
)
