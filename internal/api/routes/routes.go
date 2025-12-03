package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xphp/go-url-shortener/internal/api/handlers"
)

func SetupRoutes(app *fiber.App, urlHandler *handlers.URLHandler) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "System is healthy",
		})
	})

	api := app.Group("/api")
	api.Post("/shorten", urlHandler.ShortenURL)

	app.Get("/:shortCode", urlHandler.RedirectURL)
}
