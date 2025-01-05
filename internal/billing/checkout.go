package billing

import (
	"fmt"

	"github.com/avptp/brain/internal/generated/data"
	"github.com/stripe/stripe-go/v81"
	session "github.com/stripe/stripe-go/v81/checkout/session"
)

func (b *StripeBiller) CreateCheckoutSession(p *data.Person) (string, error) {
	s, err := session.New(&stripe.CheckoutSessionParams{
		Customer: p.StripeID,
		Mode:     stripe.String("subscription"),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(b.cfg.StripePriceID),
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String(
			fmt.Sprintf("%s/billing/success", b.cfg.FrontUrl),
		),
		CancelURL: stripe.String(
			fmt.Sprintf("%s/billing/cancel", b.cfg.FrontUrl),
		),
	})

	if err != nil {
		return "", err
	}

	return s.URL, nil
}
