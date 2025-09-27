package config

import (
	"crypto/tls"

	"github.com/hibiken/asynq"
)

type Valkey struct {
	Addr     string `required:"true,min=3"`
	Password string `required:"true"`
	Db       int
}

func (v Valkey) ToRedisOpts() asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Addr:     v.Addr,
		Password: v.Password,
		DB:       v.Db,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}
}
