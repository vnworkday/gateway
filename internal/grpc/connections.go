package grpc

import (
	"context"
	"os"

	"github.com/vnworkday/gateway/internal/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

type ConnectionParams struct {
	fx.In
	fx.Lifecycle
	Config *config.Cfg `name:"Config"`
	Logger *zap.Logger
}

func NewAccountConnection(params ConnectionParams) *grpc.ClientConn {
	return newConnection(params, params.Config.GRPCAccountServiceURI)
}

func newConnection(params ConnectionParams, targetURI string) *grpc.ClientConn {
	cfg := params.Config

	cred, err := clientCredentials("", targetURI, os.Getenv("PROFILE"))
	if err != nil {
		params.Logger.Panic("failed to create client credentials", zap.Error(err))
	}

	conn, err := grpc.NewClient(targetURI,
		grpc.WithTransportCredentials(cred),
		grpc.WithChainUnaryInterceptor(
			withLoggingInterceptor(params.Logger),
			withRetryInterceptor(),
		),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(clientMaxMessageSize(cfg.GRPCMaxMessageSizeMB)),
			grpc.UseCompressor(gzip.Name),
		),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		grpc.WithKeepaliveParams(clientKeepaliveParams(cfg.GRPCKeepaliveTime, cfg.GRPCKeepaliveTimeout)),
	)
	if err != nil {
		params.Logger.Panic("failed to create grpc connection", zap.Error(err))
	}

	params.Lifecycle.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return conn
}
