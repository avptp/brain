package billing

import (
	"fmt"

	"github.com/avptp/brain/internal/generated/data"
	"github.com/stripe/stripe-go/v78"
	session "github.com/stripe/stripe-go/v78/billingportal/session"
)

func (b *StripeBiller) CreatePortalSession(p *data.Person) (string, error) {
	s, err := session.New(&stripe.BillingPortalSessionParams{
		Customer: p.StripeID,
		ReturnURL: stripe.String(
			fmt.Sprintf("%s/billing", b.cfg.FrontUrl),
		),
	})

	if err != nil {
		return "", err
	}

	return s.URL, nil
}
