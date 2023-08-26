package messaging_test

import (
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/messaging/templates"
	"github.com/stretchr/testify/mock"
)

type MockedMessenger struct {
	mock.Mock
}

func (m *MockedMessenger) Send(t templates.Template, p *data.Person) error {
	args := m.Called(t, p)

	return args.Error(0)
}
