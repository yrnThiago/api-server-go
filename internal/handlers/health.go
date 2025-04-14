package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseMsg struct {
	Message string `json:"message"`
}

type HealthHandler struct{}

func NewHealthHandlers() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Ping(c *fiber.Ctx) error {
	return c.JSON(&ResponseMsg{"pong"})
}
