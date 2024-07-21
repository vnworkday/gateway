package http

import (
	"regexp"
	"time"

	"github.com/vnworkday/gateway/internal/common/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/vnworkday/gateway/internal/common/log"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/vnworkday/gateway/internal/conf"
	"go.uber.org/zap"
)

func LoggingMiddleware(logger *zap.Logger, _ *conf.Conf) fiber.Handler {
	skippedPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\\+\\.(html|css|js|png|json)`),
	}

	return func(ctx *fiber.Ctx) (err error) {
		method, path := ctx.Method(), ctx.Path()

		// Check if the request should be skipped based on path pattern
		for _, pattern := range skippedPatterns {
			if pattern.MatchString(path) {
				return ctx.Next()
			}
		}

		defer func(begin time.Time) {
			fields := log.CommonHTTPFields(path, method)
			fields = append(fields, zap.Duration(log.FieldHTTPDuration, time.Since(begin)))

			if err != nil {
				var e *fiber.Error
				var httpStatus int

				if errors.As(err, &e) {
					httpStatus = e.Code
				} else if gerr, ok := status.FromError(err); ok && gerr.Code() != codes.Unknown {
					httpStatus = model.ToHTTPStatus(model.FromGRPCError(gerr))
				} else {
					httpStatus = fiber.StatusInternalServerError
				}

				fields = append(fields, zap.Int(log.FieldHTTPStatus, httpStatus), zap.Error(err))

				logger.Error("logging middleware", fields...)
			} else {
				logger.Info("logging middleware", fields...)
			}
		}(time.Now())

		return ctx.Next()
	}
}
