// Package routes provides the application's HTTP routes.
package routes

import (
	metricrouter "github.com/vnworkday/gateway/internal/routes/metric"
	"go.uber.org/fx"
)

var Module = fx.Module("routers",
	metricrouter.Module,
)
