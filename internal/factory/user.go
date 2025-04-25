package factory

import (
	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/handlers"
	"github.com/yrnThiago/api-server-go/internal/infra/repository"
	"github.com/yrnThiago/api-server-go/internal/routes"
	usecase "github.com/yrnThiago/api-server-go/internal/usecase/user"
)

type UserFactory struct {
	Repository usecase.IUserRepository
	Usecase    *usecase.UserUseCase
	Handler    *handlers.UserHandlers
	Router     *routes.UserRouter
}

func NewUserFactory() *UserFactory {
	userRepository := NewUserRepository()
	userUseCase := NewUserUseCase(userRepository)
	userHandlers := NewUserHandlers(userUseCase)
	userRouter := NewUserRouter(userHandlers)

	return &UserFactory{
		Repository: userRepository,
		Usecase:    userUseCase,
		Handler:    userHandlers,
		Router:     userRouter,
	}
}

func NewUserRepository() usecase.IUserRepository {
	return repository.NewUserRepositoryMysql(config.DB)
}

func NewUserUseCase(repo usecase.IUserRepository) *usecase.UserUseCase {
	return usecase.NewUserUseCase(repo)
}

func NewUserHandlers(usecase *usecase.UserUseCase) *handlers.UserHandlers {
	return handlers.NewUserHandlers(usecase)
}

func NewUserRouter(handlers *handlers.UserHandlers) *routes.UserRouter {
	return routes.NewUserRouter(handlers)
}
