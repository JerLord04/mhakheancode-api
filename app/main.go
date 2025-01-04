package main

import (
	"mhakheancode/api/bootstrap"
	loadentities "mhakheancode/api/load_entities"
	"mhakheancode/api/routs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config := bootstrap.App()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "mhakheancode-api",
	})

	app.Use(cors.New())

	loadentities.MigrationExistData(&config.PostgresDB)

	routs.Setup(app, &config.PostgresDB)

	app.Listen(":3001")

}
