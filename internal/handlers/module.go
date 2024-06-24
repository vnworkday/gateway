package handlers

import (
	tenanthandler "github.com/vnworkday/gateway/internal/handlers/account"
	mischandler "github.com/vnworkday/gateway/internal/handlers/misc"
	"github.com/vnworkday/gateway/internal/utils"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Provide(
		// Metric handlers
		utils.FxNamedRegister(mischandler.NewHealthHandler, "health"),

		// Account handlers
		utils.FxNamedRegister(tenanthandler.NewTenantHandler, "tenant"),
	)
}
