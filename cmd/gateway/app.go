package app

import (
	"os"

	"github.com/vnworkday/gateway/internal/config"
	"github.com/vnworkday/gateway/internal/handlers"
	"github.com/vnworkday/gateway/internal/http"
	"github.com/vnworkday/gateway/internal/routes"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Start() {
	fx.New(
		fx.Provide(zapLogger()),
		fx.WithLogger(fxEventLogger()),
		fx.Decorate(func(logger *zap.Logger) *zap.Logger {
			return logger.WithLazy(zap.String("service", "gateway"))
		}),
		fx.Module("config",
			fx.Decorate(func(logger *zap.Logger) *zap.Logger {
				return logger.Named("config")
			}),
			config.Register(),
		),
		fx.Module("handlers",
			fx.Decorate(func(logger *zap.Logger) *zap.Logger {
				return logger.Named("handlers")
			}),
			handlers.Register(),
		),
		fx.Module("routers",
			fx.Decorate(func(logger *zap.Logger) *zap.Logger {
				return logger.Named("routers")
			}),
			routes.Register(),
		),
		fx.Invoke(http.NewServer),
	).Run()
}

func fxEventLogger() any {
	if os.Getenv("PROFILE") != "local" {
		return func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: logger.WithOptions(zap.IncreaseLevel(zapcore.WarnLevel)),
			}
		}
	}

	return func() fxevent.Logger {
		return &fxevent.ConsoleLogger{
			W: os.Stdout,
		}
	}
}

func zapLogger() func(options ...zap.Option) (*zap.Logger, error) {
	if os.Getenv("PROFILE") == "local" {
		return zap.NewDevelopment
	}

	return zap.NewProduction
}
