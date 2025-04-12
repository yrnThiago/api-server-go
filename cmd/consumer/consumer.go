package consumer

import (
	"context"
	"log"

	"github.com/nats-io/nats.go/jetstream"

	"github.com/yrnThiago/api-server-go/config"
)

var ConsumerContext jetstream.ConsumeContext
var Consumer jetstream.Consumer

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
		log.Printf("Received message: %s", string(msg.Subject()))
		msg.Ack()
	})
	if err != nil {
		log.Fatal(err)
	}
}
