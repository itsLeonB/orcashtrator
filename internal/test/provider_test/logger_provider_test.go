package provider_test

import (
	"testing"

	"github.com/itsLeonB/orcashtrator/internal/provider"
	"github.com/stretchr/testify/assert"
)

func TestProvideLogger(t *testing.T) {
	t.Run("debug environment", func(t *testing.T) {
		logger := provider.ProvideLogger("debug")

		assert.NotNil(t, logger)
	})

	t.Run("release environment", func(t *testing.T) {
		logger := provider.ProvideLogger("release")

		assert.NotNil(t, logger)
	})
}
