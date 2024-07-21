package account

import (
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Provide(
		fx.Annotate(
			NewConnection,
			fx.ResultTags(`name:"grpc_account_connection"`),
			fx.OnStop(OnStop),
		),
	)
}
