// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/provider/types/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xc7, 0x5b, 0x44, 0x0f, 0x41, 0x44, 0x82, 0x20, 0xf6, 0x10, 0xf1, 0x26, 0x53, 0x13, 0xd4,
	0x27, 0xd0, 0xcd, 0x83, 0x87, 0x21, 0x6c, 0x9e, 0xbc, 0x48, 0x96, 0xfe, 0xe8, 0x82, 0x7f, 0x52,
	0x93, 0x4c, 0xe7, 0x5b, 0xf8, 0x2a, 0xbe, 0x85, 0xc7, 0x1d, 0x3d, 0x4a, 0xfb, 0x22, 0x92, 0x6e,
	0x2b, 0xed, 0x9c, 0x6b, 0x7b, 0xfb, 0x1d, 0x3e, 0x7c, 0x3f, 0xdf, 0x6f, 0x20, 0x68, 0x6f, 0xcc,
	0x62, 0xad, 0x5e, 0x65, 0x08, 0x9a, 0xd9, 0xf7, 0x18, 0x0c, 0xb3, 0x63, 0x1a, 0x6b, 0x65, 0x15,
	0xde, 0x55, 0x9a, 0x4b, 0x31, 0xe4, 0xf2, 0x99, 0xba, 0x8b, 0xce, 0xb9, 0x60, 0x27, 0x52, 0x91,
	0xca, 0x18, 0xe6, 0xae, 0x29, 0x1e, 0x1c, 0x2c, 0x49, 0xba, 0x0f, 0x8d, 0x1a, 0x69, 0x01, 0xab,
	0x10, 0x65, 0x84, 0x96, 0xb1, 0x9d, 0x21, 0xfb, 0xcb, 0x10, 0x2b, 0xb8, 0x99, 0x65, 0x9c, 0x7d,
	0xae, 0xa3, 0xb5, 0xae, 0x89, 0xf0, 0x1b, 0xc2, 0x6d, 0x0d, 0xdc, 0xc2, 0xc5, 0x75, 0x87, 0x5b,
	0xde, 0xcf, 0x3c, 0x98, 0xd2, 0x7f, 0x4a, 0xd3, 0xae, 0x89, 0xfe, 0xf2, 0xc1, 0x69, 0x33, 0xbe,
	0x07, 0x06, 0xbf, 0xa0, 0xed, 0xab, 0x50, 0xda, 0x92, 0xf6, 0x78, 0x55, 0xcc, 0x22, 0x1d, 0xb0,
	0x26, 0xb4, 0x53, 0xe6, 0x5b, 0x6f, 0x34, 0x17, 0x8f, 0xd0, 0xcf, 0x1e, 0xac, 0xce, 0xd6, 0x22,
	0x5f, 0x67, 0x6b, 0x91, 0x2f, 0x6c, 0x2d, 0x69, 0x2b, 0xb7, 0x96, 0xa4, 0xac, 0x09, 0xed, 0x94,
	0x0f, 0x68, 0x6b, 0xda, 0xe5, 0x16, 0x8c, 0x6d, 0x73, 0x03, 0xb8, 0x55, 0xdd, 0x7b, 0xce, 0x06,
	0x27, 0xf5, 0x59, 0x27, 0x03, 0xb4, 0xe9, 0x3a, 0xe4, 0xaa, 0xc3, 0xaa, 0xb6, 0xb9, 0xe8, 0xa8,
	0x2e, 0xd9, 0x03, 0x73, 0xd9, 0xf9, 0x4a, 0x88, 0x3f, 0x49, 0x88, 0xff, 0x93, 0x10, 0xff, 0x23,
	0x25, 0xde, 0x24, 0x25, 0xde, 0x77, 0x4a, 0xbc, 0xbb, 0x56, 0x24, 0xed, 0x70, 0x34, 0xa0, 0x42,
	0x3d, 0xb1, 0x3c, 0x30, 0xbb, 0xd8, 0xe2, 0x47, 0x18, 0x6c, 0x64, 0x1f, 0xe0, 0xfc, 0x37, 0x00,
	0x00, 0xff, 0xff, 0xaf, 0xe2, 0x2d, 0xe7, 0xb3, 0x03, 0x00, 0x00,
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
