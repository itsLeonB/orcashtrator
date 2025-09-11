package debt

import (
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID             uuid.UUID
	ProfileID      uuid.UUID
	Type           appconstant.DebtTransactionType
	Action         appconstant.Action
	Amount         decimal.Decimal
	TransferMethod string
	Description    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

type GroupExpenseData struct {
	ID               uuid.UUID `validate:"required"`
	PayerProfileID   uuid.UUID `validate:"required"`
	CreatorProfileID uuid.UUID `validate:"required"`
	Description      string
	Participants     []ExpenseParticipantData `validate:"required,min=1,dive"`
}

type ExpenseParticipantData struct {
	ProfileID   uuid.UUID       `validate:"required"`
	ShareAmount decimal.Decimal `validate:"required"`
}
