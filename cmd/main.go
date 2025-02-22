package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"news/config"
	"news/internal/pages"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println("dbConf: ", dbConf)

	app := fiber.New()
	app.Use(recover.New())

	pages.NewHandler(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
