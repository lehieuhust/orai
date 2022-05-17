// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/airesult/types/query_reward.proto

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

// QueryRewardReq is the request type for the Query/QueryReward RPC method
type QueryRewardReq struct {
	BlockHeight string `protobuf:"bytes,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *QueryRewardReq) Reset()         { *m = QueryRewardReq{} }
func (m *QueryRewardReq) String() string { return proto.CompactTextString(m) }
func (*QueryRewardReq) ProtoMessage()    {}
func (*QueryRewardReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_eda362c64f8ca806, []int{0}
}
func (m *QueryRewardReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRewardReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRewardReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRewardReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRewardReq.Merge(m, src)
}
func (m *QueryRewardReq) XXX_Size() int {
	return m.Size()
}
func (m *QueryRewardReq) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRewardReq.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRewardReq proto.InternalMessageInfo

func (m *QueryRewardReq) GetBlockHeight() string {
	if m != nil {
		return m.BlockHeight
	}
	return ""
}

// QueryRewardRes is the response type for the Query/QueryReward RPC method
type QueryRewardRes struct {
	Reward Reward `protobuf:"bytes,1,opt,name=reward,proto3" json:"reward" json:"block_reward"`
}

func (m *QueryRewardRes) Reset()         { *m = QueryRewardRes{} }
func (m *QueryRewardRes) String() string { return proto.CompactTextString(m) }
func (*QueryRewardRes) ProtoMessage()    {}
func (*QueryRewardRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_eda362c64f8ca806, []int{1}
}
func (m *QueryRewardRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRewardRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRewardRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRewardRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRewardRes.Merge(m, src)
}
func (m *QueryRewardRes) XXX_Size() int {
	return m.Size()
}
func (m *QueryRewardRes) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRewardRes.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRewardRes proto.InternalMessageInfo

func (m *QueryRewardRes) GetReward() Reward {
	if m != nil {
		return m.Reward
	}
	return Reward{}
}

func init() {
	proto.RegisterType((*QueryRewardReq)(nil), "oraichain.orai.airesult.QueryRewardReq")
	proto.RegisterType((*QueryRewardRes)(nil), "oraichain.orai.airesult.QueryRewardRes")
}

func init() {
	proto.RegisterFile("x/airesult/types/query_reward.proto", fileDescriptor_eda362c64f8ca806)
}

var fileDescriptor_eda362c64f8ca806 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xae, 0xd0, 0x4f, 0xcc,
	0x2c, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x2c, 0x4d,
	0x2d, 0xaa, 0x8c, 0x2f, 0x4a, 0x2d, 0x4f, 0x2c, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0xcf, 0x2f, 0x4a, 0xcc, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x03, 0xb1, 0xf4, 0x60, 0x3a,
	0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x6a, 0xf4, 0x41, 0x2c, 0x88, 0x72, 0x29, 0x4c, 0x33,
	0xc1, 0x24, 0x8a, 0x99, 0x4a, 0x96, 0x5c, 0x7c, 0x81, 0x20, 0x9b, 0x82, 0xc0, 0x82, 0x41, 0xa9,
	0x85, 0x42, 0x8a, 0x5c, 0x3c, 0x49, 0x39, 0xf9, 0xc9, 0xd9, 0xf1, 0x19, 0xa9, 0x99, 0xe9, 0x19,
	0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xdc, 0x60, 0x31, 0x0f, 0xb0, 0x90, 0x15, 0xcb,
	0x8c, 0x05, 0xf2, 0x8c, 0x4a, 0x69, 0x68, 0x5a, 0x8b, 0x85, 0x42, 0xb8, 0xd8, 0x20, 0x86, 0x83,
	0x35, 0x71, 0x1b, 0xc9, 0xeb, 0xe1, 0x70, 0xb1, 0x1e, 0x44, 0x8f, 0x93, 0xf4, 0x89, 0x7b, 0xf2,
	0x0c, 0x9f, 0xee, 0xc9, 0x0b, 0x67, 0x15, 0xe7, 0xe7, 0x59, 0x29, 0x41, 0xec, 0x85, 0x18, 0xa1,
	0x14, 0x04, 0x35, 0xcb, 0xc9, 0xe5, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2,
	0xb4, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xe1, 0x36, 0x81, 0x59,
	0xfa, 0xe8, 0x7e, 0x4f, 0x62, 0x03, 0xfb, 0xd7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x53, 0x29,
	0xdc, 0xff, 0x6a, 0x01, 0x00, 0x00,
}

func (m *QueryRewardReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRewardReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRewardReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BlockHeight) > 0 {
		i -= len(m.BlockHeight)
		copy(dAtA[i:], m.BlockHeight)
		i = encodeVarintQueryReward(dAtA, i, uint64(len(m.BlockHeight)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryRewardRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRewardRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRewardRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Reward.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryReward(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQueryReward(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryReward(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRewardReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BlockHeight)
	if l > 0 {
		n += 1 + l + sovQueryReward(uint64(l))
	}
	return n
}

func (m *QueryRewardRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Reward.Size()
	n += 1 + l + sovQueryReward(uint64(l))
	return n
}

func sovQueryReward(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryReward(x uint64) (n int) {
	return sovQueryReward(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRewardReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryReward
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
			return fmt.Errorf("proto: QueryRewardReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRewardReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryReward
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
				return ErrInvalidLengthQueryReward
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHeight = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueryReward
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQueryReward
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
func (m *QueryRewardRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryReward
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
			return fmt.Errorf("proto: QueryRewardRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRewardRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryReward
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
				return ErrInvalidLengthQueryReward
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueryReward
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQueryReward
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
func skipQueryReward(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryReward
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
					return 0, ErrIntOverflowQueryReward
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
					return 0, ErrIntOverflowQueryReward
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
				return 0, ErrInvalidLengthQueryReward
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryReward
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryReward
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryReward        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryReward          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryReward = fmt.Errorf("proto: unexpected end of group")
)
