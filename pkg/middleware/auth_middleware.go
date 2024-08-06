package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(c *fiber.Ctx) error {
	sessionToken := c.Cookies("session_token")

	if sessionToken == "" {
		return c.Redirect("/login")
	}

	return c.Next()
}