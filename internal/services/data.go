package services

import (
	"fmt"
	"net/url"

	"entgo.io/ent/dialect"
	"github.com/avptp/brain/internal/config"
	"github.com/avptp/brain/internal/generated/data"
	_ "github.com/lib/pq"
	"github.com/sarulabs/dingo/v4"
)

const Data = "data"

var DataDef = dingo.Def{
	Name: Data,
	Build: func(cfg *config.Config) (*data.Client, error) {
		query := url.Values{
			"sslmode": {cfg.CockroachDBTLSMode},
		}

		if cfg.CockroachDBTLSMode != "disable" {
			query.Add("sslrootcert", cfg.CockroachDBTLSCA)
		}

		url := url.URL{
			Scheme: "postgresql",
			User: url.UserPassword(
				cfg.CockroachDBUser,
				cfg.CockroachDBPassword,
			),
			Host: fmt.Sprintf("%s:%s",
				cfg.CockroachDBHost,
				cfg.CockroachDBPort,
			),
			Path:     cfg.CockroachDBDatabase,
			RawQuery: query.Encode(),
		}

		client, err := data.Open(dialect.Postgres, url.String())

		if err != nil {
			return nil, err
		}

		return client, nil
	},
	Close: func(d *data.Client) error {
		return d.Close()
	},
}
