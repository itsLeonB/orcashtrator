package mapper_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/domain/profile"
	"github.com/itsLeonB/orcashtrator/internal/mapper"
	"github.com/stretchr/testify/assert"
)

func TestProfileToResponse(t *testing.T) {
	now := time.Now()
	profileDomain := profile.Profile{
		ID:          uuid.New(),
		UserID:      uuid.New(),
		Name:        "Test User",
		IsAnonymous: false,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   time.Time{},
	}

	result := mapper.ProfileToResponse(profileDomain)

	assert.Equal(t, profileDomain.ID, result.ID)
	assert.Equal(t, profileDomain.UserID, result.UserID)
	assert.Equal(t, profileDomain.Name, result.Name)
	assert.Equal(t, profileDomain.IsAnonymous, result.IsAnonymous)
	assert.Equal(t, profileDomain.CreatedAt, result.CreatedAt)
	assert.Equal(t, profileDomain.UpdatedAt, result.UpdatedAt)
	assert.Equal(t, profileDomain.DeletedAt, result.DeletedAt)
}
