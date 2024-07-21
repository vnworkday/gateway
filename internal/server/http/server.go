package http

import (
	"encoding/json"
	"time"

	"github.com/gofiber/swagger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/vnworkday/gateway/internal/common/log"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/vnworkday/gateway/internal/common/model"
	"github.com/vnworkday/gateway/internal/common/port"
	"github.com/vnworkday/gateway/internal/conf"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

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
	server.Use(cors.New())
	server.Use(compress.New())
	server.Use(helmet.New())
	server.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(ctx *fiber.Ctx, err any) {
			fields := append(log.CommonHTTPFields(ctx.Path(), ctx.Method()), zap.Any(log.FieldHTTPError, err))

			if props.Config.Profile != profileLocal {
				fields = append(fields, zap.Stack(log.FieldStack))
			}

			props.Logger.Error("recovered", fields...)
		},
	}))

	for _, router := range props.Routers {
		path := props.Config.HTTPPathPrefix + router.Path()
		server.Route(path, router.Register)
	}

	server.Add(fiber.MethodGet, "/swagger/*", swagger.HandlerDefault)

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
			message := err.Error()

			var ferr *fiber.Error
			if errors.As(err, &ferr) {
				code = model.FromFiberError(ferr)
				message = ferr.Message
			} else if gerr, ok := status.FromError(err); ok && gerr.Code() != codes.Unknown {
				code = model.FromGRPCError(gerr)
				message = gerr.Message()
			}

			body, _ := json.Marshal(model.NewError(code, message))

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
	logger.Info("server: starting", zap.String(log.FieldProtocol, log.ProtocolHTTP))

	go func() {
		if err := server.Listen(":3000"); err != nil {
			logger.Panic("http: failed to start server", zap.String(log.FieldProtocol, "http"), zap.Error(err))
		}
	}()

	return nil
}

func OnStop(server *fiber.App, logger *zap.Logger) error {
	logger.Info("server: shutting down", zap.String(log.FieldProtocol, log.ProtocolHTTP))

	return server.Shutdown()
}
