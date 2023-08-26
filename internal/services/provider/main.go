package provider

import (
	"github.com/avptp/brain/internal/services"
	"github.com/sarulabs/dingo/v4"
)

type Provider struct {
	dingo.BaseProvider
}

func (p *Provider) Load() error {
	return p.AddDefSlice([]dingo.Def{
		services.CaptchaDef,
		services.ConfigDef,
		services.DataDef,
		services.I18nDef,
		services.IPStrategyDef,
		services.LoggerDef,
		services.MessengerDef,
		services.SesDef,
	})
}
