package billing

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/avptp/brain/internal/config"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/stripe/stripe-go/v78"
)

type Biller interface {
	// data.Client is not a struct property because function needs to receive a transactional one
	PreparePerson(context.Context, *data.Client, *data.Person) error
	SyncPerson(*data.Person) error
	CreateCheckoutSession(*data.Person) (string, error)
	CreatePortalSession(*data.Person) (string, error)
	WebhookHandler() http.Handler
}

type StripeBiller struct {
	cfg  *config.Config
	log  *slog.Logger
	data *data.Client
}

func NewBiller(cfg *config.Config, log *slog.Logger, data *data.Client) Biller {
	stripe.Key = cfg.StripeApiSecret

	return &StripeBiller{
		cfg,
		log,
		data,
	}
}
