package publisher

import (
	"context"
	"fmt"
	"log"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/yrnThiago/api-server-go/config"
	"go.uber.org/zap"
)

func Publish(orderId string) {
	ctx := context.Background()
	_, err := config.JS.Publish(ctx, fmt.Sprintf("orders.%s", orderId), []byte("new order"))
	if err != nil {
		log.Println(err)
	}

	config.Logger.Info(
		"publishing new order",
		zap.String("order id", orderId),
	)
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
