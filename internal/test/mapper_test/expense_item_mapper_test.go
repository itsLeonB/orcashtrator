package mapper_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/mapper"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewExpenseItemRequestToData(t *testing.T) {
	req := dto.NewExpenseItemRequest{
		UserProfileID:  uuid.New(),
		GroupExpenseID: uuid.New(),
		Name:           "Test Item",
		Amount:         decimal.NewFromInt(100),
		Quantity:       2,
	}

	result := mapper.NewExpenseItemRequestToData(req)

	assert.Equal(t, req.Name, result.Name)
	assert.Equal(t, req.Amount, result.Amount)
	assert.Equal(t, req.Quantity, result.Quantity)
}

func TestUpdateExpenseItemRequestToData(t *testing.T) {
	participants := []dto.ItemParticipantRequest{
		{
			ProfileID: uuid.New(),
			Share:     decimal.NewFromInt(50),
		},
		{
			ProfileID: uuid.New(),
			Share:     decimal.NewFromInt(50),
		},
	}

	req := dto.UpdateExpenseItemRequest{
		UserProfileID:  uuid.New(),
		ID:             uuid.New(),
		GroupExpenseID: uuid.New(),
		Name:           "Updated Item",
		Amount:         decimal.NewFromInt(200),
		Quantity:       3,
		Participants:   participants,
	}

	result := mapper.UpdateExpenseItemRequestToData(req)

	assert.Equal(t, req.Name, result.Name)
	assert.Equal(t, req.Amount, result.Amount)
	assert.Equal(t, req.Quantity, result.Quantity)
	assert.Len(t, result.Participants, 2)
	assert.Equal(t, participants[0].ProfileID, result.Participants[0].ProfileID)
	assert.Equal(t, participants[0].Share, result.Participants[0].Share)
	assert.Equal(t, participants[1].ProfileID, result.Participants[1].ProfileID)
	assert.Equal(t, participants[1].Share, result.Participants[1].Share)
}
