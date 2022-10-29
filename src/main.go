package main

import (
	"os"

	"github.com/RubenPari/feat_eminem/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic(errEnv)
	}

	routes.SetUpRoutes(app)

	errServer := app.Listen(":" + os.Getenv("PORT"))
	if errServer != nil {
		panic(errServer)
	}
}
