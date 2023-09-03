package resolvers

import (
	"github.com/avptp/brain/internal/auth"
	"github.com/avptp/brain/internal/config"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/messaging"
	"github.com/go-redis/redis_rate/v10"
)

type Resolver struct {
	captcha   auth.Captcha
	cfg       *config.Config
	data      *data.Client
	limiter   *redis_rate.Limiter
	messenger messaging.Messenger
}

func NewResolver(
	captcha auth.Captcha,
	cfg *config.Config,
	data *data.Client,
	limiter *redis_rate.Limiter,
	messenger messaging.Messenger,
) *Resolver {
	return &Resolver{
		captcha,
		cfg,
		data,
		limiter,
		messenger,
	}
}
