package router

import (
	"github.com/awesomebfm/minepanel/pkg/auth"
	"github.com/awesomebfm/minepanel/pkg/database"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	listenAddr string
	app        *fiber.App
	db         *database.Database
	ath        *auth.Auth
}

func NewRouter(
	listenAddr string,
	config fiber.Config,
	db *database.Database,
	ath *auth.Auth,
) *Router {
	app := fiber.New(config)

	return &Router{
		listenAddr: listenAddr,
		app:        app,
		db:         db,
		ath:        ath,
	}
}

func (r *Router) Listen() error {
	// Register routes
	r.app.Static("/", "./public")

	r.RegisterFrontendRoutes()
	r.RegisterAuthRoutes()

	return r.app.Listen(r.listenAddr)
}
