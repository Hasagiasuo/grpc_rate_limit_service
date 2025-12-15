package mredis

import (
	"context"
	"fmt"
	"rlservice/internal/config"
	"rlservice/pkg/logger"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	log    *logger.Logger
	client *redis.Client
}

func NewRedisStorage(log *logger.Logger, cfg *config.RedisConfig) (*RedisStorage, error) {
	const op = "mredis.NewRedisStorage"
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Address, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Db,
		Protocol: cfg.Protocol,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Error(op, fmt.Sprintf("cannot init redis storage: %v", err))
		return nil, ErrCannotInitRedis
	}
	log.Info(op, pong)
	return &RedisStorage{
		log:    log,
		client: client,
	}, nil
}

func (rs *RedisStorage) Incr(ctx context.Context, key string, cost int64) (int64, error) {
	return rs.client.IncrBy(ctx, key, cost).Result()
}

func (rs *RedisStorage) ExpireOnce(ctx context.Context, key string, ttl time.Duration) error {
	return rs.client.Expire(ctx, key, ttl).Err()
}

func (rs *RedisStorage) GetTTL(ctx context.Context, key string) (time.Duration, error) {
	return rs.client.TTL(ctx, key).Result()
}
