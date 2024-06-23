package handlers

import (
	tenanthandler "github.com/vnworkday/gateway/internal/handlers/account"
	metrichandler "github.com/vnworkday/gateway/internal/handlers/metric"
	"github.com/vnworkday/gateway/internal/utils"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Provide(
		// Metric handlers
		utils.FxNamedRegister(metrichandler.NewHealthHandler, "health"),

		// Account handlers
		utils.FxNamedRegister(tenanthandler.NewTenantHandler, "tenant"),
	)
}
