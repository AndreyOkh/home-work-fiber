package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	slogfiber "github.com/samber/slog-fiber"
	"news/config"
	"news/internal/pages"
	"news/pkg/logger"
)

func main() {
	config.Init()

	logConfig := config.NewLogConfig()
	config.NewDatabaseConfig()

	loggerConfig := logger.NewLogger(logConfig)

	engine := html.New("./html", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(recover.New())
	app.Use(slogfiber.New(loggerConfig))

	pages.NewHandler(app, loggerConfig)

	if err := app.Listen(":3000"); err != nil {
		loggerConfig.Error(err.Error())
	}
}
