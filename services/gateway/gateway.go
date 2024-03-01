package gateway

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"opengate/cache"
	"opengate/constants"
	"opengate/models/dao"
	"opengate/utils"
	"regexp"

	"github.com/bappaapp/goutils/logger"
	"github.com/gin-gonic/gin"
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

// completed this function
func (s *Service) getConfig(ctx context.Context, urlPath string) *dao.Config {
	for _, c := range s.srvConfigs {
		r, err := regexp.Compile(c.ServiceConfig.Regex)
		if err != nil {
			logger.Error(ctx, "invalid regular expression in config: %v", c)
			continue
		}

		match := r.FindString(urlPath)
		if match != "" {
			return c
		}
	}
	return nil
}

func (s *Service) HandleRequest(ctx *gin.Context) error {
	cfg := s.getConfig(ctx, ctx.Param("path"))
	if cfg == nil {
		return utils.NewCustomError(http.StatusNotFound, "unknown service")
	}

	//If authorization true then verify authorization
	if cfg.ServiceConfig.Authorization {
		err := s.authorizeAndModifyRequest(ctx)
		if err != nil {
			return err
		}
	}

	remote, err := url.Parse(cfg.ServiceConfig.Endpoint)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "invalid endpoint config in db")
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = ctx.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.Method = ctx.Request.Method
		req.URL.Path = ctx.Param("path")
	}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
	return nil
}
