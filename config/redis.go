package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
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
	RedisClient = NewRedisClient("localhost:6379", "", 0)
}

func (r *Redis) Set(ctx context.Context, key, value string) {
	err := r.Client.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (r *Redis) Get(ctx context.Context, key string) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("key", val)
}
