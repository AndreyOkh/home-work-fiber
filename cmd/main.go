package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
	"news/config"
	"news/internal/pages"
	"news/pkg/logger"
)

func main() {
	config.Init()

	logConfig := config.NewLogConfig()
	dbConf := config.NewDatabaseConfig()

	loggerConfig := logger.NewLogger(logConfig)
	loggerConfig.Info("Database config", "dbConf", dbConf)

	app := fiber.New()
	app.Use(recover.New())
	app.Use(slogfiber.New(loggerConfig))

	pages.NewHandler(app, loggerConfig)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
