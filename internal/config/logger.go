package config

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func LoggerInit() {
	Logger, _ = zap.NewProduction()
	defer Logger.Sync()
}
