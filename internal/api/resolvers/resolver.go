package resolvers

import (
	"github.com/avptp/brain/internal/auth"
	"github.com/avptp/brain/internal/config"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/messaging"
)

type Resolver struct {
	cfg       *config.Config
	captcha   auth.Captcha
	data      *data.Client
	messenger messaging.Messenger
}

func NewResolver(cfg *config.Config, captcha auth.Captcha, data *data.Client, messenger messaging.Messenger) *Resolver {
	return &Resolver{
		cfg,
		captcha,
		data,
		messenger,
	}
}
