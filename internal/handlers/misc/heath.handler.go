package misc

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
	// Add any dependencies here
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Readiness(ctx *fiber.Ctx) error {
	return errors.Join(ctx.JSON(fiber.Map{
		"status": "ok",
	}))
}
