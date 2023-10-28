package resolvers_test

import (
	"github.com/avptp/brain/internal/api/reporting"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/mock"
)

func (t *TestSuite) TestBilling() {
	const createCheckoutSessionMutation = `
		mutation() {
			createBillingCheckoutSession() {
				checkoutSessionUrl
			}
		}
	`

	type createCheckoutSession struct {
		CreateBillingCheckoutSession struct {
			CheckoutSessionURL string
		}
	}

	t.Run("create_checkout_session", func() {
		authenticated, _, _, _, _ := t.authenticate()

		url := gofakeit.URL()

		t.biller.On(
			"PreparePerson",
			mock.Anything,
			mock.IsType(&data.Client{}),
			mock.IsType(&data.Person{}),
		).Return(nil).Once()

		t.biller.On(
			"CreateCheckoutSession",
			mock.IsType(&data.Person{}),
		).Return(
			url,
			nil,
		).Once()

		var response createCheckoutSession
		err := t.api.Post(
			createCheckoutSessionMutation,
			&response,
			authenticated,
		)

		t.NoError(err)
		t.biller.AssertExpectations(t.T())

		t.Equal(url, response.CreateBillingCheckoutSession.CheckoutSessionURL)
	})

	//TODO: test create_checkout_with_person_who_cannot_subscribe

	t.Run("create_checkout_session_without_authentication", func() {
		var response createCheckoutSession
		err := t.api.Post(
			createCheckoutSessionMutation,
			&response,
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrUnauthenticated.Message)
	})

	const createPortalSessionMutation = `
		mutation() {
			createBillingPortalSession() {
				portalSessionUrl
			}
		}
	`

	type createPortalSession struct {
		CreateBillingPortalSession struct {
			PortalSessionURL string
		}
	}

	t.Run("create_portal_session", func() {
		authenticated, _, _, _, _ := t.authenticate()

		url := gofakeit.URL()

		t.biller.On(
			"PreparePerson",
			mock.Anything,
			mock.IsType(&data.Client{}),
			mock.IsType(&data.Person{}),
		).Return(nil).Once()

		t.biller.On(
			"CreatePortalSession",
			mock.IsType(&data.Person{}),
		).Return(
			url,
			nil,
		).Once()

		var response createPortalSession
		err := t.api.Post(
			createPortalSessionMutation,
			&response,
			authenticated,
		)

		t.NoError(err)
		t.biller.AssertExpectations(t.T())

		t.Equal(url, response.CreateBillingPortalSession.PortalSessionURL)
	})

	t.Run("create_portal_session_without_authentication", func() {
		var response createPortalSession
		err := t.api.Post(
			createPortalSessionMutation,
			&response,
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrUnauthenticated.Message)
	})
}
