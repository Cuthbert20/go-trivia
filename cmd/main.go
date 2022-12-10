package main

import (
	"github.com/Cuthbert20/go-trivia/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3001")
}
