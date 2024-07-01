package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"slices"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

const (
	defaultMaxMessageSize = 4

	defaultKeepaliveTime    = 10
	defaultKeepaliveTimeout = 5

	defaultBackoffInterval = 100

	sizeMB = 1 << (10 * 2) //nolint:mnd
)

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
