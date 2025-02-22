package pages

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &Handler{
		router: router,
	}
	h.router.Get("/", h.mainPage)
}

func (h *Handler) mainPage(c *fiber.Ctx) error {
	return c.SendString("Home Page")
}
