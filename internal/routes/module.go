package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vnworkday/gateway/internal/routes/account"
	"github.com/vnworkday/gateway/internal/routes/misc"
	"github.com/vnworkday/gateway/internal/utils"
	"go.uber.org/fx"
)

type Router interface {
	Register(router fiber.Router)
	Path() string
}

func Register() fx.Option {
	return fx.Provide(
		// Misc routers
		registerRouter(misc.NewHealthRouter),
		registerRouter(misc.NewSwaggerRouter),

		// Account routers
		registerRouter(account.NewTenantRouter),
	)
}

func registerRouter(f any) any {
	return utils.FxGroupedRegister(f, "routers", new(Router))
}
