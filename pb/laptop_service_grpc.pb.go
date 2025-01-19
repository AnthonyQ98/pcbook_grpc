// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: laptop_service.proto

package pb

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
	LaptopService_CreateLaptop_FullMethodName = "/pcbook.LaptopService/CreateLaptop"
	LaptopService_SearchLaptop_FullMethodName = "/pcbook.LaptopService/SearchLaptop"
	LaptopService_UploadImage_FullMethodName  = "/pcbook.LaptopService/UploadImage"
	LaptopService_RateLaptop_FullMethodName   = "/pcbook.LaptopService/RateLaptop"
)

// LaptopServiceClient is the client API for LaptopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LaptopServiceClient interface {
	CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error)
	SearchLaptop(ctx context.Context, in *SearchLaptopRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SearchLaptopResponse], error)
	UploadImage(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadImageRequest, UploadImageResponse], error)
	RateLaptop(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[RateLaptopRequest, RateLaptopResponse], error)
}

type laptopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLaptopServiceClient(cc grpc.ClientConnInterface) LaptopServiceClient {
	return &laptopServiceClient{cc}
}

func (c *laptopServiceClient) CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLaptopResponse)
	err := c.cc.Invoke(ctx, LaptopService_CreateLaptop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *laptopServiceClient) SearchLaptop(ctx context.Context, in *SearchLaptopRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SearchLaptopResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &LaptopService_ServiceDesc.Streams[0], LaptopService_SearchLaptop_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SearchLaptopRequest, SearchLaptopResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LaptopService_SearchLaptopClient = grpc.ServerStreamingClient[SearchLaptopResponse]

func (c *laptopServiceClient) UploadImage(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadImageRequest, UploadImageResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &LaptopService_ServiceDesc.Streams[1], LaptopService_UploadImage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadImageRequest, UploadImageResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LaptopService_UploadImageClient = grpc.ClientStreamingClient[UploadImageRequest, UploadImageResponse]

func (c *laptopServiceClient) RateLaptop(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[RateLaptopRequest, RateLaptopResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &LaptopService_ServiceDesc.Streams[2], LaptopService_RateLaptop_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[RateLaptopRequest, RateLaptopResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LaptopService_RateLaptopClient = grpc.BidiStreamingClient[RateLaptopRequest, RateLaptopResponse]

// LaptopServiceServer is the server API for LaptopService service.
// All implementations must embed UnimplementedLaptopServiceServer
// for forward compatibility.
type LaptopServiceServer interface {
	CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error)
	SearchLaptop(*SearchLaptopRequest, grpc.ServerStreamingServer[SearchLaptopResponse]) error
	UploadImage(grpc.ClientStreamingServer[UploadImageRequest, UploadImageResponse]) error
	RateLaptop(grpc.BidiStreamingServer[RateLaptopRequest, RateLaptopResponse]) error
	mustEmbedUnimplementedLaptopServiceServer()
}

// UnimplementedLaptopServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLaptopServiceServer struct{}

func (UnimplementedLaptopServiceServer) CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLaptop not implemented")
}
func (UnimplementedLaptopServiceServer) SearchLaptop(*SearchLaptopRequest, grpc.ServerStreamingServer[SearchLaptopResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SearchLaptop not implemented")
}
func (UnimplementedLaptopServiceServer) UploadImage(grpc.ClientStreamingServer[UploadImageRequest, UploadImageResponse]) error {
	return status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedLaptopServiceServer) RateLaptop(grpc.BidiStreamingServer[RateLaptopRequest, RateLaptopResponse]) error {
	return status.Errorf(codes.Unimplemented, "method RateLaptop not implemented")
}
func (UnimplementedLaptopServiceServer) mustEmbedUnimplementedLaptopServiceServer() {}
func (UnimplementedLaptopServiceServer) testEmbeddedByValue()                       {}

// UnsafeLaptopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LaptopServiceServer will
// result in compilation errors.
type UnsafeLaptopServiceServer interface {
	mustEmbedUnimplementedLaptopServiceServer()
}

func RegisterLaptopServiceServer(s grpc.ServiceRegistrar, srv LaptopServiceServer) {
	// If the following call pancis, it indicates UnimplementedLaptopServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LaptopService_ServiceDesc, srv)
}

func _LaptopService_CreateLaptop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLaptopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LaptopService_CreateLaptop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, req.(*CreateLaptopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LaptopService_SearchLaptop_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchLaptopRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LaptopServiceServer).SearchLaptop(m, &grpc.GenericServerStream[SearchLaptopRequest, SearchLaptopResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LaptopService_SearchLaptopServer = grpc.ServerStreamingServer[SearchLaptopResponse]

func _LaptopService_UploadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LaptopServiceServer).UploadImage(&grpc.GenericServerStream[UploadImageRequest, UploadImageResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LaptopService_UploadImageServer = grpc.ClientStreamingServer[UploadImageRequest, UploadImageResponse]

func _LaptopService_RateLaptop_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LaptopServiceServer).RateLaptop(&grpc.GenericServerStream[RateLaptopRequest, RateLaptopResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LaptopService_RateLaptopServer = grpc.BidiStreamingServer[RateLaptopRequest, RateLaptopResponse]

// LaptopService_ServiceDesc is the grpc.ServiceDesc for LaptopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LaptopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pcbook.LaptopService",
	HandlerType: (*LaptopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLaptop",
			Handler:    _LaptopService_CreateLaptop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchLaptop",
			Handler:       _LaptopService_SearchLaptop_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadImage",
			Handler:       _LaptopService_UploadImage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RateLaptop",
			Handler:       _LaptopService_RateLaptop_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "laptop_service.proto",
}
