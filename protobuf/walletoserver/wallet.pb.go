// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wallet.proto

/*
	Package walletoserver is a generated protocol buffer package.

	It is generated from these files:
		wallet.proto

	It has these top-level messages:
		Fee
		GetFeeRequest
		GetFreeReply
		TransferRequest
		TransferReply
*/
package walletoserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 费用
type Fee struct {
	AssetId string `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Fee     uint32 `protobuf:"varint,2,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *Fee) Reset()                    { *m = Fee{} }
func (m *Fee) String() string            { return proto.CompactTextString(m) }
func (*Fee) ProtoMessage()               {}
func (*Fee) Descriptor() ([]byte, []int) { return fileDescriptorWallet, []int{0} }

func (m *Fee) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *Fee) GetFee() uint32 {
	if m != nil {
		return m.Fee
	}
	return 0
}

// 获取费用请求
type GetFeeRequest struct {
	Assets []string `protobuf:"bytes,1,rep,name=assets" json:"assets,omitempty"`
}

func (m *GetFeeRequest) Reset()                    { *m = GetFeeRequest{} }
func (m *GetFeeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFeeRequest) ProtoMessage()               {}
func (*GetFeeRequest) Descriptor() ([]byte, []int) { return fileDescriptorWallet, []int{1} }

func (m *GetFeeRequest) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

// 获取费用响应
type GetFreeReply struct {
	Ok     bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Fees   []*Fee `protobuf:"bytes,3,rep,name=fees" json:"fees,omitempty"`
}

func (m *GetFreeReply) Reset()                    { *m = GetFreeReply{} }
func (m *GetFreeReply) String() string            { return proto.CompactTextString(m) }
func (*GetFreeReply) ProtoMessage()               {}
func (*GetFreeReply) Descriptor() ([]byte, []int) { return fileDescriptorWallet, []int{2} }

func (m *GetFreeReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *GetFreeReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *GetFreeReply) GetFees() []*Fee {
	if m != nil {
		return m.Fees
	}
	return nil
}

