package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/yrnThiago/gdlp-go/internal/api"
	"github.com/yrnThiago/gdlp-go/internal/cmd/pub"
	"github.com/yrnThiago/gdlp-go/internal/cmd/sub"
	"github.com/yrnThiago/gdlp-go/internal/config"
	"github.com/yrnThiago/gdlp-go/internal/domain"
	"github.com/yrnThiago/gdlp-go/internal/handlers"
	"github.com/yrnThiago/gdlp-go/internal/infra/repository"
	"github.com/yrnThiago/gdlp-go/internal/usecase"
)

func main() {
	config.LoadEnv()
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

	db, err := gorm.Open(mysql.Open(config.GetDatabaseUrl()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Maybe this would be better in another place right??
	db.Migrator().AutoMigrate(&domain.Product{}, &domain.Order{}, &domain.OrderItems{})
	repositoryProducts := repository.NewProductRepositoryMysql(db)
	productUseCase := usecase.NewProductUseCase(repositoryProducts)
	productHandlers := handlers.NewProductHandlers(productUseCase)

	repositoryOrders := repository.NewOrderRepositoryMysql(db)
	orderUseCase := usecase.NewOrderUseCase(repositoryOrders)
	orderHandlers := handlers.NewOrderHandlers(orderUseCase)

	go api.CreateServer(productHandlers, orderHandlers)

	go sub.ReceiveMessage(msgChan, os.Getenv("NEW_ORDERS_TOPIC"))

	for msg := range msgChan {
		var order *domain.Order

		err = json.Unmarshal(msg.Data, &order)
		if err != nil {
			return
		}

		fmt.Println(order)
		repositoryOrders.UpdateById(order, map[string]any{"Status": "Pagamento Aprovado"})
	}
}
