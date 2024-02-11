package services

import (
	"context"
	"opengate/services/config"
	"opengate/services/gateway"
	"opengate/services/ping"
)

type Repository interface {
	ping.Repository
	config.Repository
	gateway.Repository
}

type ServiceFactory struct {
	pingService    *ping.Service
	configService  *config.Service
	gatewayService *gateway.Service
}

func NewServiceFactory(ctx context.Context, cfg Config, repo Repository) *ServiceFactory {
	factory := ServiceFactory{}
	factory.pingService = ping.NewService(ctx, repo)
	factory.configService = config.NewService(ctx, repo)
	factory.gatewayService = gateway.NewService(ctx, repo)
	return &factory
}

func (sf *ServiceFactory) GetPingService() *ping.Service {
	return sf.pingService
}

func (sf *ServiceFactory) GetConfigService() *config.Service {
	return sf.configService
}

func (sf *ServiceFactory) GetGatewayService() *gateway.Service {
	return sf.gatewayService
}
