package fiber

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/internal/config"
	configroutes "github.com/yrnThiago/api-server-go/internal/config/routes"
	"github.com/yrnThiago/api-server-go/internal/middlewares"
)

func Init() {
	app := fiber.New()

	app.Use(requestid.New(requestid.Config{
		Next:       nil,
		Header:     fiber.HeaderXRequestID,
		Generator:  utils.UUID,
		ContextKey: "requestid",
	}))

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
