// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: get_usdt/get_usdt.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GetUsdt_GetRates_FullMethodName = "/get_usdt.GetUsdt/GetRates"
)

// GetUsdtClient is the client API for GetUsdt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetUsdtClient interface {
	GetRates(ctx context.Context, in *GetRatesRequest, opts ...grpc.CallOption) (*GetRatesResponse, error)
}

type getUsdtClient struct {
	cc grpc.ClientConnInterface
}

func NewGetUsdtClient(cc grpc.ClientConnInterface) GetUsdtClient {
	return &getUsdtClient{cc}
}

func (c *getUsdtClient) GetRates(ctx context.Context, in *GetRatesRequest, opts ...grpc.CallOption) (*GetRatesResponse, error) {
	out := new(GetRatesResponse)
	err := c.cc.Invoke(ctx, GetUsdt_GetRates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetUsdtServer is the server API for GetUsdt service.
// All implementations must embed UnimplementedGetUsdtServer
// for forward compatibility
type GetUsdtServer interface {
	GetRates(context.Context, *GetRatesRequest) (*GetRatesResponse, error)
	mustEmbedUnimplementedGetUsdtServer()
}

// UnimplementedGetUsdtServer must be embedded to have forward compatible implementations.
type UnimplementedGetUsdtServer struct {
}

func (UnimplementedGetUsdtServer) GetRates(context.Context, *GetRatesRequest) (*GetRatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRates not implemented")
}
func (UnimplementedGetUsdtServer) mustEmbedUnimplementedGetUsdtServer() {}

// UnsafeGetUsdtServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetUsdtServer will
// result in compilation errors.
type UnsafeGetUsdtServer interface {
	mustEmbedUnimplementedGetUsdtServer()
}

func RegisterGetUsdtServer(s grpc.ServiceRegistrar, srv GetUsdtServer) {
	s.RegisterService(&GetUsdt_ServiceDesc, srv)
}

func _GetUsdt_GetRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetUsdtServer).GetRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetUsdt_GetRates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetUsdtServer).GetRates(ctx, req.(*GetRatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetUsdt_ServiceDesc is the grpc.ServiceDesc for GetUsdt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetUsdt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "get_usdt.GetUsdt",
	HandlerType: (*GetUsdtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRates",
			Handler:    _GetUsdt_GetRates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "get_usdt/get_usdt.proto",
}
