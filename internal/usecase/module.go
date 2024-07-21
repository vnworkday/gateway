package usecase

import (
	"github.com/vnworkday/gateway/internal/usecase/account/tenant"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Module("use_case",
		tenant.Register(),
	)
}
