package router

import "github.com/gofiber/fiber/v2"

type Router struct {
	listenAddr string
	app *fiber.App
}

func NewRouter(listenAddr string, config fiber.Config) *Router {
	app := fiber.New(config)

	return &Router{
		listenAddr: listenAddr,
		app: app,
	}
}

func (r *Router) Listen() error {
	// Register routes
	r.app.Static("/", "./public")

	r.RegisterFrontendRoutes()

	return r.app.Listen(r.listenAddr)
}