// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/api/api_auth_webauthn.proto

package api

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type WebauthnState struct {
	Credentials []WebauthnCredential `protobuf:"bytes,1,rep,name=credentials,proto3" json:"credentials" xml:"credential"`
}

func (m *WebauthnState) Reset()         { *m = WebauthnState{} }
func (m *WebauthnState) String() string { return proto.CompactTextString(m) }
func (*WebauthnState) ProtoMessage()    {}
func (*WebauthnState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0236aafa0ec1587d, []int{0}
}
func (m *WebauthnState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WebauthnState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WebauthnState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WebauthnState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebauthnState.Merge(m, src)
}
func (m *WebauthnState) XXX_Size() int {
	return m.ProtoSize()
}
func (m *WebauthnState) XXX_DiscardUnknown() {
	xxx_messageInfo_WebauthnState.DiscardUnknown(m)
}

var xxx_messageInfo_WebauthnState proto.InternalMessageInfo

type WebauthnCredential struct {
	ID            string    `protobuf:"bytes,1,opt,name=ID,proto3" json:"id" xml:"id"`
	RpId          string    `protobuf:"bytes,2,opt,name=rp_id,json=rpId,proto3" json:"rpId" xml:"rpId"`
	Nickname      string    `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname" xml:"nickname"`
	PublicKeyCose string    `protobuf:"bytes,4,opt,name=public_key_cose,json=publicKeyCose,proto3" json:"publicKeyCose" xml:"publicKeyCose"`
	SignCount     uint32    `protobuf:"varint,5,opt,name=sign_count,json=signCount,proto3" json:"signCount" xml:"signCount"`
	Transports    []string  `protobuf:"bytes,6,rep,name=transports,proto3" json:"transports" xml:"transport"`
	RequireUv     bool      `protobuf:"varint,7,opt,name=require_uv,json=requireUv,proto3" json:"requireUv" xml:"requireUv"`
	CreateTime    time.Time `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3,stdtime" json:"createTime" xml:"createTime"`
	LastUseTime   time.Time `protobuf:"bytes,9,opt,name=last_use_time,json=lastUseTime,proto3,stdtime" json:"lastUseTime" xml:"lastUseTime"`
}

func (m *WebauthnCredential) Reset()         { *m = WebauthnCredential{} }
func (m *WebauthnCredential) String() string { return proto.CompactTextString(m) }
func (*WebauthnCredential) ProtoMessage()    {}
func (*WebauthnCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_0236aafa0ec1587d, []int{1}
}
func (m *WebauthnCredential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WebauthnCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WebauthnCredential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WebauthnCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebauthnCredential.Merge(m, src)
}
func (m *WebauthnCredential) XXX_Size() int {
	return m.ProtoSize()
}
func (m *WebauthnCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_WebauthnCredential.DiscardUnknown(m)
}

var xxx_messageInfo_WebauthnCredential proto.InternalMessageInfo

func init() {
	proto.RegisterType((*WebauthnState)(nil), "api.WebauthnState")
	proto.RegisterType((*WebauthnCredential)(nil), "api.WebauthnCredential")
}

func init() { proto.RegisterFile("lib/api/api_auth_webauthn.proto", fileDescriptor_0236aafa0ec1587d) }

