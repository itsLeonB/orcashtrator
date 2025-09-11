package expenseitem

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/domain"
	"github.com/shopspring/decimal"
)

type ExpenseItem struct {
	GroupExpenseID uuid.UUID
	ExpenseItemData
	domain.AuditMetadata
}

func (ei ExpenseItem) ProfileIDs() []uuid.UUID {
	return ezutil.MapSlice(ei.Participants, func(ip ItemParticipant) uuid.UUID { return ip.ProfileID })
}

type ExpenseItemData struct {
	Name         string            `validate:"required,min=3"`
	Amount       decimal.Decimal   `validate:"required"`
	Quantity     int               `validate:"required,min=1"`
	Participants []ItemParticipant `validate:"dive"`
}

type ItemParticipant struct {
	ProfileID uuid.UUID       `validate:"required"`
	Share     decimal.Decimal `validate:"required"`
}
