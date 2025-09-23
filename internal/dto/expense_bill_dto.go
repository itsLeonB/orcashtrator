package dto

import (
	"io"
	"time"

	"github.com/google/uuid"
)

type NewExpenseBillRequest struct {
	CreatorProfileID uuid.UUID
	PayerProfileID   uuid.UUID
	ImageReader      io.ReadCloser
	ContentType      string
	Filename         string
	FileSize         int64
}

type ExpenseBillResponse struct {
	ID                 uuid.UUID `json:"id"`
	CreatorProfileID   uuid.UUID `json:"creatorProfileId"`
	PayerProfileID     uuid.UUID `json:"payerProfileId"`
	ImageURL           string    `json:"imageUrl"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	DeletedAt          time.Time `json:"deletedAt,omitzero"`
	IsCreatedByUser    bool      `json:"isCreatedByUser"`
	IsPaidByUser       bool      `json:"isPaidByUser"`
	CreatorProfileName string    `json:"creatorProfileName"`
	PayerProfileName   string    `json:"payerProfileName"`
}
