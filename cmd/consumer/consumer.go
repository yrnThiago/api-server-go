package consumer

import (
	"context"
	"strings"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
)

type Consumer struct {
	Js          jetstream.JetStream
	Ctx         context.Context
	Config      jetstream.ConsumerConfig
	ConsumerCtx jetstream.Consumer
}

func NewConsumer(name, durable, filterSubject string) *Consumer {
	return &Consumer{
		Js:  config.JS,
		Ctx: context.Background(),
		Config: jetstream.ConsumerConfig{
			Name:          name,
			Durable:       durable,
			FilterSubject: filterSubject,
			AckPolicy:     jetstream.AckExplicitPolicy,
			DeliverPolicy: jetstream.DeliverAllPolicy,
		},
	}
}

func (c *Consumer) CreateStream() {
	stream, err := c.Js.Stream(c.Ctx, "orders")
	if err != nil {
		config.Logger.Fatal("err", zap.Error(err))
	}

	c.ConsumerCtx, err = stream.CreateOrUpdateConsumer(c.Ctx, c.Config)
	if err != nil {
		config.Logger.Fatal("err", zap.Error(err))
	}
}

func (c *Consumer) ConsumeSubject() {
	_, err := c.ConsumerCtx.Consume(func(msg jetstream.Msg) {
		orderID := strings.Replace(string(msg.Subject()), "orders.", "", 1)

		config.Logger.Info(
			"new order received",
			zap.String("order id", orderID),
		)

		msg.Ack()
	})
	if err != nil {
		config.Logger.Fatal("err", zap.Error(err))
	}
}

func ConsumerInit() {
	ordersConsumer := NewConsumer("order_processor", "order_processor", "orders.>")
	ordersConsumer.CreateStream()
	ordersConsumer.ConsumeSubject()

	config.Logger.Info(
		"consumers successfully initialized",
	)
}
