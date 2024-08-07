package main

import (
	"log"
	"os"
	"time"

	"github.com/awesomebfm/minepanel/pkg/auth"
	"github.com/awesomebfm/minepanel/pkg/database"
	"github.com/awesomebfm/minepanel/pkg/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Database
	db, err := database.NewDatabase(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("error connection to database: %q", err.Error())
	}
	defer db.Close()

	// Auth
	ath := auth.NewAuth(
		&auth.Params{
			Memory:      64 * 1024,
			Iterations:  3,
			Parallelism: 2,
			SaltLength:  16,
			KeyLength:   32,
		},
		90*24*time.Hour,
	)

	// Router
	r := router.NewRouter(":8080", fiber.Config{}, db, ath)

	err = r.Listen()
	if err != nil {
		log.Fatalf("error starting router: %q", err.Error())
		return
	}
}
