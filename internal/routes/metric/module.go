package metricrouter

import (
	"github.com/vnworkday/gateway/internal/http"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewHealthRouter,
		fx.As(new(http.Router)),
		fx.ResultTags(`group:"routers"`),
	),
)
