package tenant

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Port struct {
	handler *Handler
}

type PortParams struct {
	fx.In
	Handler *Handler
}

func NewPort(params PortParams) *Port {
	return &Port{
		handler: params.Handler,
	}
}

func (r *Port) Register(router fiber.Router) {
	router.Add(fiber.MethodGet, "", r.handler.ListTenants)
	router.Add(fiber.MethodGet, "/:id", r.handler.GetTenant)
	router.Add(fiber.MethodPost, "", r.handler.CreateTenant)
	router.Add(fiber.MethodPut, "/:id", r.handler.UpdateTenant)
}

func (r *Port) Path() string {
	return "/tenants"
}
