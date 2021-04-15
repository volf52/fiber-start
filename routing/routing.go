package routing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/volf52/fiber-start/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/:name?", handlers.RenderName)

	api := app.Group("/api/v1")

	api.Get("/:name?", handlers.HelloName)
}
