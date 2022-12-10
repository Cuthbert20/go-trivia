package main

import (
	"github.com/Cuthbert20/go-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", handlers.ListFacts)
	//STU TODO need to work on this one.
	app.Get("/:id", handlers.ListFact)

	app.Post("/fact", handlers.CreateFact)

}
