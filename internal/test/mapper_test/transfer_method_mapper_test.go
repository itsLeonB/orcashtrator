package mapper_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/domain/debt"
	"github.com/itsLeonB/orcashtrator/internal/mapper"
	"github.com/stretchr/testify/assert"
)

func TestTransferMethodToResponse(t *testing.T) {
	now := time.Now()
	transferMethod := debt.TransferMethod{
		ID:        uuid.New(),
		Name:      "bank_transfer",
		Display:   "Bank Transfer",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: time.Time{},
	}

	result := mapper.TransferMethodToResponse(transferMethod)

	assert.Equal(t, transferMethod.ID, result.ID)
	assert.Equal(t, transferMethod.Name, result.Name)
	assert.Equal(t, transferMethod.Display, result.Display)
	assert.Equal(t, transferMethod.CreatedAt, result.CreatedAt)
	assert.Equal(t, transferMethod.UpdatedAt, result.UpdatedAt)
	assert.Equal(t, transferMethod.DeletedAt, result.DeletedAt)
}
