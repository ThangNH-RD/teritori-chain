// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: teritori/airdrop/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgSetAllocation defines an sdk.Msg type that set airdrop allocation
type MsgSetAllocation struct {
	Sender     string            `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Allocation AirdropAllocation `protobuf:"bytes,2,opt,name=allocation,proto3" json:"allocation"`
}

func (m *MsgSetAllocation) Reset()         { *m = MsgSetAllocation{} }
func (m *MsgSetAllocation) String() string { return proto.CompactTextString(m) }
func (*MsgSetAllocation) ProtoMessage()    {}
func (*MsgSetAllocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbdab318d176f45, []int{0}
}
func (m *MsgSetAllocation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetAllocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetAllocation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetAllocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetAllocation.Merge(m, src)
}
func (m *MsgSetAllocation) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetAllocation) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetAllocation.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetAllocation proto.InternalMessageInfo

func (m *MsgSetAllocation) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgSetAllocation) GetAllocation() AirdropAllocation {
	if m != nil {
		return m.Allocation
	}
	return AirdropAllocation{}
}

// MsgSetAllocationResponse defines the Msg/SetAllocation response type.
type MsgSetAllocationResponse struct {
}

func (m *MsgSetAllocationResponse) Reset()         { *m = MsgSetAllocationResponse{} }
func (m *MsgSetAllocationResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetAllocationResponse) ProtoMessage()    {}
func (*MsgSetAllocationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbdab318d176f45, []int{1}
}
func (m *MsgSetAllocationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetAllocationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetAllocationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetAllocationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetAllocationResponse.Merge(m, src)
}
func (m *MsgSetAllocationResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetAllocationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetAllocationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetAllocationResponse proto.InternalMessageInfo

// MsgClaimAllocation defines an sdk.Msg type that claims airdrop allocation
type MsgClaimAllocation struct {
	Address       string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey        string `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	RewardAddress string `protobuf:"bytes,3,opt,name=reward_address,json=rewardAddress,proto3" json:"reward_address,omitempty"`
	Signature     string `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *MsgClaimAllocation) Reset()         { *m = MsgClaimAllocation{} }
func (m *MsgClaimAllocation) String() string { return proto.CompactTextString(m) }
func (*MsgClaimAllocation) ProtoMessage()    {}
func (*MsgClaimAllocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbdab318d176f45, []int{2}
}
func (m *MsgClaimAllocation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimAllocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimAllocation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimAllocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimAllocation.Merge(m, src)
}
func (m *MsgClaimAllocation) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimAllocation) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimAllocation.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimAllocation proto.InternalMessageInfo

// MsgClaimAllocationResponse defines the Msg/ClaimAllocation response type.
type MsgClaimAllocationResponse struct {
}

func (m *MsgClaimAllocationResponse) Reset()         { *m = MsgClaimAllocationResponse{} }
func (m *MsgClaimAllocationResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClaimAllocationResponse) ProtoMessage()    {}
func (*MsgClaimAllocationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbdab318d176f45, []int{3}
}
func (m *MsgClaimAllocationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimAllocationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimAllocationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimAllocationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimAllocationResponse.Merge(m, src)
}
func (m *MsgClaimAllocationResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimAllocationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimAllocationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimAllocationResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSetAllocation)(nil), "teritori.airdrop.v1beta1.MsgSetAllocation")
	proto.RegisterType((*MsgSetAllocationResponse)(nil), "teritori.airdrop.v1beta1.MsgSetAllocationResponse")
	proto.RegisterType((*MsgClaimAllocation)(nil), "teritori.airdrop.v1beta1.MsgClaimAllocation")
	proto.RegisterType((*MsgClaimAllocationResponse)(nil), "teritori.airdrop.v1beta1.MsgClaimAllocationResponse")
}

func init() { proto.RegisterFile("teritori/airdrop/v1beta1/tx.proto", fileDescriptor_2fbdab318d176f45) }

var fileDescriptor_2fbdab318d176f45 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3b, 0xaf, 0xd3, 0x30,
	0x1c, 0xc5, 0x63, 0xee, 0x55, 0x2f, 0x35, 0xba, 0x80, 0x2c, 0x04, 0x21, 0xba, 0x4a, 0x4b, 0x24,
	0xa4, 0xf2, 0x8a, 0xd5, 0xc0, 0xc4, 0xd6, 0x22, 0x86, 0x0a, 0x2a, 0x44, 0xe8, 0xc4, 0x52, 0x39,
	0x8d, 0x71, 0x2d, 0xd2, 0x38, 0xb2, 0x1d, 0x68, 0x06, 0x76, 0x46, 0x26, 0xe6, 0x7e, 0x9c, 0x8e,
	0x1d, 0x99, 0x10, 0x6a, 0x17, 0x16, 0xbe, 0x03, 0x22, 0x8f, 0x3e, 0x55, 0x04, 0x5b, 0x8e, 0xff,
	0xbf, 0xff, 0x39, 0x47, 0xb1, 0xe1, 0x1d, 0x4d, 0x25, 0xd7, 0x42, 0x72, 0x4c, 0xb8, 0x0c, 0xa5,
	0x48, 0xf0, 0x87, 0x76, 0x40, 0x35, 0x69, 0x63, 0x3d, 0x75, 0x13, 0x29, 0xb4, 0x40, 0x66, 0x85,
	0xb8, 0x25, 0xe2, 0x96, 0x88, 0x75, 0x83, 0x09, 0x26, 0x72, 0x08, 0xff, 0xf9, 0x2a, 0x78, 0xeb,
	0x36, 0x13, 0x82, 0x45, 0x14, 0xe7, 0x2a, 0x48, 0xdf, 0x61, 0x12, 0x67, 0xe5, 0xe8, 0xde, 0xd1,
	0x34, 0x12, 0x45, 0x62, 0x44, 0x34, 0x17, 0x71, 0x81, 0x3a, 0x9f, 0xe0, 0xf5, 0xbe, 0x62, 0x6f,
	0xa8, 0xee, 0xac, 0x27, 0xe8, 0x26, 0xac, 0x29, 0x1a, 0x87, 0x54, 0x9a, 0xa0, 0x09, 0x5a, 0x75,
	0xbf, 0x54, 0xe8, 0x35, 0x84, 0x9b, 0x7d, 0xf3, 0x52, 0x13, 0xb4, 0xae, 0x78, 0x0f, 0xdc, 0x63,
	0xb5, 0xdd, 0x4e, 0xa1, 0x37, 0xc6, 0xdd, 0xd3, 0xf9, 0xf7, 0x86, 0xe1, 0x6f, 0x99, 0x38, 0x16,
	0x34, 0xf7, 0xe3, 0x7d, 0xaa, 0x12, 0x11, 0x2b, 0xea, 0x7c, 0x05, 0x10, 0xf5, 0x15, 0x7b, 0x16,
	0x11, 0x3e, 0xd9, 0x6a, 0x67, 0xc2, 0x33, 0x12, 0x86, 0x92, 0x2a, 0x55, 0xd6, 0xab, 0x24, 0xba,
	0x05, 0xcf, 0x92, 0x34, 0x18, 0xbe, 0xa7, 0x59, 0x5e, 0xae, 0xee, 0xd7, 0x92, 0x34, 0x78, 0x41,
	0x33, 0x74, 0x17, 0x5e, 0x95, 0xf4, 0x23, 0x91, 0xe1, 0xb0, 0xda, 0x3c, 0xc9, 0xe7, 0xe7, 0xc5,
	0x69, 0xa7, 0xdc, 0xbf, 0x80, 0x75, 0xc5, 0x59, 0x4c, 0x74, 0x2a, 0xa9, 0x79, 0x9a, 0x13, 0x9b,
	0x83, 0xa7, 0x97, 0x3f, 0xcf, 0x1a, 0xc6, 0xcf, 0x59, 0xc3, 0x70, 0x2e, 0xa0, 0x75, 0xd8, 0xab,
	0xaa, 0xed, 0xfd, 0x02, 0xf0, 0xa4, 0xaf, 0x18, 0x4a, 0xe1, 0xb5, 0xfd, 0xea, 0x0f, 0x8f, 0xff,
	0xac, 0x43, 0x43, 0xeb, 0xc9, 0xff, 0xd0, 0x55, 0x3c, 0x12, 0xf0, 0x7c, 0xf7, 0x36, 0xef, 0xff,
	0xd5, 0x66, 0x87, 0xb5, 0xbc, 0x7f, 0x67, 0xab, 0xc0, 0xee, 0xcb, 0xf9, 0xd2, 0x06, 0x8b, 0xa5,
	0x0d, 0x7e, 0x2c, 0x6d, 0xf0, 0x65, 0x65, 0x1b, 0x8b, 0x95, 0x6d, 0x7c, 0x5b, 0xd9, 0xc6, 0x5b,
	0x8f, 0x71, 0x3d, 0x4e, 0x03, 0x77, 0x24, 0x26, 0x78, 0xf0, 0xdc, 0xef, 0x0d, 0x5e, 0xf9, 0x3d,
	0x5c, 0x05, 0x3c, 0x1a, 0x8d, 0x09, 0x8f, 0xf1, 0x74, 0xfd, 0x44, 0x75, 0x96, 0x50, 0x15, 0xd4,
	0xf2, 0x67, 0xf9, 0xf8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xd5, 0x0c, 0x66, 0x31, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// ClaimAllocation defines a method to claim allocation
	ClaimAllocation(ctx context.Context, in *MsgClaimAllocation, opts ...grpc.CallOption) (*MsgClaimAllocationResponse, error)
	// SetAllocation defines a method to set allocation
	SetAllocation(ctx context.Context, in *MsgSetAllocation, opts ...grpc.CallOption) (*MsgSetAllocationResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ClaimAllocation(ctx context.Context, in *MsgClaimAllocation, opts ...grpc.CallOption) (*MsgClaimAllocationResponse, error) {
	out := new(MsgClaimAllocationResponse)
	err := c.cc.Invoke(ctx, "/teritori.airdrop.v1beta1.Msg/ClaimAllocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetAllocation(ctx context.Context, in *MsgSetAllocation, opts ...grpc.CallOption) (*MsgSetAllocationResponse, error) {
	out := new(MsgSetAllocationResponse)
	err := c.cc.Invoke(ctx, "/teritori.airdrop.v1beta1.Msg/SetAllocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// ClaimAllocation defines a method to claim allocation
	ClaimAllocation(context.Context, *MsgClaimAllocation) (*MsgClaimAllocationResponse, error)
	// SetAllocation defines a method to set allocation
	SetAllocation(context.Context, *MsgSetAllocation) (*MsgSetAllocationResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ClaimAllocation(ctx context.Context, req *MsgClaimAllocation) (*MsgClaimAllocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimAllocation not implemented")
}
func (*UnimplementedMsgServer) SetAllocation(ctx context.Context, req *MsgSetAllocation) (*MsgSetAllocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAllocation not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ClaimAllocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimAllocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimAllocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teritori.airdrop.v1beta1.Msg/ClaimAllocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimAllocation(ctx, req.(*MsgClaimAllocation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetAllocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetAllocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetAllocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teritori.airdrop.v1beta1.Msg/SetAllocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetAllocation(ctx, req.(*MsgSetAllocation))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "teritori.airdrop.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClaimAllocation",
			Handler:    _Msg_ClaimAllocation_Handler,
		},
		{
			MethodName: "SetAllocation",
			Handler:    _Msg_SetAllocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teritori/airdrop/v1beta1/tx.proto",
}

func (m *MsgSetAllocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetAllocation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetAllocation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Allocation.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetAllocationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetAllocationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetAllocationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgClaimAllocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimAllocation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimAllocation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RewardAddress) > 0 {
		i -= len(m.RewardAddress)
		copy(dAtA[i:], m.RewardAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RewardAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimAllocationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimAllocationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimAllocationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSetAllocation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Allocation.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSetAllocationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgClaimAllocation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RewardAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgClaimAllocationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetAllocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetAllocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetAllocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Allocation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetAllocationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetAllocationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetAllocationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgClaimAllocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClaimAllocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimAllocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgClaimAllocationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClaimAllocationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimAllocationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
