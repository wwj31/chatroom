// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server.proto

package msg

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

type SessionBroken struct {
	Session string `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (m *SessionBroken) Reset()         { *m = SessionBroken{} }
func (m *SessionBroken) String() string { return proto.CompactTextString(m) }
func (*SessionBroken) ProtoMessage()    {}
func (*SessionBroken) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}
func (m *SessionBroken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionBroken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SessionBroken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SessionBroken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionBroken.Merge(m, src)
}
func (m *SessionBroken) XXX_Size() int {
	return m.Size()
}
func (m *SessionBroken) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionBroken.DiscardUnknown(m)
}

var xxx_messageInfo_SessionBroken proto.InternalMessageInfo

func (m *SessionBroken) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

type GateMsgWrapper struct {
	GateSession string `protobuf:"bytes,1,opt,name=gateSession,proto3" json:"gateSession,omitempty"`
	MsgName     string `protobuf:"bytes,2,opt,name=msgName,proto3" json:"msgName,omitempty"`
	Data        []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *GateMsgWrapper) Reset()         { *m = GateMsgWrapper{} }
func (m *GateMsgWrapper) String() string { return proto.CompactTextString(m) }
func (*GateMsgWrapper) ProtoMessage()    {}
func (*GateMsgWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}
func (m *GateMsgWrapper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GateMsgWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GateMsgWrapper.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GateMsgWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GateMsgWrapper.Merge(m, src)
}
func (m *GateMsgWrapper) XXX_Size() int {
	return m.Size()
}
func (m *GateMsgWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GateMsgWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GateMsgWrapper proto.InternalMessageInfo

func (m *GateMsgWrapper) GetGateSession() string {
	if m != nil {
		return m.GateSession
	}
	return ""
}

func (m *GateMsgWrapper) GetMsgName() string {
	if m != nil {
		return m.MsgName
	}
	return ""
}

func (m *GateMsgWrapper) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type UserEnterRoomReq struct {
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Session string `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
}

func (m *UserEnterRoomReq) Reset()         { *m = UserEnterRoomReq{} }
func (m *UserEnterRoomReq) String() string { return proto.CompactTextString(m) }
func (*UserEnterRoomReq) ProtoMessage()    {}
func (*UserEnterRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}
func (m *UserEnterRoomReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserEnterRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserEnterRoomReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserEnterRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEnterRoomReq.Merge(m, src)
}
func (m *UserEnterRoomReq) XXX_Size() int {
	return m.Size()
}
func (m *UserEnterRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEnterRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserEnterRoomReq proto.InternalMessageInfo

func (m *UserEnterRoomReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserEnterRoomReq) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

type UserEnterRoomResp struct {
	ErrCode INNER_ERROR `protobuf:"varint,1,opt,name=errCode,proto3,enum=msg.INNER_ERROR" json:"errCode,omitempty"`
	History []*ChatMsg  `protobuf:"bytes,2,rep,name=history,proto3" json:"history,omitempty"`
}

func (m *UserEnterRoomResp) Reset()         { *m = UserEnterRoomResp{} }
func (m *UserEnterRoomResp) String() string { return proto.CompactTextString(m) }
func (*UserEnterRoomResp) ProtoMessage()    {}
func (*UserEnterRoomResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{3}
}
func (m *UserEnterRoomResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserEnterRoomResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserEnterRoomResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserEnterRoomResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEnterRoomResp.Merge(m, src)
}
func (m *UserEnterRoomResp) XXX_Size() int {
	return m.Size()
}
func (m *UserEnterRoomResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEnterRoomResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserEnterRoomResp proto.InternalMessageInfo

func (m *UserEnterRoomResp) GetErrCode() INNER_ERROR {
	if m != nil {
		return m.ErrCode
	}
	return INNER_ERROR_IN_SUCCESS
}

func (m *UserEnterRoomResp) GetHistory() []*ChatMsg {
	if m != nil {
		return m.History
	}
	return nil
}

type EnterRoomNotifyGateway struct {
	Session string `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Room    string `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
}

func (m *EnterRoomNotifyGateway) Reset()         { *m = EnterRoomNotifyGateway{} }
func (m *EnterRoomNotifyGateway) String() string { return proto.CompactTextString(m) }
func (*EnterRoomNotifyGateway) ProtoMessage()    {}
func (*EnterRoomNotifyGateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{4}
}
func (m *EnterRoomNotifyGateway) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EnterRoomNotifyGateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EnterRoomNotifyGateway.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EnterRoomNotifyGateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterRoomNotifyGateway.Merge(m, src)
}
func (m *EnterRoomNotifyGateway) XXX_Size() int {
	return m.Size()
}
func (m *EnterRoomNotifyGateway) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterRoomNotifyGateway.DiscardUnknown(m)
}

