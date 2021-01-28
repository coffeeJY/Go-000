package main

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signalCh)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return listenServer(ctx, ":8889")
	})

	g.Go(func() error {
		return checkExitSignal(ctx, cancel, signalCh)
	})

	if err := g.Wait(); err != nil {
		log.Printf("%+v\n", err)
	}

	log.Println("All servers have exit success!!")
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

func listenServer(ctx context.Context, addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("net listen err : %v \n", err)
		return errors.Wrap(err, "net Listen 8889 error")
	}
	log.Println("listening 8889...")
	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				log.Println("accept error:", err)
				return
			}

			go HandleConn(ctx, conn)
		}
	}()
	for {
		<-ctx.Done()
		return ctx.Err()
	}

}

func HandleConn(ctx context.Context, c net.Conn) {
	log.Printf("create a conn... \n")
	closeCh := make(chan struct{})
	bufCh := make(chan []byte, 1)
	defer func() {
		log.Println("close Conn")
		c.Close()
		close(bufCh)
		close(closeCh)
	}()

	// read from the connection
	go readBuf(c, bufCh, closeCh)

	// write to the connection
	var buf []byte
	for {
		select {
		case buf = <-bufCh:
			writeBuf(c, buf, closeCh)
		case <-closeCh:
			return
		case <-ctx.Done():
			log.Printf("exit conn")
			return
		}
	}
}

func readBuf(c net.Conn, bufCh chan<- []byte, closeCh chan struct{}) {
	var (
		buf = make([]byte, 64)
		n   int
		err error
	)
	for {
		time.Sleep(time.Second * 1)
		n, err = c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			log.Printf("send err to closeCh \n")
			closeCh <- struct{}{}
			return
		} else {
			log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
			bufCh <- buf[:n]
		}
	}
}

func writeBuf(c net.Conn, buf []byte, closeCh chan struct{}) {
	var (
		n   int
		err error
	)

	n, err = c.Write(buf)
	if err != nil {
		log.Println("conn write error:", err)
		closeCh <- struct{}{}
	} else {
		log.Printf("write %d bytes, content is %s\n", n, string(buf[:n]))
	}
}
