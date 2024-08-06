package main

import (
	"log"

	"github.com/awesomebfm/minepanel/pkg/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	r := router.NewRouter(":8080", fiber.Config{})

	err := r.Listen()
	if err != nil {
		log.Fatalf("error starting router: %q", err.Error())
		return
	}
}
