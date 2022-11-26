package services

import (
	"github.com/sarulabs/dingo/v4"
	"go.uber.org/zap"
)

const Logger = "logger"

var LoggerDef = dingo.Def{
	Name: Logger,
	Build: func() (*zap.SugaredLogger, error) {
		log, err := zap.NewProduction()

		if err != nil {
			return nil, err
		}

		return log.Sugar(), nil
	},
	Close: func(log *zap.SugaredLogger) error {
		// intentionally ignoring error here, see https://github.com/uber-go/zap/issues/328
		_ = log.Sync()

		return nil
	},
}
