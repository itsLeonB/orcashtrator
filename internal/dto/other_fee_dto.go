package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/shopspring/decimal"
)

type FeeParticipantResponse struct {
	ProfileName string          `json:"profileName"`
	ProfileID   uuid.UUID       `json:"profileId"`
	ShareAmount decimal.Decimal `json:"shareAmount"`
	IsUser      bool            `json:"isUser"`
}

type FeeCalculationMethodInfo struct {
	Name        string `json:"name"`
	Display     string `json:"display"`
	Description string `json:"description"`
}

type OtherFeeResponse struct {
	ID                uuid.UUID                        `json:"id"`
	Name              string                           `json:"name"`
	Amount            decimal.Decimal                  `json:"amount"`
	CalculationMethod appconstant.FeeCalculationMethod `json:"calculationMethod"`
	CreatedAt         time.Time                        `json:"createdAt"`
	UpdatedAt         time.Time                        `json:"updatedAt"`
	DeletedAt         time.Time                        `json:"deletedAt,omitzero"`
	Participants      []FeeParticipantResponse         `json:"participants,omitempty"`
}

type NewOtherFeeRequest struct {
	GroupExpenseID    uuid.UUID                        `json:"-"`
	Name              string                           `json:"name" binding:"required,min=3"`
	Amount            decimal.Decimal                  `json:"amount" binding:"required"`
	CalculationMethod appconstant.FeeCalculationMethod `json:"calculationMethod" binding:"required"`
}

type UpdateOtherFeeRequest struct {
	ID                uuid.UUID                        `json:"-"`
	GroupExpenseID    uuid.UUID                        `json:"-"`
	Name              string                           `json:"name" binding:"required,min=3"`
	Amount            decimal.Decimal                  `json:"amount" binding:"required"`
	CalculationMethod appconstant.FeeCalculationMethod `json:"calculationMethod" binding:"required"`
}
