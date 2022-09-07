package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/amirhossein-programmer/cmd/web/pkg/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "fiber",
		AppName:       "My App v1.0.1",
	})
	setupRoutes(app)
	log.Fatal(app.Listen(":4000"))
}
