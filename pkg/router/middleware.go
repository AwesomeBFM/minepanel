package router

import (
	"errors"
	"github.com/awesomebfm/minepanel/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"time"
)

func (r *Router) AuthMiddleware(c *fiber.Ctx) error {
	sessionToken := c.Cookies("session_token")

	if sessionToken == "" {
		c.ClearCookie("session_token")
		return c.Redirect("/login")
	}

	// decode session token
	id, secret, err := auth.DecodeSession(sessionToken)
	if err != nil {
		c.ClearCookie("session_token")
		return c.SendFile("./templates/500.html")
	}

	// Fetch session from database
	session, err := r.db.FindSessionById(id)
	if err != nil {
		c.ClearCookie("session_token")
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Redirect("/login")
		}
		return c.SendFile("./templates/500.html")
	}

	matches, err := r.ath.HashMatches(secret, session.HashedSecret)
	if err != nil || !matches {
		c.ClearCookie("session_token")
		return c.Redirect("/login")
	}

	currentTime := time.Now()
	if session.CreatedAt.After(currentTime) || session.ExpiresAt.Before(currentTime) {
		c.ClearCookie("session_token")
		return c.Redirect("/login")
	}

	return c.Next()
}

func (r *Router) ReverseAuthMiddleware(c *fiber.Ctx) error {
	sessionToken := c.Cookies("session_token")

	if sessionToken != "" {
		return c.Redirect("/")
	}

	return c.Next()
}
