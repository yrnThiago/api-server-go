package factory

import (
	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/handlers"
	"github.com/yrnThiago/api-server-go/internal/infra/repository"
	"github.com/yrnThiago/api-server-go/internal/routes"
	orderUC "github.com/yrnThiago/api-server-go/internal/usecase/order"
	productUC "github.com/yrnThiago/api-server-go/internal/usecase/product"
)

type OrderFactory struct {
	Repository orderUC.IOrderRepository
	Usecase    *orderUC.OrderUseCase
	Handler    *handlers.OrderHandlers
	Router     *routes.OrderRouter
}

func NewOrderFactory() *OrderFactory {
	productFactory := NewProductFactory()
	productRepository := productFactory.Repository

	repositoryOrders := NewOrderRepository()
	orderUseCase := NewOrderUseCase(repositoryOrders, productRepository)
	orderHandlers := NewOrderHandlers(orderUseCase)
	orderRouter := NewOrderRouter(orderHandlers)

	return &OrderFactory{
		Repository: repositoryOrders,
		Usecase:    orderUseCase,
		Handler:    orderHandlers,
		Router:     orderRouter,
	}
}

func NewOrderRepository() orderUC.IOrderRepository {
	return repository.NewOrderRepositoryMysql(config.DB)
}

func NewOrderUseCase(repo orderUC.IOrderRepository, productRepo productUC.IProductRepository) *orderUC.OrderUseCase {
	return orderUC.NewOrderUseCase(repo, productRepo)
}

func NewOrderHandlers(usecase *orderUC.OrderUseCase) *handlers.OrderHandlers {
	return handlers.NewOrderHandlers(usecase)
}

func NewOrderRouter(handlers *handlers.OrderHandlers) *routes.OrderRouter {
	return routes.NewOrderRouter(handlers)
}
