package dto

import (
	"time"

	"github.com/google/uuid"
)

type ProfileResponse struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"userId"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitzero"`
	IsAnonymous bool      `json:"isAnonymous"`
}

type UpdateProfileRequest struct {
	Name string `json:"name" binding:"required,min=3,max=255"`
}
