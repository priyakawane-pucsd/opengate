package controller

import (
	"context"
	"fmt"
	"log"
	authconfig "opengate/controller/authConfig"
	"opengate/controller/config"
	"opengate/controller/gateway"
	"opengate/controller/ping"
	"opengate/controller/swagger"
	"opengate/services"

	"github.com/bappaapp/goutils/logger"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Port         int
	GinModeDebug bool
	Ping         ping.Config
	Config       config.Config
	Gateway      gateway.Config
	AuthConfig   authconfig.Config
}

type Controller struct {
	config     *Config
	srvFactory *services.ServiceFactory
}

func NewController(ctx context.Context, cfg *Config, srvFactory *services.ServiceFactory) *Controller {
	return &Controller{config: cfg, srvFactory: srvFactory}
}

func (c *Controller) Listen(ctx context.Context) error {
	if !c.config.GinModeDebug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	//registering controllers
	ping.NewPingController(ctx, &c.config.Ping, c.srvFactory.GetPingService()).Register(router)
	swagger.NewSwaggerController(ctx).Register(router)
	config.NewConfigController(ctx, &c.config.Config, c.srvFactory.GetConfigService()).Register(router)
	gateway.NewGatewayController(ctx, &c.config.Gateway, c.srvFactory.GetGatewayService()).Register(router)
	authconfig.NewAuthConfigController(ctx, &c.config.AuthConfig, c.srvFactory.GetConfigService()).Register(router)

	logger.Info(ctx, "swagger link: http://localhost:%d/opengate/swagger/index.html", c.config.Port)
	log.Printf("HTTP server started listening on :%d", c.config.Port)
	return router.Run(fmt.Sprintf(":%d", c.config.Port))
}
