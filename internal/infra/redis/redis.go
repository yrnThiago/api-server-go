package infra

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis struct {
	Client      *redis.Client
	Logger      *zap.Logger
	RateLimiter *RateLimiter
	IsUp        bool
}

type RateLimiter struct {
	client  *redis.Client
	limit   int
	window  time.Duration
	context context.Context
}

func NewRateLimiter(client *redis.Client, limit int, window time.Duration, ctx context.Context) *RateLimiter {
	return &RateLimiter{
		client:  client,
		limit:   limit,
		window:  window,
		context: ctx,
	}
}

func NewRedisClient(addr, password string, db int, logger *zap.Logger) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	},
	)

	rateLimiter := NewRateLimiter(client, 10, 1*time.Minute, context.Background())

	return &Redis{
		Client:      client,
		Logger:      logger,
		RateLimiter: rateLimiter,
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

func (r *Redis) Allow(key string) bool {
	pipe := r.Client.TxPipeline()
	incr := pipe.Incr(r.RateLimiter.context, key)
	pipe.Expire(r.RateLimiter.context, key, r.RateLimiter.window)

	_, err := pipe.Exec(r.RateLimiter.context)
	if err != nil {
		r.Logger.Panic("failed to exec pipe rate limit")
		return false
	}

	return incr.Val() <= int64(r.RateLimiter.limit)
}
