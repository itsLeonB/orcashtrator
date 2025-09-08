package debt

import (
	"time"

	"github.com/google/uuid"
)

type TransferMethod struct {
	ID        uuid.UUID
	Name      string
	Display   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
