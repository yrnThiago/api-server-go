package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yrnThiago/api-server-go/config"
	"go.uber.org/zap"
)

type RedisCfg struct {
	Client      *redis.Client
	RateLimiter *RateLimiter
	IsUp        bool
}

type RateLimiter struct {
	client  *redis.Client
	limit   int
	window  time.Duration
	context context.Context
}

var Redis *RedisCfg

func Init() {
	Redis = NewRedis(config.Env.RdbAddress, config.Env.RdbPassword, config.Env.RdbDB, config.Env.RateLimit, config.Env.RateLimitWindow)

	_, err := Redis.Ping(context.Background())
	if err != nil {
		config.Logger.Warn(
			"redis did not pong",
			zap.Error(err),
		)

		Redis.IsUp = false
		return
	}

	Redis.IsUp = true
	config.Logger.Info(
		"Redis successfully initialized",
		zap.String("addr", config.Env.RdbAddress),
	)
}
func NewRateLimiter(client *redis.Client, limit int, window time.Duration, ctx context.Context) *RateLimiter {
	return &RateLimiter{
		client:  client,
		limit:   limit,
		window:  window,
		context: ctx,
	}
}

func NewRedis(addr, password string, db, limit int, window time.Duration) *RedisCfg {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	},
	)

	rateLimiter := NewRateLimiter(client, limit, window, context.Background())

	return &RedisCfg{
		Client:      client,
		RateLimiter: rateLimiter,
	}
}

func (r *RedisCfg) Ping(ctx context.Context) (string, error) {
	return r.Client.Ping(ctx).Result()
}

func (r *RedisCfg) Set(ctx context.Context, key, val string, ttl time.Duration) {
	err := r.Client.Set(ctx, key, val, ttl).Err()
	if err != nil {
		config.Logger.Warn(
			"set key/val",
			zap.Error(err),
		)

		return
	}

	config.Logger.Info(
		"set redis key/val",
		zap.String("key", key),
	)
}

func (r *RedisCfg) Get(ctx context.Context, key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		config.Logger.Warn(
			"get key/val",
			zap.Error(err),
		)

		return "", err
	}

	config.Logger.Info(
		"get redis key/val",
		zap.String("key", key),
	)
	return val, err
}

func (r *RedisCfg) Del(ctx context.Context, key string) error {
	err := r.Client.Del(ctx, key).Err()
	if err != nil {
		config.Logger.Info(
			"error del redis key",
			zap.String("key", key),
		)
		return err
	}

	config.Logger.Info(
		"key/val deleted",
		zap.String("key", key),
	)
	return nil
}

func (r *RedisCfg) Allow(key string) bool {
	pipe := r.Client.TxPipeline()
	incr := pipe.Incr(r.RateLimiter.context, key)
	pipe.Expire(r.RateLimiter.context, key, r.RateLimiter.window)

	_, err := pipe.Exec(r.RateLimiter.context)
	if err != nil {
		config.Logger.Panic("failed to exec pipe rate limit")
		return false
	}

	return incr.Val() <= int64(r.RateLimiter.limit)
}
