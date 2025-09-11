package provider_test

import (
	"testing"

	"github.com/itsLeonB/orcashtrator/internal/config"
	"github.com/itsLeonB/orcashtrator/internal/provider"
	"github.com/stretchr/testify/assert"
)

func TestProvideLogger(t *testing.T) {
	t.Run("debug environment", func(t *testing.T) {
		cfg := config.App{
			Name: "TestApp",
			Env:  "debug",
		}
		
		logger := provider.ProvideLogger(cfg)
		
		assert.NotNil(t, logger)
	})

	t.Run("release environment", func(t *testing.T) {
		cfg := config.App{
			Name: "TestApp",
			Env:  "release",
		}
		
		logger := provider.ProvideLogger(cfg)
		
		assert.NotNil(t, logger)
	})
}
