package expenseitem

import "github.com/google/uuid"

type AddRequest struct {
	ProfileID      uuid.UUID `validate:"required"`
	GroupExpenseID uuid.UUID `validate:"required"`
	ExpenseItemData
}

type GetDetailsRequest struct {
	ID             uuid.UUID `validate:"required"`
	GroupExpenseID uuid.UUID `validate:"required"`
}

type UpdateRequest struct {
	ProfileID      uuid.UUID `validate:"required"`
	ID             uuid.UUID `validate:"required"`
	GroupExpenseID uuid.UUID `validate:"required"`
	ExpenseItemData
}

type RemoveRequest struct {
	ProfileID      uuid.UUID `validate:"required"`
	ID             uuid.UUID `validate:"required"`
	GroupExpenseID uuid.UUID `validate:"required"`
}
