package config

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var RedisClient *Redis

type Redis struct {
	Client *redis.Client
}

func NewRedisClient(addr, password string, db int) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &Redis{
		Client: client,
	}
}

func RedisInit() {
	rdbDb, _ := strconv.Atoi(Env.RDB_DB)
	RedisClient = NewRedisClient(Env.RDB_ADDRESS, Env.RDB_PASSWORD, rdbDb)

	_, err := RedisClient.Ping(context.Background())
	if err != nil {
		Logger.Warn(
			"redis did not pong",
			zap.Error(err),
		)

		return
	}

	Logger.Info(
		"Redis successfully initialized",
		zap.String("addr", Env.RDB_ADDRESS),
	)

}

func (r *Redis) Ping(ctx context.Context) (string, error) {
	return r.Client.Ping(ctx).Result()
}

func (r *Redis) Set(ctx context.Context, key, val string) {
	err := r.Client.Set(ctx, key, val, 0).Err()
	if err != nil {
		Logger.Warn(
			"set key/val",
			zap.Error(err),
		)

		return
	}

	Logger.Info(
		"set redis key/val",
		zap.String("key", key),
		zap.String("val", val),
	)
}

func (r *Redis) Get(ctx context.Context, key string) string {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		Logger.Warn(
			"get key/val",
			zap.Error(err),
		)
	}

	Logger.Info(
		"get redis key/val",
		zap.String("key", key),
		zap.String("val", val),
	)
	return val
}
