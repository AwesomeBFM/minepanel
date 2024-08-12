package router

import (
	"github.com/gofiber/fiber/v2"
)

func (r *Router) RegisterFrontendRoutes() {
	r.app.Get("/login", r.ReverseAuthMiddleware, r.handleGetLogin)
	r.app.Get("/500", r.handleGet500)
	r.app.Get("/", r.AuthMiddleware, r.handleGetDashboard)
}

// GET / (Dashboard, 303 if !auth)
func (r *Router) handleGetDashboard(c *fiber.Ctx) error {
	return c.SendFile("./templates/dashboard.html")
}

// Get /login
func (r *Router) handleGetLogin(c *fiber.Ctx) error {
	return c.Render("./templates/login.html", fiber.Map{
		"UsernameError": "",
		"PasswordError": "",
	})
}

func (r *Router) handleGet500(c *fiber.Ctx) error {
	return c.SendFile("./templates/500.html")
}
