package services

import (
	"context"
	"opengate/services/config"
	"opengate/services/ping"
)

type Repository interface {
	ping.Repository
	config.Repository
}

type ServiceFactory struct {
	pingService   *ping.Service
	configService *config.Service
}

func NewServiceFactory(ctx context.Context, cfg Config, repo Repository) *ServiceFactory {
	factory := ServiceFactory{}
	factory.pingService = ping.NewService(ctx, repo)
	factory.configService = config.NewService(ctx, repo)
	return &factory
}

func (sf *ServiceFactory) GetPingService() *ping.Service {
	return sf.pingService
}

func (sf *ServiceFactory) GetConfigService() *config.Service {
	return sf.configService
}
