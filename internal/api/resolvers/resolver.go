package resolvers

import (
	"github.com/avptp/brain/internal/auth"
	"github.com/avptp/brain/internal/billing"
	"github.com/avptp/brain/internal/config"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/messaging"
	"github.com/go-redis/redis_rate/v10"
)

type Resolver struct {
	biller    billing.Biller
	captcha   auth.Captcha
	cfg       *config.Config
	data      *data.Client
	limiter   *redis_rate.Limiter
	messenger messaging.Messenger
}

func NewResolver(
	biller billing.Biller,
	captcha auth.Captcha,
	cfg *config.Config,
	data *data.Client,
	limiter *redis_rate.Limiter,
	messenger messaging.Messenger,
) *Resolver {
	return &Resolver{
		biller,
		captcha,
		cfg,
		data,
		limiter,
		messenger,
	}
}
