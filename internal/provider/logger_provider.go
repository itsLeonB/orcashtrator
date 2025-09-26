package provider

import (
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/config"
)

func ProvideLogger(env string) ezutil.Logger {
	minLevel := 0
	if env == "release" {
		minLevel = 1
	}

	return ezutil.NewSimpleLogger(config.AppName, true, minLevel)
}
