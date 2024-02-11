package gateway

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Service struct {
	repo Repository
}

type Repository interface {
}

func NewService(ctx context.Context, repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) HandleRequest(ctx *gin.Context) {
	remote, err := url.Parse("http://localhost:8081")
	if err != nil {
		panic(err)
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
}
