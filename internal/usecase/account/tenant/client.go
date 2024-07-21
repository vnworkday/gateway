package tenant

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
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	timeout time.Duration
	conn    *grpc.ClientConn
}

type GRPCClientParams struct {
	fx.In
	Config *conf.Conf
	Conn   *grpc.ClientConn `name:"grpc_account_connection"`
}

func NewGRPCClient(params GRPCClientParams) *GRPCClient {
	return &GRPCClient{
		timeout: params.Config.GRPCCallTimeout,
		conn:    params.Conn,
	}
}

func (c *GRPCClient) CreateTenant(
	ctx context.Context,
	in *CreateTenantRequest,
	encodeRequest converter.ConvertFunc[CreateTenantRequest, tenantv1.CreateTenantRequest],
	decodeResponse converter.ConvertFunc[tenantv1.CreateTenantResponse, CreateTenantResponse],
	opts ...grpctransport.ClientOption,
) (*CreateTenantResponse, error) {
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

	return util.SafeCast[*CreateTenantResponse](resp), err
}

func (c *GRPCClient) GetTenant(
	ctx context.Context,
	in *GetTenantRequest,
	encodeRequest converter.ConvertFunc[GetTenantRequest, tenantv1.GetTenantRequest],
	decodeResponse converter.ConvertFunc[tenantv1.GetTenantResponse, GetTenantResponse],
	opts ...grpctransport.ClientOption,
) (*GetTenantResponse, error) {
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

	return util.SafeCast[*GetTenantResponse](resp), err
}

func (c *GRPCClient) ListTenants(
	ctx context.Context,
	in *ListTenantsRequest,
	encodeRequest converter.ConvertFunc[ListTenantsRequest, tenantv1.ListTenantsRequest],
	decodeResponse converter.ConvertFunc[tenantv1.ListTenantsResponse, ListTenantsResponse],
	opts ...grpctransport.ClientOption,
) (*ListTenantsResponse, error) {
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
	if err != nil {
		return nil, err
	}

	return util.SafeCast[*ListTenantsResponse](resp), err
}

func (c *GRPCClient) UpdateTenant(
	ctx context.Context,
	in *UpdateTenantRequest,
	encodeRequest converter.ConvertFunc[UpdateTenantRequest, tenantv1.UpdateTenantRequest],
	decodeResponse converter.ConvertFunc[tenantv1.UpdateTenantResponse, UpdateTenantResponse],
	opts ...grpctransport.ClientOption,
) (*UpdateTenantResponse, error) {
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

	return util.SafeCast[*UpdateTenantResponse](resp), err
}
