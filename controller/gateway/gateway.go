package gateway

import (
	"context"
)

type Config struct {
}

type GatewayController struct {
	cfg *Config
}

func NewPingController(ctx context.Context, cfg *Config) *GatewayController {
	return &GatewayController{cfg: cfg}
}

func (pc *GatewayController) Register() {

}
