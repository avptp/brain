package services

import (
	"net"

	"github.com/avptp/brain/internal/config"
	"github.com/realclientip/realclientip-go"
	"github.com/sarulabs/dingo/v4"
)

const IPStrategy = "ipStrategy"

var IPStrategyDef = dingo.Def{
	Name: IPStrategy,
	Build: func(cfg *config.Config) (realclientip.Strategy, error) {
		// Parse trusted ranges from configuration
		var ranges []net.IPNet

		for _, str := range cfg.HttpTrustedProxies {
			_, addr, err := net.ParseCIDR(str)

			if err != nil {
				return nil, err
			}

			ranges = append(ranges, *addr)
		}

		// Create strategy
		strategy, err := realclientip.NewRightmostTrustedRangeStrategy("x-forwarded-for", ranges)

		if err != nil {
			return nil, err
		}

		return strategy, nil
	},
}
