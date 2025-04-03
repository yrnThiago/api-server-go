package configroutes

import (
	"github.com/go-chi/chi/v5"

	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/handlers"
	"github.com/yrnThiago/api-server-go/internal/infra/repository"
	"github.com/yrnThiago/api-server-go/internal/routes"
	"github.com/yrnThiago/api-server-go/internal/usecase"
)

func ProductRouter() chi.Router {
	repositoryProducts := repository.NewProductRepositoryMysql(config.DB)
	productUseCase := usecase.NewProductUseCase(repositoryProducts)
	productHandlers := handlers.NewProductHandlers(productUseCase)
	productRouter := routes.NewProductRouter(productHandlers)

	return productRouter.GetRouter()
}
