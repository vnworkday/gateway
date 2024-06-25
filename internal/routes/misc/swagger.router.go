package misc

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type SwaggerRouter struct {
	// No dependency
}

func NewSwaggerRouter() *SwaggerRouter {
	return &SwaggerRouter{}
}

func (r *SwaggerRouter) Register(router fiber.Router) {
	router.Add(fiber.MethodGet, "/*", swagger.HandlerDefault)
	router.Add(fiber.MethodGet, "/*", swagger.New(swagger.Config{
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

func (r *SwaggerRouter) Path() string {
	return "/swagger"
}
