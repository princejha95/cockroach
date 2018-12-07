// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: build/info.proto

package build

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Info describes build information for this CockroachDB binary.
type Info struct {
	GoVersion       string `protobuf:"bytes,1,opt,name=go_version,json=goVersion" json:"go_version"`
	Tag             string `protobuf:"bytes,2,opt,name=tag" json:"tag"`
	Time            string `protobuf:"bytes,3,opt,name=time" json:"time"`
	Revision        string `protobuf:"bytes,4,opt,name=revision" json:"revision"`
	CgoCompiler     string `protobuf:"bytes,5,opt,name=cgo_compiler,json=cgoCompiler" json:"cgo_compiler"`
	CgoTargetTriple string `protobuf:"bytes,10,opt,name=cgo_target_triple,json=cgoTargetTriple" json:"cgo_target_triple"`
	Platform        string `protobuf:"bytes,6,opt,name=platform" json:"platform"`
	Distribution    string `protobuf:"bytes,7,opt,name=distribution" json:"distribution"`
	Type            string `protobuf:"bytes,8,opt,name=type" json:"type"`
	Channel         string `protobuf:"bytes,9,opt,name=channel" json:"channel"`
	// dependencies exists to allow tests that run against old clusters
	// to unmarshal JSON containing this field. The tag is unimportant,
	// but the field name must remain unchanged.
	//
	// alternatively, we could set jsonpb.Unmarshaler.AllowUnknownFields
	// to true in httputil.doJSONRequest, but that comes at the expense
	// of run-time type checking, which is nice to have.
	Dependencies         *string  `protobuf:"bytes,10000,opt,name=dependencies" json:"dependencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_info_6ca0bae8d1d234a0, []int{0}
}
func (m *Info) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(dst, src)
}
func (m *Info) XXX_Size() int {
	return m.Size()
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Info)(nil), "cockroach.build.Info")
}
func (m *Info) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Info) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.GoVersion)))
	i += copy(dAtA[i:], m.GoVersion)
	dAtA[i] = 0x12
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.Tag)))
	i += copy(dAtA[i:], m.Tag)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.Time)))
	i += copy(dAtA[i:], m.Time)
	dAtA[i] = 0x22
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.Revision)))
	i += copy(dAtA[i:], m.Revision)
	dAtA[i] = 0x2a
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.CgoCompiler)))
	i += copy(dAtA[i:], m.CgoCompiler)
	dAtA[i] = 0x32
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.Platform)))
	i += copy(dAtA[i:], m.Platform)
	dAtA[i] = 0x3a
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.Distribution)))
	i += copy(dAtA[i:], m.Distribution)
	dAtA[i] = 0x42
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.Type)))
	i += copy(dAtA[i:], m.Type)
	dAtA[i] = 0x4a
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.Channel)))
	i += copy(dAtA[i:], m.Channel)
	dAtA[i] = 0x52
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.CgoTargetTriple)))
	i += copy(dAtA[i:], m.CgoTargetTriple)
	if m.Dependencies != nil {
		dAtA[i] = 0x82
		i++
		dAtA[i] = 0xf1
		i++
		dAtA[i] = 0x4
		i++
		i = encodeVarintInfo(dAtA, i, uint64(len(*m.Dependencies)))
		i += copy(dAtA[i:], *m.Dependencies)
	}
	return i, nil
}

func encodeVarintInfo(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Info) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GoVersion)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.Tag)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.Time)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.Revision)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.CgoCompiler)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.Platform)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.Distribution)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.Type)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.Channel)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.CgoTargetTriple)
	n += 1 + l + sovInfo(uint64(l))
	if m.Dependencies != nil {
		l = len(*m.Dependencies)
		n += 3 + l + sovInfo(uint64(l))
	}
	return n
}

func sovInfo(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInfo(x uint64) (n int) {
	return sovInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Info) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Revision = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CgoCompiler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CgoCompiler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Platform", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Platform = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Distribution", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Distribution = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CgoTargetTriple", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CgoTargetTriple = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10000:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dependencies", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Dependencies = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfo
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
func skipInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInfo
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
					return 0, ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInfo
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthInfo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInfo
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipInfo(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthInfo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInfo   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("build/info.proto", fileDescriptor_info_6ca0bae8d1d234a0) }

var fileDescriptor_info_6ca0bae8d1d234a0 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xb1, 0x4e, 0x32, 0x41,
	0x14, 0x85, 0xd9, 0x9f, 0xe5, 0x07, 0xae, 0x24, 0xe8, 0xc4, 0x98, 0x89, 0xc5, 0x42, 0xa4, 0x90,
	0x0a, 0x7c, 0x06, 0xac, 0x6c, 0x2c, 0x08, 0xb1, 0xb0, 0x21, 0xcb, 0xec, 0x65, 0x98, 0xb8, 0xcc,
	0x9d, 0x0c, 0x03, 0x09, 0x0f, 0x60, 0xcf, 0x63, 0x51, 0x5a, 0x5a, 0x19, 0x5d, 0x5f, 0xc4, 0xec,
	0x20, 0x66, 0xe8, 0x26, 0xdf, 0x77, 0x72, 0xe6, 0x24, 0x17, 0xce, 0x67, 0x6b, 0x95, 0x67, 0x43,
	0xa5, 0xe7, 0x34, 0x30, 0x96, 0x1c, 0xb1, 0xb6, 0x20, 0xf1, 0x62, 0x29, 0x15, 0x8b, 0x81, 0x77,
	0xd7, 0x97, 0x92, 0x24, 0x79, 0x37, 0x2c, 0x5f, 0x87, 0xd8, 0xcd, 0x6b, 0x15, 0xe2, 0x07, 0x3d,
	0x27, 0xd6, 0x03, 0x90, 0x34, 0xdd, 0xa0, 0x5d, 0x29, 0xd2, 0x3c, 0xea, 0x46, 0xfd, 0xe6, 0x28,
	0xde, 0x7f, 0x74, 0x2a, 0xe3, 0xa6, 0xa4, 0xa7, 0x03, 0x66, 0x57, 0x50, 0x75, 0xa9, 0xe4, 0xff,
	0x02, 0x5b, 0x02, 0xc6, 0x21, 0x76, 0x6a, 0x89, 0xbc, 0x1a, 0x08, 0x4f, 0x58, 0x17, 0x1a, 0x16,
	0x37, 0xca, 0x97, 0xc6, 0x81, 0xfd, 0xa3, 0xec, 0x16, 0x5a, 0x42, 0xd2, 0x54, 0xd0, 0xd2, 0xa8,
	0x1c, 0x2d, 0xaf, 0x05, 0xa9, 0x33, 0x21, 0xe9, 0xfe, 0x57, 0x94, 0x55, 0x26, 0x4f, 0xdd, 0x9c,
	0xec, 0x92, 0xff, 0x0f, 0xab, 0x8e, 0x94, 0xf5, 0xa1, 0x95, 0xa9, 0x95, 0xb3, 0x6a, 0xb6, 0x76,
	0xe5, 0x87, 0xf5, 0x20, 0x75, 0x62, 0xfc, 0xe0, 0xad, 0x41, 0xde, 0x38, 0x19, 0xbc, 0x35, 0xc8,
	0x12, 0xa8, 0x8b, 0x45, 0xaa, 0x35, 0xe6, 0xbc, 0x19, 0xc8, 0x23, 0x64, 0x77, 0x70, 0x51, 0xce,
	0x75, 0xa9, 0x95, 0xe8, 0xa6, 0xce, 0x2a, 0x93, 0x23, 0x87, 0x20, 0xd9, 0x16, 0x92, 0x26, 0xde,
	0x4e, 0xbc, 0x64, 0x3d, 0x68, 0x65, 0x68, 0x50, 0x67, 0xa8, 0x85, 0xc2, 0x15, 0xdf, 0x3d, 0x96,
	0xe9, 0xf1, 0x09, 0x1c, 0x75, 0xf6, 0x5f, 0x49, 0x65, 0x5f, 0x24, 0xd1, 0x5b, 0x91, 0x44, 0xef,
	0x45, 0x12, 0x7d, 0x16, 0x49, 0xb4, 0xfb, 0x4e, 0x2a, 0xcf, 0x35, 0x7f, 0xbe, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xec, 0xe6, 0x1c, 0x81, 0xe2, 0x01, 0x00, 0x00,
}
