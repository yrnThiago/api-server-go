package config

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
)

var (
	NC *nats.Conn
	JS jetstream.JetStream
)

func NatsInit() {
	var err error
	NC, err = nats.Connect("connect.ngs.global", nats.UserCredentials("./user.creds"))
	if err != nil {
		Logger.Fatal(
			"nats connection",
			zap.Error(err),
		)
	}

	JS, err = jetstream.New(NC)
	if err != nil {
		Logger.Fatal(
			"jetstream",
			zap.Error(err),
		)
	}

	Logger.Info(
		"Nats successfully initialized",
	)
}

func CloseNatsConections() {
	Logger.Info("Closing all nats connections")
	NC.Close()
	JS.Conn().Close()
	Logger.Info("Connections closed")
}
