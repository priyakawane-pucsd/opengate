package ping

import (
	"context"
	"opengate/models/dto"
	"opengate/utils"

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

// Ping handles the ping endpoint.
// @Summary Pings the server.
// @Description Pings the server and returns "Okay" if successful.
// @Tags Ping
// @Produce json
// @Success 200 {object} dto.PingResponse
// @Error 500 utils.CustomError
// @Router /opengate/ping/ [get]
func (pc *PingController) Ping(ctx *gin.Context) {
	if pc.cfg.DB {
		err := pc.service.Ping(ctx)
		if err != nil {
			utils.WriteError(ctx, err)
			return
		}
	}
	utils.WriteResponse(ctx, dto.PingResponse{StatusCode: 200, Message: "Okay"})
}
