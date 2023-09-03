package services

import (
	"github.com/avptp/brain/internal/api/resolvers"
	"github.com/avptp/brain/internal/auth"
	"github.com/avptp/brain/internal/config"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/messaging"
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
)

const Resolver = "resolver"

var ResolverDef = dingo.Def{
	Name:  Resolver,
	Scope: di.App,
	Build: func(cfg *config.Config, captcha auth.Captcha, data *data.Client, messenger messaging.Messenger) (*resolvers.Resolver, error) {
		return resolvers.NewResolver(
			cfg,
			captcha,
			data,
			messenger,
		), nil
	},
}
