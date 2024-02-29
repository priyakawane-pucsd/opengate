package config

import (
	"log"
	"opengate/cache"
	"opengate/controller"
	"opengate/repository"
	"opengate/services"

	"github.com/spf13/viper"
)

type Config struct {
	Name       string
	Build      string
	Controller controller.Config
	Service    services.Config
	Repository repository.Config
	Cache      cache.Config
}

func GetConfig(env, path string) (*Config, error) {
	// Set the configuration file name and path based on the environment
	viper.SetConfigName(env)
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
		return nil, err
	}

	// Unmarshal the configuration into the Config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling configuration: %v", err)
		return nil, err
	}

	return &config, nil
}
