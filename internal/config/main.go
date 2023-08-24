package config

type Config struct {
	Environment string `env:"APP_ENV" envDefault:"production"`
	Debug       bool   `env:"APP_DEBUG" envDefault:"false"`

	HttpPort           string   `env:"HTTP_PORT" envDefault:"8000"`
	HttpTrustedProxies []string `env:"HTTP_TRUSTED_PROXIES"`

	CockroachDBHost     string `env:"COCKROACHDB_HOST"`
	CockroachDBPort     string `env:"COCKROACHDB_PORT" envDefault:"26257"`
	CockroachDBUser     string `env:"COCKROACHDB_USER"`
	CockroachDBPassword string `env:"COCKROACHDB_PASSWORD"`
	CockroachDBDatabase string `env:"COCKROACHDB_DATABASE"`
	CockroachDBTLSMode  string `env:"COCKROACHDB_TLS_MODE" envDefault:"require"`
	CockroachDBTLSCA    string `env:"COCKROACHDB_TLS_CA"`

	HcaptchaSecret string `env:"HCAPTCHA_SECRET"`
}
