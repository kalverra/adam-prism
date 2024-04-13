package main_test

import (
	"github.com/go-resty/resty/v2"
	"github.com/test-go/testify/mock"
)

// MockRestyClient is a mock for the Resty client
type MockRestyClient struct {
	mock.Mock
	*resty.Client
}

// NewMockRestyClient creates a new mock instance
func NewMockRestyClient() *MockRestyClient {
	return &MockRestyClient{
		Client: resty.New(),
	}
}

// R is a mockable method for creating a request
func (m *MockRestyClient) R() *resty.Request {
	args := m.Called()
	return args.Get(0).(*resty.Request)
}
