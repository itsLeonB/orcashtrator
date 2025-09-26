package provider

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/meq"
	"github.com/itsLeonB/orcashtrator/internal/config"
)

type Provider struct {
	Logger ezutil.Logger
	meq.DB
	*Clients
	*Services
}

func All(configs config.Config) *Provider {
	logger := ProvideLogger(configs.Env)
	db := meq.NewAsynqDB(logger, configs.ToRedisOpts())
	clients := ProvideClients(configs.ServiceClient, validator.New(), logger)
	queues := ProvideQueues(logger, db)

	return &Provider{
		DB:       db,
		Logger:   logger,
		Clients:  clients,
		Services: ProvideServices(clients, logger, configs.Storage, queues),
	}
}

func (p *Provider) Shutdown() error {
	var err error
	if p.DB != nil {
		if e := p.DB.Shutdown(); e != nil {
			err = errors.Join(err, e)
		}
	}
	if p.Clients != nil {
		if e := p.Clients.Shutdown(); e != nil {
			err = errors.Join(err, e)
		}
	}
	return err
}
