package redis

import (
	"context"
	"github.com/Sabyradinov/golang-hex-layout/config"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/cache"
	"github.com/go-redis/redis/v8"
	"time"
)

// redisClient redis client instance
type redisClient struct {
	client *redis.Client
}

// New constructor to return redis client instance
func New(cfg *config.Redis) (cache.ICache, error) {
	cacheInstance := &redisClient{}
	cacheInstance.client = redis.NewClient(&redis.Options{
		Addr:     cfg.GetHostPort(),
		DB:       cfg.DatabaseNumber,
		Password: cfg.Password})

	if err := cacheInstance.client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return cacheInstance, nil
}

// RegisterMetrics function to register prometheus metrics
func (c *redisClient) RegisterMetrics() {
	// No need to implement
}

// Get return redis value by key
func (c *redisClient) Get(ctx context.Context, key string) (res string, err error) {
	res, err = c.client.Get(ctx, key).Result()
	return
}

// Set value to redis
func (c *redisClient) Set(ctx context.Context, key string, val interface{}, ttl time.Duration) (err error) {
	err = c.client.Set(ctx, key, val, ttl).Err()
	return
}

// Del delete value from redis by key
func (c *redisClient) Del(ctx context.Context, key string) (err error) {
	err = c.client.Del(ctx, key).Err()
	return
}

// Exists check value on redis with key
func (c *redisClient) Exists(ctx context.Context, key string) (res bool) {
	cnt, err := c.client.Exists(ctx, key).Result()
	if err != nil {
		res = false
		return
	}

	if cnt > 0 {
		res = true
		return
	}

	return
}
