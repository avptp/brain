package services

import (
	"log/slog"

	"github.com/avptp/brain/internal/billing"
	"github.com/avptp/brain/internal/config"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
)

const Biller = "biller"

var BillerDef = dingo.Def{
	Name:  Biller,
	Scope: di.App,
	Build: func(cfg *config.Config, log *slog.Logger, data *data.Client) (billing.Biller, error) {
		return billing.NewBiller(
			cfg,
			log,
			data,
		), nil
	},
}
