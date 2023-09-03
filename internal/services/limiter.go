package services

import (
	"github.com/avptp/brain/internal/config"
	"github.com/go-redis/redis_rate/v10"
	libredis "github.com/redis/go-redis/v9"
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
)

const Limiter = "limiter"

var LimiterDef = dingo.Def{
	Name:  Limiter,
	Scope: di.App,
	Build: func(cfg *config.Config, redis *libredis.Client) (*redis_rate.Limiter, error) {
		limiter := redis_rate.NewLimiter(redis)

		return limiter, nil
	},
}
