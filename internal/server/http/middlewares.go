package http

import (
	"regexp"
	"time"

	"github.com/vnworkday/gateway/internal/common/log"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/vnworkday/gateway/internal/conf"
	"go.uber.org/zap"
)

func LoggingMiddleware(logger *zap.Logger, config *conf.Conf) fiber.Handler {
	skippedByName := map[string][]string{
		fiber.MethodGet: {
			"/favicon.ico",
		},
	}

	skippedByPattern := map[string][]string{
		fiber.MethodGet: {
			config.HTTPPathPrefix + "/swagger",
		},
	}

	return func(ctx *fiber.Ctx) (err error) {
		method := ctx.Method()
		path := ctx.Path()

		if skipped, ok := skippedByName[method]; ok {
			for _, name := range skipped {
				if path == name {
					return ctx.Next()
				}
			}
		}

		if skipped, ok := skippedByPattern[method]; ok {
			for _, pattern := range skipped {
				if matched, _ := regexp.MatchString(pattern, path); matched {
					return ctx.Next()
				}
			}
		}

		defer func(begin time.Time) {
			fields := append(log.CommonHTTPFields(ctx.Path(), ctx.Method()), zap.Duration(log.FieldDuration, time.Since(begin)))

			if err != nil {
				var e *fiber.Error
				if errors.As(err, &e) {
					fields = append(fields, zap.Int(log.FieldStatus, e.Code), zap.Error(err))
				} else {
					fields = append(fields, zap.Int(log.FieldStatus, fiber.StatusInternalServerError), zap.Error(err))
				}
			} else {
				fields = append(fields, zap.Int(log.FieldStatus, fiber.StatusOK))
			}

			logger.Info("logging middleware", fields...)
		}(time.Now())

		return ctx.Next()
	}
}
