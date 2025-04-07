package configroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/handlers"
	"github.com/yrnThiago/api-server-go/internal/infra/repository"
	"github.com/yrnThiago/api-server-go/internal/routes"
	"github.com/yrnThiago/api-server-go/internal/usecase"
)

func OrderRouter() *fiber.App {
	repositoryOrders := repository.NewOrderRepositoryMysql(config.DB)
	orderUseCase := usecase.NewOrderUseCase(repositoryOrders)
	orderHandlers := handlers.NewOrderHandlers(orderUseCase)
	orderRouter := routes.NewOrderRouter(orderHandlers)

	return orderRouter.GetRouter()
}
