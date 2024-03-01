package gateway

import (
	"context"
	"opengate/cache"
	"opengate/constants"
	"opengate/models/dao"
	"regexp"

	"github.com/bappaapp/goutils/logger"
)

type SrvConf struct {
	Config       *dao.Config
	CompileRegex *regexp.Regexp
}

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
	s := &Service{repo: repo, cache: c}
	if err := s.populateConfigs(ctx); err != nil {
		logger.Panic(ctx, "failed to populating configs:- %v", err.Error())
		return nil
	}
	return s
}

func (s *Service) populateConfigs(ctx context.Context) error {
	if err := s.populateServiceConfig(ctx); err != nil {
		return err
	}
	if err := s.populateAuthConfig(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Service) populateServiceConfig(ctx context.Context) error {
	configs, err := s.repo.GetAllConfigs(ctx)
	if err != nil {
		logger.Error(ctx, "getting service configs")
		return err
	}
	s.srvConfigs = configs
	return nil
}

func (s *Service) populateAuthConfig(ctx context.Context) error {
	auth, err := s.repo.GetConfigById(ctx, constants.AUTH_CONFIG)
	if err != nil {
		logger.Error(ctx, "failed to get authConfig configs")
		return err
	}
	s.authConfig = auth
	return nil
}
