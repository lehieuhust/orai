// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/provider/types/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() { proto.RegisterFile("x/provider/types/tx.proto", fileDescriptor_853b5c7cbf7be1e9) }

var fileDescriptor_853b5c7cbf7be1e9 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xac, 0xd0, 0x2f, 0x28,
	0xca, 0x2f, 0xcb, 0x4c, 0x49, 0x2d, 0xd2, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0xa9, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcf, 0x2f, 0x4a, 0xcc, 0x4c, 0xce, 0x48, 0xcc, 0xcc,
	0xd3, 0x03, 0xb1, 0xf4, 0x60, 0xea, 0xa4, 0x14, 0xb1, 0xe8, 0x89, 0x4f, 0x29, 0xce, 0x2f, 0x2d,
	0x4a, 0x4e, 0x85, 0xe8, 0xc5, 0xae, 0x24, 0xbf, 0x38, 0xb9, 0x28, 0xb3, 0xa0, 0x04, 0xaa, 0x44,
	0x1e, 0x9b, 0x92, 0x92, 0xe4, 0xc4, 0x62, 0xa8, 0x19, 0x46, 0x1b, 0x59, 0xb9, 0x98, 0x7d, 0x8b,
	0xd3, 0x85, 0xca, 0xb9, 0x84, 0x9c, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x1d, 0x3d, 0x5d, 0x12, 0x4b,
	0x12, 0x83, 0xc1, 0xf6, 0x08, 0xe9, 0xe9, 0xe1, 0x70, 0x9e, 0x9e, 0x6f, 0x71, 0x3a, 0xa6, 0x7a,
	0x29, 0x43, 0xd2, 0xd4, 0x07, 0xa5, 0x16, 0x0b, 0x15, 0x72, 0x09, 0xb8, 0xa6, 0x64, 0x96, 0xa0,
	0x58, 0xab, 0x83, 0xcf, 0x18, 0x74, 0xd5, 0x52, 0xfa, 0xa4, 0xa8, 0x06, 0x59, 0x09, 0xf7, 0xab,
	0x7f, 0x51, 0x62, 0x72, 0x4e, 0x6a, 0x30, 0x38, 0xc0, 0x88, 0xf1, 0x2b, 0xb2, 0x7a, 0x62, 0xfc,
	0x8a, 0xac, 0x1e, 0xc9, 0xaf, 0x28, 0xd6, 0x12, 0xf4, 0x2b, 0x8a, 0xa5, 0xfa, 0xa4, 0xa8, 0x06,
	0x59, 0x99, 0xcd, 0xc5, 0x07, 0x71, 0x4b, 0x48, 0x6a, 0x71, 0x89, 0x73, 0x62, 0x71, 0xaa, 0x90,
	0x16, 0x61, 0x77, 0xc3, 0xd4, 0x4a, 0xe9, 0x12, 0xaf, 0x16, 0x64, 0x59, 0x2a, 0x17, 0x0f, 0xc8,
	0x0d, 0x70, 0xab, 0x34, 0x08, 0xb9, 0x16, 0x6e, 0x91, 0x36, 0xb1, 0x2a, 0x83, 0x52, 0x8b, 0x9d,
	0x5c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x2b, 0x3d, 0xb3, 0x24,
	0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x6e, 0x20, 0x98, 0xa5, 0x8f, 0x9e, 0x11, 0x92,
	0xd8, 0xc0, 0x19, 0xc0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x45, 0x77, 0x7c, 0x55, 0x9d, 0x03,
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
	// Create a new data source
	CreateAIDataSource(ctx context.Context, in *MsgCreateAIDataSource, opts ...grpc.CallOption) (*MsgCreateAIDataSourceRes, error)
	// Edit an existing data source
	EditAIDataSource(ctx context.Context, in *MsgEditAIDataSource, opts ...grpc.CallOption) (*MsgEditAIDataSourceRes, error)
	// Create a new oracle script
	CreateOracleScript(ctx context.Context, in *MsgCreateOracleScript, opts ...grpc.CallOption) (*MsgCreateOracleScriptRes, error)
	// Edit an existing oracle script
	EditOracleScript(ctx context.Context, in *MsgEditOracleScript, opts ...grpc.CallOption) (*MsgEditOracleScriptRes, error)
	// Create a new test case
	CreateTestCase(ctx context.Context, in *MsgCreateTestCase, opts ...grpc.CallOption) (*MsgCreateTestCaseRes, error)
	// Edit an existing test case
	EditTestCase(ctx context.Context, in *MsgEditTestCase, opts ...grpc.CallOption) (*MsgEditTestCaseRes, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateAIDataSource(ctx context.Context, in *MsgCreateAIDataSource, opts ...grpc.CallOption) (*MsgCreateAIDataSourceRes, error) {
	out := new(MsgCreateAIDataSourceRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.provider.Msg/CreateAIDataSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EditAIDataSource(ctx context.Context, in *MsgEditAIDataSource, opts ...grpc.CallOption) (*MsgEditAIDataSourceRes, error) {
	out := new(MsgEditAIDataSourceRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.provider.Msg/EditAIDataSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateOracleScript(ctx context.Context, in *MsgCreateOracleScript, opts ...grpc.CallOption) (*MsgCreateOracleScriptRes, error) {
	out := new(MsgCreateOracleScriptRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.provider.Msg/CreateOracleScript", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EditOracleScript(ctx context.Context, in *MsgEditOracleScript, opts ...grpc.CallOption) (*MsgEditOracleScriptRes, error) {
	out := new(MsgEditOracleScriptRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.provider.Msg/EditOracleScript", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateTestCase(ctx context.Context, in *MsgCreateTestCase, opts ...grpc.CallOption) (*MsgCreateTestCaseRes, error) {
	out := new(MsgCreateTestCaseRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.provider.Msg/CreateTestCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EditTestCase(ctx context.Context, in *MsgEditTestCase, opts ...grpc.CallOption) (*MsgEditTestCaseRes, error) {
	out := new(MsgEditTestCaseRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.provider.Msg/EditTestCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Create a new data source
	CreateAIDataSource(context.Context, *MsgCreateAIDataSource) (*MsgCreateAIDataSourceRes, error)
	// Edit an existing data source
	EditAIDataSource(context.Context, *MsgEditAIDataSource) (*MsgEditAIDataSourceRes, error)
	// Create a new oracle script
	CreateOracleScript(context.Context, *MsgCreateOracleScript) (*MsgCreateOracleScriptRes, error)
	// Edit an existing oracle script
	EditOracleScript(context.Context, *MsgEditOracleScript) (*MsgEditOracleScriptRes, error)
	// Create a new test case
	CreateTestCase(context.Context, *MsgCreateTestCase) (*MsgCreateTestCaseRes, error)
	// Edit an existing test case
	EditTestCase(context.Context, *MsgEditTestCase) (*MsgEditTestCaseRes, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateAIDataSource(ctx context.Context, req *MsgCreateAIDataSource) (*MsgCreateAIDataSourceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAIDataSource not implemented")
}
func (*UnimplementedMsgServer) EditAIDataSource(ctx context.Context, req *MsgEditAIDataSource) (*MsgEditAIDataSourceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAIDataSource not implemented")
}
func (*UnimplementedMsgServer) CreateOracleScript(ctx context.Context, req *MsgCreateOracleScript) (*MsgCreateOracleScriptRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOracleScript not implemented")
}
func (*UnimplementedMsgServer) EditOracleScript(ctx context.Context, req *MsgEditOracleScript) (*MsgEditOracleScriptRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditOracleScript not implemented")
}
func (*UnimplementedMsgServer) CreateTestCase(ctx context.Context, req *MsgCreateTestCase) (*MsgCreateTestCaseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTestCase not implemented")
}
func (*UnimplementedMsgServer) EditTestCase(ctx context.Context, req *MsgEditTestCase) (*MsgEditTestCaseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditTestCase not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateAIDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateAIDataSource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateAIDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.provider.Msg/CreateAIDataSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateAIDataSource(ctx, req.(*MsgCreateAIDataSource))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EditAIDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEditAIDataSource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EditAIDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.provider.Msg/EditAIDataSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EditAIDataSource(ctx, req.(*MsgEditAIDataSource))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateOracleScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateOracleScript)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateOracleScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.provider.Msg/CreateOracleScript",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateOracleScript(ctx, req.(*MsgCreateOracleScript))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EditOracleScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEditOracleScript)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EditOracleScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.provider.Msg/EditOracleScript",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EditOracleScript(ctx, req.(*MsgEditOracleScript))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateTestCase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.provider.Msg/CreateTestCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateTestCase(ctx, req.(*MsgCreateTestCase))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EditTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEditTestCase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EditTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.provider.Msg/EditTestCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EditTestCase(ctx, req.(*MsgEditTestCase))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oraichain.orai.provider.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAIDataSource",
			Handler:    _Msg_CreateAIDataSource_Handler,
		},
		{
			MethodName: "EditAIDataSource",
			Handler:    _Msg_EditAIDataSource_Handler,
		},
		{
			MethodName: "CreateOracleScript",
			Handler:    _Msg_CreateOracleScript_Handler,
		},
		{
			MethodName: "EditOracleScript",
			Handler:    _Msg_EditOracleScript_Handler,
		},
		{
			MethodName: "CreateTestCase",
			Handler:    _Msg_CreateTestCase_Handler,
		},
		{
			MethodName: "EditTestCase",
			Handler:    _Msg_EditTestCase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/provider/types/tx.proto",
}
