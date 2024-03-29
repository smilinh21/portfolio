package main

import (
	"Portfolio/database"
	"Portfolio/handlers"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

const defaultPort = ":8000"

func main() {
	database.Connect()

	if err := database.Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := fiber.New()

	app.Use(middlewareCors)

	app.Post("/submit", handlers.HandleSubmit)

	if err := app.Listen(defaultPort); err != nil {
		log.Fatal(err)
	}

}

func middlewareCors(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "*")
	if c.Method() == "OPTIONS" {
		return c.SendStatus(fiber.StatusOK)
	}
	return c.Next()
}
