package services

import (
	"github.com/avptp/brain/internal/api/auth"
	"github.com/avptp/brain/internal/api/resolvers"
	"github.com/avptp/brain/internal/billing"
	"github.com/avptp/brain/internal/config"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/messaging"
	"github.com/go-redis/redis_rate/v10"
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
)

const Resolver = "resolver"

var ResolverDef = dingo.Def{
	Name:  Resolver,
	Scope: di.App,
	Build: func(
		biller billing.Biller,
		captcha auth.Captcha,
		cfg *config.Config,
		data *data.Client,
		limiter *redis_rate.Limiter,
		messenger messaging.Messenger,
	) (*resolvers.Resolver, error) {
		return resolvers.NewResolver(
			biller,
			captcha,
			cfg,
			data,
			limiter,
			messenger,
		), nil
	},
}
