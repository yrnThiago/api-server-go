package infra

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var RedisClient *Redis

type Redis struct {
	Client *redis.Client
	Logger *zap.Logger
}

func NewRedisClient(addr, password string, db int, logger *zap.Logger) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	},
	)

	return &Redis{
		Client: client,
		Logger: logger,
	}
}

func (r *Redis) Ping(ctx context.Context) (string, error) {
	return r.Client.Ping(ctx).Result()
}

func (r *Redis) Set(ctx context.Context, key, val string) {
	err := r.Client.Set(ctx, key, val, 0).Err()
	if err != nil {
		r.Logger.Warn(
			"set key/val",
			zap.Error(err),
		)

		return
	}

	r.Logger.Info(
		"set redis key/val",
		zap.String("key", key),
		zap.String("val", val),
	)
}

func (r *Redis) Get(ctx context.Context, key string) string {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		r.Logger.Warn(
			"get key/val",
			zap.Error(err),
		)
	}

	r.Logger.Info(
		"get redis key/val",
		zap.String("key", key),
		zap.String("val", val),
	)
	return val
}
