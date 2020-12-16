package grpc

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"log"
	"net"
	"time"
)

// Server is a gRPC server wrapper.
type Server struct {
	*grpc.Server
	addr string
	opts serverOptions
}

// NewServer creates a gRPC server by options.
func NewServer(addr string, opts ...ServerOption) *Server {
	options := serverOptions{}
	for _, op := range opts {
		op(&options)
	}
	srv := &Server{
		addr: addr,
		opts: options,
	}
	srv.Server = grpc.NewServer(options.grpcOpts...)
	return srv
}

// Start start the gRPC server.
func (s *Server) Start(ctx context.Context) error {

	errChan := make(chan error, 1)
	go func() {
		<-ctx.Done()
		shutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.Stop(shutCtx); err != nil {
			errChan <- errors.Wrap(status.Error(codes.FailedPrecondition, "grpc graceful stop error"), err.Error())
		}
		log.Printf("The server graceful exit ：%v \n", s.addr)
		close(errChan)
	}()

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf(" Listen server err: %s,%s", err.Error(), s.addr))
	}
	return s.Serve(lis)
}

// Stop stop the gRPC server.
func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	return nil
}
