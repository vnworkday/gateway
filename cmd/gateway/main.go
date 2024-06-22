package main

import (
	"github.com/vnworkday/gateway/internal/handlers"
	"github.com/vnworkday/gateway/internal/http"
	"github.com/vnworkday/gateway/internal/routes"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(zap.NewProduction),
		fx.Decorate(func(logger *zap.Logger) *zap.Logger {
			return logger.WithLazy(zap.String("service", "gateway"))
		}),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.NopLogger
		}),
		handlers.Module,
		routes.Module,
		http.Module,
		fx.Invoke(func(server http.Server) {}),
	).Run()
}
