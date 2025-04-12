package config

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var (
	NC *nats.Conn
	JS jetstream.JetStream
)

func NatsInit() {
	var err error
	NC, err = nats.Connect("connect.ngs.global", nats.UserCredentials("./user.creds"))
	if err != nil {
		log.Fatal(err)
	}

	// defer nc.Close()
	JS, err = jetstream.New(NC)
	if err != nil {
		log.Fatal(err)
	}
}

func CloseNatsConections() {
	Logger.Info("Closing all nats connections")
	NC.Close()
	JS.Conn().Close()
	Logger.Info("Connections closed")
}
