package config

import (
	"os"
	"time"

	"github.com/vnworkday/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type CfgParams struct {
	fx.In
	Logger *zap.Logger
}

type Cfg struct {
	AppName string `config:"app_name"`

	GRPCMaxMessageSizeMB int           `config:"grpc_max_message_size_mb"`
	GRPCKeepaliveTime    int           `config:"grpc_keepalive_time"`
	GRPCKeepaliveTimeout int           `config:"grpc_keepalive_timeout"`
	GRPCCallTimeout      time.Duration `config:"grpc_call_timeout"`

	GRPCAccountServiceURI string `config:"grpc_account_service_uri"`
}

func NewConfig(params CfgParams) *Cfg {
	cfg := new(Cfg)

	cfgBuilder := config.FromEnv()

	if os.Getenv("PROFILE") == "local" {
		cfgBuilder.FromFile(".env")
	}

	if err := cfgBuilder.MapTo(cfg); err != nil {
		params.Logger.Panic("Failed to read config", zap.Error(err))
	}

	return cfg
}
