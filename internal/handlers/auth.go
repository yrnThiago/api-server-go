package handlers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	infra "github.com/yrnThiago/api-server-go/internal/infra/redis"
	"github.com/yrnThiago/api-server-go/internal/usecase/user"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

var WRONG_CREDENTIALS_ERR = "wrong credentials"

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

	output, err := h.UserUseCase.GetByLogin(userInputDto.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": WRONG_CREDENTIALS_ERR})
	}

	if !utils.CheckPasswordHash(userInputDto.Password, output.Password) {
		config.Logger.Warn(WRONG_CREDENTIALS_ERR)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": WRONG_CREDENTIALS_ERR})
	}

	c.Locals(utils.UserIdKeyCtx, output.ID)

	config.Logger.Info(
		"user logged in",
		zap.String("user id", c.Locals(utils.UserIdKeyCtx).(string)),
	)

	authToken, err := utils.GenerateJWT(output.ID)
	if err != nil {
		config.Logger.Fatal(
			"jwt token not generated",
			zap.Error(err),
		)

		return err
	}

	userJson, err := json.Marshal(output)
	if err != nil {
		config.Logger.Fatal(
			"marshal user json",
			zap.Error(err),
		)

		return err
	}

	infra.Redis.Set(context.Background(), "user-"+output.ID, string(userJson), 0)

	cookie := &fiber.Cookie{}
	cookie.Name = config.Env.COOKIE_NAME
	cookie.Value = utils.BEARER_KEY + authToken
	cookie.Expires = time.Now().Add(365 * 24 * time.Hour)
	cookie.Secure = false
	cookie.HTTPOnly = true
	cookie.Path = "/"

	utils.SetCookie(c, cookie)
	return c.JSON(fiber.Map{"message": "user logged in"})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	cookie := &fiber.Cookie{}
	cookie.Name = config.Env.COOKIE_NAME
	cookie.Value = "deleted"
	cookie.Expires = time.Now().Add(-3 * time.Second)
	cookie.HTTPOnly = true

	utils.SetCookie(c, cookie)
	return c.JSON(fiber.Map{"message": "user logout"})
}
