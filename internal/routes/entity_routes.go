package routes

import (
	"api-puri/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterEntityRoutes(router fiber.Router, h *handlers.EntityHandler) {
	entities := router.Group("/entities")

	entities.Get("/", h.Get)
	entities.Get("/map", h.GetMapPoints)
	entities.Get("/:id", h.Detail)
	entities.Post("/", h.Create)
	entities.Put("/:id", h.Update)
	entities.Delete("/:id", h.Delete)
}
