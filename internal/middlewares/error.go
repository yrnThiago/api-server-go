package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
	"github.com/yrnThiago/api-server-go/internal/utils"
	"go.uber.org/zap"
)

func ErrorMiddleware(c *fiber.Ctx) error {
	// Executa o próximo handler/middleware
	err := c.Next()

	// Verifica se foi registrado um erro no contexto
	contextErrorVal := c.Locals(string(keys.ErrorKey))
	if contextErrorVal != nil {
		contextError := contextErrorVal.(*utils.ErrorInfo)

		// Log do erro
		config.Logger.Info("error occurred",
			zap.Int("status", contextError.StatusCode),
			zap.String("message", contextError.Message),
		)

		// Responde com JSON de erro
		return c.Status(contextError.StatusCode).JSON(contextError)
	}

	// Se tiver erro retornado por c.Next(), devolve também
	if err != nil {
		return err
	}

	return nil
}
