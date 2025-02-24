package pages

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
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
	content := Content{
		Categories: []Categories{
			{Path: "/food", Name: "#Еда"},
			{Path: "/animals", Name: "#Животные"},
			{Path: "/cars", Name: "#Машины"},
			{Path: "/sport", Name: "#Спорт"},
			{Path: "/music", Name: "#Музыка"},
			{Path: "/tech", Name: "#Технологии"},
			{Path: "/other", Name: "#Прочее"},
		},
	}
	return c.Render("category", content)
}
