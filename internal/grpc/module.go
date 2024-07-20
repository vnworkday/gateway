package grpc

import (
	"github.com/vnworkday/common/pkg/ioc"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Provide(
		ioc.RegisterWithName(NewAccountConnection, "grpc_account_connection"),
		ioc.RegisterWithName(NewTenantClient, "grpc_tenant_client"),
	)
}
