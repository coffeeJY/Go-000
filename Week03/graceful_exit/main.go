package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	//"io"
	"log"
	"net/http"

	errors "github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signalCh)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return openServer(ctx, ":8080", appServerHandler)
	})

	g.Go(func() error {
		return openServer(ctx, ":8081", debugServerHandler)
	})

	g.Go(func() error {
		return checkExitSignal(ctx, cancel, signalCh)
	})

	if err := g.Wait(); err != nil {
		log.Printf("%+v\n", err)
	}

	log.Println("All servers have exit success!!")
}

func openServer(ctx context.Context, addr string, handler http.HandlerFunc) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	errChan := make(chan error, 1)
	go func() {
		<-ctx.Done()
		shutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutCtx); err != nil {
			errChan <- errors.Wrap(http.ErrServerClosed, err.Error())
		}
		log.Printf("The server graceful exit ：%v \n", addr)
		close(errChan)
	}()

	log.Printf("The server is listening : %s\n", addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return errors.Wrap(err, fmt.Sprintf("%s, server : %s", err.Error(), addr))
	}
	return <-errChan
}

func checkExitSignal(ctx context.Context, cancel context.CancelFunc, signalCh <-chan os.Signal) error {
	select {
	case <-signalCh:
		log.Println("Get signal and start exit...")
		cancel()
		return errors.Wrapf(context.Canceled, "get exit signal")
	case <-ctx.Done():
		return nil
	}
}

func appServerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is appServer !")
}

func debugServerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is debugServer !")
}
