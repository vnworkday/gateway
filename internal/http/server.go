package http

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"time"

	"go.uber.org/fx"
)

// Server defines the interface for an HTTP server.
// It is sealed to allow for easy implementation of the interface.
type Server interface {
	// Listen serves HTTP requests from the given addr.
	Listen(addr string) error
	// Shutdown gracefully shuts down the server without interrupting any active connections.
	Shutdown() error
}

type ServerProps struct {
	fx.In
	fx.Lifecycle
	Routers []Router `group:"routers"`
}

func NewHTTPServer(props ServerProps) Server {
	server := configureServer()

	props.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.Listen(":3000"); err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown()
		},
	})

	return server
}

func configureServer() *fiber.App {
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

		// The system is designed for low latency, exceeding 1 second should be considered as a performance issue and requires further investigation.
		ReadTimeout:  time.Second, // NOTE: We may want to fine tune this later
		WriteTimeout: time.Second, // NOTE: We may want to fine tune this later

		ReduceMemoryUsage: false, // NOTE: We may want to enable this later if we have memory issues. Please note enabling this may cause CPU usage to increase.

		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
			return ctx.Status(code).JSON(map[string]interface{}{
				"code":    InternalError,
				"title":   "Internal Server Error",
				"message": "Something went wrong",
				"details": nil,
			})
		},
	})

	return app
}
