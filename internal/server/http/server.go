package http

import (
	"encoding/json"
	"github.com/vnworkday/gateway/internal/common/model"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/vnworkday/gateway/internal/common/port"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/vnworkday/gateway/internal/conf"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const profileLocal = "local"

type ServerProps struct {
	fx.In
	Logger  *zap.Logger
	Config  *conf.Conf
	Routers []port.Router `group:"routers"`
}

func NewServer(props ServerProps) *fiber.App {
	server := buildHTTPServer(props.Config)

	server.Use(LoggingMiddleware(props.Logger, props.Config))
	server.Use(healthcheck.New())
	server.Use(cors.New(cors.Config{
		AllowOrigins: "",
		AllowOriginsFunc: func(origin string) bool {
			var matched bool
			if matched, _ = regexp.MatchString(`^https://.*\.vnworkday\.com$`, origin); matched {
				return true
			}

			props.Logger.Debug("CORS origin not allowed", zap.String("origin", origin))

			return false
		},
	}))
	server.Use(compress.New())
	server.Use(helmet.New())
	server.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(ctx *fiber.Ctx, err any) {
			fields := []zap.Field{
				zap.String("path", ctx.Path()),
				zap.String("method", ctx.Method()),
				zap.Any("error", err),
			}

			if props.Config.Profile != profileLocal {
				fields = append(fields, zap.Stack("stack"))
			}

			props.Logger.Error("panic recovered", fields...)
		},
	}))

	for _, router := range props.Routers {
		path := props.Config.HTTPPathPrefix + router.Path()
		server.Route(path, router.Register)
	}

	return server
}

func buildHTTPServer(config *conf.Conf) *fiber.App {
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
		ReadTimeout:  config.HTTPRequestReadTimeout,   // NOTE: We may want to fine tune this later
		WriteTimeout: config.HTTPResponseWriteTimeout, // NOTE: We may want to fine tune this later

		ReduceMemoryUsage: false, // NOTE: We may want to enable this later if we have memory issues.

		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := model.CodeErrInternal

			var ferr *fiber.Error
			if errors.As(err, &ferr) {
				code = model.FromFiberError(ferr)
			}

			body, _ := json.Marshal(model.NewError(code))

			ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

			if ferr != nil {
				ctx.Status(ferr.Code).Response().SetBody(body)
			} else {
				ctx.Status(model.ToHTTPStatus(code)).Response().SetBody(body)
			}

			return nil
		},

		EnablePrintRoutes:     config.Profile == profileLocal,
		DisableStartupMessage: config.Profile != profileLocal,
	})

	return app
}

func OnStart(server *fiber.App, logger *zap.Logger) error {
	logger.Info("Starting server", zap.String("address", ":3000"))

	go func() {
		if err := server.Listen(":3000"); err != nil {
			logger.Panic("Failed to start server", zap.Error(err))
		}
	}()

	return nil
}

func OnStop(server *fiber.App) error {
	return server.Shutdown()
}
