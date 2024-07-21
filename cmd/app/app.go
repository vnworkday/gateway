package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vnworkday/common/pkg/log"
	"github.com/vnworkday/gateway/internal/grpc"
	"github.com/vnworkday/gateway/internal/logger"
	"github.com/vnworkday/gateway/internal/server"
	"github.com/vnworkday/gateway/internal/usecase"

	"github.com/vnworkday/gateway/internal/conf"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Start() {
	fx.New(
		conf.Register(),
		logger.Register(),
		usecase.Register(),
		fx.Module("grpc",
			fx.Decorate(func(logger *zap.Logger) *zap.Logger {
				return logger.Named("grpc")
			}),
			grpc.Register(),
		),
		server.Register(),
		fx.WithLogger(log.NewFxEvent),
		fx.Invoke(func(*fiber.App) {}),
	).Run()
}
