package usecase

import (
	"github.com/vnworkday/gateway/internal/usecase/account/tenant"
	"github.com/vnworkday/gateway/internal/usecase/misc"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Module("use_case",
		tenant.Register(),
		misc.Register(),
	)
}
