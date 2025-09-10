package domain_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestAuditMetadata(t *testing.T) {
	id := uuid.New()
	now := time.Now()
	
	metadata := domain.AuditMetadata{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: time.Time{},
	}
	
	assert.Equal(t, id, metadata.ID)
	assert.Equal(t, now, metadata.CreatedAt)
	assert.Equal(t, now, metadata.UpdatedAt)
	assert.True(t, metadata.DeletedAt.IsZero())
}
