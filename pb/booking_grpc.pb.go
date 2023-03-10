// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: booking.proto

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

// BookingServiceClient is the client API for BookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingServiceClient interface {
	CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error)
	CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*BookingRespnose, error)
	SearchBooking(ctx context.Context, in *SearchBookingRequest, opts ...grpc.CallOption) (*BookingInfor, error)
	SearchBookingId(ctx context.Context, in *SearchBookingByIdRequest, opts ...grpc.CallOption) (*ListBooking, error)
}

type bookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceClient(cc grpc.ClientConnInterface) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error) {
	out := new(CreateBookingResponse)
	err := c.cc.Invoke(ctx, "/booking_asm.BookingService/CreateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*BookingRespnose, error) {
	out := new(BookingRespnose)
	err := c.cc.Invoke(ctx, "/booking_asm.BookingService/CancelBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) SearchBooking(ctx context.Context, in *SearchBookingRequest, opts ...grpc.CallOption) (*BookingInfor, error) {
	out := new(BookingInfor)
	err := c.cc.Invoke(ctx, "/booking_asm.BookingService/SearchBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) SearchBookingId(ctx context.Context, in *SearchBookingByIdRequest, opts ...grpc.CallOption) (*ListBooking, error) {
	out := new(ListBooking)
	err := c.cc.Invoke(ctx, "/booking_asm.BookingService/SearchBookingId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
// All implementations must embed UnimplementedBookingServiceServer
// for forward compatibility
type BookingServiceServer interface {
	CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error)
	CancelBooking(context.Context, *CancelBookingRequest) (*BookingRespnose, error)
	SearchBooking(context.Context, *SearchBookingRequest) (*BookingInfor, error)
	SearchBookingId(context.Context, *SearchBookingByIdRequest) (*ListBooking, error)
	mustEmbedUnimplementedBookingServiceServer()
}

// UnimplementedBookingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookingServiceServer struct {
}

func (UnimplementedBookingServiceServer) CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedBookingServiceServer) CancelBooking(context.Context, *CancelBookingRequest) (*BookingRespnose, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBooking not implemented")
}
func (UnimplementedBookingServiceServer) SearchBooking(context.Context, *SearchBookingRequest) (*BookingInfor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBooking not implemented")
}
func (UnimplementedBookingServiceServer) SearchBookingId(context.Context, *SearchBookingByIdRequest) (*ListBooking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBookingId not implemented")
}
func (UnimplementedBookingServiceServer) mustEmbedUnimplementedBookingServiceServer() {}

// UnsafeBookingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServiceServer will
// result in compilation errors.
type UnsafeBookingServiceServer interface {
	mustEmbedUnimplementedBookingServiceServer()
}

func RegisterBookingServiceServer(s grpc.ServiceRegistrar, srv BookingServiceServer) {
	s.RegisterService(&BookingService_ServiceDesc, srv)
}

func _BookingService_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking_asm.BookingService/CreateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateBooking(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking_asm.BookingService/CancelBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CancelBooking(ctx, req.(*CancelBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_SearchBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).SearchBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking_asm.BookingService/SearchBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).SearchBooking(ctx, req.(*SearchBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_SearchBookingId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBookingByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).SearchBookingId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking_asm.BookingService/SearchBookingId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).SearchBookingId(ctx, req.(*SearchBookingByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingService_ServiceDesc is the grpc.ServiceDesc for BookingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking_asm.BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _BookingService_CreateBooking_Handler,
		},
		{
			MethodName: "CancelBooking",
			Handler:    _BookingService_CancelBooking_Handler,
		},
		{
			MethodName: "SearchBooking",
			Handler:    _BookingService_SearchBooking_Handler,
		},
		{
			MethodName: "SearchBookingId",
			Handler:    _BookingService_SearchBookingId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}
