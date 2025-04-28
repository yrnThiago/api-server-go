package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/utils"
	"go.uber.org/zap"
)

var (
	NC *nats.Conn
	JS jetstream.JetStream
)

func Init() {
	var err error
	natsURL := getNatsURL()
	NC, err = nats.Connect(natsURL)
	if err != nil {
		config.Logger.Fatal(
			"nats connection",
			zap.Error(err),
		)
	}

	JS, err = jetstream.New(NC)
	if err != nil {
		config.Logger.Fatal(
			"jetstream",
			zap.Error(err),
		)
	}

	config.Logger.Info(
		"Nats successfully initialized",
	)
}

func getNatsURL() string {
	if utils.IsEmpty(config.Env.NATS_URL) {
		return nats.DefaultURL
	}

	return config.Env.NATS_URL
}

func CloseAllConections() {
	config.Logger.Info("Closing all nats connections")
	NC.Close()
	JS.Conn().Close()
	config.Logger.Info("Connections closed")
}
