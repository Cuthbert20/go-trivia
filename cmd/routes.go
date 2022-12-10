package main

import (
	"github.com/Cuthbert20/go-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", handlers.Home)

}
