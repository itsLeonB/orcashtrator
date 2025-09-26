package config

import "time"

type App struct {
	Name       string        `default:"Orcashtrator"`
	Env        string        `default:"debug"`
	Port       string        `default:"8080"`
	Timeout    time.Duration `default:"10s"`
	ClientUrls []string      `split_words:"true"`
}
