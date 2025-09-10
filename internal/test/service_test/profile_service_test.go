package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/domain/profile"
	"github.com/itsLeonB/orcashtrator/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProfileClient struct {
	mock.Mock
}

func (m *MockProfileClient) Get(ctx context.Context, id uuid.UUID) (profile.Profile, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(profile.Profile), args.Error(1)
}

func (m *MockProfileClient) Create(ctx context.Context, req profile.CreateRequest) (profile.Profile, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(profile.Profile), args.Error(1)
}

func (m *MockProfileClient) GetByIDs(ctx context.Context, ids uuid.UUIDs) ([]profile.Profile, error) {
	args := m.Called(ctx, ids)
	return args.Get(0).([]profile.Profile), args.Error(1)
}

func TestProfileService_GetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockClient := new(MockProfileClient)
		svc := service.NewProfileService(mockClient)
		
		id := uuid.New()
		now := time.Now()
		profileData := profile.Profile{
			ID:          id,
			UserID:      uuid.New(),
			Name:        "Test User",
			IsAnonymous: false,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		
		mockClient.On("Get", mock.Anything, id).Return(profileData, nil)
		
		result, err := svc.GetByID(context.Background(), id)
		
		assert.NoError(t, err)
		assert.Equal(t, profileData.ID, result.ID)
		assert.Equal(t, profileData.UserID, result.UserID)
		assert.Equal(t, profileData.Name, result.Name)
		assert.Equal(t, profileData.IsAnonymous, result.IsAnonymous)
		
		mockClient.AssertExpectations(t)
	})

	t.Run("client error", func(t *testing.T) {
		mockClient := new(MockProfileClient)
		svc := service.NewProfileService(mockClient)
		
		id := uuid.New()
		mockClient.On("Get", mock.Anything, id).Return(profile.Profile{}, assert.AnError)
		
		result, err := svc.GetByID(context.Background(), id)
		
		assert.Error(t, err)
		assert.Empty(t, result.ID)
		
		mockClient.AssertExpectations(t)
	})
}

func TestProfileService_GetNames(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockClient := new(MockProfileClient)
		svc := service.NewProfileService(mockClient)
		
		id1, id2 := uuid.New(), uuid.New()
		ids := uuid.UUIDs{id1, id2}
		
		profiles := []profile.Profile{
			{ID: id1, Name: "User 1"},
			{ID: id2, Name: "User 2"},
		}
		
		mockClient.On("GetByIDs", mock.Anything, ids).Return(profiles, nil)
		
		result, err := svc.GetNames(context.Background(), ids)
		
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "User 1", result[id1])
		assert.Equal(t, "User 2", result[id2])
		
		mockClient.AssertExpectations(t)
	})

	t.Run("client error", func(t *testing.T) {
		mockClient := new(MockProfileClient)
		svc := service.NewProfileService(mockClient)
		
		ids := uuid.UUIDs{uuid.New()}
		mockClient.On("GetByIDs", mock.Anything, ids).Return([]profile.Profile{}, assert.AnError)
		
		result, err := svc.GetNames(context.Background(), ids)
		
		assert.Error(t, err)
		assert.Nil(t, result)
		
		mockClient.AssertExpectations(t)
	})
}

func TestProfileService_GetByIDs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockClient := new(MockProfileClient)
		svc := service.NewProfileService(mockClient)
		
		id1, id2 := uuid.New(), uuid.New()
		ids := uuid.UUIDs{id1, id2}
		now := time.Now()
		
		profiles := []profile.Profile{
			{ID: id1, Name: "User 1", CreatedAt: now, UpdatedAt: now},
			{ID: id2, Name: "User 2", CreatedAt: now, UpdatedAt: now},
		}
		
		mockClient.On("GetByIDs", mock.Anything, ids).Return(profiles, nil)
		
		result, err := svc.GetByIDs(context.Background(), ids)
		
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "User 1", result[id1].Name)
		assert.Equal(t, "User 2", result[id2].Name)
		
		mockClient.AssertExpectations(t)
	})

	t.Run("client error", func(t *testing.T) {
		mockClient := new(MockProfileClient)
		svc := service.NewProfileService(mockClient)
		
		ids := uuid.UUIDs{uuid.New()}
		mockClient.On("GetByIDs", mock.Anything, ids).Return([]profile.Profile{}, assert.AnError)
		
		result, err := svc.GetByIDs(context.Background(), ids)
		
		assert.Error(t, err)
		assert.Nil(t, result)
		
		mockClient.AssertExpectations(t)
	})
}
