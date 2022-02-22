// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/rpc/sentry/relaypeer.proto

package sentry

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

// RelayPeerServiceClient is the client API for RelayPeerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelayPeerServiceClient interface {
	RelayPeerHelloRPC(ctx context.Context, opts ...grpc.CallOption) (RelayPeerService_RelayPeerHelloRPCClient, error)
	RelayPeerProbeRPC(ctx context.Context, opts ...grpc.CallOption) (RelayPeerService_RelayPeerProbeRPCClient, error)
	RelayPeerSurveyRPC(ctx context.Context, opts ...grpc.CallOption) (RelayPeerService_RelayPeerSurveyRPCClient, error)
}

type relayPeerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelayPeerServiceClient(cc grpc.ClientConnInterface) RelayPeerServiceClient {
	return &relayPeerServiceClient{cc}
}

func (c *relayPeerServiceClient) RelayPeerHelloRPC(ctx context.Context, opts ...grpc.CallOption) (RelayPeerService_RelayPeerHelloRPCClient, error) {
	stream, err := c.cc.NewStream(ctx, &RelayPeerService_ServiceDesc.Streams[0], "/rafay.dev.sentry.rpc.RelayPeerService/RelayPeerHelloRPC", opts...)
	if err != nil {
		return nil, err
	}
	x := &relayPeerServiceRelayPeerHelloRPCClient{stream}
	return x, nil
}

type RelayPeerService_RelayPeerHelloRPCClient interface {
	Send(*PeerHelloRequest) error
	Recv() (*PeerHelloResponse, error)
	grpc.ClientStream
}

type relayPeerServiceRelayPeerHelloRPCClient struct {
	grpc.ClientStream
}

