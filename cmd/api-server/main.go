package main

import (
	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/infra/nats"
	"github.com/yrnThiago/api-server-go/internal/infra/redis"
	"github.com/yrnThiago/api-server-go/internal/server"
)

func main() {
	config.Init()
	config.LoggerInit()
	config.DatabaseInit()

	nats.Init()
	nats.PublisherInit()
	nats.ConsumerInit()

	redis.Init()

	server.Init()
}
