package gateway

import (
	"context"
	"opengate/cache"
	"opengate/constants"
	"opengate/models/dao"

	"github.com/bappaapp/goutils/logger"
)

type Service struct {
	repo       Repository
	srvConfigs []*dao.Config
	authConfig *dao.Config
	cache      cache.Cache
}

type Repository interface {
	GetAllConfigs(ctx context.Context) ([]*dao.Config, error)
	GetConfigById(ctx context.Context, id string) (*dao.Config, error)
}

func NewService(ctx context.Context, repo Repository, c cache.Cache) *Service {
	configs, err := repo.GetAllConfigs(ctx)
	if err != nil {
		logger.Panic(ctx, "getting service configs")
		return nil
	}

	auth, err := repo.GetConfigById(ctx, constants.AUTH_CONFIG)
	if err != nil {
		logger.Panic(ctx, "failed to get authConfig configs")
		return nil
	}
	return &Service{repo: repo, srvConfigs: configs, authConfig: auth, cache: c}
}
