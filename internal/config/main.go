package config

import "time"

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

	RedisAddress  string `env:"REDIS_ADDRESS"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDatabase int    `env:"REDIS_DATABASE"`

	HcaptchaSecret string `env:"HCAPTCHA_SECRET"`

	MailSource  string `env:"MAIL_SOURCE"`
	MailReplyTo string `env:"MAIL_REPLY_TO"`

	AwsRegion    string `env:"AWS_REGION"`
	AwsKeyId     string `env:"AWS_KEY_ID"`
	AwsKeySecret string `env:"AWS_KEY_SECRET"`

	AuthenticationRateLimit        int           `env:"AUTHENTICATION_RATE_LIMIT" envDefault:"5"` // per email and hour
	AuthorizationMaxAge            time.Duration `env:"AUTHORIZATION_MAX_AGE" envDefault:"24h"`
	AuthorizationEmailRateLimit    int           `env:"AUTHORIZATION_EMAIL_RATE_LIMIT" envDefault:"2"`    // per email and hour
	AuthorizationPasswordRateLimit int           `env:"AUTHORIZATION_PASSWORD_RATE_LIMIT" envDefault:"2"` // per email and hour

	FrontUrl                       string `env:"FRONT_URL"`
	FrontEmailAuthorizationPath    string `env:"FRONT_EMAIL_AUTHORIZATION_PATH"`
	FrontPasswordAuthorizationPath string `env:"FRONT_PASSWORD_AUTHORIZATION_PATH"`

	OrgName string `env:"ORG_NAME"`
	OrgLogo string `env:"ORG_LOGO"`

	StripeApiSecret      string `env:"STRIPE_API_SECRET"`
	StripeEndpointSecret string `env:"STRIPE_ENDPOINT_SECRET"`
	StripePriceID        string `env:"STRIPE_PRICE_ID"`
}
