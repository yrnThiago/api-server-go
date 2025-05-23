package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/entity"
	"github.com/yrnThiago/api-server-go/internal/infra/redis"
	"github.com/yrnThiago/api-server-go/internal/infra/repository"
	offerUseCase "github.com/yrnThiago/api-server-go/internal/usecase/offer"
	usecase "github.com/yrnThiago/api-server-go/internal/usecase/payment"
)

const (
	OrdersSubject = "orders"
	OffersSubject = "offers"
)

var OrdersFilter = fmt.Sprintf("%s.>", OrdersSubject)
var OffersAcceptedFilter = fmt.Sprintf("%s.accepted.>", OffersSubject)
var OffersDeclinedFilter = fmt.Sprintf("%s.declined.>", OffersSubject)
var OffersPendingFilter = fmt.Sprintf("%s.pending.>", OffersSubject)

type Consumer struct {
	Js          jetstream.JetStream
	Ctx         context.Context
	Config      jetstream.ConsumerConfig
	ConsumerCtx jetstream.Consumer
}

type OffersConsumer struct {
	ConsumerCfg   *Consumer
	OffersUseCase *offerUseCase.OfferUseCase
}

type OrdersConsumer struct {
	ConsumerCfg    *Consumer
	PaymentUseCase *usecase.PaymentUseCase
}

func NewOrdersConsumer(name, durable, filterSubject string, paymentUseCase *usecase.PaymentUseCase) *OrdersConsumer {
	return &OrdersConsumer{
		ConsumerCfg: &Consumer{
			Js:  JS,
			Ctx: context.Background(),
			Config: jetstream.ConsumerConfig{
				Name:          name,
				Durable:       durable,
				FilterSubject: filterSubject,
				AckPolicy:     jetstream.AckExplicitPolicy,
				DeliverPolicy: jetstream.DeliverAllPolicy,
			},
		},
		PaymentUseCase: paymentUseCase,
	}
}

func NewOffersConsumer(name, durable, filterSubject string, offersUseCase *offerUseCase.OfferUseCase) *OffersConsumer {
	return &OffersConsumer{
		ConsumerCfg: &Consumer{
			Js:  JS,
			Ctx: context.Background(),
			Config: jetstream.ConsumerConfig{
				Name:          name,
				Durable:       durable,
				FilterSubject: filterSubject,
				AckPolicy:     jetstream.AckExplicitPolicy,
				DeliverPolicy: jetstream.DeliverAllPolicy,
			},
		},
		OffersUseCase: offersUseCase,
	}
}

func ConsumerInit() {
	repositoryOrders := repository.NewOrderRepositoryMysql(config.DB)
	paymentUseCase := usecase.NewPaymentUseCase(repositoryOrders)

	ordersConsumer := NewOrdersConsumer("order_processor", "order_processor", OrdersFilter, paymentUseCase)
	ordersConsumer.ConsumerCfg.CreateStream(OrdersSubject)
	ordersConsumer.HandlingNewOrders()

	repositoryOffers := repository.NewOfferRepositoryMysql(config.DB)
	offersUseCase := offerUseCase.NewOfferUseCase(repositoryOffers)

	offersConsumer := NewOffersConsumer("offer_answer_processor", "offer_answer_processor", OffersAcceptedFilter, offersUseCase)
	offersConsumer.ConsumerCfg.CreateStream(OffersSubject)
	offersConsumer.HandlingOffers()

	config.Logger.Info(
		"consumers successfully initialized",
	)
}

func (c *OrdersConsumer) HandlingNewOrders() {
	_, err := c.ConsumerCfg.ConsumerCtx.Consume(func(msg jetstream.Msg) {
		orderID := getIdFromMsg(msg, OrdersSubject)
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

func (o *OffersConsumer) HandlingOffers() {
	_, err := o.ConsumerCfg.ConsumerCtx.Consume(func(msg jetstream.Msg) {
		offerId := getIdFromMsg(msg, OffersAcceptedFilter)
		msg.Ack()

		config.Logger.Info(
			"new offer accepted",
			zap.String("offer id", offerId),
		)

		offer, _ := o.OffersUseCase.OfferRepository.GetById(offerId)
		offer.SetAcceptedStatus()
		o.OffersUseCase.OfferRepository.UpdateById(offer)

		offer.Product.SetOfferPrice(offer.Price)
		offerProductJson, _ := json.Marshal(offer.Product)

		go redis.Redis.Set(context.Background(), redis.GetOfferCacheId(offer.BuyerID, offer.ProductID), string(offerProductJson), config.Env.OfferExpiresAt)

		config.Logger.Info(
			"order status updated to accepted",
			zap.String("offer id", offerId),
		)
	})

	if err != nil {
		config.Logger.Fatal("err", zap.Error(err))
	}
}

func (c *Consumer) CreateStream(subject string) {
	stream, err := c.Js.Stream(c.Ctx, subject)
	if err != nil {
		config.Logger.Fatal("err", zap.Error(err))
	}

	c.ConsumerCtx, err = stream.CreateOrUpdateConsumer(c.Ctx, c.Config)
	if err != nil {
		config.Logger.Fatal("err", zap.Error(err))
	}
}

func getIdFromMsg(msg jetstream.Msg, filter string) string {
	return strings.Replace(string(msg.Subject()), fmt.Sprintf("%s.", filter), "", 1)
}
