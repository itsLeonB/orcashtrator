package profile

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Name      string
	Avatar    string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
