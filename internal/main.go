package main

import (
	"fmt"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"

	"github.com/yrnThiago/gdlp-go/internal/api"
	"github.com/yrnThiago/gdlp-go/internal/cmd/pub"
	"github.com/yrnThiago/gdlp-go/internal/cmd/sub"
	"github.com/yrnThiago/gdlp-go/internal/domain"
	"github.com/yrnThiago/gdlp-go/internal/handlers"
	"github.com/yrnThiago/gdlp-go/internal/infra/repository"
	"github.com/yrnThiago/gdlp-go/internal/usecase"
)

func main() {
	api.LoadEnv()
	api.CreateLogger()

	// Can u please make a proper palce to config NATs
	opts := &server.Options{}
	ns, err := server.NewServer(opts)
	if err != nil {
		panic(err)
	}
	go ns.Start()

	if !ns.ReadyForConnections(4 * time.Second) {
		panic("not ready for connection")
	}

	msgChan := make(chan *nats.Msg)

	pub.PublisherInit()
	sub := sub.Connect()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Maybe this would be better in another place right??
	db.AutoMigrate(&domain.Product{}, &domain.Order{}, &domain.OrderItems{})
	repositoryProducts := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repositoryProducts)
	listProductUseCase := usecase.NewListProductsCase(repositoryProducts)
	productHandlers := handlers.NewProductHandlers(createProductUseCase, listProductUseCase)

	repositoryOrders := repository.NewOrderRepositoryMysql(db)
	createOrderUseCase := usecase.NewCreateOrderUseCase(repositoryOrders)
	listOrderUseCase := usecase.NewListOrdersCase(repositoryOrders)
	orderHandlers := handlers.NewOrderHandlers(createOrderUseCase, listOrderUseCase)

	go api.CreateServer(productHandlers, orderHandlers)

	go sub.ReceiveMessage(msgChan, os.Getenv("NEW_ORDERS_TOPIC"))

	for msg := range msgChan {
		fmt.Println(string(msg.Data))
	}
}
