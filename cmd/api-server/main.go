package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/infra/nats"
	infra "github.com/yrnThiago/api-server-go/internal/infra/redis"
	"github.com/yrnThiago/api-server-go/internal/server"
)

func main() {
	config.Init()
	config.LoggerInit()
	config.DatabaseInit()

	nats.Init()
	nats.PublisherInit()
	nats.ConsumerInit()

	infra.RedisInit()

	go server.Init()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	nats.CloseAllConections()
}
