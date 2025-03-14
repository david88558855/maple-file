// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/file/service.proto

package file

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

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	List(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error)
	Move(ctx context.Context, in *MoveFileRequest, opts ...grpc.CallOption) (*MoveFileResponse, error)
	Copy(ctx context.Context, in *CopyFileRequest, opts ...grpc.CallOption) (*CopyFileResponse, error)
	Mkdir(ctx context.Context, in *MkdirFileRequest, opts ...grpc.CallOption) (*MkdirFileResponse, error)
	Rename(ctx context.Context, in *RenameFileRequest, opts ...grpc.CallOption) (*RenameFileResponse, error)
	Remove(ctx context.Context, in *RemoveFileRequest, opts ...grpc.CallOption) (*RemoveFileResponse, error)
	Upload(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadClient, error)
	Download(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (FileService_DownloadClient, error)
	Preview(ctx context.Context, in *PreviewFileRequest, opts ...grpc.CallOption) (FileService_PreviewClient, error)
	ListRepos(ctx context.Context, in *ListReposRequest, opts ...grpc.CallOption) (*ListReposResponse, error)
	CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*CreateRepoResponse, error)
	UpdateRepo(ctx context.Context, in *UpdateRepoRequest, opts ...grpc.CallOption) (*UpdateRepoResponse, error)
	DeleteRepo(ctx context.Context, in *DeleteRepoRequest, opts ...grpc.CallOption) (*DeleteRepoResponse, error)
	TestRepo(ctx context.Context, in *TestRepoRequest, opts ...grpc.CallOption) (*TestRepoResponse, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) List(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error) {
	out := new(ListFilesResponse)
	err := c.cc.Invoke(ctx, "/api.file.FileService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) Move(ctx context.Context, in *MoveFileRequest, opts ...grpc.CallOption) (*MoveFileResponse, error) {
	out := new(MoveFileResponse)
	err := c.cc.Invoke(ctx, "/api.file.FileService/Move", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) Copy(ctx context.Context, in *CopyFileRequest, opts ...grpc.CallOption) (*CopyFileResponse, error) {
	out := new(CopyFileResponse)
	err := c.cc.Invoke(ctx, "/api.file.FileService/Copy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) Mkdir(ctx context.Context, in *MkdirFileRequest, opts ...grpc.CallOption) (*MkdirFileResponse, error) {
	out := new(MkdirFileResponse)
	err := c.cc.Invoke(ctx, "/api.file.FileService/Mkdir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) Rename(ctx context.Context, in *RenameFileRequest, opts ...grpc.CallOption) (*RenameFileResponse, error) {
	out := new(RenameFileResponse)
	err := c.cc.Invoke(ctx, "/api.file.FileService/Rename", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) Remove(ctx context.Context, in *RemoveFileRequest, opts ...grpc.CallOption) (*RemoveFileResponse, error) {
	out := new(RemoveFileResponse)
	err := c.cc.Invoke(ctx, "/api.file.FileService/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[0], "/api.file.FileService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceUploadClient{stream}
	return x, nil
}

type FileService_UploadClient interface {
	Send(*FileRequest) error
	CloseAndRecv() (*FileResponse, error)
	grpc.ClientStream
}

type fileServiceUploadClient struct {
	grpc.ClientStream
}

func (x *fileServiceUploadClient) Send(m *FileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceUploadClient) CloseAndRecv() (*FileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceClient) Download(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (FileService_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[1], "/api.file.FileService/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileService_DownloadClient interface {
	Recv() (*DownloadFileResponse, error)
	grpc.ClientStream
}

type fileServiceDownloadClient struct {
	grpc.ClientStream
}

func (x *fileServiceDownloadClient) Recv() (*DownloadFileResponse, error) {
	m := new(DownloadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceClient) Preview(ctx context.Context, in *PreviewFileRequest, opts ...grpc.CallOption) (FileService_PreviewClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[2], "/api.file.FileService/Preview", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServicePreviewClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileService_PreviewClient interface {
	Recv() (*PreviewFileResponse, error)
	grpc.ClientStream
}

type fileServicePreviewClient struct {
	grpc.ClientStream
}

func (x *fileServicePreviewClient) Recv() (*PreviewFileResponse, error) {
	m := new(PreviewFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceClient) ListRepos(ctx context.Context, in *ListReposRequest, opts ...grpc.CallOption) (*ListReposResponse, error) {
	out := new(ListReposResponse)
	err := c.cc.Invoke(ctx, "/api.file.FileService/ListRepos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*CreateRepoResponse, error) {
	out := new(CreateRepoResponse)
	err := c.cc.Invoke(ctx, "/api.file.FileService/CreateRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) UpdateRepo(ctx context.Context, in *UpdateRepoRequest, opts ...grpc.CallOption) (*UpdateRepoResponse, error) {
	out := new(UpdateRepoResponse)
	err := c.cc.Invoke(ctx, "/api.file.FileService/UpdateRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) DeleteRepo(ctx context.Context, in *DeleteRepoRequest, opts ...grpc.CallOption) (*DeleteRepoResponse, error) {
	out := new(DeleteRepoResponse)
	err := c.cc.Invoke(ctx, "/api.file.FileService/DeleteRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) TestRepo(ctx context.Context, in *TestRepoRequest, opts ...grpc.CallOption) (*TestRepoResponse, error) {
	out := new(TestRepoResponse)
	err := c.cc.Invoke(ctx, "/api.file.FileService/TestRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	List(context.Context, *ListFilesRequest) (*ListFilesResponse, error)
	Move(context.Context, *MoveFileRequest) (*MoveFileResponse, error)
	Copy(context.Context, *CopyFileRequest) (*CopyFileResponse, error)
	Mkdir(context.Context, *MkdirFileRequest) (*MkdirFileResponse, error)
	Rename(context.Context, *RenameFileRequest) (*RenameFileResponse, error)
	Remove(context.Context, *RemoveFileRequest) (*RemoveFileResponse, error)
	Upload(FileService_UploadServer) error
	Download(*DownloadFileRequest, FileService_DownloadServer) error
	Preview(*PreviewFileRequest, FileService_PreviewServer) error
	ListRepos(context.Context, *ListReposRequest) (*ListReposResponse, error)
	CreateRepo(context.Context, *CreateRepoRequest) (*CreateRepoResponse, error)
	UpdateRepo(context.Context, *UpdateRepoRequest) (*UpdateRepoResponse, error)
	DeleteRepo(context.Context, *DeleteRepoRequest) (*DeleteRepoResponse, error)
	TestRepo(context.Context, *TestRepoRequest) (*TestRepoResponse, error)
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) List(context.Context, *ListFilesRequest) (*ListFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedFileServiceServer) Move(context.Context, *MoveFileRequest) (*MoveFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (UnimplementedFileServiceServer) Copy(context.Context, *CopyFileRequest) (*CopyFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Copy not implemented")
}
func (UnimplementedFileServiceServer) Mkdir(context.Context, *MkdirFileRequest) (*MkdirFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mkdir not implemented")
}
func (UnimplementedFileServiceServer) Rename(context.Context, *RenameFileRequest) (*RenameFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rename not implemented")
}
func (UnimplementedFileServiceServer) Remove(context.Context, *RemoveFileRequest) (*RemoveFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedFileServiceServer) Upload(FileService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedFileServiceServer) Download(*DownloadFileRequest, FileService_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedFileServiceServer) Preview(*PreviewFileRequest, FileService_PreviewServer) error {
	return status.Errorf(codes.Unimplemented, "method Preview not implemented")
}
func (UnimplementedFileServiceServer) ListRepos(context.Context, *ListReposRequest) (*ListReposResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepos not implemented")
}
func (UnimplementedFileServiceServer) CreateRepo(context.Context, *CreateRepoRequest) (*CreateRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepo not implemented")
}
func (UnimplementedFileServiceServer) UpdateRepo(context.Context, *UpdateRepoRequest) (*UpdateRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRepo not implemented")
}
func (UnimplementedFileServiceServer) DeleteRepo(context.Context, *DeleteRepoRequest) (*DeleteRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepo not implemented")
}
func (UnimplementedFileServiceServer) TestRepo(context.Context, *TestRepoRequest) (*TestRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestRepo not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.file.FileService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).List(ctx, req.(*ListFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.file.FileService/Move",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).Move(ctx, req.(*MoveFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_Copy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).Copy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.file.FileService/Copy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).Copy(ctx, req.(*CopyFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_Mkdir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MkdirFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).Mkdir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.file.FileService/Mkdir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).Mkdir(ctx, req.(*MkdirFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).Rename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.file.FileService/Rename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).Rename(ctx, req.(*RenameFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.file.FileService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).Remove(ctx, req.(*RemoveFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).Upload(&fileServiceUploadServer{stream})
}

type FileService_UploadServer interface {
	SendAndClose(*FileResponse) error
	Recv() (*FileRequest, error)
	grpc.ServerStream
}

type fileServiceUploadServer struct {
	grpc.ServerStream
}

func (x *fileServiceUploadServer) SendAndClose(m *FileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceUploadServer) Recv() (*FileRequest, error) {
	m := new(FileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileService_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServiceServer).Download(m, &fileServiceDownloadServer{stream})
}

type FileService_DownloadServer interface {
	Send(*DownloadFileResponse) error
	grpc.ServerStream
}

type fileServiceDownloadServer struct {
	grpc.ServerStream
}

func (x *fileServiceDownloadServer) Send(m *DownloadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FileService_Preview_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PreviewFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServiceServer).Preview(m, &fileServicePreviewServer{stream})
}

type FileService_PreviewServer interface {
	Send(*PreviewFileResponse) error
	grpc.ServerStream
}

type fileServicePreviewServer struct {
	grpc.ServerStream
}

func (x *fileServicePreviewServer) Send(m *PreviewFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FileService_ListRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReposRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).ListRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.file.FileService/ListRepos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).ListRepos(ctx, req.(*ListReposRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_CreateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).CreateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.file.FileService/CreateRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).CreateRepo(ctx, req.(*CreateRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_UpdateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UpdateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.file.FileService/UpdateRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UpdateRepo(ctx, req.(*UpdateRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_DeleteRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).DeleteRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.file.FileService/DeleteRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).DeleteRepo(ctx, req.(*DeleteRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_TestRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).TestRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.file.FileService/TestRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).TestRepo(ctx, req.(*TestRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.file.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _FileService_List_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _FileService_Move_Handler,
		},
		{
			MethodName: "Copy",
			Handler:    _FileService_Copy_Handler,
		},
		{
			MethodName: "Mkdir",
			Handler:    _FileService_Mkdir_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _FileService_Rename_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _FileService_Remove_Handler,
		},
		{
			MethodName: "ListRepos",
			Handler:    _FileService_ListRepos_Handler,
		},
		{
			MethodName: "CreateRepo",
			Handler:    _FileService_CreateRepo_Handler,
		},
		{
			MethodName: "UpdateRepo",
			Handler:    _FileService_UpdateRepo_Handler,
		},
		{
			MethodName: "DeleteRepo",
			Handler:    _FileService_DeleteRepo_Handler,
		},
		{
			MethodName: "TestRepo",
			Handler:    _FileService_TestRepo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _FileService_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Download",
			Handler:       _FileService_Download_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Preview",
			Handler:       _FileService_Preview_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/file/service.proto",
}
