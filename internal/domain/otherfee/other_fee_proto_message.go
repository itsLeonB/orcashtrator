package otherfee

import "github.com/google/uuid"

type AddRequest struct {
	ProfileID      uuid.UUID `validate:"required"`
	GroupExpenseID uuid.UUID `validate:"required"`
	OtherFeeData
}

type UpdateRequest struct {
	ProfileID      uuid.UUID `validate:"required"`
	ID             uuid.UUID `validate:"required"`
	GroupExpenseID uuid.UUID `validate:"required"`
	OtherFeeData
}

type RemoveRequest struct {
	ProfileID      uuid.UUID `validate:"required"`
	ID             uuid.UUID `validate:"required"`
	GroupExpenseID uuid.UUID `validate:"required"`
}
