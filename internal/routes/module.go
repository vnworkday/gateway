package routes

import (
	"github.com/vnworkday/gateway/internal/http"
	accountrouter "github.com/vnworkday/gateway/internal/routes/account"
	metricrouter "github.com/vnworkday/gateway/internal/routes/metric"
	"github.com/vnworkday/gateway/internal/utils"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Provide(
		// Metric routers
		utils.FxGroupedRegister(metricrouter.NewHealthRouter, "routers", new(http.Router)),

		// Account routers
		utils.FxGroupedRegister(accountrouter.NewTenantRouter, "routers", new(http.Router)),
	)
}