var fileDescriptor_0236aafa0ec1587d = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0xed, 0xa4, 0x2d, 0xc9, 0x45, 0xa1, 0x70, 0x48, 0x60, 0x75, 0xf0, 0x99, 0x1b, 0xc0,
	0x02, 0xe4, 0x48, 0x65, 0xeb, 0x82, 0x70, 0x3a, 0x10, 0xb1, 0x20, 0x43, 0x85, 0xc4, 0x62, 0xd9,
	0xce, 0xe1, 0x9c, 0xea, 0xd8, 0xc6, 0x77, 0x2e, 0xe4, 0x53, 0xd0, 0x8f, 0xc0, 0xc7, 0xc9, 0x98,
	0x91, 0xe9, 0x50, 0x9b, 0xcd, 0x63, 0x3e, 0x01, 0xf2, 0xf9, 0x4f, 0x5c, 0x31, 0x30, 0x44, 0xb1,
	0x7f, 0xcf, 0xf3, 0x3e, 0xef, 0x3b, 0x3c, 0x06, 0x28, 0xa2, 0xfe, 0xc4, 0x4b, 0x69, 0xf9, 0x73,
	0xbd, 0x9c, 0x2f, 0xdc, 0xef, 0xc4, 0x2f, 0xff, 0x63, 0x2b, 0xcd, 0x12, 0x9e, 0xc0, 0xbe, 0x97,
	0xd2, 0x13, 0x14, 0x26, 0x49, 0x18, 0x91, 0x89, 0x44, 0x7e, 0xfe, 0x75, 0xc2, 0xe9, 0x92, 0x30,
	0xee, 0x2d, 0xd3, 0xca, 0x85, 0x19, 0x18, 0x7f, 0xae, 0xe7, 0x3e, 0x72, 0x8f, 0x13, 0xe8, 0x83,
	0x51, 0x90, 0x91, 0x39, 0x89, 0x39, 0xf5, 0x22, 0xa6, 0xa9, 0x46, 0xdf, 0x1c, 0x9d, 0x3e, 0xb1,
	0xbc, 0x94, 0x5a, 0x8d, 0x71, 0xda, 0xea, 0xf6, 0xf3, 0xb5, 0x40, 0x4a, 0x21, 0x50, 0x77, 0x66,
	0x27, 0xd0, 0x83, 0x1f, 0xcb, 0xe8, 0x0c, 0xef, 0x19, 0x76, 0xba, 0x06, 0xfc, 0xf3, 0x10, 0xc0,
	0x7f, 0xc3, 0x20, 0x06, 0xbd, 0xd9, 0xb9, 0xa6, 0x1a, 0xaa, 0x39, 0xb4, 0x61, 0x21, 0x50, 0x8f,
	0xce, 0x77, 0x02, 0x0d, 0x64, 0x16, 0x9d, 0x63, 0xa7, 0x37, 0x3b, 0x87, 0x2f, 0xc1, 0x61, 0x96,
	0xba, 0x74, 0xae, 0xf5, 0xa4, 0xed, 0x71, 0x21, 0xd0, 0x41, 0x96, 0xce, 0x4a, 0x23, 0x90, 0xc6,
	0xf2, 0x05, 0x3b, 0x92, 0xc1, 0x33, 0x30, 0x88, 0x69, 0x70, 0x19, 0x7b, 0x4b, 0xa2, 0xf5, 0xa5,
	0x5f, 0x2f, 0x04, 0x6a, 0xd9, 0x4e, 0xa0, 0xfb, 0x72, 0xa6, 0x01, 0xd8, 0x69, 0x35, 0xf8, 0x01,
	0x1c, 0xa7, 0xb9, 0x1f, 0xd1, 0xc0, 0xbd, 0x24, 0x2b, 0x37, 0x48, 0x18, 0xd1, 0x0e, 0x64, 0x84,
	0x59, 0x08, 0x34, 0xae, 0xa4, 0xf7, 0x64, 0x35, 0x4d, 0x58, 0x99, 0xf3, 0x48, 0xe6, 0xdc, 0xa1,
	0xd8, 0xb9, 0xeb, 0x82, 0x6f, 0x00, 0x60, 0x34, 0x8c, 0xdd, 0x20, 0xc9, 0x63, 0xae, 0x1d, 0x1a,
	0xaa, 0x39, 0xb6, 0x8d, 0x42, 0xa0, 0x61, 0x49, 0xa7, 0x25, 0xdc, 0x09, 0x74, 0x2c, 0x83, 0x5a,
	0x82, 0x9d, 0xbd, 0x0a, 0xdf, 0x02, 0xc0, 0x33, 0x2f, 0x66, 0x69, 0x92, 0x71, 0xa6, 0x1d, 0x19,
	0x7d, 0x73, 0x68, 0x3f, 0x2d, 0x04, 0xea, 0xd0, 0x36, 0xa1, 0x45, 0xd8, 0xe9, 0xc8, 0xe5, 0x0d,
	0x19, 0xf9, 0x96, 0xd3, 0x8c, 0xb8, 0xf9, 0x95, 0x76, 0xcf, 0x50, 0xcd, 0x41, 0x75, 0x43, 0x4d,
	0x2f, 0xae, 0xda, 0x84, 0x96, 0x60, 0x67, 0xaf, 0xc2, 0x50, 0xd6, 0xc3, 0xe3, 0xc4, 0x2d, 0x9b,
	0xa4, 0x0d, 0x0c, 0xd5, 0x1c, 0x9d, 0x9e, 0x58, 0x55, 0xcd, 0xac, 0xa6, 0x66, 0xd6, 0xa7, 0xa6,
	0x66, 0xf6, 0x8b, 0xba, 0x21, 0xa0, 0x1a, 0x2b, 0x85, 0x6e, 0x41, 0x6a, 0x84, 0xaf, 0xff, 0x20,
	0xd5, 0xe9, 0x78, 0x60, 0x04, 0xc6, 0x91, 0xc7, 0xb8, 0x9b, 0xb3, 0x7a, 0xd5, 0xf0, 0xbf, 0xab,
	0x5e, 0x35, 0x65, 0x2c, 0x07, 0x2f, 0x58, 0xb3, 0xeb, 0xa1, 0xdc, 0xd5, 0x61, 0xd5, 0xb2, 0xae,
	0xcb, 0x7e, 0xb7, 0xbe, 0xd1, 0x95, 0xcd, 0x8d, 0xae, 0xac, 0x6f, 0x75, 0x75, 0x73, 0xab, 0xab,
	0xd7, 0x5b, 0x5d, 0xf9, 0xb5, 0xd5, 0xd5, 0xcd, 0x56, 0x57, 0x7e, 0x6f, 0x75, 0xe5, 0xcb, 0xb3,
	0x90, 0xf2, 0x45, 0xee, 0x5b, 0x41, 0xb2, 0x9c, 0xb0, 0x55, 0x1c, 0xf0, 0x05, 0x8d, 0xc3, 0xce,
	0x53, 0xfd, 0x39, 0xfa, 0x47, 0xf2, 0xb0, 0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x89,
	0x6d, 0xa3, 0xa0, 0x03, 0x00, 0x00,
}

