package publisher

import (
	"context"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
)

var OrdersPub *Pub

type Pub struct {
	NatsConn *nats.Conn
	Js       jetstream.JetStream
	Ctx      context.Context
	Config   jetstream.StreamConfig
}

func NewPub(name, description, subject string) *Pub {
	return &Pub{
		NatsConn: config.NC,
		Js:       config.JS,
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

func (p *Pub) Publish(msg string) {
	_, err := p.Js.Publish(p.Ctx, fmt.Sprintf("orders.%s", msg), []byte("new order"))
	if err != nil {
		log.Println(err)
	}

	config.Logger.Info(
		"publishing new order",
		zap.String("order id", msg),
	)
}

func (p *Pub) CreateStream() {
	_, err := p.Js.CreateOrUpdateStream(p.Ctx, p.Config)
	if err != nil {
		log.Fatal(err)
	}
}

func PubInit() {
	OrdersPub = NewPub("orders", "Msgs for orders", "orders.>")
	OrdersPub.CreateStream()
}
