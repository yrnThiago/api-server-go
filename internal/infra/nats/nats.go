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
			"unable to connect with nats server",
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
	if utils.IsEmpty(config.Env.NatsURL) {
		return nats.DefaultURL
	}

	return config.Env.NatsURL
}

func CloseAllConections() {
	config.Logger.Info("Closing all nats connections")
	NC.Close()
	JS.Conn().Close()
	config.Logger.Info("Connections closed")
}
