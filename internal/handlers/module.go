package handlers

import (
	"github.com/vnworkday/common/pkg/ioc"
	"github.com/vnworkday/gateway/internal/handlers/account"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Provide(
		// Account handlers
		ioc.RegisterWithName(account.NewTenantHandler, "tenant"),
	)
}
