package routes

import (
	"github.com/vnworkday/gateway/internal/http"
	accountrouter "github.com/vnworkday/gateway/internal/routes/account"
	miscrouter "github.com/vnworkday/gateway/internal/routes/misc"
	"github.com/vnworkday/gateway/internal/utils"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Provide(
		// Miscellaneous routers
		registerRouter(miscrouter.NewHealthRouter),
		registerRouter(miscrouter.NewSwaggerRouter),

		// Account routers
		registerRouter(accountrouter.NewTenantRouter),
	)
}

func registerRouter(f any) any {
	return utils.FxGroupedRegister(f, "routers", new(http.Router))
}
