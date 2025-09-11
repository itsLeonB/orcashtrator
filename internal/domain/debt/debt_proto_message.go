package debt

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/shopspring/decimal"
)

type RecordNewTransactionRequest struct {
	UserProfileID    uuid.UUID          `validate:"required"`
	FriendProfileID  uuid.UUID          `validate:"required"`
	Action           appconstant.Action `validate:"required"`
	Amount           decimal.Decimal    `validate:"required"`
	TransferMethodID uuid.UUID          `validate:"required"`
	Description      string
}

type GetAllByProfileIDsRequest struct {
	UserProfileID   uuid.UUID `validate:"required"`
	FriendProfileID uuid.UUID `validate:"required"`
}
