// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/provider/types/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

func init() { proto.RegisterFile("x/provider/types/query.proto", fileDescriptor_68661d423e743651) }

var fileDescriptor_68661d423e743651 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4d, 0x8b, 0x1a, 0x31,
	0x18, 0xc7, 0x4d, 0xa1, 0x1e, 0x42, 0x5f, 0x73, 0x91, 0x0e, 0x76, 0x94, 0xd0, 0x16, 0x6b, 0xdb,
	0x09, 0xb4, 0xdf, 0xa0, 0xf5, 0x52, 0x28, 0x94, 0xd6, 0x9e, 0x7a, 0x29, 0x71, 0x4c, 0x35, 0xa0,
	0x93, 0x31, 0x4f, 0x2c, 0x95, 0xd2, 0xcb, 0x1e, 0xf6, 0xb0, 0x97, 0x5d, 0x5c, 0xf6, 0x3b, 0xed,
	0x51, 0xd8, 0xcb, 0x1e, 0x17, 0xdd, 0x2f, 0xb0, 0xdf, 0x60, 0x49, 0xc6, 0xf1, 0x65, 0x70, 0x5c,
	0xbd, 0x89, 0xfc, 0xf2, 0x3c, 0xbf, 0xff, 0x7f, 0x12, 0x5c, 0xfe, 0xcb, 0x62, 0xad, 0xfe, 0xc8,
	0xb6, 0xd0, 0xcc, 0x8c, 0x62, 0x01, 0x6c, 0x30, 0x14, 0x7a, 0x14, 0xc4, 0x5a, 0x19, 0x45, 0x4a,
	0x4a, 0x73, 0x19, 0x76, 0xb9, 0x8c, 0x02, 0xfb, 0x2b, 0x48, 0x51, 0xaf, 0xdc, 0x51, 0xaa, 0xd3,
	0x13, 0x8c, 0xc7, 0x92, 0xf1, 0x28, 0x52, 0x86, 0x1b, 0xa9, 0x22, 0x48, 0x8e, 0x79, 0x2f, 0x36,
	0x0f, 0xfd, 0xd5, 0x06, 0x35, 0xd4, 0xa1, 0xb8, 0x83, 0x52, 0x10, 0x6a, 0x19, 0x9b, 0x39, 0x45,
	0x73, 0x28, 0x13, 0x72, 0x98, 0x4f, 0x7a, 0x7f, 0x53, 0xc4, 0xf7, 0xbf, 0xd9, 0x7f, 0xc9, 0x18,
	0xe1, 0x47, 0x0d, 0x6e, 0x78, 0xd3, 0x2d, 0xfa, 0x1c, 0xfd, 0x56, 0xa4, 0x1e, 0xe4, 0x84, 0x08,
	0xd6, 0xc1, 0xef, 0x62, 0xe0, 0xed, 0xce, 0x02, 0xad, 0x1d, 0x5c, 0x5c, 0x9f, 0xde, 0xa3, 0xa4,
	0xca, 0x2c, 0xb9, 0x94, 0x6c, 0x73, 0xc3, 0x93, 0x9c, 0xec, 0x5f, 0xc4, 0xfb, 0xe2, 0x3f, 0x39,
	0x46, 0xf8, 0xf1, 0x17, 0x09, 0x66, 0x39, 0x03, 0xc8, 0x9b, 0xdc, 0x4d, 0x19, 0xd2, 0x6a, 0xed,
	0x01, 0x03, 0xa5, 0xce, 0xab, 0x4c, 0xbc, 0x5c, 0x2f, 0x20, 0x67, 0x08, 0x3f, 0xf9, 0xaa, 0x79,
	0xd8, 0x13, 0x4d, 0xd7, 0xb5, 0x2b, 0xea, 0x6d, 0xee, 0x96, 0x2c, 0x6a, 0x9d, 0xf6, 0xa1, 0x81,
	0xbe, 0x74, 0x52, 0x15, 0xf2, 0x3c, 0x23, 0x35, 0xff, 0xd6, 0x69, 0x53, 0x63, 0x84, 0x9f, 0xda,
	0x48, 0xab, 0x23, 0x80, 0xbc, 0xdb, 0x1a, 0x7f, 0x8d, 0xb5, 0x66, 0x7b, 0xe1, 0x40, 0x2b, 0x4e,
	0xed, 0x19, 0x29, 0x6d, 0x56, 0x03, 0x72, 0x84, 0xf0, 0x83, 0x1f, 0x02, 0xcc, 0x27, 0x0e, 0xc9,
	0x8d, 0xaa, 0xe5, 0x2e, 0x58, 0xc5, 0xac, 0xca, 0xae, 0x24, 0xd0, 0x57, 0xce, 0xa2, 0x4a, 0xfc,
	0x8c, 0x85, 0x11, 0xe0, 0x6e, 0x7a, 0xda, 0xd0, 0x21, 0xc2, 0x0f, 0x6d, 0x8c, 0xf4, 0x3c, 0x90,
	0xd7, 0x5b, 0xe3, 0x2e, 0x38, 0xab, 0xb3, 0x33, 0x9a, 0xdf, 0x4a, 0xea, 0xf3, 0xb1, 0x71, 0x3e,
	0xf5, 0xd1, 0x64, 0xea, 0xa3, 0xab, 0xa9, 0x8f, 0x4e, 0x66, 0x7e, 0x61, 0x32, 0xf3, 0x0b, 0x97,
	0x33, 0xbf, 0xf0, 0xb3, 0xde, 0x91, 0xa6, 0x3b, 0x6c, 0x05, 0xa1, 0xea, 0xb3, 0xc5, 0xbe, 0x64,
	0x4c, 0xf6, 0x2d, 0xb7, 0x8a, 0xee, 0x01, 0x7f, 0xb8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x32, 0x28,
	0xf5, 0xce, 0x87, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// DataSourceInfo gets the data source meta data
	DataSourceInfo(ctx context.Context, in *DataSourceInfoReq, opts ...grpc.CallOption) (*DataSourceInfoRes, error)
	// ListDataSources gets the list of data sources
	ListDataSources(ctx context.Context, in *ListDataSourcesReq, opts ...grpc.CallOption) (*ListDataSourcesRes, error)
	// OracleScriptInfo gets the oracle script meta data
	OracleScriptInfo(ctx context.Context, in *OracleScriptInfoReq, opts ...grpc.CallOption) (*OracleScriptInfoRes, error)
	// ListOracleScripts gets the list of oracle scripts
	ListOracleScripts(ctx context.Context, in *ListOracleScriptsReq, opts ...grpc.CallOption) (*ListOracleScriptsRes, error)
	// TestCaseInfo gets the test case meta data
	TestCaseInfo(ctx context.Context, in *TestCaseInfoReq, opts ...grpc.CallOption) (*TestCaseInfoRes, error)
	// ListTestCases gets the list of test cases
	ListTestCases(ctx context.Context, in *ListTestCasesReq, opts ...grpc.CallOption) (*ListTestCasesRes, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DataSourceInfo(ctx context.Context, in *DataSourceInfoReq, opts ...grpc.CallOption) (*DataSourceInfoRes, error) {
	out := new(DataSourceInfoRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.provider.Query/DataSourceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListDataSources(ctx context.Context, in *ListDataSourcesReq, opts ...grpc.CallOption) (*ListDataSourcesRes, error) {
	out := new(ListDataSourcesRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.provider.Query/ListDataSources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) OracleScriptInfo(ctx context.Context, in *OracleScriptInfoReq, opts ...grpc.CallOption) (*OracleScriptInfoRes, error) {
	out := new(OracleScriptInfoRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.provider.Query/OracleScriptInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListOracleScripts(ctx context.Context, in *ListOracleScriptsReq, opts ...grpc.CallOption) (*ListOracleScriptsRes, error) {
	out := new(ListOracleScriptsRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.provider.Query/ListOracleScripts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TestCaseInfo(ctx context.Context, in *TestCaseInfoReq, opts ...grpc.CallOption) (*TestCaseInfoRes, error) {
	out := new(TestCaseInfoRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.provider.Query/TestCaseInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListTestCases(ctx context.Context, in *ListTestCasesReq, opts ...grpc.CallOption) (*ListTestCasesRes, error) {
	out := new(ListTestCasesRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.provider.Query/ListTestCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// DataSourceInfo gets the data source meta data
	DataSourceInfo(context.Context, *DataSourceInfoReq) (*DataSourceInfoRes, error)
	// ListDataSources gets the list of data sources
	ListDataSources(context.Context, *ListDataSourcesReq) (*ListDataSourcesRes, error)
	// OracleScriptInfo gets the oracle script meta data
	OracleScriptInfo(context.Context, *OracleScriptInfoReq) (*OracleScriptInfoRes, error)
	// ListOracleScripts gets the list of oracle scripts
	ListOracleScripts(context.Context, *ListOracleScriptsReq) (*ListOracleScriptsRes, error)
	// TestCaseInfo gets the test case meta data
	TestCaseInfo(context.Context, *TestCaseInfoReq) (*TestCaseInfoRes, error)
	// ListTestCases gets the list of test cases
	ListTestCases(context.Context, *ListTestCasesReq) (*ListTestCasesRes, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DataSourceInfo(ctx context.Context, req *DataSourceInfoReq) (*DataSourceInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataSourceInfo not implemented")
}
func (*UnimplementedQueryServer) ListDataSources(ctx context.Context, req *ListDataSourcesReq) (*ListDataSourcesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDataSources not implemented")
}
func (*UnimplementedQueryServer) OracleScriptInfo(ctx context.Context, req *OracleScriptInfoReq) (*OracleScriptInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OracleScriptInfo not implemented")
}
func (*UnimplementedQueryServer) ListOracleScripts(ctx context.Context, req *ListOracleScriptsReq) (*ListOracleScriptsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOracleScripts not implemented")
}
func (*UnimplementedQueryServer) TestCaseInfo(ctx context.Context, req *TestCaseInfoReq) (*TestCaseInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestCaseInfo not implemented")
}
func (*UnimplementedQueryServer) ListTestCases(ctx context.Context, req *ListTestCasesReq) (*ListTestCasesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTestCases not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DataSourceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSourceInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DataSourceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.provider.Query/DataSourceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DataSourceInfo(ctx, req.(*DataSourceInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListDataSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataSourcesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListDataSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.provider.Query/ListDataSources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListDataSources(ctx, req.(*ListDataSourcesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_OracleScriptInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OracleScriptInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).OracleScriptInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.provider.Query/OracleScriptInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).OracleScriptInfo(ctx, req.(*OracleScriptInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListOracleScripts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOracleScriptsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListOracleScripts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.provider.Query/ListOracleScripts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListOracleScripts(ctx, req.(*ListOracleScriptsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TestCaseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestCaseInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TestCaseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.provider.Query/TestCaseInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TestCaseInfo(ctx, req.(*TestCaseInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListTestCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTestCasesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListTestCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.provider.Query/ListTestCases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListTestCases(ctx, req.(*ListTestCasesReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oraichain.orai.provider.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DataSourceInfo",
			Handler:    _Query_DataSourceInfo_Handler,
		},
		{
			MethodName: "ListDataSources",
			Handler:    _Query_ListDataSources_Handler,
		},
		{
			MethodName: "OracleScriptInfo",
			Handler:    _Query_OracleScriptInfo_Handler,
		},
		{
			MethodName: "ListOracleScripts",
			Handler:    _Query_ListOracleScripts_Handler,
		},
		{
			MethodName: "TestCaseInfo",
			Handler:    _Query_TestCaseInfo_Handler,
		},
		{
			MethodName: "ListTestCases",
			Handler:    _Query_ListTestCases_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/provider/types/query.proto",
}