func (m *WebauthnState) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WebauthnState) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WebauthnState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Credentials) > 0 {
		for iNdEx := len(m.Credentials) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Credentials[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApiAuthWebauthn(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *WebauthnCredential) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WebauthnCredential) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WebauthnCredential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastUseTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastUseTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintApiAuthWebauthn(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x4a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreateTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreateTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintApiAuthWebauthn(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x42
	if m.RequireUv {
		i--
		if m.RequireUv {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.Transports) > 0 {
		for iNdEx := len(m.Transports) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Transports[iNdEx])
			copy(dAtA[i:], m.Transports[iNdEx])
			i = encodeVarintApiAuthWebauthn(dAtA, i, uint64(len(m.Transports[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.SignCount != 0 {
		i = encodeVarintApiAuthWebauthn(dAtA, i, uint64(m.SignCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.PublicKeyCose) > 0 {
		i -= len(m.PublicKeyCose)
		copy(dAtA[i:], m.PublicKeyCose)
		i = encodeVarintApiAuthWebauthn(dAtA, i, uint64(len(m.PublicKeyCose)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Nickname) > 0 {
		i -= len(m.Nickname)
		copy(dAtA[i:], m.Nickname)
		i = encodeVarintApiAuthWebauthn(dAtA, i, uint64(len(m.Nickname)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RpId) > 0 {
		i -= len(m.RpId)
		copy(dAtA[i:], m.RpId)
		i = encodeVarintApiAuthWebauthn(dAtA, i, uint64(len(m.RpId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintApiAuthWebauthn(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApiAuthWebauthn(dAtA []byte, offset int, v uint64) int {
	offset -= sovApiAuthWebauthn(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WebauthnState) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Credentials) > 0 {
		for _, e := range m.Credentials {
			l = e.ProtoSize()
			n += 1 + l + sovApiAuthWebauthn(uint64(l))
		}
	}
	return n
}

func (m *WebauthnCredential) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovApiAuthWebauthn(uint64(l))
	}
	l = len(m.RpId)
	if l > 0 {
		n += 1 + l + sovApiAuthWebauthn(uint64(l))
	}
	l = len(m.Nickname)
	if l > 0 {
		n += 1 + l + sovApiAuthWebauthn(uint64(l))
	}
	l = len(m.PublicKeyCose)
	if l > 0 {
		n += 1 + l + sovApiAuthWebauthn(uint64(l))
	}
	if m.SignCount != 0 {
		n += 1 + sovApiAuthWebauthn(uint64(m.SignCount))
	}
	if len(m.Transports) > 0 {
		for _, s := range m.Transports {
			l = len(s)
			n += 1 + l + sovApiAuthWebauthn(uint64(l))
		}
	}
	if m.RequireUv {
		n += 2
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreateTime)
	n += 1 + l + sovApiAuthWebauthn(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastUseTime)
	n += 1 + l + sovApiAuthWebauthn(uint64(l))
	return n
}

func sovApiAuthWebauthn(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApiAuthWebauthn(x uint64) (n int) {
	return sovApiAuthWebauthn(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WebauthnState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiAuthWebauthn
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
			return fmt.Errorf("proto: WebauthnState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WebauthnState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credentials", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiAuthWebauthn
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
				return ErrInvalidLengthApiAuthWebauthn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApiAuthWebauthn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credentials = append(m.Credentials, WebauthnCredential{})
			if err := m.Credentials[len(m.Credentials)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiAuthWebauthn(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApiAuthWebauthn
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
func (m *WebauthnCredential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiAuthWebauthn
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
			return fmt.Errorf("proto: WebauthnCredential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WebauthnCredential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiAuthWebauthn
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
				return ErrInvalidLengthApiAuthWebauthn
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiAuthWebauthn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiAuthWebauthn
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
				return ErrInvalidLengthApiAuthWebauthn
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiAuthWebauthn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RpId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nickname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiAuthWebauthn
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
				return ErrInvalidLengthApiAuthWebauthn
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiAuthWebauthn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nickname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKeyCose", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiAuthWebauthn
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
				return ErrInvalidLengthApiAuthWebauthn
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiAuthWebauthn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKeyCose = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignCount", wireType)
			}
			m.SignCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiAuthWebauthn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transports", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiAuthWebauthn
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
				return ErrInvalidLengthApiAuthWebauthn
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiAuthWebauthn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transports = append(m.Transports, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequireUv", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiAuthWebauthn
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
			m.RequireUv = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiAuthWebauthn
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
				return ErrInvalidLengthApiAuthWebauthn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApiAuthWebauthn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreateTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUseTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiAuthWebauthn
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
				return ErrInvalidLengthApiAuthWebauthn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApiAuthWebauthn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastUseTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiAuthWebauthn(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApiAuthWebauthn
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
func skipApiAuthWebauthn(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApiAuthWebauthn
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
					return 0, ErrIntOverflowApiAuthWebauthn
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
					return 0, ErrIntOverflowApiAuthWebauthn
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
				return 0, ErrInvalidLengthApiAuthWebauthn
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApiAuthWebauthn
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApiAuthWebauthn
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApiAuthWebauthn        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApiAuthWebauthn          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApiAuthWebauthn = fmt.Errorf("proto: unexpected end of group")
)
