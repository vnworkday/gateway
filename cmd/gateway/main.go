package main

import (
	"github.com/vnworkday/gateway/internal/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			zap.NewProduction,
			http.NewHTTPServer,
		),
		fx.Invoke(func(app http.Server) {}),
	).Run()
}
