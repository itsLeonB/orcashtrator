package service_test

import (
	"context"
	"testing"

	"github.com/itsLeonB/orcashtrator/internal/domain/auth"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthClient struct {
	mock.Mock
}

func (m *MockAuthClient) Register(ctx context.Context, req auth.RegisterRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockAuthClient) Login(ctx context.Context, req auth.LoginRequest) (auth.LoginResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(auth.LoginResponse), args.Error(1)
}

func (m *MockAuthClient) VerifyToken(ctx context.Context, token string) (bool, map[string]any, error) {
	args := m.Called(ctx, token)
	return args.Bool(0), args.Get(1).(map[string]any), args.Error(2)
}

func TestAuthService_Register(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockClient := new(MockAuthClient)
		svc := service.NewAuthService(mockClient)
		
		req := dto.RegisterRequest{
			Email:                "test@example.com",
			Password:             "password123",
			PasswordConfirmation: "password123",
		}
		
		expectedAuthReq := auth.RegisterRequest{
			Email:                req.Email,
			Password:             req.Password,
			PasswordConfirmation: req.PasswordConfirmation,
		}
		
		mockClient.On("Register", mock.Anything, expectedAuthReq).Return(nil)
		
		err := svc.Register(context.Background(), req)
		
		assert.NoError(t, err)
		mockClient.AssertExpectations(t)
	})

	t.Run("client error", func(t *testing.T) {
		mockClient := new(MockAuthClient)
		svc := service.NewAuthService(mockClient)
		
		req := dto.RegisterRequest{
			Email:                "test@example.com",
			Password:             "password123",
			PasswordConfirmation: "password123",
		}
		
		mockClient.On("Register", mock.Anything, mock.Anything).Return(assert.AnError)
		
		err := svc.Register(context.Background(), req)
		
		assert.Error(t, err)
		mockClient.AssertExpectations(t)
	})
}

func TestAuthService_Login(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockClient := new(MockAuthClient)
		svc := service.NewAuthService(mockClient)
		
		req := dto.LoginRequest{
			Email:    "test@example.com",
			Password: "password123",
		}
		
		expectedAuthReq := auth.LoginRequest{
			Email:    req.Email,
			Password: req.Password,
		}
		
		authResponse := auth.LoginResponse{
			Type:  "Bearer",
			Token: "test-token",
		}
		
		mockClient.On("Login", mock.Anything, expectedAuthReq).Return(authResponse, nil)
		
		result, err := svc.Login(context.Background(), req)
		
		assert.NoError(t, err)
		assert.Equal(t, authResponse.Type, result.Type)
		assert.Equal(t, authResponse.Token, result.Token)
		mockClient.AssertExpectations(t)
	})

	t.Run("client error", func(t *testing.T) {
		mockClient := new(MockAuthClient)
		svc := service.NewAuthService(mockClient)
		
		req := dto.LoginRequest{
			Email:    "test@example.com",
			Password: "password123",
		}
		
		mockClient.On("Login", mock.Anything, mock.Anything).Return(auth.LoginResponse{}, assert.AnError)
		
		result, err := svc.Login(context.Background(), req)
		
		assert.Error(t, err)
		assert.Empty(t, result.Type)
		assert.Empty(t, result.Token)
		mockClient.AssertExpectations(t)
	})
}

func TestAuthService_VerifyToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockClient := new(MockAuthClient)
		svc := service.NewAuthService(mockClient)
		
		token := "test-token"
		claims := map[string]any{"user_id": "123"}
		
		mockClient.On("VerifyToken", mock.Anything, token).Return(true, claims, nil)
		
		valid, resultClaims, err := svc.VerifyToken(context.Background(), token)
		
		assert.NoError(t, err)
		assert.True(t, valid)
		assert.Equal(t, claims, resultClaims)
		mockClient.AssertExpectations(t)
	})

	t.Run("invalid token", func(t *testing.T) {
		mockClient := new(MockAuthClient)
		svc := service.NewAuthService(mockClient)
		
		token := "invalid-token"
		
		mockClient.On("VerifyToken", mock.Anything, token).Return(false, map[string]any{}, nil)
		
		valid, resultClaims, err := svc.VerifyToken(context.Background(), token)
		
		assert.NoError(t, err)
		assert.False(t, valid)
		assert.Empty(t, resultClaims)
		mockClient.AssertExpectations(t)
	})

	t.Run("client error", func(t *testing.T) {
		mockClient := new(MockAuthClient)
		svc := service.NewAuthService(mockClient)
		
		token := "test-token"
		
		mockClient.On("VerifyToken", mock.Anything, token).Return(false, map[string]any{}, assert.AnError)
		
		valid, resultClaims, err := svc.VerifyToken(context.Background(), token)
		
		assert.Error(t, err)
		assert.False(t, valid)
		assert.Empty(t, resultClaims)
		mockClient.AssertExpectations(t)
	})
}
