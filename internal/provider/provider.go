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

func All(configs config.Config) (*Provider, error) {
	logger := ProvideLogger(configs.Env)
	db := meq.NewAsynqDB(logger, configs.ToRedisOpts())
	clients := ProvideClients(configs, validator.New(), logger)
	queues, err := ProvideQueues(logger, db)
	if err != nil {
		return nil, err
	}
	services, err := ProvideServices(clients, logger, configs, queues)
	if err != nil {
		return nil, err
	}

	return &Provider{
		DB:       db,
		Logger:   logger,
		Clients:  clients,
		Services: services,
	}, nil
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
