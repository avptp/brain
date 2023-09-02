package commands

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/avptp/brain/internal/generated/container"
	"github.com/avptp/brain/internal/transport"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Serve(ctx context.Context, ctn *container.Container) error {
	log := ctn.GetLogger()
	cfg := ctn.GetConfig()

	// Initialize HTTP server
	srv := &http.Server{
		Addr: ":" + cfg.HttpPort,
		Handler: h2c.NewHandler(
			transport.Mux(ctn),
			&http2.Server{},
		),
	}

	// Start HTTP server and keep listening to stop signals
	errChan := make(chan error, 1)
	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM,
	)

	go func() {
		err := srv.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	log.Info(
		"started listening HTTP requests",
		"address", srv.Addr,
	)

	defer func() {
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			errChan <- err
		}

		log.Info("stopped")
	}()

	select {
	case err := <-errChan:
		return err
	case <-signalChan:
		log.Info("stopping...")
	}

	return nil
}
