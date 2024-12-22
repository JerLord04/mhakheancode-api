package main

import (
	"mhakheancode/api/routs"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "mhakheancode-api",
	})

	routs.Setup(app)

	app.Listen(":3000")

}
