package tenant

import (
	"github.com/vnworkday/common/pkg/ioc"
	"github.com/vnworkday/gateway/internal/common/port"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Provide(
		ioc.RegisterWithGroup(NewPort, "routers", new(port.Router)),
		ioc.RegisterWithName(NewHandler),
	)
}
