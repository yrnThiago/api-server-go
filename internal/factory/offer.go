package factory

import (
	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/handlers"
	"github.com/yrnThiago/api-server-go/internal/infra/repository"
	"github.com/yrnThiago/api-server-go/internal/routes"
	usecase "github.com/yrnThiago/api-server-go/internal/usecase/offer"
)

type OfferFactory struct {
	Repository usecase.IOfferRepository
	Usecase    *usecase.OfferUseCase
	Handler    *handlers.OfferHandlers
	Router     *routes.OfferRouter
}

func NewOfferFactory() *OfferFactory {
	repositoryOffers := NewOfferRepository()
	offerUseCase := NewOfferUseCase(repositoryOffers)
	offerHandlers := NewOfferHandlers(offerUseCase)
	offerRouter := NewOfferRouter(offerHandlers)

	return &OfferFactory{
		Repository: repositoryOffers,
		Usecase:    offerUseCase,
		Handler:    offerHandlers,
		Router:     offerRouter,
	}
}

func NewOfferRepository() usecase.IOfferRepository {
	return repository.NewOfferRepositoryMysql(config.DB)
}

func NewOfferUseCase(repo usecase.IOfferRepository) *usecase.OfferUseCase {
	return usecase.NewOfferUseCase(repo)
}

func NewOfferHandlers(usecase *usecase.OfferUseCase) *handlers.OfferHandlers {
	return handlers.NewOfferHandlers(usecase)
}

func NewOfferRouter(handlers *handlers.OfferHandlers) *routes.OfferRouter {
	return routes.NewOfferRouter(handlers)
}
