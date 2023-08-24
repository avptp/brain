package services

import (
	"log/slog"
	"os"

	"github.com/avptp/brain/internal/config"
	"github.com/sarulabs/dingo/v4"
)

const Logger = "logger"

var LoggerDef = dingo.Def{
	Name: Logger,
	Build: func(cfg *config.Config) (*slog.Logger, error) {
		level := slog.LevelInfo

		if cfg.Debug {
			level = slog.LevelDebug
		}

		log := slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: level,
			}),
		)

		return log, nil
	},
}
