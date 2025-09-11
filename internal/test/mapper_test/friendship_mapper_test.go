package mapper_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/domain/friendship"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/mapper"
	"github.com/stretchr/testify/assert"
)

func TestFriendshipToResponse(t *testing.T) {
	now := time.Now()
	fs := friendship.Friendship{
		ID:          uuid.New(),
		Type:        appconstant.Real,
		ProfileID:   uuid.New(),
		ProfileName: "Test Friend",
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   time.Time{},
	}

	result := mapper.FriendshipToResponse(fs)

	assert.Equal(t, fs.ID, result.ID)
	assert.Equal(t, fs.Type, result.Type)
	assert.Equal(t, fs.ProfileID, result.ProfileID)
	assert.Equal(t, fs.ProfileName, result.ProfileName)
	assert.Equal(t, fs.CreatedAt, result.CreatedAt)
	assert.Equal(t, fs.UpdatedAt, result.UpdatedAt)
	assert.Equal(t, fs.DeletedAt, result.DeletedAt)
}

func TestMapToFriendDetailsResponse(t *testing.T) {
	userProfileID := uuid.New()
	now := time.Now()
	
	friendDetails := friendship.FriendDetails{
		ID:        uuid.New(),
		ProfileID: uuid.New(),
		Name:      "Test Friend",
		Type:      appconstant.Anonymous,
		Email:     "test@example.com",
		Phone:     "+1234567890",
		Avatar:    "avatar.jpg",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: time.Time{},
	}

	debtTransactions := []dto.DebtTransactionResponse{
		{
			ID:        uuid.New(),
			ProfileID: userProfileID,
			Type:      appconstant.Lend,
			Action:    appconstant.LendAction,
		},
	}

	result, err := mapper.MapToFriendDetailsResponse(userProfileID, friendDetails, debtTransactions)

	assert.NoError(t, err)
	assert.Equal(t, friendDetails.ID, result.Friend.ID)
	assert.Equal(t, friendDetails.ProfileID, result.Friend.ProfileID)
	assert.Equal(t, friendDetails.Name, result.Friend.Name)
	assert.Equal(t, friendDetails.Type, result.Friend.Type)
	assert.Equal(t, friendDetails.Email, result.Friend.Email)
	assert.Equal(t, friendDetails.Phone, result.Friend.Phone)
	assert.Equal(t, friendDetails.Avatar, result.Friend.Avatar)
	assert.Equal(t, friendDetails.CreatedAt, result.Friend.CreatedAt)
	assert.Equal(t, friendDetails.UpdatedAt, result.Friend.UpdatedAt)
	assert.Equal(t, friendDetails.DeletedAt, result.Friend.DeletedAt)
	assert.Equal(t, debtTransactions, result.Transactions)
	assert.NotNil(t, result.Balance)
}

func TestMapToFriendDetailsResponse_NilTransactions(t *testing.T) {
	userProfileID := uuid.New()
	now := time.Now()
	
	friendDetails := friendship.FriendDetails{
		ID:        uuid.New(),
		ProfileID: uuid.New(),
		Name:      "Test Friend",
		Type:      appconstant.Anonymous,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := mapper.MapToFriendDetailsResponse(userProfileID, friendDetails, nil)

	assert.NoError(t, err)
	assert.Empty(t, result.Transactions)
	assert.NotNil(t, result.Balance)
}