var xxx_messageInfo_EnterRoomNotifyGateway proto.InternalMessageInfo

func (m *EnterRoomNotifyGateway) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *EnterRoomNotifyGateway) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

func init() {
	proto.RegisterType((*SessionBroken)(nil), "msg.SessionBroken")
	proto.RegisterType((*GateMsgWrapper)(nil), "msg.GateMsgWrapper")
	proto.RegisterType((*UserEnterRoomReq)(nil), "msg.UserEnterRoomReq")
	proto.RegisterType((*UserEnterRoomResp)(nil), "msg.UserEnterRoomResp")
	proto.RegisterType((*EnterRoomNotifyGateway)(nil), "msg.EnterRoomNotifyGateway")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xc1, 0x4e, 0x32, 0x31,
	0x18, 0x64, 0x81, 0xfc, 0xe4, 0x2f, 0x48, 0xb0, 0x07, 0xb3, 0xe1, 0xd0, 0x6c, 0xf6, 0x60, 0xd0,
	0x03, 0x26, 0xf8, 0x02, 0x06, 0x82, 0xc6, 0x03, 0x6b, 0x52, 0x62, 0x4c, 0xbc, 0x68, 0x95, 0xcf,
	0x42, 0x4c, 0xb7, 0xeb, 0xf7, 0x35, 0x1a, 0xde, 0xc2, 0xc7, 0xf2, 0xc8, 0xd1, 0xa3, 0x81, 0x17,
	0x31, 0xdd, 0x05, 0x03, 0x26, 0xde, 0x66, 0x66, 0x67, 0xa7, 0xd3, 0x29, 0x6b, 0x10, 0xe0, 0x2b,
	0x60, 0x37, 0x43, 0xeb, 0x2c, 0xaf, 0x18, 0xd2, 0xed, 0x3a, 0x20, 0xda, 0xb5, 0xd2, 0x6e, 0x3e,
	0x4e, 0x95, 0x43, 0x6b, 0x4d, 0xc1, 0xe3, 0x23, 0xb6, 0x37, 0x06, 0xa2, 0x99, 0x4d, 0xfb, 0x68,
	0x9f, 0x21, 0xe5, 0x21, 0xab, 0x51, 0x21, 0x84, 0x41, 0x14, 0x74, 0xfe, 0xcb, 0x0d, 0x8d, 0xef,
	0x59, 0xf3, 0x42, 0x39, 0x18, 0x91, 0xbe, 0x41, 0x95, 0x65, 0x80, 0x3c, 0x62, 0x75, 0xad, 0x1c,
	0x8c, 0x77, 0xfc, 0xdb, 0x92, 0x4f, 0x33, 0xa4, 0x13, 0x65, 0x20, 0x2c, 0x17, 0x69, 0x6b, 0xca,
	0x39, 0xab, 0x4e, 0x94, 0x53, 0x61, 0x25, 0x0a, 0x3a, 0x0d, 0x99, 0xe3, 0xf8, 0x8c, 0xb5, 0xae,
	0x09, 0x70, 0x98, 0x3a, 0x40, 0x69, 0xad, 0x91, 0xf0, 0xe2, 0x7d, 0xa9, 0xff, 0xbd, 0x08, 0xcf,
	0xf1, 0x76, 0xc7, 0xf2, 0x6e, 0x47, 0xcd, 0xf6, 0x7f, 0x25, 0x50, 0xc6, 0x8f, 0x59, 0x0d, 0x10,
	0x07, 0x76, 0x52, 0xa4, 0x34, 0x7b, 0xad, 0xae, 0x21, 0xdd, 0xbd, 0x4c, 0x92, 0xa1, 0xbc, 0x1b,
	0x4a, 0x79, 0x25, 0xe5, 0xc6, 0xc0, 0x0f, 0x59, 0x6d, 0x3a, 0x23, 0x67, 0x71, 0x1e, 0x96, 0xa3,
	0x4a, 0xa7, 0xde, 0x6b, 0xe4, 0xde, 0xc1, 0x54, 0xb9, 0x11, 0x69, 0xb9, 0xf9, 0x18, 0x9f, 0xb3,
	0x83, 0x9f, 0x43, 0x12, 0xeb, 0x66, 0x4f, 0x73, 0xbf, 0xcd, 0x9b, 0x9a, 0xff, 0x3d, 0xa0, 0xbf,
	0x8a, 0x5f, 0x7e, 0xdd, 0x39, 0xc7, 0x7d, 0xf1, 0xb1, 0x14, 0xc1, 0x62, 0x29, 0x82, 0xaf, 0xa5,
	0x08, 0xde, 0x57, 0xa2, 0xb4, 0x58, 0x89, 0xd2, 0xe7, 0x4a, 0x94, 0x6e, 0xab, 0x27, 0x86, 0xf4,
	0xc3, 0xbf, 0xfc, 0x99, 0x4e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xc4, 0x93, 0xea, 0xd8,
	0x01, 0x00, 0x00,
}

