package commands

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/avptp/brain/internal/generated/container"
	"github.com/avptp/brain/internal/jobs"
)

func Work(ctx context.Context, ctn *container.Container) error {
	log := ctn.GetLogger()

	// Initialize job scheduler
	scheduler := ctn.GetScheduler()

	for _, job := range jobs.All {
		_, err := scheduler.Add(
			job(ctx, ctn),
		)

		if err != nil {
			return err
		}
	}

	// Wait listening to stop signals
	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM,
	)

	log.Info("started waiting for jobs")

	<-signalChan

	log.Info("stopping...")

	return nil
}
