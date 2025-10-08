package config

import "time"

const AppName = "Orcashtrator"

type App struct {
	Env                     string        `default:"debug"`
	Port                    string        `default:"8080"`
	Timeout                 time.Duration `default:"10s"`
	ClientUrls              []string      `split_words:"true"`
	RegisterVerificationUrl string        `split_words:"true" required:"true"`
	ResetPasswordUrl        string        `split_words:"true" required:"true"`
}
