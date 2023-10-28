package billing_test

import (
	"context"
	"net/http"

	"github.com/avptp/brain/internal/generated/data"
	"github.com/stretchr/testify/mock"
)

type MockedBiller struct {
	mock.Mock
}

func (m *MockedBiller) PreparePerson(ctx context.Context, d *data.Client, p *data.Person) error {
	args := m.Called(ctx, d, p)

	return args.Error(0)
}

func (m *MockedBiller) SyncPerson(p *data.Person) error {
	args := m.Called(p)

	return args.Error(0)
}

func (m *MockedBiller) CreateCheckoutSession(p *data.Person) (string, error) {
	args := m.Called(p)

	return args.String(0), args.Error(1)
}

func (m *MockedBiller) CreatePortalSession(p *data.Person) (string, error) {
	args := m.Called(p)

	return args.String(0), args.Error(1)
}

func (m *MockedBiller) WebhookHandler() http.Handler {
	args := m.Called()

	return args.Get(0).(http.Handler)
}
