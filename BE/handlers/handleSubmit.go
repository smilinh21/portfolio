package handlers

import (
	"Portfolio/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

type FormData struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// Handle form submission
func HandleSubmit(c *fiber.Ctx) error {
	formData := new(FormData)

	if err := c.BodyParser(formData); err != nil {
		return err
	}

	_, err := database.Client.FormData.Create().
		SetName(formData.Name).
		SetEmail(formData.Email).
		SetMessage(formData.Message).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(c.Context())

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Form data received successfully",
	})
}
