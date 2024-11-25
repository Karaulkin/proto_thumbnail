// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: thumbnail/thumbnail.proto

package thumbnailv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Thumbnail_GetThumbnail_FullMethodName = "/thumbnail.Thumbnail/GetThumbnail"
)

// ThumbnailClient is the client API for Thumbnail service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThumbnailClient interface {
	GetThumbnail(ctx context.Context, in *ThumbnailRequest, opts ...grpc.CallOption) (*ThumbnailResponse, error)
}

type thumbnailClient struct {
	cc grpc.ClientConnInterface
}

func NewThumbnailClient(cc grpc.ClientConnInterface) ThumbnailClient {
	return &thumbnailClient{cc}
}

func (c *thumbnailClient) GetThumbnail(ctx context.Context, in *ThumbnailRequest, opts ...grpc.CallOption) (*ThumbnailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ThumbnailResponse)
	err := c.cc.Invoke(ctx, Thumbnail_GetThumbnail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThumbnailServer is the server API for Thumbnail service.
// All implementations must embed UnimplementedThumbnailServer
// for forward compatibility.
type ThumbnailServer interface {
	GetThumbnail(context.Context, *ThumbnailRequest) (*ThumbnailResponse, error)
	mustEmbedUnimplementedThumbnailServer()
}

// UnimplementedThumbnailServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedThumbnailServer struct{}

func (UnimplementedThumbnailServer) GetThumbnail(context.Context, *ThumbnailRequest) (*ThumbnailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThumbnail not implemented")
}
func (UnimplementedThumbnailServer) mustEmbedUnimplementedThumbnailServer() {}
func (UnimplementedThumbnailServer) testEmbeddedByValue()                   {}

// UnsafeThumbnailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThumbnailServer will
// result in compilation errors.
type UnsafeThumbnailServer interface {
	mustEmbedUnimplementedThumbnailServer()
}

func RegisterThumbnailServer(s grpc.ServiceRegistrar, srv ThumbnailServer) {
	// If the following call pancis, it indicates UnimplementedThumbnailServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Thumbnail_ServiceDesc, srv)
}

func _Thumbnail_GetThumbnail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThumbnailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThumbnailServer).GetThumbnail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Thumbnail_GetThumbnail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThumbnailServer).GetThumbnail(ctx, req.(*ThumbnailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Thumbnail_ServiceDesc is the grpc.ServiceDesc for Thumbnail service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Thumbnail_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thumbnail.Thumbnail",
	HandlerType: (*ThumbnailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetThumbnail",
			Handler:    _Thumbnail_GetThumbnail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thumbnail/thumbnail.proto",
}
