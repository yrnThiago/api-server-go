package factory

import (
	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/handlers"
	"github.com/yrnThiago/api-server-go/internal/infra/repository"
	"github.com/yrnThiago/api-server-go/internal/routes"
	offerUc "github.com/yrnThiago/api-server-go/internal/usecase/offer"
	usecase "github.com/yrnThiago/api-server-go/internal/usecase/product"
)

type ProductFactory struct {
	Repository usecase.IProductRepository
	Usecase    *usecase.ProductUseCase
	Handler    *handlers.ProductHandlers
	Router     *routes.ProductRouter
}

func NewProductFactory() *ProductFactory {
	repositoryOffers := NewOfferRepository()
	offerUseCase := NewOfferUseCase(repositoryOffers)


	repositoryProducts := NewProductRepository()
	productUseCase := NewProductUseCase(repositoryProducts, offerUseCase)
	productHandlers := NewProductHandlers(productUseCase)
	productRouter := NewProductRouter(productHandlers)

	return &ProductFactory{
		Repository: repositoryProducts,
		Usecase:    productUseCase,
		Handler:    productHandlers,
		Router:     productRouter,
	}
}

func NewProductRepository() usecase.IProductRepository {
	return repository.NewProductRepositoryMysql(config.DB)
}

func NewProductUseCase(repo usecase.IProductRepository, offerUseCase *offerUc.OfferUseCase) *usecase.ProductUseCase {
	return usecase.NewProductUseCase(repo, offerUseCase)
}

func NewProductHandlers(usecase *usecase.ProductUseCase) *handlers.ProductHandlers {
	return handlers.NewProductHandlers(usecase)
}

func NewProductRouter(handlers *handlers.ProductHandlers) *routes.ProductRouter {
	return routes.NewProductRouter(handlers)
}
