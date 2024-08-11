package router

import (
	"github.com/awesomebfm/minepanel/pkg/auth"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) RegisterUserRoutes() {
	r.app.Group("/users")

	r.app.Post("", r.handleCreateUser)
}

type CreateUserRequest struct {
	Username string      `json:"username"`
	Password string      `json:"password"`
	Roles    []auth.Role `json:"roles"`
}

func (r *Router) handleCreateUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"error": "not implemented",
	})
}
