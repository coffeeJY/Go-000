package grpc

import (
	"google.golang.org/grpc"
)

// ServerOption is gRPC server option.
type ServerOption func(o *serverOptions)

type serverOptions struct {
	grpcOpts []grpc.ServerOption
}

func ServerOptions(s ...grpc.ServerOption) ServerOption {
	return func(o *serverOptions) { o.grpcOpts = append(o.grpcOpts, s...) }
}
