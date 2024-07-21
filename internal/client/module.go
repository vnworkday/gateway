package client

import (
	"github.com/vnworkday/gateway/internal/client/grpc"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Module("client",
		grpc.Register(),
	)
}
