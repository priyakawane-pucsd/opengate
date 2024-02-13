package gateway

import (
	"context"
	"opengate/utils"

	"github.com/gin-gonic/gin"
)

type Config struct {
	DB bool
}

type GatewayController struct {
	cfg     *Config
	service Service
}

type Service interface {
	HandleRequest(ctx *gin.Context) error
}

func NewGatewayController(ctx context.Context, cfg *Config, s Service) *GatewayController {
	return &GatewayController{cfg: cfg, service: s}
}

func (gc *GatewayController) Register(router gin.IRouter) {
	router.Any("/api/*path", gc.handleGatewayRequest)
}

func (gc *GatewayController) handleGatewayRequest(ctx *gin.Context) {
	err := gc.service.HandleRequest(ctx)
	if err != nil {
		utils.WriteError(ctx, err)
	}
}
