package account

import (
	grpcutil "github.com/vnworkday/gateway/internal/common/grpc"
	"github.com/vnworkday/gateway/internal/conf"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

type ConnectionParams struct {
	fx.In
	Config *conf.Conf
	Logger *zap.Logger
}

func NewConnection(params ConnectionParams) *grpc.ClientConn {
	cred, err := grpcutil.ClientCredentials("", params.Config.GRPCAccountServiceURI, params.Config.Profile)
	if err != nil {
		params.Logger.Panic("grpc: failed to create client credentials", zap.Error(err))
	}

	conn, err := grpc.NewClient(params.Config.GRPCAccountServiceURI,
		grpc.WithTransportCredentials(cred),
		grpc.WithChainUnaryInterceptor(
			grpcutil.WithLoggingInterceptor(params.Logger),
			grpcutil.WithRetryInterceptor(),
		),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(grpcutil.ClientMaxMessageSize(params.Config.GRPCMaxMessageSizeMB)),
			grpc.UseCompressor(gzip.Name),
		),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		grpc.WithKeepaliveParams(grpcutil.ClientKeepaliveParams(
			params.Config.GRPCKeepaliveTime,
			params.Config.GRPCKeepaliveTimeout,
		)),
	)
	if err != nil {
		params.Logger.Panic("grpc: failed to create grpc connection", zap.Error(err))
	}

	return conn
}

type OnStopParams struct {
	fx.In
	Conn *grpc.ClientConn `name:"grpc_account_connection"`
}

func OnStop(params OnStopParams) error {
	return params.Conn.Close()
}
