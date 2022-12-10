package handlers

import (
	"github.com/Cuthbert20/go-trivia/database"
	"github.com/Cuthbert20/go-trivia/models"
	"github.com/gofiber/fiber/v2"
)

// Get Home handler
func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

// Post Fact Handler
func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	//error handling
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}
