package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"slices"
	"time"

	"github.com/pkg/errors"
	"github.com/vnworkday/gateway/internal/conf"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/keepalive"
)

type ConnectionParams struct {
	fx.In
	Config *conf.Conf
	Logger *zap.Logger
}

func NewConnection(params ConnectionParams, targetURI string) *grpc.ClientConn {
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

	return conn
}

func clientKeepaliveParams(
	keepaliveTime int,
	keepaliveTimeout int,
) keepalive.ClientParameters {
	params := keepalive.ClientParameters{
		Time:                defaultKeepaliveTime * time.Second,
		Timeout:             defaultKeepaliveTimeout * time.Second,
		PermitWithoutStream: true,
	}

	if keepaliveTime > 0 {
		params.Time = time.Duration(keepaliveTime) * time.Second
	}

	if keepaliveTimeout > 0 {
		params.Timeout = time.Duration(keepaliveTimeout) * time.Second
	}

	return params
}

func clientMaxMessageSize(
	maxMessageSize int,
) int {
	if maxMessageSize > 0 {
		return maxMessageSize * sizeMB
	}

	return defaultMaxMessageSize * sizeMB
}

func clientCredentials(
	certPerm string,
	clientAddr string,
	profile string,
) (credentials.TransportCredentials, error) {
	if !shouldUseTLSConnection(clientAddr) {
		return insecure.NewCredentials(), nil
	}

	tlsConfig, err := newTLSConfig(certPerm, profile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create tls Config")
	}

	return credentials.NewTLS(tlsConfig), nil
}

func shouldUseTLSConnection(addr string) bool {
	plaintextHosts := []string{
		"localhost",
		"127.0.0.1",
	}

	return slices.Contains(plaintextHosts, addr)
}

func newTLSConfig(certPerm string, profile string) (*tls.Config, error) {
	tlsConfig := tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	if certPerm == "" {
		switch profile {
		case "local":
			tlsConfig.InsecureSkipVerify = true //nolint:gosec

			return &tlsConfig, nil
		default:
			return nil, errors.New("failed to load tls certificate from secrets")
		}
	}

	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load system cert pool")
	}

	if ok := certPool.AppendCertsFromPEM([]byte(certPerm)); !ok {
		return nil, errors.New("failed to append certs from PEM")
	}

	tlsConfig.RootCAs = certPool
	tlsConfig.ClientCAs = certPool

	return &tlsConfig, nil
}

type OnStopParams struct {
	fx.In
	Conn *grpc.ClientConn `name:"grpc_account_connection"`
}

func OnStop(params OnStopParams) error {
	return params.Conn.Close()
}
