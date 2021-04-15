package handlers

import "github.com/gofiber/fiber/v2"

func RenderName(c *fiber.Ctx) error {
	name := c.Params("name", "Human")

	return c.Render("index", fiber.Map{
		"Name": name,
	})
}

func HelloName(c *fiber.Ctx) error {
	name := c.Params("name", "Human")

	return c.SendString("Hello, " + name)
}
