package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/greatdaveo/fiber-api/database"
	"github.com/greatdaveo/fiber-api/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func setupRoutes(app *fiber.App) {
	// Welcome endpoint
	app.Get("/api", welcome)
	// User endpoint
	app.Post("/api/users", routes.CreateUser)
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)
	
	log.Fatal(app.Listen(":3000"))
}
