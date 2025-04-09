package main

import (
	"encoding/json"
	"fmt"

	"github.com/yrnThiago/api-server-go/cmd/pub"
	"github.com/yrnThiago/api-server-go/cmd/sub"
	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/fiber"
	"github.com/yrnThiago/api-server-go/internal/models"
)

func main() {
	config.Init()
	config.DatabaseInit()
	config.LoggerInit()
	config.NatsInit()

	go fiber.Init()

	pub.PublisherInit()
	sub := sub.Connect()

	go sub.ReceiveMessage(config.MsgChan, config.Env.NEW_ORDERS_TOPIC)

	for msg := range config.MsgChan {
		var order *models.Order

		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			return
		}

		fmt.Println(order)
	}
}
