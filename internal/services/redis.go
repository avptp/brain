package services

import (
	"github.com/avptp/brain/internal/config"
	"github.com/redis/go-redis/v9"
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
)

const Redis = "redis"

var RedisDef = dingo.Def{
	Name:  Redis,
	Scope: di.App,
	Build: func(cfg *config.Config) (*redis.Client, error) {
		return redis.NewClient(&redis.Options{
			Addr:     cfg.RedisAddress,
			Password: cfg.RedisPassword,
			DB:       cfg.RedisDatabase,
		}), nil
	},
}
