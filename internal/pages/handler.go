package pages

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

type Handler struct {
	router fiber.Router
	log    *slog.Logger
}

func NewHandler(router fiber.Router, log *slog.Logger) {
	h := &Handler{
		router: router,
		log:    log,
	}
	h.router.Get("/", h.mainPage)
}

func (h *Handler) mainPage(c *fiber.Ctx) error {
	h.log.Warn("Test")
	return c.SendString("Home Page")
}
