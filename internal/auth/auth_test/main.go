package auth_test

import (
	"github.com/stretchr/testify/mock"
)

type MockedCaptcha struct {
	mock.Mock
}

func (m *MockedCaptcha) Verify(token string) bool {
	args := m.Called(token)

	return args.Bool(0)
}
