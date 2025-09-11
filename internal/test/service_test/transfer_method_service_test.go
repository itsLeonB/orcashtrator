package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/domain/debt"
	"github.com/itsLeonB/orcashtrator/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransferMethodClient struct {
	mock.Mock
}

func (m *MockTransferMethodClient) GetAll(ctx context.Context) ([]debt.TransferMethod, error) {
	args := m.Called(ctx)
	return args.Get(0).([]debt.TransferMethod), args.Error(1)
}

func TestTransferMethodService_GetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockClient := new(MockTransferMethodClient)
		svc := service.NewTransferMethodService(mockClient)
		
		now := time.Now()
		transferMethods := []debt.TransferMethod{
			{
				ID:        uuid.New(),
				Name:      "bank_transfer",
				Display:   "Bank Transfer",
				CreatedAt: now,
				UpdatedAt: now,
			},
			{
				ID:        uuid.New(),
				Name:      "cash",
				Display:   "Cash",
				CreatedAt: now,
				UpdatedAt: now,
			},
		}
		
		mockClient.On("GetAll", mock.Anything).Return(transferMethods, nil)
		
		result, err := svc.GetAll(context.Background())
		
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, transferMethods[0].ID, result[0].ID)
		assert.Equal(t, transferMethods[0].Name, result[0].Name)
		assert.Equal(t, transferMethods[0].Display, result[0].Display)
		assert.Equal(t, transferMethods[1].ID, result[1].ID)
		assert.Equal(t, transferMethods[1].Name, result[1].Name)
		assert.Equal(t, transferMethods[1].Display, result[1].Display)
		
		mockClient.AssertExpectations(t)
	})

	t.Run("client error", func(t *testing.T) {
		mockClient := new(MockTransferMethodClient)
		svc := service.NewTransferMethodService(mockClient)
		
		mockClient.On("GetAll", mock.Anything).Return([]debt.TransferMethod{}, assert.AnError)
		
		result, err := svc.GetAll(context.Background())
		
		assert.Error(t, err)
		assert.Nil(t, result)
		
		mockClient.AssertExpectations(t)
	})
}
