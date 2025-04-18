package config

import (
	"context"
	"strconv"

	infra "github.com/yrnThiago/api-server-go/internal/infra/redis"
	"go.uber.org/zap"
)

var RedisClient *infra.Redis

func RedisInit() {
	rdbDb, _ := strconv.Atoi(Env.RDB_DB)
	RedisClient = infra.NewRedisClient(Env.RDB_ADDRESS, Env.RDB_PASSWORD, rdbDb, Logger)

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
