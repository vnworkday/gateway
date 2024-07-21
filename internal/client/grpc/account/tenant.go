package account

import (
	"context"
	"time"

	"buf.build/gen/go/ntduycs/vnworkday/grpc/go/account/tenant/v1/tenantv1grpc"
	tenantv1 "buf.build/gen/go/ntduycs/vnworkday/protocolbuffers/go/account/tenant/v1"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/vnworkday/gateway/internal/common/converter"
	"github.com/vnworkday/gateway/internal/common/port"
	"github.com/vnworkday/gateway/internal/common/util"
	"github.com/vnworkday/gateway/internal/conf"
	"github.com/vnworkday/gateway/internal/usecase/account/tenant"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type TenantGRPCClient struct {
	timeout time.Duration
	conn    *grpc.ClientConn
}

type TenantGRPCClientParams struct {
	fx.In
	Config *conf.Conf
	Conn   *grpc.ClientConn `name:"grpc_account_connection"`
}

func NewTenantGRPCClient(params TenantGRPCClientParams) *TenantGRPCClient {
	return &TenantGRPCClient{
		timeout: params.Config.GRPCCallTimeout,
		conn:    params.Conn,
	}
}

func (c *TenantGRPCClient) CreateTenant(
	ctx context.Context,
	in *tenant.CreateTenantRequest,
	encodeRequest converter.ConvertFunc[tenant.CreateTenantRequest, tenantv1.CreateTenantRequest],
	decodeResponse converter.ConvertFunc[tenantv1.CreateTenantResponse, *tenant.CreateTenantResponse],
	opts ...grpctransport.ClientOption,
) (*tenant.CreateTenantResponse, error) {
	gRPCCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	resp, err := port.NewGRPCClient(
		c.conn,
		tenantv1grpc.TenantService_ServiceDesc.ServiceName,
		tenantv1grpc.TenantService_CreateTenant_FullMethodName,
		encodeRequest,
		decodeResponse,
		opts...,
	)(gRPCCtx, in)

	return util.SafeCast[*tenant.CreateTenantResponse](resp), err
}

func (c *TenantGRPCClient) GetTenant(
	ctx context.Context,
	in *tenant.GetTenantRequest,
	encodeRequest converter.ConvertFunc[tenant.GetTenantRequest, tenantv1.GetTenantRequest],
	decodeResponse converter.ConvertFunc[tenantv1.GetTenantResponse, *tenant.GetTenantResponse],
	opts ...grpctransport.ClientOption,
) (*tenant.GetTenantResponse, error) {
	gRPCCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	resp, err := port.NewGRPCClient(
		c.conn,
		tenantv1grpc.TenantService_ServiceDesc.ServiceName,
		tenantv1grpc.TenantService_GetTenant_FullMethodName,
		encodeRequest,
		decodeResponse,
		opts...,
	)(gRPCCtx, in)

	return util.SafeCast[*tenant.GetTenantResponse](resp), err
}

func (c *TenantGRPCClient) ListTenants(
	ctx context.Context,
	in *tenant.ListTenantsRequest,
	encodeRequest converter.ConvertFunc[tenant.ListTenantsRequest, tenantv1.ListTenantsRequest],
	decodeResponse converter.ConvertFunc[tenantv1.ListTenantsResponse, *tenant.ListTenantsResponse],
	opts ...grpctransport.ClientOption,
) (*tenant.ListTenantsResponse, error) {
	gRPCCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	resp, err := port.NewGRPCClient(
		c.conn,
		tenantv1grpc.TenantService_ServiceDesc.ServiceName,
		tenantv1grpc.TenantService_ListTenants_FullMethodName,
		encodeRequest,
		decodeResponse,
		opts...,
	)(gRPCCtx, in)

	return util.SafeCast[*tenant.ListTenantsResponse](resp), err
}

func (c *TenantGRPCClient) UpdateTenant(
	ctx context.Context,
	in *tenant.UpdateTenantRequest,
	encodeRequest converter.ConvertFunc[tenant.UpdateTenantRequest, tenantv1.UpdateTenantRequest],
	decodeResponse converter.ConvertFunc[tenantv1.UpdateTenantResponse, *tenant.UpdateTenantResponse],
	opts ...grpctransport.ClientOption,
) (*tenant.UpdateTenantResponse, error) {
	gRPCCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	resp, err := port.NewGRPCClient(
		c.conn,
		tenantv1grpc.TenantService_ServiceDesc.ServiceName,
		tenantv1grpc.TenantService_UpdateTenant_FullMethodName,
		encodeRequest,
		decodeResponse,
		opts...,
	)(gRPCCtx, in)

	return util.SafeCast[*tenant.UpdateTenantResponse](resp), err
}
