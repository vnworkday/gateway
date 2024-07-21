package account

import (
	"github.com/vnworkday/common/pkg/ioc"
	grpcutil "github.com/vnworkday/gateway/internal/common/grpc"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Provide(
		fx.Annotate(
			grpcutil.NewConnection,
			fx.ResultTags(`name:"grpc_account_connection"`),
			fx.OnStop(grpcutil.OnStop),
		),
		ioc.RegisterWithName(NewTenantGRPCClient),
	)
}
