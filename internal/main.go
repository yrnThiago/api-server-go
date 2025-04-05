package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"

	"github.com/yrnThiago/api-server-go/internal/chiserver"
	"github.com/yrnThiago/api-server-go/internal/cmd/pub"
	"github.com/yrnThiago/api-server-go/internal/cmd/sub"
	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/models"
)

func main() {
	config.Init()
	config.DatabaseInit()
	config.LoggerInit()

	go chiserver.Init()

	// Can u please make a proper palce to config NATs
	opts := &server.Options{}
	ns, err := server.NewServer(opts)
	if err != nil {
		panic(err)
	}
	go ns.Start()

	if !ns.ReadyForConnections(4 * time.Second) {
		panic("not ready for connection")
	}

	msgChan := make(chan *nats.Msg)

	pub.PublisherInit()
	sub := sub.Connect()

	go sub.ReceiveMessage(msgChan, os.Getenv("NEW_ORDERS_TOPIC"))

	for msg := range msgChan {
		var order *models.Order

		err = json.Unmarshal(msg.Data, &order)
		if err != nil {
			return
		}

		fmt.Println(order)
	}
}
