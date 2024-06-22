package metrichandler

import (
	"github.com/gofiber/fiber/v2"
)

type HealthHandler interface {
	Check(ctx *fiber.Ctx) error
}

type healthHandler struct {
}

func NewHealthHandler() HealthHandler {
	return healthHandler{}
}

func (h healthHandler) Check(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status": "ok",
	})
}
