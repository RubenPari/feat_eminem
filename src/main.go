package main

import (
	"os"

	"github.com/RubenPari/feat_eminem/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetUpRoutes(app)

	errServer := app.Listen(":" + os.Getenv("PORT"))
	if errServer != nil {
		panic(errServer)
	}
}
