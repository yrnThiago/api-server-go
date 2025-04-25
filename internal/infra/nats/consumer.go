package nats

import (
	"context"
	"strings"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/entity"
	"github.com/yrnThiago/api-server-go/internal/infra/repository"
	"github.com/yrnThiago/api-server-go/internal/usecase/payment"
)

type Consumer struct {
	Js             jetstream.JetStream
	Ctx            context.Context
	Config         jetstream.ConsumerConfig
	ConsumerCtx    jetstream.Consumer
	PaymentUseCase *payment.PaymentUseCase
}

func NewConsumer(name, durable, filterSubject string, paymentUseCase *payment.PaymentUseCase) *Consumer {
	return &Consumer{
		Js:  JS,
		Ctx: context.Background(),
		Config: jetstream.ConsumerConfig{
			Name:          name,
			Durable:       durable,
			FilterSubject: filterSubject,
			AckPolicy:     jetstream.AckExplicitPolicy,
			DeliverPolicy: jetstream.DeliverAllPolicy,
		},
		PaymentUseCase: paymentUseCase,
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

func (c *Consumer) HandlingNewOrders() {
	_, err := c.ConsumerCtx.Consume(func(msg jetstream.Msg) {
		orderID := strings.Replace(string(msg.Subject()), "orders.", "", 1)
		msg.Ack()

		config.Logger.Info(
			"new order received",
			zap.String("order id", orderID),
		)

		order, _ := c.PaymentUseCase.OrderRepository.GetById(orderID)
		if !c.PaymentUseCase.IsOrderPaymentValid(order) {
			order.SetStatus(entity.Canceled)
			c.PaymentUseCase.OrderRepository.UpdateById(order)

			return
		}

		order.SetStatus(entity.Pending)
		paymentRes := c.PaymentUseCase.GeneratePayment(order)
		config.Logger.Info(
			"payment generated",
			zap.String("result", paymentRes),
			zap.String("order id", order.ID),
		)
	})
	if err != nil {
		config.Logger.Fatal("err", zap.Error(err))
	}
}

func ConsumerInit() {
	repositoryOrders := repository.NewOrderRepositoryMysql(config.DB)
	paymentUseCase := payment.NewPaymentUseCase(repositoryOrders)

	ordersConsumer := NewConsumer("order_processor", "order_processor", "orders.>", paymentUseCase)
	ordersConsumer.CreateStream()
	ordersConsumer.HandlingNewOrders()

	config.Logger.Info(
		"consumers successfully initialized",
	)
}
