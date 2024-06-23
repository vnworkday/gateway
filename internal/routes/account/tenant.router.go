package accountrouter

import (
	"github.com/gofiber/fiber/v2"
	accounthandler "github.com/vnworkday/gateway/internal/handlers/account"
	"go.uber.org/fx"
)

type TenantRouter struct {
	handler *accounthandler.TenantHandler
}

type TenantRouterParams struct {
	fx.In
	Handler *accounthandler.TenantHandler `name:"tenant"`
}

func NewTenantRouter(params TenantRouterParams) *TenantRouter {
	return &TenantRouter{
		handler: params.Handler,
	}
}

func (r *TenantRouter) Register(router fiber.Router) {
	router.Add(fiber.MethodGet, "", r.handler.ListTenants)
	router.Add(fiber.MethodGet, "/:id", r.handler.GetTenant)
	router.Add(fiber.MethodPost, "", r.handler.CreateTenant)
	router.Add(fiber.MethodPut, `/:id<minLen(1)\>`, r.handler.UpdateTenant)
}

func (r *TenantRouter) Path() string {
	return "/tenants"
}
