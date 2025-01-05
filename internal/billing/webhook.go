package billing

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/avptp/brain/internal/generated/data/person"
	"github.com/avptp/brain/internal/generated/data/privacy"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/webhook"
)

func (b *StripeBiller) WebhookHandler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// Fail if wrong method
		if req.Method != "POST" {
			fail(res, http.StatusMethodNotAllowed)
			return
		}

		// Read body
		payload, err := io.ReadAll(req.Body)

		if err != nil {
			fail(res, http.StatusServiceUnavailable)
			return
		}

		// Verify signature
		event, err := webhook.ConstructEvent(
			payload,
			req.Header.Get("Stripe-Signature"),
			b.cfg.StripeEndpointSecret,
		)

		if err != nil {
			fail(res, http.StatusUnauthorized)
			return
		}

		// Handle event
		allowCtx := privacy.DecisionContext(context.Background(), privacy.Allow)

		switch event.Type {
		case "customer.subscription.created",
			"customer.subscription.deleted",
			"customer.subscription.paused",
			"customer.subscription.pending_update_applied",
			"customer.subscription.pending_update_expired",
			"customer.subscription.resumed",
			"customer.subscription.trial_will_end",
			"customer.subscription.updated":
			var subscription stripe.Subscription
			err := json.Unmarshal(event.Data.Raw, &subscription)

			if err != nil {
				fail(res, http.StatusBadRequest)
				return
			}

			_, err = b.data.Person.
				Update().
				Where(person.StripeIDEQ(subscription.Customer.ID)).
				SetSubscribed(subscription.Status == stripe.SubscriptionStatusActive).
				Save(allowCtx)

			if err != nil {
				fail(res, http.StatusInternalServerError)
				b.log.Error(
					err.Error(),
				)
				return
			}
		default:
			fail(res, http.StatusBadRequest)
			return
		}

		// Return success
		res.WriteHeader(http.StatusOK)
	})
}

func fail(res http.ResponseWriter, code int) {
	http.Error(
		res,
		http.StatusText(code),
		code,
	)
}
