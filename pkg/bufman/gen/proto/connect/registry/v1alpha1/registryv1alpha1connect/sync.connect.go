// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: registry/v1alpha1/sync.proto

package registryv1alpha1connect

import (
	context "context"
	errors "errors"
	v1alpha1 "github.com/apache/dubbo-kubernetes/pkg/bufman/gen/proto/go/registry/v1alpha1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion1_7_0

const (
	// SyncServiceName is the fully-qualified name of the SyncService service.
	SyncServiceName = "bufman.dubbo.apache.org.registry.v1alpha1.SyncService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SyncServiceGetGitSyncPointProcedure is the fully-qualified name of the SyncService's
	// GetGitSyncPoint RPC.
	SyncServiceGetGitSyncPointProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.SyncService/GetGitSyncPoint"
	// SyncServiceSyncGitCommitProcedure is the fully-qualified name of the SyncService's SyncGitCommit
	// RPC.
	SyncServiceSyncGitCommitProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.SyncService/SyncGitCommit"
)

// SyncServiceClient is a client for the bufman.dubbo.apache.org.registry.v1alpha1.SyncService
// service.
type SyncServiceClient interface {
	// GetGitSyncPoint retrieves the Git sync point for the named repository
	// on the specified branch.
	GetGitSyncPoint(context.Context, *connect_go.Request[v1alpha1.GetGitSyncPointRequest]) (*connect_go.Response[v1alpha1.GetGitSyncPointResponse], error)
	// SyncGitCommit syncs a Git commit containing a module to a named repository.
	SyncGitCommit(context.Context, *connect_go.Request[v1alpha1.SyncGitCommitRequest]) (*connect_go.Response[v1alpha1.SyncGitCommitResponse], error)
}

// NewSyncServiceClient constructs a client for the
// bufman.dubbo.apache.org.registry.v1alpha1.SyncService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSyncServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) SyncServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &syncServiceClient{
		getGitSyncPoint: connect_go.NewClient[v1alpha1.GetGitSyncPointRequest, v1alpha1.GetGitSyncPointResponse](
			httpClient,
			baseURL+SyncServiceGetGitSyncPointProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		syncGitCommit: connect_go.NewClient[v1alpha1.SyncGitCommitRequest, v1alpha1.SyncGitCommitResponse](
			httpClient,
			baseURL+SyncServiceSyncGitCommitProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyIdempotent),
			connect_go.WithClientOptions(opts...),
		),
	}
}

// syncServiceClient implements SyncServiceClient.
type syncServiceClient struct {
	getGitSyncPoint *connect_go.Client[v1alpha1.GetGitSyncPointRequest, v1alpha1.GetGitSyncPointResponse]
	syncGitCommit   *connect_go.Client[v1alpha1.SyncGitCommitRequest, v1alpha1.SyncGitCommitResponse]
}

// GetGitSyncPoint calls bufman.dubbo.apache.org.registry.v1alpha1.SyncService.GetGitSyncPoint.
func (c *syncServiceClient) GetGitSyncPoint(ctx context.Context, req *connect_go.Request[v1alpha1.GetGitSyncPointRequest]) (*connect_go.Response[v1alpha1.GetGitSyncPointResponse], error) {
	return c.getGitSyncPoint.CallUnary(ctx, req)
}

// SyncGitCommit calls bufman.dubbo.apache.org.registry.v1alpha1.SyncService.SyncGitCommit.
func (c *syncServiceClient) SyncGitCommit(ctx context.Context, req *connect_go.Request[v1alpha1.SyncGitCommitRequest]) (*connect_go.Response[v1alpha1.SyncGitCommitResponse], error) {
	return c.syncGitCommit.CallUnary(ctx, req)
}

// SyncServiceHandler is an implementation of the
// bufman.dubbo.apache.org.registry.v1alpha1.SyncService service.
type SyncServiceHandler interface {
	// GetGitSyncPoint retrieves the Git sync point for the named repository
	// on the specified branch.
	GetGitSyncPoint(context.Context, *connect_go.Request[v1alpha1.GetGitSyncPointRequest]) (*connect_go.Response[v1alpha1.GetGitSyncPointResponse], error)
	// SyncGitCommit syncs a Git commit containing a module to a named repository.
	SyncGitCommit(context.Context, *connect_go.Request[v1alpha1.SyncGitCommitRequest]) (*connect_go.Response[v1alpha1.SyncGitCommitResponse], error)
}

// NewSyncServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSyncServiceHandler(svc SyncServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	syncServiceGetGitSyncPointHandler := connect_go.NewUnaryHandler(
		SyncServiceGetGitSyncPointProcedure,
		svc.GetGitSyncPoint,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	syncServiceSyncGitCommitHandler := connect_go.NewUnaryHandler(
		SyncServiceSyncGitCommitProcedure,
		svc.SyncGitCommit,
		connect_go.WithIdempotency(connect_go.IdempotencyIdempotent),
		connect_go.WithHandlerOptions(opts...),
	)
	return "/bufman.dubbo.apache.org.registry.v1alpha1.SyncService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SyncServiceGetGitSyncPointProcedure:
			syncServiceGetGitSyncPointHandler.ServeHTTP(w, r)
		case SyncServiceSyncGitCommitProcedure:
			syncServiceSyncGitCommitHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSyncServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSyncServiceHandler struct{}

func (UnimplementedSyncServiceHandler) GetGitSyncPoint(context.Context, *connect_go.Request[v1alpha1.GetGitSyncPointRequest]) (*connect_go.Response[v1alpha1.GetGitSyncPointResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.SyncService.GetGitSyncPoint is not implemented"))
}

func (UnimplementedSyncServiceHandler) SyncGitCommit(context.Context, *connect_go.Request[v1alpha1.SyncGitCommitRequest]) (*connect_go.Response[v1alpha1.SyncGitCommitResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.SyncService.SyncGitCommit is not implemented"))
}
