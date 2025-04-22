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

func NewRedis(addr, password string, db, limit int, window time.Duration, logger *zap.Logger) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	},
	)

	rateLimiter := NewRateLimiter(client, limit, window, context.Background())

	return &Redis{
		Client:      client,
		Logger:      logger,
		RateLimiter: rateLimiter,
	}
}

func (r *Redis) Ping(ctx context.Context) (string, error) {
	return r.Client.Ping(ctx).Result()
}

func (r *Redis) Set(ctx context.Context, key, val string, ttl time.Duration) {
	err := r.Client.Set(ctx, key, val, ttl).Err()
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
	)
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		r.Logger.Warn(
			"get key/val",
			zap.Error(err),
		)

		return "", err
	}

	r.Logger.Info(
		"get redis key/val",
		zap.String("key", key),
	)
	return val, err
}

func (r *Redis) Del(ctx context.Context, key string) error {
	err := r.Client.Del(ctx, key).Err()
	if err != nil {
		r.Logger.Info(
			"error del redis key",
			zap.String("key", key),
		)
		return err
	}

	r.Logger.Info(
		"key/val deleted",
		zap.String("key", key),
	)
	return nil
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
