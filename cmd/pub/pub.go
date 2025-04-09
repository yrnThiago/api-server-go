package pub

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/usecase"
)

var Pub *nats.Conn

func PublisherInit() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	Pub = nc
}

func SendMessage(order *usecase.OrderOutputDto) {
	orderStr, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Erro JSON")
	}

	fmt.Sprintf("New order: %s", order.ID)
	err = Pub.Publish(config.Env.NEW_ORDERS_TOPIC, []byte(orderStr))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)

	// defer Pub.Drain()
}
