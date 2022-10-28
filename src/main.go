package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic(errEnv)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	errServer := app.Listen(":" + os.Getenv("PORT"))
	if errServer != nil {
		panic(errServer)
	}
}
