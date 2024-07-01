package config

import (
	"os"

	"github.com/pkg/errors"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type CfgParams struct {
	fx.In
	Logger *zap.Logger
}

func NewConfig(params CfgParams) *viper.Viper {
	conf := viper.New()
	conf.SetDefault("PROFILE", "local")
	conf.SetConfigFile(".env")
	conf.AddConfigPath(".")
	conf.AddConfigPath("..")
	conf.AddConfigPath("$HOME/.config")
	conf.AutomaticEnv()

	if err := conf.ReadInConfig(); err != nil && os.Getenv("PROFILE") == "local" {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			params.Logger.Error("Failed to read config", zap.Error(err))
		} else {
			params.Logger.Panic("Failed to read config", zap.Error(err))
		}
	}

	return conf
}
