package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App
	Valkey
	Storage
	ServiceClient
}

func Load() Config {
	var app App
	envconfig.MustProcess("APP", &app)

	var valkey Valkey
	envconfig.MustProcess("VALKEY", &valkey)

	var storage Storage
	envconfig.MustProcess("STORAGE", &storage)

	var svcClient ServiceClient
	envconfig.MustProcess("SERVICE_CLIENT", &svcClient)

	return Config{
		app,
		valkey,
		storage,
		svcClient,
	}
}
