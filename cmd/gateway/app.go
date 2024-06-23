package app

import (
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
		fx.Provide(zap.NewProduction),
		fx.Decorate(func(logger *zap.Logger) *zap.Logger {
			return logger.WithLazy(zap.String("service", "gateway"))
		}),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: logger.WithOptions(zap.IncreaseLevel(zapcore.WarnLevel)),
			}
		}),
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
