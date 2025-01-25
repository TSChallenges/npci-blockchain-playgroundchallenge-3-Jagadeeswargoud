// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: orderer/cluster.proto

package orderer

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
	Cluster_Step_FullMethodName = "/orderer.Cluster/Step"
)

// ClusterClient is the client API for Cluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Cluster defines communication between cluster members.
type ClusterClient interface {
	// Step passes an implementation-specific message to another cluster member.
	Step(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StepRequest, StepResponse], error)
}

type clusterClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterClient(cc grpc.ClientConnInterface) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) Step(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StepRequest, StepResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Cluster_ServiceDesc.Streams[0], Cluster_Step_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StepRequest, StepResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Cluster_StepClient = grpc.BidiStreamingClient[StepRequest, StepResponse]

// ClusterServer is the server API for Cluster service.
// All implementations should embed UnimplementedClusterServer
// for forward compatibility.
//
// Cluster defines communication between cluster members.
type ClusterServer interface {
	// Step passes an implementation-specific message to another cluster member.
	Step(grpc.BidiStreamingServer[StepRequest, StepResponse]) error
}

// UnimplementedClusterServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClusterServer struct{}

func (UnimplementedClusterServer) Step(grpc.BidiStreamingServer[StepRequest, StepResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Step not implemented")
}
func (UnimplementedClusterServer) testEmbeddedByValue() {}

// UnsafeClusterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServer will
// result in compilation errors.
type UnsafeClusterServer interface {
	mustEmbedUnimplementedClusterServer()
}

func RegisterClusterServer(s grpc.ServiceRegistrar, srv ClusterServer) {
	// If the following call pancis, it indicates UnimplementedClusterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Cluster_ServiceDesc, srv)
}

func _Cluster_Step_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClusterServer).Step(&grpc.GenericServerStream[StepRequest, StepResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Cluster_StepServer = grpc.BidiStreamingServer[StepRequest, StepResponse]

// Cluster_ServiceDesc is the grpc.ServiceDesc for Cluster service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cluster_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orderer.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Step",
			Handler:       _Cluster_Step_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "orderer/cluster.proto",
}
