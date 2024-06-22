package http

import "github.com/gofiber/fiber/v2"

type Router interface {
	Register(router fiber.Router)
	Path() string
}
