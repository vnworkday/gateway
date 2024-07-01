package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"

	"github.com/gofiber/fiber/v2/utils"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
)

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
			panic(fmt.Sprintf("unknown level %v", lvl))
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
		default:
			logFields = append(logFields, zap.Any(key, val))
		}
	}

	return logFields
}

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
