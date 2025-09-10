package mapper_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/domain/debt"
	"github.com/itsLeonB/orcashtrator/internal/dto"
	"github.com/itsLeonB/orcashtrator/internal/mapper"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestDebtTransactionToResponse(t *testing.T) {
	now := time.Now()
	dt := debt.Transaction{
		ID:             uuid.New(),
		ProfileID:      uuid.New(),
		Type:           appconstant.Lend,
		Action:         appconstant.LendAction,
		Amount:         decimal.NewFromInt(100),
		TransferMethod: "bank_transfer",
		Description:    "Test transaction",
		CreatedAt:      now,
		UpdatedAt:      now,
		DeletedAt:      time.Time{},
	}

	result := mapper.DebtTransactionToResponse(dt)

	assert.Equal(t, dt.ID, result.ID)
	assert.Equal(t, dt.ProfileID, result.ProfileID)
	assert.Equal(t, dt.Type, result.Type)
	assert.Equal(t, dt.Action, result.Action)
	assert.Equal(t, dt.Amount, result.Amount)
	assert.Equal(t, dt.TransferMethod, result.TransferMethod)
	assert.Equal(t, dt.Description, result.Description)
	assert.Equal(t, dt.CreatedAt, result.CreatedAt)
	assert.Equal(t, dt.UpdatedAt, result.UpdatedAt)
	assert.Equal(t, dt.DeletedAt, result.DeletedAt)
}

func TestMapToFriendBalanceSummary(t *testing.T) {
	userProfileID := uuid.New()

	t.Run("empty transactions", func(t *testing.T) {
		result := mapper.MapToFriendBalanceSummary(userProfileID, []dto.DebtTransactionResponse{})

		assert.True(t, result.TotalOwedToYou.IsZero())
		assert.True(t, result.TotalYouOwe.IsZero())
		assert.True(t, result.NetBalance.IsZero())
		assert.Equal(t, "IDR", result.CurrencyCode)
	})

	t.Run("lend transactions", func(t *testing.T) {
		transactions := []dto.DebtTransactionResponse{
			{
				Type:   appconstant.Lend,
				Action: appconstant.LendAction,
				Amount: decimal.NewFromInt(100),
			},
			{
				Type:   appconstant.Lend,
				Action: appconstant.BorrowAction,
				Amount: decimal.NewFromInt(50),
			},
		}

		result := mapper.MapToFriendBalanceSummary(userProfileID, transactions)

		assert.Equal(t, decimal.NewFromInt(100), result.TotalOwedToYou)
		assert.Equal(t, decimal.NewFromInt(50), result.TotalYouOwe)
		assert.Equal(t, decimal.NewFromInt(50), result.NetBalance)
	})

	t.Run("repay transactions", func(t *testing.T) {
		transactions := []dto.DebtTransactionResponse{
			{
				Type:   appconstant.Lend,
				Action: appconstant.LendAction,
				Amount: decimal.NewFromInt(100),
			},
			{
				Type:   appconstant.Repay,
				Action: appconstant.ReceiveAction,
				Amount: decimal.NewFromInt(30),
			},
			{
				Type:   appconstant.Repay,
				Action: appconstant.ReturnAction,
				Amount: decimal.NewFromInt(20),
			},
		}

		result := mapper.MapToFriendBalanceSummary(userProfileID, transactions)

		assert.Equal(t, decimal.NewFromInt(70), result.TotalOwedToYou)
		assert.Equal(t, decimal.NewFromInt(-20), result.TotalYouOwe)
		assert.Equal(t, decimal.NewFromInt(90), result.NetBalance)
	})
}
