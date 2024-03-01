package gateway

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"opengate/utils"

	"github.com/gin-gonic/gin"
)

func (s *Service) HandleRequest(ctx *gin.Context) error {
	cfg := s.getServiceConfig(ctx, ctx.Param("path"))
	if cfg == nil {
		return utils.NewCustomError(http.StatusNotFound, "unknown service")
	}

	apiCfg := s.getApiConfig(ctx, ctx.Param("path"), cfg)
	if apiCfg == nil {
		return utils.NewCustomError(http.StatusNotFound, "unknown service")
	}

	//If authorization true then verify authorization
	if apiCfg.Authorization {
		err := s.authorizeAndModifyRequest(ctx)
		if err != nil {
			return err
		}
	}

	remote, err := url.Parse(cfg.Endpoint)
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
