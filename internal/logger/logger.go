package logger

import (
	"github.com/vnworkday/common/pkg/log"
	logfield "github.com/vnworkday/gateway/internal/common/log"
	"github.com/vnworkday/gateway/internal/conf"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In
	Config *conf.Conf
}

func New(params Params) (*zap.Logger, error) {
	serviceLogger, err := log.NewLogger()
	if err != nil {
		return nil, err
	}

	serviceLogger = serviceLogger.WithLazy(zap.String(logfield.FieldService, params.Config.ServiceName))

	return serviceLogger, nil
}
