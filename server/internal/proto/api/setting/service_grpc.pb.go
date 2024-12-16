// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/setting/service.proto

package setting

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

// SystemServiceClient is the client API for SystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemServiceClient interface {
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	ResetSetting(ctx context.Context, in *ResetSettingRequest, opts ...grpc.CallOption) (*ResetSettingResponse, error)
	UpdateSetting(ctx context.Context, in *UpdateSettingRequest, opts ...grpc.CallOption) (*UpdateSettingResponse, error)
	GetSetting(ctx context.Context, in *GetSettingRequest, opts ...grpc.CallOption) (*GetSettingResponse, error)
}

type systemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemServiceClient(cc grpc.ClientConnInterface) SystemServiceClient {
	return &systemServiceClient{cc}
}

func (c *systemServiceClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/api.setting.SystemService/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) ResetSetting(ctx context.Context, in *ResetSettingRequest, opts ...grpc.CallOption) (*ResetSettingResponse, error) {
	out := new(ResetSettingResponse)
	err := c.cc.Invoke(ctx, "/api.setting.SystemService/ResetSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UpdateSetting(ctx context.Context, in *UpdateSettingRequest, opts ...grpc.CallOption) (*UpdateSettingResponse, error) {
	out := new(UpdateSettingResponse)
	err := c.cc.Invoke(ctx, "/api.setting.SystemService/UpdateSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) GetSetting(ctx context.Context, in *GetSettingRequest, opts ...grpc.CallOption) (*GetSettingResponse, error) {
	out := new(GetSettingResponse)
	err := c.cc.Invoke(ctx, "/api.setting.SystemService/GetSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemServiceServer is the server API for SystemService service.
// All implementations must embed UnimplementedSystemServiceServer
// for forward compatibility
type SystemServiceServer interface {
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	ResetSetting(context.Context, *ResetSettingRequest) (*ResetSettingResponse, error)
	UpdateSetting(context.Context, *UpdateSettingRequest) (*UpdateSettingResponse, error)
	GetSetting(context.Context, *GetSettingRequest) (*GetSettingResponse, error)
	mustEmbedUnimplementedSystemServiceServer()
}

// UnimplementedSystemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemServiceServer struct {
}

func (UnimplementedSystemServiceServer) Info(context.Context, *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedSystemServiceServer) ResetSetting(context.Context, *ResetSettingRequest) (*ResetSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetSetting not implemented")
}
func (UnimplementedSystemServiceServer) UpdateSetting(context.Context, *UpdateSettingRequest) (*UpdateSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSetting not implemented")
}
func (UnimplementedSystemServiceServer) GetSetting(context.Context, *GetSettingRequest) (*GetSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSetting not implemented")
}
func (UnimplementedSystemServiceServer) mustEmbedUnimplementedSystemServiceServer() {}

// UnsafeSystemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemServiceServer will
// result in compilation errors.
type UnsafeSystemServiceServer interface {
	mustEmbedUnimplementedSystemServiceServer()
}

func RegisterSystemServiceServer(s grpc.ServiceRegistrar, srv SystemServiceServer) {
	s.RegisterService(&SystemService_ServiceDesc, srv)
}

func _SystemService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.setting.SystemService/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_ResetSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).ResetSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.setting.SystemService/ResetSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).ResetSetting(ctx, req.(*ResetSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UpdateSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UpdateSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.setting.SystemService/UpdateSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UpdateSetting(ctx, req.(*UpdateSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_GetSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).GetSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.setting.SystemService/GetSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).GetSetting(ctx, req.(*GetSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemService_ServiceDesc is the grpc.ServiceDesc for SystemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.setting.SystemService",
	HandlerType: (*SystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _SystemService_Info_Handler,
		},
		{
			MethodName: "ResetSetting",
			Handler:    _SystemService_ResetSetting_Handler,
		},
		{
			MethodName: "UpdateSetting",
			Handler:    _SystemService_UpdateSetting_Handler,
		},
		{
			MethodName: "GetSetting",
			Handler:    _SystemService_GetSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/setting/service.proto",
}
