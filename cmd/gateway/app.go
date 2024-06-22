package app

import (
	metrichandler "github.com/vnworkday/gateway/internal/handlers/metric"
	"github.com/vnworkday/gateway/internal/http"
	metricrouter "github.com/vnworkday/gateway/internal/routes/metric"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Start() {
	fx.New(
		fx.Provide(zap.NewProduction),
		fx.Decorate(func(logger *zap.Logger) *zap.Logger {
			return logger.WithLazy(zap.String("service", "gateway"))
		}),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: logger.WithOptions(zap.IncreaseLevel(zapcore.WarnLevel)),
			}
		}),
		fx.Module("handlers",
			fx.Decorate(func(logger *zap.Logger) *zap.Logger {
				return logger.Named("handlers")
			}),
			fx.Provide(
				namedRegister(metrichandler.NewHealthHandler, "health"),
			),
		),
		fx.Module("routers",
			fx.Decorate(func(logger *zap.Logger) *zap.Logger {
				return logger.Named("routers")
			}),
			fx.Provide(
				groupedRegister(metricrouter.NewHealthRouter, "routers", new(http.Router)),
			),
		),
		fx.Invoke(http.NewServer),
	).Run()
}

func namedRegister(constructor interface{}, name string, params ...string) any {
	paramTags := make([]string, 0, len(params))

	for _, p := range params {
		paramTags = append(paramTags, `name:"`+p+`"`)
	}

	return fx.Annotate(
		constructor,
		fx.ParamTags(paramTags...),
		fx.ResultTags(`name:"`+name+`"`),
	)
}

func groupedRegister(constructor interface{}, group string, g interface{}) any {
	return fx.Annotate(
		constructor,
		fx.As(g),
		fx.ResultTags(`group:"`+group+`"`),
	)
}
