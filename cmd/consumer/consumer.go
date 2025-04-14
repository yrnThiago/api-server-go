package consumer

import (
	"context"
	"log"
	"strings"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
)

type Cons struct {
	Js     jetstream.JetStream
	Ctx    context.Context
	Config jetstream.ConsumerConfig
}

var Consumer jetstream.Consumer
var ConsumerContext jetstream.ConsumeContext

func NewConsumer(name, durable, filterSubject string) *Cons {
	return &Cons{
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

func (c *Cons) CreateStream() {
	stream, err := c.Js.Stream(c.Ctx, "orders")
	if err != nil {
		log.Fatal(err)
	}

	Consumer, err = stream.CreateOrUpdateConsumer(c.Ctx, c.Config)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Cons) ConsumeSubject() {
	var err error
	ConsumerContext, err = Consumer.Consume(func(msg jetstream.Msg) {
		orderID := strings.Replace(string(msg.Subject()), c.Config.FilterSubject, "", 1)

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

func ConsumerInit() {
	ordersConsumer := NewConsumer("order_processor", "order_processor", "orders.>")
	ordersConsumer.CreateStream()
	ordersConsumer.ConsumeSubject()
}
