package main

import (
	"golang-fiber-sqlite/controllers"
	"golang-fiber-sqlite/initializers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())

	micro.Route("/notes", func(router fiber.Router) {
		router.Post("", controllers.CreateNoteHandler)
		router.Get("", controllers.FindNotes)

	})

	micro.Route("/notes/:noteId", func(router fiber.Router) {
		router.Patch("", controllers.UpdateNote)
	})

	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, SQLite, and GORM",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
