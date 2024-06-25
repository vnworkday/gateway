package handlers

import (
	"github.com/vnworkday/gateway/internal/handlers/account"
	"github.com/vnworkday/gateway/internal/utils"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Provide(
		// Account handlers
		utils.FxNamedRegister(account.NewTenantHandler, "tenant"),
	)
}
