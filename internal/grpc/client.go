package grpc

import (
	"context"
	"time"

	"github.com/vnworkday/gateway/internal/conf"

	"buf.build/gen/go/ntduycs/vnworkday/grpc/go/account/tenant/v1/tenantv1grpc"
	tenantv1 "buf.build/gen/go/ntduycs/vnworkday/protocolbuffers/go/account/tenant/v1"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type TenantClientParams struct {
	fx.In
	Config *conf.Conf
	Conn   *grpc.ClientConn
}

type TenantClient struct {
	client  tenantv1grpc.TenantServiceClient
	timeout time.Duration
}

func NewTenantClient(params TenantClientParams) *TenantClient {
	return &TenantClient{
		client:  tenantv1grpc.NewTenantServiceClient(params.Conn),
		timeout: params.Config.GRPCCallTimeout,
	}
}

func (c *TenantClient) CreateTenant(
	ctx context.Context,
	in *tenantv1.CreateTenantRequest,
	md metadata.MD,
	opts ...grpc.CallOption,
) (*tenantv1.CreateTenantResponse, error) {
	grpcCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	grpcCtx = metadata.NewOutgoingContext(grpcCtx, md)

	return c.client.CreateTenant(grpcCtx, in, opts...)
}

func (c *TenantClient) GetTenant(
	ctx context.Context,
	in *tenantv1.GetTenantRequest,
	md metadata.MD,
	opts ...grpc.CallOption,
) (*tenantv1.GetTenantResponse, error) {
	grpcCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	grpcCtx = metadata.NewOutgoingContext(grpcCtx, md)

	return c.client.GetTenant(grpcCtx, in, opts...)
}

func (c *TenantClient) ListTenants(
	ctx context.Context,
	in *tenantv1.ListTenantsRequest,
	md metadata.MD,
	opts ...grpc.CallOption,
) (*tenantv1.ListTenantsResponse, error) {
	grpcCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	grpcCtx = metadata.NewOutgoingContext(grpcCtx, md)

	return c.client.ListTenants(grpcCtx, in, opts...)
}

func (c *TenantClient) UpdateTenant(
	ctx context.Context,
	in *tenantv1.UpdateTenantRequest,
	md metadata.MD,
	opts ...grpc.CallOption,
) (*tenantv1.UpdateTenantResponse, error) {
	grpcCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	grpcCtx = metadata.NewOutgoingContext(grpcCtx, md)

	return c.client.UpdateTenant(grpcCtx, in, opts...)
}
