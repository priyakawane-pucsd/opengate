package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"opengate/config"
	"opengate/controller"
	"opengate/repository"
	"opengate/services"
)

var (
	configPath string
	env        string
)

func init() {
	flag.StringVar(&configPath, "configPath", "./", "Path to the configuration file")
	flag.StringVar(&env, "env", "dev", "environment")
	flag.Parse()
}

func main() {
	// Use the specified configPath or the default value if not provided
	fmt.Printf("Using Env: %s\n & Config file path : %s", env, configPath)

	// Get the configuration
	cfg, err := config.GetConfig(env, configPath)
	if err != nil {
		log.Fatalf("Error getting configuration: %v", err)
	}

	bytes, _ := json.MarshalIndent(cfg, "", "    ")
	log.Printf("config: \n %s", bytes)

	ctx := context.Background()

	repo := repository.NewRepository(ctx, cfg.Repository)
	srvFactory := services.NewServiceFactory(ctx, cfg.Service, repo)
	ctrl := controller.NewController(ctx, &cfg.Controller, srvFactory)
	if err = ctrl.Listen(ctx); err != nil {
		log.Panic(err)
	}
}
