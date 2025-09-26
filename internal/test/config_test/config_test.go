package config_test

import (
	"os"
	"testing"
	"time"

	"github.com/itsLeonB/orcashtrator/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	t.Run("load with environment variables", func(t *testing.T) {
		_ = os.Setenv("APP_NAME", "TestApp")
		_ = os.Setenv("APP_ENV", "test")
		_ = os.Setenv("APP_PORT", "9090")
		_ = os.Setenv("APP_TIMEOUT", "5s")
		_ = os.Setenv("SERVICE_CLIENT_BILLSPLITTR_HOST", "http://billsplittr")
		_ = os.Setenv("SERVICE_CLIENT_COCOON_HOST", "http://cocoon")
		_ = os.Setenv("SERVICE_CLIENT_DREX_HOST", "http://drex")
		_ = os.Setenv("SERVICE_CLIENT_STORTR_HOST", "http://stortr")
		_ = os.Setenv("STORAGE_BUCKET_NAME_EXPENSE_BILL", "bills")

		defer func() {
			_ = os.Unsetenv("APP_NAME")
			_ = os.Unsetenv("APP_ENV")
			_ = os.Unsetenv("APP_PORT")
			_ = os.Unsetenv("APP_TIMEOUT")
			_ = os.Unsetenv("SERVICE_CLIENT_BILLSPLITTR_HOST")
			_ = os.Unsetenv("SERVICE_CLIENT_COCOON_HOST")
			_ = os.Unsetenv("SERVICE_CLIENT_DREX_HOST")
			_ = os.Unsetenv("SERVICE_CLIENT_STORTR_HOST")
			_ = os.Unsetenv("STORAGE_BUCKET_NAME_EXPENSE_BILL")
		}()

		cfg := config.Load()

		assert.Equal(t, "TestApp", cfg.Name)
		assert.Equal(t, "test", cfg.Env)
		assert.Equal(t, "9090", cfg.Port)
		assert.Equal(t, 5*time.Second, cfg.Timeout)
		assert.Equal(t, "http://billsplittr", cfg.BillsplittrHost)
		assert.Equal(t, "http://cocoon", cfg.CocoonHost)
		assert.Equal(t, "http://drex", cfg.DrexHost)
	})

	t.Run("load with defaults", func(t *testing.T) {
		_ = os.Setenv("SERVICE_CLIENT_BILLSPLITTR_HOST", "http://billsplittr")
		_ = os.Setenv("SERVICE_CLIENT_COCOON_HOST", "http://cocoon")
		_ = os.Setenv("SERVICE_CLIENT_DREX_HOST", "http://drex")
		_ = os.Setenv("SERVICE_CLIENT_STORTR_HOST", "http://stortr")
		_ = os.Setenv("STORAGE_BUCKET_NAME_EXPENSE_BILL", "bills")

		defer func() {
			_ = os.Unsetenv("SERVICE_CLIENT_BILLSPLITTR_HOST")
			_ = os.Unsetenv("SERVICE_CLIENT_COCOON_HOST")
			_ = os.Unsetenv("SERVICE_CLIENT_DREX_HOST")
			_ = os.Unsetenv("SERVICE_CLIENT_STORTR_HOST")
			_ = os.Unsetenv("STORAGE_BUCKET_NAME_EXPENSE_BILL")
		}()

		cfg := config.Load()

		assert.Equal(t, "Orcashtrator", cfg.Name)
		assert.Equal(t, "debug", cfg.Env)
		assert.Equal(t, "8080", cfg.Port)
		assert.Equal(t, 10*time.Second, cfg.Timeout)
	})
}