// 转账操作请求
type TransferRequest struct {
	To      string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	AssetId string `protobuf:"bytes,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Amount  uint32 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Memo    string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (m *TransferRequest) Reset()                    { *m = TransferRequest{} }
func (m *TransferRequest) String() string            { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()               {}
func (*TransferRequest) Descriptor() ([]byte, []int) { return fileDescriptorWallet, []int{3} }

func (m *TransferRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TransferRequest) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *TransferRequest) GetAmount() uint32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransferRequest) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

// 转账操作回复
type TransferReply struct {
	Ok     bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (m *TransferReply) Reset()                    { *m = TransferReply{} }
func (m *TransferReply) String() string            { return proto.CompactTextString(m) }
func (*TransferReply) ProtoMessage()               {}
func (*TransferReply) Descriptor() ([]byte, []int) { return fileDescriptorWallet, []int{4} }

func (m *TransferReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *TransferReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*Fee)(nil), "walletoserver.Fee")
	proto.RegisterType((*GetFeeRequest)(nil), "walletoserver.GetFeeRequest")
	proto.RegisterType((*GetFreeReply)(nil), "walletoserver.GetFreeReply")
	proto.RegisterType((*TransferRequest)(nil), "walletoserver.TransferRequest")
	proto.RegisterType((*TransferReply)(nil), "walletoserver.TransferReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Wallet service

type WalletClient interface {
	// GetFees 获取费用
	GetFees(ctx context.Context, in *GetFeeRequest, opts ...grpc.CallOption) (*GetFreeReply, error)
	// Transfer 转账操作
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferReply, error)
}

type walletClient struct {
	cc *grpc.ClientConn
}

func NewWalletClient(cc *grpc.ClientConn) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) GetFees(ctx context.Context, in *GetFeeRequest, opts ...grpc.CallOption) (*GetFreeReply, error) {
	out := new(GetFreeReply)
	err := grpc.Invoke(ctx, "/walletoserver.Wallet/GetFees", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferReply, error) {
	out := new(TransferReply)
	err := grpc.Invoke(ctx, "/walletoserver.Wallet/Transfer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Wallet service

type WalletServer interface {
	// GetFees 获取费用
	GetFees(context.Context, *GetFeeRequest) (*GetFreeReply, error)
	// Transfer 转账操作
	Transfer(context.Context, *TransferRequest) (*TransferReply, error)
}

func RegisterWalletServer(s *grpc.Server, srv WalletServer) {
	s.RegisterService(&_Wallet_serviceDesc, srv)
}

func _Wallet_GetFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletoserver.Wallet/GetFees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetFees(ctx, req.(*GetFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletoserver.Wallet/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Wallet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "walletoserver.Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFees",
			Handler:    _Wallet_GetFees_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _Wallet_Transfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet.proto",
}

func (m *Fee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fee) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AssetId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.AssetId)))
		i += copy(dAtA[i:], m.AssetId)
	}
	if m.Fee != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintWallet(dAtA, i, uint64(m.Fee))
	}
	return i, nil
}

func (m *GetFeeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetFeeRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for _, s := range m.Assets {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *GetFreeReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetFreeReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Ok {
		dAtA[i] = 0x8
		i++
		if m.Ok {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Reason) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.Reason)))
		i += copy(dAtA[i:], m.Reason)
	}
	if len(m.Fees) > 0 {
		for _, msg := range m.Fees {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintWallet(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *TransferRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransferRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.To) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.To)))
		i += copy(dAtA[i:], m.To)
	}
	if len(m.AssetId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.AssetId)))
		i += copy(dAtA[i:], m.AssetId)
	}
	if m.Amount != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintWallet(dAtA, i, uint64(m.Amount))
	}
	if len(m.Memo) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.Memo)))
		i += copy(dAtA[i:], m.Memo)
	}
	return i, nil
}

func (m *TransferReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransferReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Ok {
		dAtA[i] = 0x8
		i++
		if m.Ok {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Reason) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.Reason)))
		i += copy(dAtA[i:], m.Reason)
	}
	return i, nil
}

func encodeVarintWallet(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Fee) Size() (n int) {
	var l int
	_ = l
	l = len(m.AssetId)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	if m.Fee != 0 {
		n += 1 + sovWallet(uint64(m.Fee))
	}
	return n
}

func (m *GetFeeRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for _, s := range m.Assets {
			l = len(s)
			n += 1 + l + sovWallet(uint64(l))
		}
	}
	return n
}

func (m *GetFreeReply) Size() (n int) {
	var l int
	_ = l
	if m.Ok {
		n += 2
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	if len(m.Fees) > 0 {
		for _, e := range m.Fees {
			l = e.Size()
			n += 1 + l + sovWallet(uint64(l))
		}
	}
	return n
}

func (m *TransferRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	l = len(m.AssetId)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovWallet(uint64(m.Amount))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	return n
}

func (m *TransferReply) Size() (n int) {
	var l int
	_ = l
	if m.Ok {
		n += 2
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	return n
}

func sovWallet(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWallet(x uint64) (n int) {
	return sovWallet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Fee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWallet
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
			return fmt.Errorf("proto: Fee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			m.Fee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fee |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWallet
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
func (m *GetFeeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWallet
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
			return fmt.Errorf("proto: GetFeeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetFeeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Assets = append(m.Assets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWallet
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
func (m *GetFreeReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWallet
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
			return fmt.Errorf("proto: GetFreeReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetFreeReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ok", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ok = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fees = append(m.Fees, &Fee{})
			if err := m.Fees[len(m.Fees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWallet
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
func (m *TransferRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWallet
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
			return fmt.Errorf("proto: TransferRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransferRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWallet
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
func (m *TransferReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWallet
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
			return fmt.Errorf("proto: TransferReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransferReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ok", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ok = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWallet
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
func skipWallet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWallet
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
					return 0, ErrIntOverflowWallet
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
					return 0, ErrIntOverflowWallet
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
				return 0, ErrInvalidLengthWallet
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWallet
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
				next, err := skipWallet(dAtA[start:])
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
	ErrInvalidLengthWallet = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWallet   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("wallet.proto", fileDescriptorWallet) }

var fileDescriptorWallet = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xbd, 0x4e, 0x33, 0x31,
	0x10, 0x8c, 0xcf, 0x51, 0x7e, 0xf6, 0xcb, 0xe5, 0x8b, 0xb6, 0x40, 0x26, 0x44, 0xa7, 0x93, 0x0b,
	0xb8, 0x2a, 0x45, 0x28, 0xe8, 0x29, 0x82, 0xa0, 0xb4, 0x90, 0xe8, 0x40, 0x87, 0xb2, 0x11, 0x28,
	0x3f, 0x0e, 0xb6, 0x03, 0xe2, 0x4d, 0x28, 0x78, 0x20, 0x4a, 0x1e, 0x01, 0x85, 0x17, 0x41, 0x76,
	0x2e, 0x42, 0x77, 0x82, 0x82, 0x6e, 0x47, 0x3b, 0x9e, 0x9d, 0x19, 0x19, 0x3a, 0x4f, 0xf9, 0x7c,
	0x4e, 0x6e, 0xb8, 0x32, 0xda, 0x69, 0x8c, 0xb7, 0x48, 0x5b, 0x32, 0x8f, 0x64, 0xe4, 0x08, 0xf8,
	0x98, 0x08, 0xf7, 0xa1, 0x95, 0x5b, 0x4b, 0xee, 0xe6, 0x7e, 0x22, 0x58, 0xca, 0xb2, 0xb6, 0x6a,
	0x06, 0x7c, 0x3e, 0xc1, 0x1e, 0xf0, 0x29, 0x91, 0x88, 0x52, 0x96, 0xc5, 0xca, 0x8f, 0xf2, 0x08,
	0xe2, 0x33, 0x72, 0x63, 0x22, 0x45, 0x0f, 0x6b, 0xb2, 0x0e, 0xf7, 0xa0, 0x11, 0xd8, 0x56, 0xb0,
	0x94, 0x67, 0x6d, 0x55, 0x20, 0x79, 0x0d, 0x1d, 0x4f, 0x34, 0x9e, 0xb9, 0x9a, 0x3f, 0x63, 0x17,
	0x22, 0x3d, 0x0b, 0xfa, 0x2d, 0x15, 0xe9, 0x99, 0x7f, 0x67, 0x28, 0xb7, 0x7a, 0x19, 0xd4, 0xdb,
	0xaa, 0x40, 0x78, 0x08, 0xf5, 0x29, 0x91, 0x15, 0x3c, 0xe5, 0xd9, 0xbf, 0x11, 0x0e, 0x4b, 0x96,
	0x87, 0xfe, 0x70, 0xd8, 0xcb, 0x3b, 0xf8, 0x7f, 0x69, 0xf2, 0xa5, 0x9d, 0x92, 0xd9, 0x59, 0xe9,
	0x42, 0xe4, 0x74, 0x11, 0x21, 0x72, 0xba, 0x14, 0x2c, 0x2a, 0x07, 0xf3, 0xae, 0x17, 0x7a, 0xbd,
	0x74, 0x82, 0x87, 0x6c, 0x05, 0x42, 0x84, 0xfa, 0x82, 0x16, 0x5a, 0xd4, 0x03, 0x3d, 0xcc, 0xf2,
	0x04, 0xe2, 0xef, 0x4b, 0x7f, 0x88, 0x32, 0x7a, 0x65, 0xd0, 0xb8, 0x0a, 0xf6, 0x71, 0x0c, 0xcd,
	0x6d, 0x6d, 0x16, 0x07, 0x95, 0x48, 0xa5, 0x3a, 0xfb, 0x07, 0x3f, 0x6c, 0x77, 0x1d, 0xca, 0x1a,
	0x5e, 0x40, 0x6b, 0xe7, 0x05, 0x93, 0x0a, 0xb5, 0x52, 0x47, 0x7f, 0xf0, 0xeb, 0x3e, 0x68, 0x9d,
	0xf6, 0xde, 0x36, 0x09, 0x7b, 0xdf, 0x24, 0xec, 0x63, 0x93, 0xb0, 0x97, 0xcf, 0xa4, 0x76, 0xdb,
	0x08, 0xdf, 0xe4, 0xf8, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x81, 0x00, 0xf8, 0x13, 0x36, 0x02, 0x00,
	0x00,
}