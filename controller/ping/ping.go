package ping

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config struct {
	DB bool
}

type PingController struct {
	cfg     *Config
	service Service
}

type Service interface {
	Ping(ctx context.Context) error
}

func NewPingController(ctx context.Context, cfg *Config, service Service) *PingController {
	return &PingController{cfg: cfg, service: service}
}

func (pc *PingController) Register(router gin.IRouter) {
	pingRouter := router.Group("/opengate/ping")
	pingRouter.GET("/", pc.Ping)
}

func (pc *PingController) Ping(ctx *gin.Context) {
	if pc.cfg.DB {
		err := pc.service.Ping(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "failed to ping db")
		}
	}
	ctx.JSON(http.StatusOK, "Okay")
}
