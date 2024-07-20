package http

import (
	"github.com/vnworkday/common/pkg/ioc"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Provide(
		fx.Annotate(
			ioc.RegisterWithName(NewServer),
			fx.OnStart(OnStart),
			fx.OnStop(OnStop),
		),
	)
}
