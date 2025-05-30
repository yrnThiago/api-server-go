package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	configroutes "github.com/yrnThiago/api-server-go/config/routes"
	"github.com/yrnThiago/api-server-go/internal/middlewares"
)

func Init() {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: middlewares.ErrorMiddleware,
		},
	)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://localhost",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	app.Use(requestid.New(requestid.Config{
		Next:       nil,
		Header:     fiber.HeaderXRequestID,
		Generator:  utils.UUID,
		ContextKey: "requestid",
	}))

	// Global middlewares
	app.Use(middlewares.LoggingMiddleware)

	// Public routes
	public := app.Group("/public")
	public.Mount("/health", configroutes.HealthRouter())
	public.Mount("/auth", configroutes.AuthRouter())

	// Private routes
	private := app.Group("/private", middlewares.AuthMiddleware, middlewares.RateLimitMiddleware)
	private.Mount("/orders", configroutes.OrderRouter())
	private.Mount("/offers", configroutes.OfferRouter())
	private.Mount("/products", configroutes.ProductRouter())
	private.Mount("/users", configroutes.UserRouter())

	// Global not found middleware
	app.Use(middlewares.NotFoundMiddleware)

	config.Logger.Info(
		"server listening",
		zap.String("port", config.Env.Port),
	)

	log.Fatal(app.Listen(":" + config.Env.Port))
}
