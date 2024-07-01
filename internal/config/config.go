package config

import (
	"os"

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
