// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: upload.proto

package chunker

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
	UploadFile(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) UploadFile(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, "/chunker.FileService/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	UploadFile(context.Context, *UploadRequest) (*UploadResponse, error)
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) UploadFile(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
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

func _FileService_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chunker.FileService/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UploadFile(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chunker.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadFile",
			Handler:    _FileService_UploadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upload.proto",
}
