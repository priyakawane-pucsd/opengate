package gateway

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"opengate/models/dao"
	"opengate/utils"
	"regexp"

	"github.com/bappaapp/goutils/logger"
	"github.com/gin-gonic/gin"
)

type Service struct {
	repo       Repository
	srvConfigs []*dao.ServiceConfig
}

type Repository interface {
	GetAllConfigs(ctx context.Context) ([]*dao.ServiceConfig, error)
}

func NewService(ctx context.Context, repo Repository) *Service {
	configs, err := repo.GetAllConfigs(ctx)
	if err != nil {
		logger.Panic(ctx, "getting service configs")
		return nil
	}
	return &Service{repo: repo, srvConfigs: configs}
}

// completed this function
func (s *Service) getConfig(urlPath string) *dao.ServiceConfig {
	var conf *dao.ServiceConfig
	for _, c := range s.srvConfigs {
		r, err := regexp.Compile(*&c.ServiceConfig.Regex)
		if err != nil {
			return nil
		}

		match := r.FindString(urlPath)
		if match != "" {
			conf = c
		}
	}
	return conf
}

func (s *Service) HandleRequest(ctx *gin.Context) error {
	cfg := s.getConfig(ctx.Param("path"))
	if cfg == nil {
		return utils.NewCustomError(http.StatusNotFound, "unknown service")
	}
	remote, err := url.Parse(cfg.ServiceConfig.Regex)
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
