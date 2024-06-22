package metricrouter

import (
	"github.com/gofiber/fiber/v2"
	metrichandler "github.com/vnworkday/gateway/internal/handlers/metric"
	"go.uber.org/fx"
)

type HealthRouterParams struct {
	fx.In
	Handler *metrichandler.HealthHandler `name:"health"`
}

func NewHealthRouter(params HealthRouterParams) *HealthRouter {
	return &HealthRouter{
		handler: params.Handler,
	}
}

type HealthRouter struct {
	handler *metrichandler.HealthHandler
}

func (h *HealthRouter) Register(router fiber.Router) {
	router.Add(fiber.MethodGet, "", h.handler.Check)
}

func (h *HealthRouter) Path() string {
	return "/health"
}
