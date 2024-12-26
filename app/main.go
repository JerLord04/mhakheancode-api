package main

import (
	"mhakheancode/api/bootstrap"
	loadentities "mhakheancode/api/load_entities"
	"mhakheancode/api/routs"

	"github.com/gofiber/fiber/v2"
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

	loadentities.MigrationExistData(&config.PostgresDB)

	routs.Setup(app, &config.PostgresDB)

	app.Listen(":3000")

}
