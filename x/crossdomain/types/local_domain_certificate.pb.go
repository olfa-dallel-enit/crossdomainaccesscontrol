// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crossdomain/local_domain_certificate.proto

package types

import (
	fmt "fmt"
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

type LocalDomainCertificate struct {
	Value   string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *LocalDomainCertificate) Reset()         { *m = LocalDomainCertificate{} }
func (m *LocalDomainCertificate) String() string { return proto.CompactTextString(m) }
func (*LocalDomainCertificate) ProtoMessage()    {}
func (*LocalDomainCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f13646ebb3084cd, []int{0}
}
func (m *LocalDomainCertificate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocalDomainCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocalDomainCertificate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LocalDomainCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalDomainCertificate.Merge(m, src)
}
func (m *LocalDomainCertificate) XXX_Size() int {
	return m.Size()
}
func (m *LocalDomainCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalDomainCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_LocalDomainCertificate proto.InternalMessageInfo

func (m *LocalDomainCertificate) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *LocalDomainCertificate) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*LocalDomainCertificate)(nil), "crossdomain.crossdomain.LocalDomainCertificate")
}

func init() {
	proto.RegisterFile("crossdomain/local_domain_certificate.proto", fileDescriptor_1f13646ebb3084cd)
}

var fileDescriptor_1f13646ebb3084cd = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4a, 0x2e, 0xca, 0x2f,
	0x2e, 0x4e, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0xd3, 0xcf, 0xc9, 0x4f, 0x4e, 0xcc, 0x89, 0x87, 0x70,
	0xe2, 0x93, 0x53, 0x8b, 0x4a, 0x32, 0xd3, 0x32, 0x93, 0x13, 0x4b, 0x52, 0xf5, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0xc4, 0x91, 0xd4, 0xea, 0x21, 0xb1, 0x95, 0x3c, 0xb8, 0xc4, 0x7c, 0x40, 0x5a,
	0x5d, 0xc0, 0x5c, 0x67, 0x84, 0x46, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x82, 0x8b, 0x3d, 0xb9, 0x28, 0x35, 0xb1,
	0x24, 0xbf, 0x48, 0x82, 0x09, 0x2c, 0x0e, 0xe3, 0x3a, 0x59, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70,
	0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x3c, 0xb2, 0x43, 0x2b, 0xf4, 0x91, 0x79, 0x25, 0x95, 0x05, 0xa9,
	0xc5, 0x49, 0x6c, 0x60, 0x47, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x51, 0x26, 0x6b, 0x69,
	0xd2, 0x00, 0x00, 0x00,
}

func (m *LocalDomainCertificate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalDomainCertificate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LocalDomainCertificate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintLocalDomainCertificate(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintLocalDomainCertificate(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLocalDomainCertificate(dAtA []byte, offset int, v uint64) int {
	offset -= sovLocalDomainCertificate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LocalDomainCertificate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovLocalDomainCertificate(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovLocalDomainCertificate(uint64(l))
	}
	return n
}

func sovLocalDomainCertificate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLocalDomainCertificate(x uint64) (n int) {
	return sovLocalDomainCertificate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LocalDomainCertificate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocalDomainCertificate
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
			return fmt.Errorf("proto: LocalDomainCertificate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalDomainCertificate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalDomainCertificate
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
				return ErrInvalidLengthLocalDomainCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocalDomainCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalDomainCertificate
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
				return ErrInvalidLengthLocalDomainCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocalDomainCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLocalDomainCertificate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLocalDomainCertificate
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
func skipLocalDomainCertificate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLocalDomainCertificate
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
					return 0, ErrIntOverflowLocalDomainCertificate
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
					return 0, ErrIntOverflowLocalDomainCertificate
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
				return 0, ErrInvalidLengthLocalDomainCertificate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLocalDomainCertificate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLocalDomainCertificate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLocalDomainCertificate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLocalDomainCertificate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLocalDomainCertificate = fmt.Errorf("proto: unexpected end of group")
)
