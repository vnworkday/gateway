package misc

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"
)

type Port struct{}

type PortParams struct {
	fx.In
}

func NewPort(_ PortParams) *Port {
	return &Port{}
}

func (h *Port) Register(router fiber.Router) {
	router.Add(fiber.MethodGet, "/swagger/*", swagger.HandlerDefault)
	router.Add(fiber.MethodGet, "/swagger/*", swagger.New(swagger.Config{
		DeepLinking:            true,
		DocExpansion:           "none",
		RequestSnippetsEnabled: true,
		QueryConfigEnabled:     true,
		SyntaxHighlight: &swagger.SyntaxHighlightConfig{
			Activate: true,
			Theme:    "monokai",
		},
	}))
}

func (h *Port) Path() string {
	return "/misc"
}
