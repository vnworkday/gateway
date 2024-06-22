package metrichandler

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewHealthHandler,
		fx.As(new(HealthHandler)),
		fx.ResultTags(`name:"health"`),
	),
)
