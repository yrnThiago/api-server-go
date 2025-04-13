package consumer

import (
	"context"
	"log"
	"strings"

	"github.com/nats-io/nats.go/jetstream"

	"github.com/yrnThiago/api-server-go/config"
	"go.uber.org/zap"
)

var (
	ConsumerContext jetstream.ConsumeContext
	Consumer        jetstream.Consumer
)

func StartOrdersConsumer() {
	ctx := context.Background()
	stream, err := config.JS.Stream(ctx, "orders")
	if err != nil {
		log.Fatal(err)
	}

	Consumer, err = stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:          "order_processor",
		Durable:       "order_processor",
		FilterSubject: "orders.>",
		AckPolicy:     jetstream.AckExplicitPolicy,
		DeliverPolicy: jetstream.DeliverAllPolicy,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ConsumeMsgs() {
	var err error
	ConsumerContext, err = Consumer.Consume(func(msg jetstream.Msg) {
		orderID := strings.Replace(string(msg.Subject()), "orders.", "", 1)

		config.Logger.Info(
			"new order received",
			zap.String("order id", orderID),
		)

		msg.Ack()
	})
	if err != nil {
		log.Fatal(err)
	}
}
