// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: relation/v1/grpc.proto

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RelationService_ListFollowings_FullMethodName = "/relation.v1.RelationService/ListFollowings"
	RelationService_ListFollowers_FullMethodName  = "/relation.v1.RelationService/ListFollowers"
	RelationService_CreateRelation_FullMethodName = "/relation.v1.RelationService/CreateRelation"
	RelationService_DeleteRelation_FullMethodName = "/relation.v1.RelationService/DeleteRelation"
)

// RelationServiceClient is the client API for RelationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// *
// RelationService is the service that provides access to user relations
// for other services.
// It is assumed that consumers pass the userID
// when making requests on their behalf
// in the context metadata (key: x-user-id).
type RelationServiceClient interface {
	ListFollowings(ctx context.Context, in *ListFollowingsRequest, opts ...grpc.CallOption) (*ListFollowingsResponse, error)
	ListFollowers(ctx context.Context, in *ListFollowersRequest, opts ...grpc.CallOption) (*ListFollowersResponse, error)
	CreateRelation(ctx context.Context, in *CreateRelationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteRelation(ctx context.Context, in *DeleteRelationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type relationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationServiceClient(cc grpc.ClientConnInterface) RelationServiceClient {
	return &relationServiceClient{cc}
}

func (c *relationServiceClient) ListFollowings(ctx context.Context, in *ListFollowingsRequest, opts ...grpc.CallOption) (*ListFollowingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFollowingsResponse)
	err := c.cc.Invoke(ctx, RelationService_ListFollowings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) ListFollowers(ctx context.Context, in *ListFollowersRequest, opts ...grpc.CallOption) (*ListFollowersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFollowersResponse)
	err := c.cc.Invoke(ctx, RelationService_ListFollowers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) CreateRelation(ctx context.Context, in *CreateRelationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RelationService_CreateRelation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) DeleteRelation(ctx context.Context, in *DeleteRelationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RelationService_DeleteRelation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServiceServer is the server API for RelationService service.
// All implementations must embed UnimplementedRelationServiceServer
// for forward compatibility.
//
// *
// RelationService is the service that provides access to user relations
// for other services.
// It is assumed that consumers pass the userID
// when making requests on their behalf
// in the context metadata (key: x-user-id).
type RelationServiceServer interface {
	ListFollowings(context.Context, *ListFollowingsRequest) (*ListFollowingsResponse, error)
	ListFollowers(context.Context, *ListFollowersRequest) (*ListFollowersResponse, error)
	CreateRelation(context.Context, *CreateRelationRequest) (*emptypb.Empty, error)
	DeleteRelation(context.Context, *DeleteRelationRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedRelationServiceServer()
}

// UnimplementedRelationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRelationServiceServer struct{}

func (UnimplementedRelationServiceServer) ListFollowings(context.Context, *ListFollowingsRequest) (*ListFollowingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFollowings not implemented")
}
func (UnimplementedRelationServiceServer) ListFollowers(context.Context, *ListFollowersRequest) (*ListFollowersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFollowers not implemented")
}
func (UnimplementedRelationServiceServer) CreateRelation(context.Context, *CreateRelationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelation not implemented")
}
func (UnimplementedRelationServiceServer) DeleteRelation(context.Context, *DeleteRelationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelation not implemented")
}
func (UnimplementedRelationServiceServer) mustEmbedUnimplementedRelationServiceServer() {}
func (UnimplementedRelationServiceServer) testEmbeddedByValue()                         {}

// UnsafeRelationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationServiceServer will
// result in compilation errors.
type UnsafeRelationServiceServer interface {
	mustEmbedUnimplementedRelationServiceServer()
}

func RegisterRelationServiceServer(s grpc.ServiceRegistrar, srv RelationServiceServer) {
	// If the following call pancis, it indicates UnimplementedRelationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RelationService_ServiceDesc, srv)
}

func _RelationService_ListFollowings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFollowingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).ListFollowings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_ListFollowings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).ListFollowings(ctx, req.(*ListFollowingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_ListFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFollowersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).ListFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_ListFollowers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).ListFollowers(ctx, req.(*ListFollowersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_CreateRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).CreateRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_CreateRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).CreateRelation(ctx, req.(*CreateRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_DeleteRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).DeleteRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_DeleteRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).DeleteRelation(ctx, req.(*DeleteRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationService_ServiceDesc is the grpc.ServiceDesc for RelationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relation.v1.RelationService",
	HandlerType: (*RelationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFollowings",
			Handler:    _RelationService_ListFollowings_Handler,
		},
		{
			MethodName: "ListFollowers",
			Handler:    _RelationService_ListFollowers_Handler,
		},
		{
			MethodName: "CreateRelation",
			Handler:    _RelationService_CreateRelation_Handler,
		},
		{
			MethodName: "DeleteRelation",
			Handler:    _RelationService_DeleteRelation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relation/v1/grpc.proto",
}