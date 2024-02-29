package cache

import (
	"context"
	"opengate/cache/redis"
	"time"
)

const (
	REDIS = "Redis"
)

type Config struct {
	Name  string
	Redis redis.Config
}
type Cache interface {
	Set(ctx context.Context, key string, value any) error
	Get(ctx context.Context, key string) (string, error)
	GetV(ctx context.Context, key string, value any) error
	SetWithTimeout(ctx context.Context, key string, value any, timeout time.Duration) error
}

func NewCache(ctx context.Context, config *Config) Cache {
	return redis.NewRedisClient(ctx, &config.Redis)
}
