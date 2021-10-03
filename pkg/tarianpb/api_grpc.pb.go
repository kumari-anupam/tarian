// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tarianpb

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

// ConfigClient is the client API for Config service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigClient interface {
	GetConstraints(ctx context.Context, in *GetConstraintsRequest, opts ...grpc.CallOption) (*GetConstraintsResponse, error)
	AddConstraint(ctx context.Context, in *AddConstraintRequest, opts ...grpc.CallOption) (*AddConstraintResponse, error)
	RemoveConstraint(ctx context.Context, in *RemoveConstraintRequest, opts ...grpc.CallOption) (*RemoveConstraintResponse, error)
	AddAction(ctx context.Context, in *AddActionRequest, opts ...grpc.CallOption) (*AddActionResponse, error)
	GetActions(ctx context.Context, in *GetActionsRequest, opts ...grpc.CallOption) (*GetActionsResponse, error)
	RemoveAction(ctx context.Context, in *RemoveActionRequest, opts ...grpc.CallOption) (*RemoveActionResponse, error)
}

type configClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigClient(cc grpc.ClientConnInterface) ConfigClient {
	return &configClient{cc}
}

func (c *configClient) GetConstraints(ctx context.Context, in *GetConstraintsRequest, opts ...grpc.CallOption) (*GetConstraintsResponse, error) {
	out := new(GetConstraintsResponse)
	err := c.cc.Invoke(ctx, "/tarianpb.api.Config/GetConstraints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) AddConstraint(ctx context.Context, in *AddConstraintRequest, opts ...grpc.CallOption) (*AddConstraintResponse, error) {
	out := new(AddConstraintResponse)
	err := c.cc.Invoke(ctx, "/tarianpb.api.Config/AddConstraint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) RemoveConstraint(ctx context.Context, in *RemoveConstraintRequest, opts ...grpc.CallOption) (*RemoveConstraintResponse, error) {
	out := new(RemoveConstraintResponse)
	err := c.cc.Invoke(ctx, "/tarianpb.api.Config/RemoveConstraint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) AddAction(ctx context.Context, in *AddActionRequest, opts ...grpc.CallOption) (*AddActionResponse, error) {
	out := new(AddActionResponse)
	err := c.cc.Invoke(ctx, "/tarianpb.api.Config/AddAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) GetActions(ctx context.Context, in *GetActionsRequest, opts ...grpc.CallOption) (*GetActionsResponse, error) {
	out := new(GetActionsResponse)
	err := c.cc.Invoke(ctx, "/tarianpb.api.Config/GetActions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) RemoveAction(ctx context.Context, in *RemoveActionRequest, opts ...grpc.CallOption) (*RemoveActionResponse, error) {
	out := new(RemoveActionResponse)
	err := c.cc.Invoke(ctx, "/tarianpb.api.Config/RemoveAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServer is the server API for Config service.
// All implementations must embed UnimplementedConfigServer
// for forward compatibility
type ConfigServer interface {
	GetConstraints(context.Context, *GetConstraintsRequest) (*GetConstraintsResponse, error)
	AddConstraint(context.Context, *AddConstraintRequest) (*AddConstraintResponse, error)
	RemoveConstraint(context.Context, *RemoveConstraintRequest) (*RemoveConstraintResponse, error)
	AddAction(context.Context, *AddActionRequest) (*AddActionResponse, error)
	GetActions(context.Context, *GetActionsRequest) (*GetActionsResponse, error)
	RemoveAction(context.Context, *RemoveActionRequest) (*RemoveActionResponse, error)
	mustEmbedUnimplementedConfigServer()
}

// UnimplementedConfigServer must be embedded to have forward compatible implementations.
type UnimplementedConfigServer struct {
}

func (UnimplementedConfigServer) GetConstraints(context.Context, *GetConstraintsRequest) (*GetConstraintsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConstraints not implemented")
}
func (UnimplementedConfigServer) AddConstraint(context.Context, *AddConstraintRequest) (*AddConstraintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddConstraint not implemented")
}
func (UnimplementedConfigServer) RemoveConstraint(context.Context, *RemoveConstraintRequest) (*RemoveConstraintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveConstraint not implemented")
}
func (UnimplementedConfigServer) AddAction(context.Context, *AddActionRequest) (*AddActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAction not implemented")
}
func (UnimplementedConfigServer) GetActions(context.Context, *GetActionsRequest) (*GetActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActions not implemented")
}
func (UnimplementedConfigServer) RemoveAction(context.Context, *RemoveActionRequest) (*RemoveActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAction not implemented")
}
func (UnimplementedConfigServer) mustEmbedUnimplementedConfigServer() {}

// UnsafeConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServer will
// result in compilation errors.
type UnsafeConfigServer interface {
	mustEmbedUnimplementedConfigServer()
}

func RegisterConfigServer(s grpc.ServiceRegistrar, srv ConfigServer) {
	s.RegisterService(&Config_ServiceDesc, srv)
}

func _Config_GetConstraints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConstraintsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).GetConstraints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarianpb.api.Config/GetConstraints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).GetConstraints(ctx, req.(*GetConstraintsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_AddConstraint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddConstraintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).AddConstraint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarianpb.api.Config/AddConstraint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).AddConstraint(ctx, req.(*AddConstraintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_RemoveConstraint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveConstraintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).RemoveConstraint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarianpb.api.Config/RemoveConstraint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).RemoveConstraint(ctx, req.(*RemoveConstraintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_AddAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).AddAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarianpb.api.Config/AddAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).AddAction(ctx, req.(*AddActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_GetActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).GetActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarianpb.api.Config/GetActions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).GetActions(ctx, req.(*GetActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_RemoveAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).RemoveAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarianpb.api.Config/RemoveAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).RemoveAction(ctx, req.(*RemoveActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Config_ServiceDesc is the grpc.ServiceDesc for Config service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Config_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tarianpb.api.Config",
	HandlerType: (*ConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConstraints",
			Handler:    _Config_GetConstraints_Handler,
		},
		{
			MethodName: "AddConstraint",
			Handler:    _Config_AddConstraint_Handler,
		},
		{
			MethodName: "RemoveConstraint",
			Handler:    _Config_RemoveConstraint_Handler,
		},
		{
			MethodName: "AddAction",
			Handler:    _Config_AddAction_Handler,
		},
		{
			MethodName: "GetActions",
			Handler:    _Config_GetActions_Handler,
		},
		{
			MethodName: "RemoveAction",
			Handler:    _Config_RemoveAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tarianpb/api.proto",
}

// EventClient is the client API for Event service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventClient interface {
	IngestEvent(ctx context.Context, in *IngestEventRequest, opts ...grpc.CallOption) (*IngestEventResponse, error)
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error)
}

type eventClient struct {
	cc grpc.ClientConnInterface
}

func NewEventClient(cc grpc.ClientConnInterface) EventClient {
	return &eventClient{cc}
}

func (c *eventClient) IngestEvent(ctx context.Context, in *IngestEventRequest, opts ...grpc.CallOption) (*IngestEventResponse, error) {
	out := new(IngestEventResponse)
	err := c.cc.Invoke(ctx, "/tarianpb.api.Event/IngestEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error) {
	out := new(GetEventsResponse)
	err := c.cc.Invoke(ctx, "/tarianpb.api.Event/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServer is the server API for Event service.
// All implementations must embed UnimplementedEventServer
// for forward compatibility
type EventServer interface {
	IngestEvent(context.Context, *IngestEventRequest) (*IngestEventResponse, error)
	GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error)
	mustEmbedUnimplementedEventServer()
}

// UnimplementedEventServer must be embedded to have forward compatible implementations.
type UnimplementedEventServer struct {
}

func (UnimplementedEventServer) IngestEvent(context.Context, *IngestEventRequest) (*IngestEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IngestEvent not implemented")
}
func (UnimplementedEventServer) GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedEventServer) mustEmbedUnimplementedEventServer() {}

// UnsafeEventServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServer will
// result in compilation errors.
type UnsafeEventServer interface {
	mustEmbedUnimplementedEventServer()
}

func RegisterEventServer(s grpc.ServiceRegistrar, srv EventServer) {
	s.RegisterService(&Event_ServiceDesc, srv)
}

func _Event_IngestEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngestEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).IngestEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarianpb.api.Event/IngestEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).IngestEvent(ctx, req.(*IngestEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarianpb.api.Event/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).GetEvents(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Event_ServiceDesc is the grpc.ServiceDesc for Event service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Event_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tarianpb.api.Event",
	HandlerType: (*EventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IngestEvent",
			Handler:    _Event_IngestEvent_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _Event_GetEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tarianpb/api.proto",
}
