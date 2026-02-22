package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

type RedisClient struct {
	Client *redis.Client
}

func NewRedis(addr, password string, db int) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           db,
		PoolSize:     10,
		MinIdleConns: 5,
	})

	// Gunakan timeout supaya tidak hang
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Cek koneksi
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &RedisClient{
		Client: rdb,
	}, nil
}

func (r *RedisClient) Set(key string, value string, ttl time.Duration) error {
	return r.Client.Set(Ctx, key, value, ttl).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.Client.Get(Ctx, key).Result()
}

func (r *RedisClient) Delete(key string) error {
	return r.Client.Del(Ctx, key).Err()
}