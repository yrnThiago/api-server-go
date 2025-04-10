package configroutes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/handlers"
	"github.com/yrnThiago/api-server-go/internal/infra/repository"
	"github.com/yrnThiago/api-server-go/internal/routes"
	"github.com/yrnThiago/api-server-go/internal/usecase"
)

func AuthRouter() *fiber.App {
	userRepository := repository.NewUserRepositoryMysql(config.DB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	authHandlers := handlers.NewAuthHandlers(userUseCase)
	authRouter := routes.NewAuthRouter(authHandlers)

	return authRouter.GetRouter()
}
