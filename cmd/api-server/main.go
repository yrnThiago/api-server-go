package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/yrnThiago/api-server-go/cmd/consumer"
	"github.com/yrnThiago/api-server-go/cmd/publisher"
	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/server"
)

func main() {
	config.Init()
	config.LoggerInit()
	config.DatabaseInit()
	config.NatsInit()
	config.RedisInit()

	publisher.PubInit()
	consumer.ConsumerInit()

	go server.Init()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	config.CloseNatsConections()
}
