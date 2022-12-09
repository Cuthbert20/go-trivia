package handlers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.SendString("Helloo,  Stu lets have some fun with Trivia!!! test")
}
