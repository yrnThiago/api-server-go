package nats

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
)

var OrdersPublisher, OffersPublisher *Publisher

type Publisher struct {
	NatsConn *nats.Conn
	Js       jetstream.JetStream
	Ctx      context.Context
	Config   jetstream.StreamConfig
}

func NewPublisher(name, description, subject string) *Publisher {
	return &Publisher{
		NatsConn: NC,
		Js:       JS,
		Ctx:      context.Background(),
		Config: jetstream.StreamConfig{
			Name:        name,
			Description: description,
			Subjects: []string{
				subject,
			},
			MaxBytes: 1024 * 1024 * 1024,
		},
	}
}

func (p *Publisher) Publish(filter, msg string) {
	_, err := p.Js.Publish(p.Ctx, fmt.Sprintf("%s.%s", filter, msg), []byte("new msg"))
	if err != nil {
		config.Logger.Warn(
			"msg cant be published",
			zap.Error(err),
		)
	}

	logMsg := fmt.Sprintf("publishing new %s", p.Config.Name)

	config.Logger.Info(
		logMsg,
		zap.String("id", msg),
	)
}

func (p *Publisher) CreateStream() {
	_, err := p.Js.CreateOrUpdateStream(p.Ctx, p.Config)
	if err != nil {
		config.Logger.Fatal(
			"publisher cant be initialized",
			zap.Error(err),
		)
	}
}

func PublisherInit() {
	OrdersPublisher = NewPublisher("orders", "Msgs for orders", "orders.>")
	OrdersPublisher.CreateStream()

	OffersPublisher = NewPublisher("offers", "Msgs for offers", "offers.>")
	OffersPublisher.CreateStream()

	config.Logger.Info(
		"publishers successfully initialized",
	)
}
