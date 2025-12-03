package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xphp/go-url-shortener/internal/core/domain"
	"github.com/xphp/go-url-shortener/internal/core/services"
)

type URLHandler struct {
	urlService *services.URLService
}

func NewURLHandler(urlService *services.URLService) *URLHandler {
	return &URLHandler{urlService: urlService}
}

func (h *URLHandler) ShortenURL(c *fiber.Ctx) error {
	var req domain.CreateURLRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.OriginalURL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "original_url is required",
		})
	}

	response, err := h.urlService.ShortenURL(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *URLHandler) RedirectURL(c *fiber.Ctx) error {
	shortCode := c.Params("shortCode")

	if shortCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "short code is required",
		})
	}

	originalURL, err := h.urlService.GetOriginalURL(c.Context(), shortCode)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}

	return c.Redirect(originalURL, fiber.StatusMovedPermanently)
}
