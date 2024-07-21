package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vnworkday/common/pkg/log"

	"github.com/vnworkday/gateway/internal/client"
	"github.com/vnworkday/gateway/internal/conf"
	"github.com/vnworkday/gateway/internal/logger"
	"github.com/vnworkday/gateway/internal/server"
	"github.com/vnworkday/gateway/internal/usecase"
	"go.uber.org/fx"
)

func Start() {
	fx.New(
		conf.Register(),
		logger.Register(),
		usecase.Register(),
		client.Register(),
		server.Register(),
		fx.WithLogger(log.NewFxEvent),
		fx.Invoke(func(*fiber.App) {}),
	).Run()
}
