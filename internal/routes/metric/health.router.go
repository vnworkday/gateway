package metricrouter

import (
	"github.com/gofiber/fiber/v2"
	metrichandler "github.com/vnworkday/gateway/internal/handlers/metric"
	"github.com/vnworkday/gateway/internal/http"
	"go.uber.org/fx"
)

type HealthRouter interface {
	http.Router
}

type HealthRouterParams struct {
	fx.In
	Handler metrichandler.HealthHandler `name:"health"`
}

func NewHealthRouter(params HealthRouterParams) HealthRouter {
	return healthRouter{
		handler: params.Handler,
	}
}

type healthRouter struct {
	handler metrichandler.HealthHandler
}

func (h healthRouter) Register(router fiber.Router) {
	router.Add(fiber.MethodGet, "", h.handler.Check)
}

func (h healthRouter) Path() string {
	return "/health"
}
