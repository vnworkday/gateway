package server

import (
	"github.com/vnworkday/gateway/internal/server/http"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Module("server",
		http.Register(),
	)
}
