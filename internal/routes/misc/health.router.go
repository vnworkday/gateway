package misc

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type HealthRouter struct{}

func NewHealthRouter() *HealthRouter {
	return &HealthRouter{}
}

func (h *HealthRouter) Register(router fiber.Router) {
	router.Add(fiber.MethodGet, "", func(ctx *fiber.Ctx) error {
		return errors.Join(ctx.JSON(fiber.Map{
			"status": "ok",
		}))
	})
}

func (h *HealthRouter) Path() string {
	return "/health"
}
