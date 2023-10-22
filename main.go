package main

import (
	"golang-fiber-sqlite/initializers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	app := fiber.New()

	app.Get("/api/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "status",
			"message": "Welcome to Golang, Fiber, SQLite, and GORM",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
