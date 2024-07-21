package log

import "go.uber.org/zap"

const (
	FieldPath     = "path"
	FieldMethod   = "method"
	FieldDuration = "duration"
	FieldStatus   = "status"
	FieldReason   = "reason"
	FieldStack    = "stack"
	FieldService  = "service"
)

func CommonHTTPFields(path, method string) []zap.Field {
	return []zap.Field{
		zap.String(FieldPath, path),
		zap.String(FieldMethod, method),
	}
}
