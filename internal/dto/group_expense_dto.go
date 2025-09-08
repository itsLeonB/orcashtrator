package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type NewGroupExpenseRequest struct {
	PayerProfileID uuid.UUID               `json:"payerProfileId"`
	TotalAmount    decimal.Decimal         `json:"totalAmount" binding:"required"`
	Subtotal       decimal.Decimal         `json:"subtotal" binding:"required"`
	Description    string                  `json:"description"`
	Items          []NewExpenseItemRequest `json:"items" binding:"required,min=1,dive"`
	OtherFees      []NewOtherFeeRequest    `json:"otherFees" binding:"dive"`
}

type GroupExpenseResponse struct {
	ID                    uuid.UUID                    `json:"id"`
	PayerProfileID        uuid.UUID                    `json:"payerProfileId"`
	PayerName             string                       `json:"payerName,omitempty"`
	PaidByUser            bool                         `json:"paidByUser"`
	TotalAmount           decimal.Decimal              `json:"totalAmount"`
	Description           string                       `json:"description"`
	Items                 []ExpenseItemResponse        `json:"items,omitempty"`
	OtherFees             []OtherFeeResponse           `json:"otherFees,omitempty"`
	CreatorProfileID      uuid.UUID                    `json:"creatorProfileId"`
	CreatorName           string                       `json:"creatorName,omitempty"`
	CreatedByUser         bool                         `json:"createdByUser"`
	Confirmed             bool                         `json:"confirmed"`
	ParticipantsConfirmed bool                         `json:"participantsConfirmed"`
	CreatedAt             time.Time                    `json:"createdAt"`
	UpdatedAt             time.Time                    `json:"updatedAt"`
	DeletedAt             time.Time                    `json:"deletedAt,omitzero"`
	Participants          []ExpenseParticipantResponse `json:"participants,omitempty"`
}

type ExpenseParticipantResponse struct {
	ProfileName string          `json:"profileName"`
	ProfileID   uuid.UUID       `json:"profileId"`
	ShareAmount decimal.Decimal `json:"shareAmount"`
	IsUser      bool            `json:"isUser"`
}
