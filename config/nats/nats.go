package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/yrnThiago/api-server-go/config"
	"go.uber.org/zap"
)

var (
	NC *nats.Conn
	JS jetstream.JetStream
)

func Init() {
	var err error
	NC, err = nats.Connect("connect.ngs.global", nats.UserCredentials("./user.creds"))
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

func CloseAllConections() {
	config.Logger.Info("Closing all nats connections")
	NC.Close()
	JS.Conn().Close()
	config.Logger.Info("Connections closed")
}
