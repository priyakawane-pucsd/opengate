package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bappaapp/goutils/logger"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

type Config struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisClient(ctx context.Context, config *Config) *RedisClient {
	// Connect to the Redis server
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,     // "localhost:6379"
		Password: config.Password, // no password by default
		DB:       config.DB,       // use the default DB
	})

	// Check if the connection was successful
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Panic(ctx, "Failed to connect to Redis: %v", err)
		return nil
	}
	fmt.Println("Connected to Redis:", pong)
	// Initialize and return a new RedisClient instance
	return &RedisClient{
		client: client,
	}
}

func (r *RedisClient) Set(ctx context.Context, key string, value any) error {
	// Implementation to set a key-value pair in Redis
	v, err := json.Marshal(value)
	if err != nil {
		logger.Error(ctx, "invalid value to set in redis %v", err.Error())
		return fmt.Errorf("invalid value to set in redis")
	}
	err = r.client.Set(context.Background(), key, string(v), 0).Err()
	if err != nil {
		logger.Error(ctx, "Failed to set key-value pair in Redis: %v", err)
	}
	return err
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	// Implementation to get a value from Redis using key
	value, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		//logger.Error(ctx, "Failed to get value from Redis: %v", err)
		return "", err
	}
	return value, nil
}

func (r *RedisClient) GetV(ctx context.Context, key string, value any) error {
	v, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		//logger.Error(ctx, "Failed to get value from Redis: %v", err)
		return err
	}
	err = json.Unmarshal([]byte(v), value)
	if err != nil {
		return fmt.Errorf("invalid value for key %v", key)
	}
	return nil
}

func (r *RedisClient) SetWithTimeout(ctx context.Context, key string, value any, timeout time.Duration) error {
	// Implementation to set a key-value pair in Redis
	v, err := json.Marshal(value)
	if err != nil {
		logger.Error(ctx, "invalid value to set in redis %v", err.Error())
		return fmt.Errorf("invalid value to set in redis")
	}
	fmt.Println("timeout>>>>>", timeout)
	err = r.client.Set(context.Background(), key, string(v), timeout*time.Minute).Err()
	if err != nil {
		logger.Error(ctx, "Failed to set key-value pair in Redis: %v", err)
	}
	return err
}
