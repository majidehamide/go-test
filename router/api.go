package router

import (
	"github.com/gofiber/fiber/v2"
)

func Routes() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is API")
	})
}
