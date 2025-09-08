package provider

import (
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/config"
)

func ProvideLogger(configs config.App) ezutil.Logger {
	minLevel := 0
	if configs.Env == "release" {
		minLevel = 1
	}

	return ezutil.NewSimpleLogger(configs.Name, true, minLevel)
}
