package pub

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/yrnThiago/gdlp-go/internal/usecase"
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
	msg := fmt.Sprintf("New order: %s", order.ID)
	err := Pub.Publish(os.Getenv("NEW_ORDERS_TOPIC"), []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)

	// defer Pub.Drain()
}