func (m *SessionBroken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionBroken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SessionBroken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Session) > 0 {
		i -= len(m.Session)
		copy(dAtA[i:], m.Session)
		i = encodeVarintServer(dAtA, i, uint64(len(m.Session)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GateMsgWrapper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GateMsgWrapper) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GateMsgWrapper) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintServer(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MsgName) > 0 {
		i -= len(m.MsgName)
		copy(dAtA[i:], m.MsgName)
		i = encodeVarintServer(dAtA, i, uint64(len(m.MsgName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GateSession) > 0 {
		i -= len(m.GateSession)
		copy(dAtA[i:], m.GateSession)
		i = encodeVarintServer(dAtA, i, uint64(len(m.GateSession)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserEnterRoomReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserEnterRoomReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserEnterRoomReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Session) > 0 {
		i -= len(m.Session)
		copy(dAtA[i:], m.Session)
		i = encodeVarintServer(dAtA, i, uint64(len(m.Session)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintServer(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserEnterRoomResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserEnterRoomResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserEnterRoomResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.History) > 0 {
		for iNdEx := len(m.History) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.History[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintServer(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.ErrCode != 0 {
		i = encodeVarintServer(dAtA, i, uint64(m.ErrCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EnterRoomNotifyGateway) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnterRoomNotifyGateway) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EnterRoomNotifyGateway) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Room) > 0 {
		i -= len(m.Room)
		copy(dAtA[i:], m.Room)
		i = encodeVarintServer(dAtA, i, uint64(len(m.Room)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Session) > 0 {
		i -= len(m.Session)
		copy(dAtA[i:], m.Session)
		i = encodeVarintServer(dAtA, i, uint64(len(m.Session)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintServer(dAtA []byte, offset int, v uint64) int {
	offset -= sovServer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SessionBroken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Session)
	if l > 0 {
		n += 1 + l + sovServer(uint64(l))
	}
	return n
}

func (m *GateMsgWrapper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GateSession)
	if l > 0 {
		n += 1 + l + sovServer(uint64(l))
	}
	l = len(m.MsgName)
	if l > 0 {
		n += 1 + l + sovServer(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovServer(uint64(l))
	}
	return n
}

func (m *UserEnterRoomReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovServer(uint64(l))
	}
	l = len(m.Session)
	if l > 0 {
		n += 1 + l + sovServer(uint64(l))
	}
	return n
}

func (m *UserEnterRoomResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ErrCode != 0 {
		n += 1 + sovServer(uint64(m.ErrCode))
	}
	if len(m.History) > 0 {
		for _, e := range m.History {
			l = e.Size()
			n += 1 + l + sovServer(uint64(l))
		}
	}
	return n
}

func (m *EnterRoomNotifyGateway) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Session)
	if l > 0 {
		n += 1 + l + sovServer(uint64(l))
	}
	l = len(m.Room)
	if l > 0 {
		n += 1 + l + sovServer(uint64(l))
	}
	return n
}

func sovServer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozServer(x uint64) (n int) {
	return sovServer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SessionBroken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
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
			return fmt.Errorf("proto: SessionBroken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionBroken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
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
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Session = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServer
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
func (m *GateMsgWrapper) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
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
			return fmt.Errorf("proto: GateMsgWrapper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GateMsgWrapper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GateSession", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
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
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GateSession = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
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
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServer
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
func (m *UserEnterRoomReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
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
			return fmt.Errorf("proto: UserEnterRoomReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserEnterRoomReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
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
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
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
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Session = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServer
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
func (m *UserEnterRoomResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
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
			return fmt.Errorf("proto: UserEnterRoomResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserEnterRoomResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrCode", wireType)
			}
			m.ErrCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrCode |= INNER_ERROR(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field History", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
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
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.History = append(m.History, &ChatMsg{})
			if err := m.History[len(m.History)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServer
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
func (m *EnterRoomNotifyGateway) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
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
			return fmt.Errorf("proto: EnterRoomNotifyGateway: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnterRoomNotifyGateway: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
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
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Session = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Room", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
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
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Room = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServer
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
func skipServer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServer
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
					return 0, ErrIntOverflowServer
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
					return 0, ErrIntOverflowServer
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
				return 0, ErrInvalidLengthServer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupServer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthServer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthServer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupServer = fmt.Errorf("proto: unexpected end of group")
)
