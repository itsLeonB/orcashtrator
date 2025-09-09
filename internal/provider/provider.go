package provider

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/config"
)

type Provider struct {
	Logger ezutil.Logger
	*Clients
	*Services
}

func All(configs config.Config) *Provider {
	logger := ProvideLogger(configs.App)
	clients := ProvideClients(configs.ServiceClient, validator.New(), logger)

	return &Provider{
		Logger:   logger,
		Clients:  clients,
		Services: ProvideServices(clients),
	}
}

func (p *Provider) Shutdown() error {
	var err error
	if p.Clients != nil {
		if e := p.Clients.Shutdown(); e != nil {
			err = errors.Join(err, e)
		}
	}
	return err
}
