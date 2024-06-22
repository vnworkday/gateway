// Package handlers provides the application's business logic.
package handlers

import (
	"github.com/vnworkday/gateway/internal/handlers/metric"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"handlers",
	metrichandler.Module,
)
