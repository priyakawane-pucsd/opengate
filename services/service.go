package services

import (
	"context"
	"opengate/services/ping"
)

type Repository interface {
	ping.Repository
}

type ServiceFactory struct {
	pingService *ping.Service
}

func NewServiceFactory(ctx context.Context, config Config, repo Repository) *ServiceFactory {
	factory := ServiceFactory{}
	factory.pingService = ping.NewService(ctx, repo)
	return &factory
}

func (sf *ServiceFactory) GetPingService() *ping.Service {
	return sf.pingService
}
