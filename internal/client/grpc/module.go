package grpc

import (
	"github.com/vnworkday/gateway/internal/client/grpc/account"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Options(
		account.Register(),
	)
}
