package router

import (
	"github.com/awesomebfm/minepanel/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) RegisterFrontendRoutes() {
	r.app.Get("/login", r.handleGetLogin)
	r.app.Get("/", middleware.AuthMiddleware, r.handleGetDashboard)
}

// GET / (Dashboard, 303 if !auth)
func (r *Router) handleGetDashboard(c *fiber.Ctx) error {
	return c.SendFile("./templates/dashboard.html")
}

// Get /login
func (r *Router) handleGetLogin(c *fiber.Ctx) error {
	return c.Render("./templates/login.html", fiber.Map{
		"BadField": false,
	})
}