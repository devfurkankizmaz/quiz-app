package config

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	redis *redis.Client
}

var redisInstance *redis.Client

func NewRedisCache(env *Env) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     env.RedisAddress,
		Password: env.RedisPass,
		DB:       env.RedisDb,
	})

	redisInstance = client

	return redisInstance
}

func (r *Redis) Ping(ctx context.Context) error {
	_, err := r.redis.Ping(ctx).Result()
	return err
}

func (r *Redis) Close() error {
	err := r.redis.Close()
	return err
}
