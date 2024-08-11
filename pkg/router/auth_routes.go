package router

import (
	"github.com/gofiber/fiber/v2"
)

func (r *Router) RegisterAuthRoutes() {
	r.app.Post("/login", r.handleLogin)
}

// POST /login
func (r *Router) handleLogin(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	userAgent := c.Get("User-Agent")
	ip := c.IP()

	// In case they think they are sneaky
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
	passwordsMatch, err := r.ath.HashMatches(password, user.HashedPassword)
	if err != nil {
		return c.SendFile("./templates/500.html")
	}
	if !passwordsMatch {
		return c.Render("./templates/login.html", fiber.Map{
			"BadField": true,
		})
	}

	// Create a session
	session, secret, err := r.ath.NewSession(user, userAgent, ip)
	if err != nil {
		return c.SendFile("./templates/500.html")
	}

	// Persist the session
	err = r.db.PersistSession(session)
	if err != nil || session.Id == 0 {
		return c.SendFile("./templates/500.html")
	}

	token := r.ath.EncodeSession(session.Id, secret)
	c.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  session.ExpiresAt,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Redirect("/")
}
