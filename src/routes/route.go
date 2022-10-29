package routes

import (
	"github.com/RubenPari/feat_eminem/src/controllers/artist"
	"github.com/RubenPari/feat_eminem/src/controllers/utils"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	// TODO: implement group routes
	app.Post("/artist/add/:id", artist.Add)

	app.Get("/utils/artist/get-id/:name", utils.GetIdByName)
	app.Get("/utils/artist/get-name/:id", utils.GetNameById)
}
