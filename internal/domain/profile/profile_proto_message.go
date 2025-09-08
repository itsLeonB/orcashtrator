package profile

import "github.com/google/uuid"

type CreateRequest struct {
	UserID uuid.UUID
	Name   string `validate:"required,min=3"`
}
