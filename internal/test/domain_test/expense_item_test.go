package domain_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/domain"
	"github.com/itsLeonB/orcashtrator/internal/domain/expenseitem"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestExpenseItem_ProfileIDs(t *testing.T) {
	profileID1 := uuid.New()
	profileID2 := uuid.New()
	
	expenseItem := expenseitem.ExpenseItem{
		GroupExpenseID: uuid.New(),
		ExpenseItemData: expenseitem.ExpenseItemData{
			Name:     "Test Item",
			Amount:   decimal.NewFromInt(100),
			Quantity: 1,
			Participants: []expenseitem.ItemParticipant{
				{
					ProfileID: profileID1,
					Share:     decimal.NewFromInt(50),
				},
				{
					ProfileID: profileID2,
					Share:     decimal.NewFromInt(50),
				},
			},
		},
		AuditMetadata: domain.AuditMetadata{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	
	profileIDs := expenseItem.ProfileIDs()
	
	assert.Len(t, profileIDs, 2)
	assert.Contains(t, profileIDs, profileID1)
	assert.Contains(t, profileIDs, profileID2)
}

func TestExpenseItem_ProfileIDs_Empty(t *testing.T) {
	expenseItem := expenseitem.ExpenseItem{
		GroupExpenseID: uuid.New(),
		ExpenseItemData: expenseitem.ExpenseItemData{
			Name:         "Test Item",
			Amount:       decimal.NewFromInt(100),
			Quantity:     1,
			Participants: []expenseitem.ItemParticipant{},
		},
		AuditMetadata: domain.AuditMetadata{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	
	profileIDs := expenseItem.ProfileIDs()
	
	assert.Empty(t, profileIDs)
}
