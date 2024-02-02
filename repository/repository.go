package repository

import (
	"context"
	"log"
	"opengate/repository/mongo"
	"opengate/services"
)

const (
	MONGO = "MONGO"
)

type Config struct {
	Name  string
	Mongo mongo.Config
}

func NewRepository(ctx context.Context, config Config) services.Repository {
	switch config.Name {
	case MONGO:
		return mongo.NewRepository(ctx, &config.Mongo)
	}
	log.Panic("invalid repository name")
	return nil
}
