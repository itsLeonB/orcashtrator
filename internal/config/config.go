package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App
	ServiceClient
}

type App struct {
	Name       string        `default:"Orcashtrator"`
	Env        string        `default:"debug"`
	Port       string        `default:"8080"`
	Timeout    time.Duration `default:"10s"`
	ClientUrls []string      `split_words:"true"`
}

type ServiceClient struct {
	BillsplittrHost string `split_words:"true" required:"true"`
	CocoonHost      string `split_words:"true" required:"true"`
	DrexHost        string `split_words:"true" required:"true"`
	StortrHost      string `split_words:"true" required:"true"`
}

func Load() Config {
	var app App
	envconfig.MustProcess("APP", &app)

	var svcClient ServiceClient
	envconfig.MustProcess("SERVICE_CLIENT", &svcClient)

	return Config{
		App:           app,
		ServiceClient: svcClient,
	}
}
