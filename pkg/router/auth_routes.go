package router

import "github.com/gofiber/fiber/v2"

func (r *Router) RegisterAuthRoutes() {
	r.app.Post("/login")
}

// POST /login
func (r *Router) handlePostLogin(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.Render("./templates/login.html", fiber.Map{
			"BadField": true,
		})
	}

	return c.Redirect("/")
}
