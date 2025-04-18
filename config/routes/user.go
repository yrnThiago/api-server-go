package configroutes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/handlers"
	"github.com/yrnThiago/api-server-go/internal/infra/repository"
	"github.com/yrnThiago/api-server-go/internal/routes"
	"github.com/yrnThiago/api-server-go/internal/usecase/user"
)

func UserRouter() *fiber.App {
	repositoryUsers := repository.NewUserRepositoryMysql(config.DB)
	userUseCase := usecase.NewUserUseCase(repositoryUsers)
	userHandlers := handlers.NewUserHandlers(userUseCase)
	userRouter := routes.NewUserRouter(userHandlers)

	return userRouter.GetRouter()
}
