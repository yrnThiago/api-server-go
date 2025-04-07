package fiber

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/internal/config"
	configroutes "github.com/yrnThiago/api-server-go/internal/config/routes"
	"github.com/yrnThiago/api-server-go/internal/middlewares"
)

func Init() {
	app := fiber.New()

	// Middlewares globais
	app.Use(middlewares.LoggingMiddleware)
	app.Use(middlewares.ErrorMiddleware)

	// Rotas públicas
	public := app.Group("/")
	public.Mount("/health", configroutes.HealthRouter())
	public.Mount("/auth", configroutes.AuthRouter())

	// Rotas privadas
	private := app.Group("/", middlewares.AuthMiddleware, middlewares.ContextMiddleware)
	private.Mount("/orders", configroutes.OrderRouter())
	private.Mount("/products", configroutes.ProductRouter())

	config.Logger.Info(
		"server listening",
		zap.String("port", config.Env.PORT),
	)

	log.Fatal(app.Listen(":" + config.Env.PORT))
}
