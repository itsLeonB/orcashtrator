package delivery_test

import (
	"context"
	"testing"
	"time"

	"github.com/itsLeonB/orcashtrator/internal/config"
	"github.com/itsLeonB/orcashtrator/internal/delivery/http"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Register(ctx context.Context, request dto.RegisterRequest) error {
	args := m.Called(ctx, request)
	return args.Error(0)
}

func (m *MockAuthService) Login(ctx context.Context, request dto.InternalLoginRequest) (dto.LoginResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(dto.LoginResponse), args.Error(1)
}

func (m *MockAuthService) VerifyToken(ctx context.Context, token string) (bool, map[string]any, error) {
	args := m.Called(ctx, token)
	return args.Bool(0), args.Get(1).(map[string]any), args.Error(2)
}

func TestMiddlewaresIntegration(t *testing.T) {
	// Since provideMiddlewares is not exported, we test the integration
	// by ensuring the HTTP setup works correctly
	cfg := config.Config{
		App: config.App{
			Env:        "debug",
			Port:       "8080",
			Timeout:    time.Second * 10,
			ClientUrls: []string{"http://localhost:3000"},
		},
		ServiceClient: config.ServiceClient{
			BillsplittrHost: "http://localhost:8081",
			CocoonHost:      "http://localhost:8082",
			DrexHost:        "http://localhost:8083",
		},
	}

	assert.NotPanics(t, func() {
		server, err := http.Setup(cfg)
		assert.NotNil(t, server)
		assert.NoError(t, err)
	})
}
