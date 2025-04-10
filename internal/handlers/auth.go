package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/keys"
	"github.com/yrnThiago/api-server-go/internal/usecase"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

type AuthHandler struct {
	UserUseCase *usecase.UserUseCase
}

func NewAuthHandlers(createUserUseCase *usecase.UserUseCase) *AuthHandler {
	return &AuthHandler{
		UserUseCase: createUserUseCase,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var userInputDto usecase.UserInputDto
	c.BodyParser(&userInputDto)

	output, err := h.UserUseCase.GetByEmail(userInputDto.Email)
	if err != nil {
		errorInfo := utils.NewErrorInfo(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
		c.Locals(string(keys.ErrorKey), errorInfo)
		return err
	}

	if !utils.CheckPasswordHash(userInputDto.Password, output.Password) {
		config.Logger.Warn("wrong credentials")
		return err
	}

	config.Logger.Info(
		"user logged in",
		zap.String("user id", output.ID),
	)

	authToken, err := utils.GenerateJWT(output.ID)
	userAuthorization := utils.BEARER_KEY + authToken
	if err != nil {
		config.Logger.Warn(
			"jwt token not generated",
			zap.Error(err),
		)
	}

	cookie := &fiber.Cookie{}
	cookie.Name = config.Env.COOKIE_NAME
	cookie.Value = userAuthorization
	cookie.Expires = time.Now().Add(365 * 24 * time.Hour)
	cookie.Secure = false
	cookie.HTTPOnly = true
	cookie.Path = "/"

	utils.SetCookie(c, cookie)
	return c.JSON(TestResponse{"user logged in"})
}
