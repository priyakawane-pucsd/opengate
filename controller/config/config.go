package config

import (
	"context"
)

type Config struct {
}

type ConfigController struct {
	cfg *Config
}

func NewPingController(ctx context.Context, cfg *Config) *ConfigController {
	return &ConfigController{cfg: cfg}
}

func (pc *ConfigController) Register() {

}
