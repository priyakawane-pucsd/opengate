package controller

import (
	"context"
	"fmt"
	"log"
	"opengate/controller/ping"
	"opengate/services"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port         int
	GinModeDebug bool
	Ping         ping.Config
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

	log.Printf("HTTP server started listening on :%d", c.config.Port)
	return router.Run(fmt.Sprintf(":%d", c.config.Port))
}
