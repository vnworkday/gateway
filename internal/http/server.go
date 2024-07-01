package http

import (
	"context"
	"time"


	"github.com/pkg/errors"

	"github.com/spf13/viper"

	"github.com/vnworkday/gateway/internal/routes"

	"github.com/vnworkday/gateway/internal/models/shared"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"go.uber.org/fx"
)

type ServerProps struct {
	fx.In
	fx.Lifecycle
	Logger  *zap.Logger
	Config  *viper.Viper    `name:"config"`
	Routers []routes.Router `group:"routers"`
}

func NewServer(props ServerProps) *fiber.App {
	svr := buildHTTPServer(props.Logger)

	registerRoutes(svr, props.Routers)

	props.Lifecycle.Append(fx.Hook{
		OnStart: onStart(svr),
		OnStop:  onStop(svr),
	})

	return svr
}

func buildHTTPServer(logger *zap.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit:               fiber.DefaultBodyLimit, // 4 * 1024 * 1024 = 4MB
		CaseSensitive:           true,
		StrictRouting:           true,
		Concurrency:             fiber.DefaultConcurrency,
		EnableIPValidation:      true,
		IdleTimeout:             time.Minute, // NOTE: We may want to fine tune this later
		EnableTrustedProxyCheck: false,       // NOTE: We may want to enable this later
		TrustedProxies:          []string{},  // NOTE: We may want to set this later after enabling EnableTrustedProxyCheck
		ProxyHeader:             "",          // NOTE: We may want to set this later after enabling EnableTrustedProxyCheck

		RequestMethods: []string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
			fiber.MethodOptions,
		},

		// The system is designed for low latency.
		// Exceeding 1 second should be considered as a performance issue and requires further investigation.
		ReadTimeout:  time.Second, // NOTE: We may want to fine tune this later
		WriteTimeout: time.Second, // NOTE: We may want to fine tune this later

		ReduceMemoryUsage: false, // NOTE: We may want to enable this later if we have memory issues.

		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := shared.CodeErrInternal
			var e *fiber.Error
			if errors.As(err, &e) {
				code = shared.FromFiberError(e)
			}

			if shared.IsClientError(code) {
				logger.Info("Client error", zap.Error(err))
			} else {
				logger.Error("Server error", zap.Error(err))
			}

			ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

			return ctx.Status(shared.ToHTTPStatus(code)).JSON(shared.NewError(code))
		},

		EnablePrintRoutes: true,
	})

	return app
}

func registerRoutes(rootRouter fiber.Router, routers []routes.Router) {
	for _, router := range routers {
		path := "/api/v1" + router.Path()
		rootRouter.Route(path, router.Register)
	}
}

func onStart(server *fiber.App) func(context.Context) error {
	return func(_ context.Context) error {
		go func() {
			if err := server.Listen(":3000"); err != nil {
				panic(err)
			}
		}()

		return nil
	}
}

func onStop(server *fiber.App) func(context.Context) error {
	return func(_ context.Context) error {
		return server.Shutdown()
	}
}
