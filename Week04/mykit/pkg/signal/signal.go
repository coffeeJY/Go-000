package signal

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func CheckExitSignal(ctx context.Context, cancel context.CancelFunc) error {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signalCh)
	select {
	case <-signalCh:
		log.Println("Get signal and start exit...")
		cancel()
		return errors.Wrapf(context.Canceled, "get exit signal")
	case <-ctx.Done():
		return nil
	}
}
