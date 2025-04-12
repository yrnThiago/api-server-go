package publisher

import (
	"context"
	"fmt"
	"log"

	"github.com/nats-io/nats.go/jetstream"

	"github.com/yrnThiago/api-server-go/config"
)

func Publish(orderId string) {
	ctx := context.Background()
	_, err := config.JS.Publish(ctx, fmt.Sprintf("orders.%s", orderId), []byte("new order"))
	if err != nil {
		log.Println(err)
	}

	log.Printf("Published message %s", orderId)
}

func StartOrdersPublisher() {
	ctx := context.Background()
	_, err := config.JS.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:        "orders",
		Description: "Messages for orders",
		Subjects: []string{
			"orders.>",
		},
		MaxBytes: 1024 * 1024 * 1024,
	})
	if err != nil {
		log.Fatal(err)
	}
}