func (x *relayPeerServiceRelayPeerHelloRPCClient) Send(m *PeerHelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *relayPeerServiceRelayPeerHelloRPCClient) Recv() (*PeerHelloResponse, error) {
	m := new(PeerHelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *relayPeerServiceClient) RelayPeerProbeRPC(ctx context.Context, opts ...grpc.CallOption) (RelayPeerService_RelayPeerProbeRPCClient, error) {
	stream, err := c.cc.NewStream(ctx, &RelayPeerService_ServiceDesc.Streams[1], "/rafay.dev.sentry.rpc.RelayPeerService/RelayPeerProbeRPC", opts...)
	if err != nil {
		return nil, err
	}
	x := &relayPeerServiceRelayPeerProbeRPCClient{stream}
	return x, nil
}

type RelayPeerService_RelayPeerProbeRPCClient interface {
	Send(*PeerProbeRequest) error
	Recv() (*PeerProbeResponse, error)
	grpc.ClientStream
}

type relayPeerServiceRelayPeerProbeRPCClient struct {
	grpc.ClientStream
}

func (x *relayPeerServiceRelayPeerProbeRPCClient) Send(m *PeerProbeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *relayPeerServiceRelayPeerProbeRPCClient) Recv() (*PeerProbeResponse, error) {
	m := new(PeerProbeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *relayPeerServiceClient) RelayPeerSurveyRPC(ctx context.Context, opts ...grpc.CallOption) (RelayPeerService_RelayPeerSurveyRPCClient, error) {
	stream, err := c.cc.NewStream(ctx, &RelayPeerService_ServiceDesc.Streams[2], "/rafay.dev.sentry.rpc.RelayPeerService/RelayPeerSurveyRPC", opts...)
	if err != nil {
		return nil, err
	}
	x := &relayPeerServiceRelayPeerSurveyRPCClient{stream}
	return x, nil
}

type RelayPeerService_RelayPeerSurveyRPCClient interface {
	Send(*PeerSurveyResponse) error
	Recv() (*PeerSurveyRequest, error)
	grpc.ClientStream
}

type relayPeerServiceRelayPeerSurveyRPCClient struct {
	grpc.ClientStream
}

func (x *relayPeerServiceRelayPeerSurveyRPCClient) Send(m *PeerSurveyResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *relayPeerServiceRelayPeerSurveyRPCClient) Recv() (*PeerSurveyRequest, error) {
	m := new(PeerSurveyRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RelayPeerServiceServer is the server API for RelayPeerService service.
// All implementations should embed UnimplementedRelayPeerServiceServer
// for forward compatibility
type RelayPeerServiceServer interface {
	RelayPeerHelloRPC(RelayPeerService_RelayPeerHelloRPCServer) error
	RelayPeerProbeRPC(RelayPeerService_RelayPeerProbeRPCServer) error
	RelayPeerSurveyRPC(RelayPeerService_RelayPeerSurveyRPCServer) error
}

// UnimplementedRelayPeerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRelayPeerServiceServer struct {
}

func (UnimplementedRelayPeerServiceServer) RelayPeerHelloRPC(RelayPeerService_RelayPeerHelloRPCServer) error {
	return status.Errorf(codes.Unimplemented, "method RelayPeerHelloRPC not implemented")
}
func (UnimplementedRelayPeerServiceServer) RelayPeerProbeRPC(RelayPeerService_RelayPeerProbeRPCServer) error {
	return status.Errorf(codes.Unimplemented, "method RelayPeerProbeRPC not implemented")
}
func (UnimplementedRelayPeerServiceServer) RelayPeerSurveyRPC(RelayPeerService_RelayPeerSurveyRPCServer) error {
	return status.Errorf(codes.Unimplemented, "method RelayPeerSurveyRPC not implemented")
}

// UnsafeRelayPeerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelayPeerServiceServer will
// result in compilation errors.
type UnsafeRelayPeerServiceServer interface {
	mustEmbedUnimplementedRelayPeerServiceServer()
}

func RegisterRelayPeerServiceServer(s grpc.ServiceRegistrar, srv RelayPeerServiceServer) {
	s.RegisterService(&RelayPeerService_ServiceDesc, srv)
}

func _RelayPeerService_RelayPeerHelloRPC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RelayPeerServiceServer).RelayPeerHelloRPC(&relayPeerServiceRelayPeerHelloRPCServer{stream})
}

type RelayPeerService_RelayPeerHelloRPCServer interface {
	Send(*PeerHelloResponse) error
	Recv() (*PeerHelloRequest, error)
	grpc.ServerStream
}

type relayPeerServiceRelayPeerHelloRPCServer struct {
	grpc.ServerStream
}

func (x *relayPeerServiceRelayPeerHelloRPCServer) Send(m *PeerHelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *relayPeerServiceRelayPeerHelloRPCServer) Recv() (*PeerHelloRequest, error) {
	m := new(PeerHelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RelayPeerService_RelayPeerProbeRPC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RelayPeerServiceServer).RelayPeerProbeRPC(&relayPeerServiceRelayPeerProbeRPCServer{stream})
}

type RelayPeerService_RelayPeerProbeRPCServer interface {
	Send(*PeerProbeResponse) error
	Recv() (*PeerProbeRequest, error)
	grpc.ServerStream
}

type relayPeerServiceRelayPeerProbeRPCServer struct {
	grpc.ServerStream
}

func (x *relayPeerServiceRelayPeerProbeRPCServer) Send(m *PeerProbeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *relayPeerServiceRelayPeerProbeRPCServer) Recv() (*PeerProbeRequest, error) {
	m := new(PeerProbeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RelayPeerService_RelayPeerSurveyRPC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RelayPeerServiceServer).RelayPeerSurveyRPC(&relayPeerServiceRelayPeerSurveyRPCServer{stream})
}

type RelayPeerService_RelayPeerSurveyRPCServer interface {
	Send(*PeerSurveyRequest) error
	Recv() (*PeerSurveyResponse, error)
	grpc.ServerStream
}

type relayPeerServiceRelayPeerSurveyRPCServer struct {
	grpc.ServerStream
}

func (x *relayPeerServiceRelayPeerSurveyRPCServer) Send(m *PeerSurveyRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *relayPeerServiceRelayPeerSurveyRPCServer) Recv() (*PeerSurveyResponse, error) {
	m := new(PeerSurveyResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RelayPeerService_ServiceDesc is the grpc.ServiceDesc for RelayPeerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelayPeerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rafay.dev.sentry.rpc.RelayPeerService",
	HandlerType: (*RelayPeerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RelayPeerHelloRPC",
			Handler:       _RelayPeerService_RelayPeerHelloRPC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RelayPeerProbeRPC",
			Handler:       _RelayPeerService_RelayPeerProbeRPC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RelayPeerSurveyRPC",
			Handler:       _RelayPeerService_RelayPeerSurveyRPC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/rpc/sentry/relaypeer.proto",
}
