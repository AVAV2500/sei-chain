// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/deposit.proto

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

type DepositInfoEntry struct {
	Creator string                                 `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator"`
	Denom   string                                 `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom"`
	Amount  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"amount" yaml:"amount"`
}

func (m *DepositInfoEntry) Reset()         { *m = DepositInfoEntry{} }
func (m *DepositInfoEntry) String() string { return proto.CompactTextString(m) }
func (*DepositInfoEntry) ProtoMessage()    {}
func (*DepositInfoEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_93caea15661b640e, []int{0}
}
func (m *DepositInfoEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositInfoEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositInfoEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositInfoEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositInfoEntry.Merge(m, src)
}
func (m *DepositInfoEntry) XXX_Size() int {
	return m.Size()
}
func (m *DepositInfoEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositInfoEntry.DiscardUnknown(m)
}

var xxx_messageInfo_DepositInfoEntry proto.InternalMessageInfo

func (m *DepositInfoEntry) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *DepositInfoEntry) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*DepositInfoEntry)(nil), "seiprotocol.seichain.dex.DepositInfoEntry")
}

func init() { proto.RegisterFile("dex/deposit.proto", fileDescriptor_93caea15661b640e) }

var fileDescriptor_93caea15661b640e = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0x31, 0x4f, 0xf3, 0x30,
	0x14, 0x8c, 0xbf, 0x4f, 0x14, 0x35, 0x08, 0x01, 0x11, 0x43, 0xd4, 0xc1, 0x46, 0x95, 0x40, 0x2c,
	0x4d, 0x06, 0x36, 0xc6, 0xaa, 0x08, 0x75, 0xed, 0xc8, 0x96, 0x3a, 0x26, 0xb5, 0xa8, 0xfd, 0xa2,
	0xd8, 0x91, 0x92, 0x7f, 0xc1, 0x1f, 0x62, 0xef, 0xd8, 0x11, 0x31, 0x58, 0x28, 0xd9, 0x32, 0xf2,
	0x0b, 0x50, 0xec, 0x44, 0x62, 0x7a, 0xf7, 0xee, 0xee, 0xbd, 0x93, 0xce, 0xbf, 0x4a, 0x59, 0x15,
	0xa7, 0x2c, 0x07, 0xc5, 0x75, 0x94, 0x17, 0xa0, 0x21, 0x08, 0x15, 0xe3, 0x16, 0x51, 0xd8, 0x47,
	0x8a, 0x71, 0xba, 0x4b, 0xb8, 0x8c, 0x52, 0x56, 0xcd, 0xae, 0x33, 0xc8, 0xc0, 0x4a, 0x71, 0x8f,
	0x9c, 0x7f, 0x76, 0xd1, 0xbf, 0x60, 0xb2, 0x14, 0xca, 0x11, 0xf3, 0x0f, 0xe4, 0x5f, 0xae, 0xdc,
	0xcb, 0xb5, 0x7c, 0x85, 0x27, 0xa9, 0x8b, 0x3a, 0xb8, 0xf5, 0x4f, 0x69, 0xc1, 0x12, 0x0d, 0x45,
	0x88, 0x6e, 0xd0, 0xfd, 0x74, 0x79, 0xd6, 0x19, 0x32, 0x52, 0x9b, 0x11, 0x04, 0xc4, 0x3f, 0x49,
	0x99, 0x04, 0x11, 0xfe, 0xb3, 0xa6, 0x69, 0x67, 0x88, 0x23, 0x36, 0x6e, 0x04, 0x89, 0x3f, 0x49,
	0x04, 0x94, 0x52, 0x87, 0xff, 0xad, 0x63, 0x7d, 0x30, 0xc4, 0xfb, 0x32, 0xe4, 0x2e, 0xe3, 0x7a,
	0x57, 0x6e, 0x23, 0x0a, 0x22, 0xa6, 0xa0, 0x04, 0xa8, 0x61, 0x2c, 0x54, 0xfa, 0x16, 0xeb, 0x3a,
	0x67, 0x2a, 0x5a, 0x31, 0xda, 0x19, 0x32, 0xdc, 0xff, 0x18, 0x72, 0x5e, 0x27, 0x62, 0xff, 0x38,
	0x77, 0xfb, 0x7c, 0x33, 0x08, 0xcb, 0xe7, 0x43, 0x83, 0xd1, 0xb1, 0xc1, 0xe8, 0xbb, 0xc1, 0xe8,
	0xbd, 0xc5, 0xde, 0xb1, 0xc5, 0xde, 0x67, 0x8b, 0xbd, 0x97, 0xc5, 0x9f, 0x10, 0xc5, 0xf8, 0x62,
	0xac, 0xc9, 0x2e, 0xb6, 0xa7, 0xb8, 0x6f, 0xb3, 0x72, 0x79, 0xdb, 0x89, 0xd5, 0x1f, 0x7e, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x6d, 0x55, 0x42, 0x28, 0x65, 0x01, 0x00, 0x00,
}

func (m *DepositInfoEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositInfoEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositInfoEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeposit(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeposit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DepositInfoEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovDeposit(uint64(l))
	return n
}

func sovDeposit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeposit(x uint64) (n int) {
	return sovDeposit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DepositInfoEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeposit
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
			return fmt.Errorf("proto: DepositInfoEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositInfoEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeposit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeposit
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
func skipDeposit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeposit
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
					return 0, ErrIntOverflowDeposit
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
					return 0, ErrIntOverflowDeposit
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
				return 0, ErrInvalidLengthDeposit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeposit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeposit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeposit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeposit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeposit = fmt.Errorf("proto: unexpected end of group")
)
