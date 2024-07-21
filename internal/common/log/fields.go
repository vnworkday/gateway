package log

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

const (
	ProtocolHTTP = "http"
	ProtocolGRPC = "grpc"
)

const (
	FieldStack    = "stack"
	FieldService  = "service"
	FieldProtocol = "protocol"
)

const (
	FieldHTTPPath     = "http.path"
	FieldHTTPMethod   = "http.method"
	FieldHTTPStatus   = "http.status"
	FieldHTTPError    = "http.error"
	FieldHTTPDuration = "http.time_ms"
)

const (
	FieldGRPCComponent = "grpc.component"
	FieldGRPCService   = "grpc.service"
	FieldGRPCMethod    = "grpc.method"
	FieldGRPCStartTime = "grpc.start_time"
	FieldGRPCCode      = "grpc.code"
	FieldGRPCError     = "grpc.error"
	FieldGRPCDuration  = "grpc.time_ms"
)

func CommonHTTPFields(path, method string) []zap.Field {
	return []zap.Field{
		zap.String(FieldProtocol, ProtocolHTTP),
		zap.String(FieldHTTPPath, path),
		zap.String(FieldHTTPMethod, method),
		zap.Int(FieldHTTPStatus, fiber.StatusOK),
	}
}

func CommonGRPCFields(service, method string, status status.Status) []zap.Field {
	return []zap.Field{
		zap.String(FieldProtocol, ProtocolGRPC),
		zap.String(FieldGRPCService, service),
		zap.String(FieldGRPCMethod, method),
		zap.Int(FieldGRPCCode, int(status.Code())),
	}
}
