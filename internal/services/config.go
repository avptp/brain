package services

import (
	"github.com/avptp/brain/internal/config"
	"github.com/caarlos0/env/v9"
	"github.com/sarulabs/dingo/v4"
)

const Config = "config"

var ConfigDef = dingo.Def{
	Name: Config,
	Build: func() (*config.Config, error) {
		cfg := &config.Config{}

		if err := env.Parse(cfg); err != nil {
			return nil, err
		}

		return cfg, nil
	},
}
