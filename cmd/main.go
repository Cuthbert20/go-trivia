package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello,  Stu lets have some fun with Trivia!!! test")
	})

	app.Listen(":3001")
}