package pages

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"news/pkg/templ_adaptor"
	"news/views"
)

type Handler struct {
	router fiber.Router
	log    *slog.Logger
}

type Categories struct {
	Name string
	Path string
}

type Content struct {
	Categories []Categories
}

func NewHandler(router fiber.Router, log *slog.Logger) {
	h := &Handler{
		router: router,
		log:    log,
	}
	h.router.Get("/", h.mainPage)
}

func (h *Handler) mainPage(c *fiber.Ctx) error {
	component := views.Main()
	return templ_adaptor.Render(c, component)
}
