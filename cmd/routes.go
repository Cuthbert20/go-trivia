package main

import "github.com/gofiber/fiber/v2"

func setupRoutes(app *filber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Helloo,  Stu lets have some fun with Trivia!!! test")
	})

}
