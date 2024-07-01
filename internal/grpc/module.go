package grpc

import (
	"github.com/vnworkday/gateway/internal/utils"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Provide(
		utils.FxNamedRegister(NewAccountConnection, "grpc_account_connection"),
	)
}
