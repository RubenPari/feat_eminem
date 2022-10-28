package routes

import (
	"github.com/RubenPari/feat-eminem/src/controllers/artist"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Post("/artist/add/:id", artist.Add)
}
