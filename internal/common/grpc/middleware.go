package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	defaultMaxMessageSize = 4

	defaultKeepaliveTime    = 10
	defaultKeepaliveTimeout = 5

	defaultBackoffInterval = 100

	sizeMB = 1 << (10 * 2) //nolint:mnd
)

func withLoggingInterceptor(logger *zap.Logger) grpc.UnaryClientInterceptor {
	return logging.UnaryClientInterceptor(interceptorLogger(logger), logging.WithLogOnEvents(
		logging.StartCall,
		logging.FinishCall,
	))
}

func withRetryInterceptor() grpc.UnaryClientInterceptor {
	return retry.UnaryClientInterceptor(retry.WithBackoff(
		retry.BackoffExponential(time.Duration(defaultBackoffInterval) * time.Millisecond),
	))
}

// InterceptorLogger adapts zap Logger to interceptor Logger.
//
// Ref: https://github.com/grpc-ecosystem/go-grpc-middleware/blob/main/interceptors/logging/examples/zap/example_test.go
func interceptorLogger(l *zap.Logger) logging.LoggerFunc {
	return func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		logFields := buildLogFields(fields...)

		logger := l.WithOptions(zap.AddCallerSkip(1)).With(logFields...)

		switch lvl {
		case logging.LevelDebug:
			logger.Debug(msg)
		case logging.LevelInfo:
			logger.Info(msg)
		case logging.LevelWarn:
			logger.Warn(msg)
		case logging.LevelError:
			logger.Error(msg)
		default:
			logger.DPanic(fmt.Sprintf("unknown level %v", lvl))
		}
	}
}

func buildLogFields(fields ...any) []zap.Field {
	logFields := make([]zap.Field, 0)

	for i := 0; i < len(fields); i += 2 {
		key := utils.ToString(fields[i])
		value := fields[i+1]

		switch val := value.(type) {
		case string:
			logFields = append(logFields, zap.String(key, val))
		case int:
			logFields = append(logFields, zap.Int(key, val))
		case bool:
			logFields = append(logFields, zap.Bool(key, val))
		case []byte:
			logFields = append(logFields, zap.ByteString(key, val))
		case time.Duration:
			logFields = append(logFields, zap.Duration(key, val))
		default:
			logFields = append(logFields, zap.Any(key, val))
		}
	}

	return logFields
}
