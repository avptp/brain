package billing

import (
	"context"

	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/generated/data/privacy"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/customer"
)

// data.Client is not a struct property because function needs to receive a transactional one
func (b *StripeBiller) PreparePerson(ctx context.Context, d *data.Client, p *data.Person) error {
	if p.StripeID != nil {
		return nil
	}

	c, err := customer.New(
		customerParams(p),
	)

	if err != nil {
		return err
	}

	allowCtx := privacy.DecisionContext(ctx, privacy.Allow)

	_, err = d.Person.
		UpdateOneID(p.ID).
		SetStripeID(c.ID).
		Save(allowCtx)

	return err
}

func (b *StripeBiller) SyncPerson(p *data.Person) error {
	if p.StripeID == nil {
		return nil
	}

	_, err := customer.Update(
		*p.StripeID,
		customerParams(p),
	)

	return err
}

func customerParams(p *data.Person) *stripe.CustomerParams {
	return &stripe.CustomerParams{
		Email: &p.Email,
		Phone: p.Phone,
		TaxIDData: []*stripe.CustomerTaxIDDataParams{
			{
				Type:  stripe.String("es_cif"),
				Value: &p.TaxID,
			},
		},
		Name: stripe.String(p.FullName()),
		PreferredLocales: []*string{
			&p.Language,
		},
		Address: &stripe.AddressParams{
			Line1:      p.Address,
			PostalCode: p.PostalCode,
			City:       p.City,
			Country:    p.Country,
		},
	}
}
