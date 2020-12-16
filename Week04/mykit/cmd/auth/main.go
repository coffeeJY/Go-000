package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	pb "mykit/api/auth/grpc"
	"mykit/internal/auth/repository"
	"mykit/internal/auth/service"
	signal "mykit/pkg/signal"
	grpcTransport "mykit/pkg/transport/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)

	// transport server
	grpcSrv := grpcTransport.NewServer(":8000")

	repo := repository.NewRepository()
	gs := service.NewService(repo)

	pb.RegisterAuthServerServer(grpcSrv.Server, gs)
	fmt.Println("Listen on " + ":8000")

	g.Go(func() error {
		return grpcSrv.Start(ctx)
	})
	g.Go(func() error {
		return signal.CheckExitSignal(ctx, cancel)
	})

	if err := g.Wait(); err != nil {
		log.Printf("%+v\n", err)
	}

	log.Println("All servers have exit success!!")
}
