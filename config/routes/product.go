package configroutes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/handlers"
	"github.com/yrnThiago/api-server-go/internal/infra/repository"
	"github.com/yrnThiago/api-server-go/internal/routes"
	"github.com/yrnThiago/api-server-go/internal/usecase/product"
)

func ProductRouter() *fiber.App {
	repositoryProducts := repository.NewProductRepositoryMysql(config.DB)
	productUseCase := usecase.NewProductUseCase(repositoryProducts)
	productHandlers := handlers.NewProductHandlers(productUseCase)
	productRouter := routes.NewProductRouter(productHandlers)

	return productRouter.GetRouter()
}
