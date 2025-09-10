package util_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestGetProfileID(t *testing.T) {
	t.Run("success with UUID", func(t *testing.T) {
		id := uuid.New()
		ctx := context.WithValue(context.Background(), appconstant.ContextProfileID, id)
		
		result, err := util.GetProfileID(ctx)
		
		assert.NoError(t, err)
		assert.Equal(t, id, result)
	})

	t.Run("success with string", func(t *testing.T) {
		id := uuid.New()
		ctx := context.WithValue(context.Background(), appconstant.ContextProfileID, id.String())
		
		result, err := util.GetProfileID(ctx)
		
		assert.NoError(t, err)
		assert.Equal(t, id, result)
	})

	t.Run("error when profileID not found", func(t *testing.T) {
		ctx := context.Background()
		
		result, err := util.GetProfileID(ctx)
		
		assert.Error(t, err)
		assert.Equal(t, uuid.Nil, result)
		assert.Contains(t, err.Error(), "profileID not found in ctx")
	})

	t.Run("error with invalid format", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), appconstant.ContextProfileID, 123)
		
		result, err := util.GetProfileID(ctx)
		
		assert.Error(t, err)
		assert.Equal(t, uuid.Nil, result)
		assert.Contains(t, err.Error(), "unknown profileID format")
	})

	t.Run("error with invalid string", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), appconstant.ContextProfileID, "invalid-uuid")
		
		result, err := util.GetProfileID(ctx)
		
		assert.Error(t, err)
		assert.Equal(t, uuid.Nil, result)
	})
}
