package gateway

import (
	"context"
	"opengate/models/dao"
	"regexp"

	"github.com/bappaapp/goutils/logger"
)

func (s *Service) getServiceConfig(ctx context.Context, urlPath string) *dao.ServiceConfig {
	for _, c := range s.srvConfigs {
		r, err := regexp.Compile(c.ServiceConfig.Regex)
		if err != nil {
			logger.Error(ctx, "invalid regular expression in config: %v", c)
			continue
		}

		if r.Match([]byte(urlPath)) {
			return c.ServiceConfig
		}
	}
	return nil
}

func (s *Service) getApiConfig(ctx context.Context, urlPath string, serviceConfig *dao.ServiceConfig) *dao.ServiceApis {
	for _, c := range serviceConfig.Apis {
		r, err := regexp.Compile(c.Regex)
		if err != nil {
			logger.Error(ctx, "invalid regular expression in config: %v", c)
			continue
		}

		if r.Match([]byte(urlPath)) {
			return &c
		}
	}
	return nil
}
