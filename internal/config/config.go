package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App
	Storage
	ServiceClient
}

func Load() Config {
	var app App
	envconfig.MustProcess("APP", &app)

	var storage Storage
	envconfig.MustProcess("STORAGE", &storage)

	var svcClient ServiceClient
	envconfig.MustProcess("SERVICE_CLIENT", &svcClient)

	return Config{
		app,
		storage,
		svcClient,
	}
}
