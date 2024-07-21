package logger

import (
	"syscall"

	"github.com/pkg/errors"
	"github.com/vnworkday/common/pkg/ioc"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Register() fx.Option {
	return fx.Provide(
		fx.Annotate(
			ioc.RegisterWithName(New),
			fx.OnStop(func(logger *zap.Logger) error {
				if err := logger.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
					return err
				}

				return nil
			}),
		),
	)
}
