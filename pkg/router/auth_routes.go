package router

import "github.com/gofiber/fiber/v2"

func (r *Router) RegisterAuthRoutes() {
	r.app.Post("/login")
}

// POST /login
func (r *Router) handlePostLogin(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Incase they think they are sneaky
	if username == "" || password == "" {
		return c.Render("./templates/login.html", fiber.Map{
			"BadField": true,
		})
	}

	// Fetch the user
	user, err := r.db.FindUserByUsername(username)
	if err != nil {
		return c.Render("./templates/login.html", fiber.Map{
			"BadField": true,
		})
	}

	// Validate password
	passwordsMatch, err := r.ath.PasswordsMatch(password, user.HashedPassword)
	if err != nil {
		return c.SendFile("./templates/500.html")
	}
	if !passwordsMatch {
		return c.Render("./templates/login.html", fiber.Map{
			"BadField": true,
		})
	}

	// Create a session

	return c.Redirect("/")
}
