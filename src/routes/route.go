package routes

import (
	"github.com/RubenPari/feat_eminem/src/controllers/artist"
	"github.com/RubenPari/feat_eminem/src/controllers/utils"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	// TODO: implement group routes
	app.Post("/artist/add/:id", artist.Add)

	// get all songs of a specifi artist
	// and add them to the database
	app.Get("/artist/get-all-songs/:id", artist.GetAllSongs)

	// filters all songs of a specifi artist
	// where Eminem is featured
	// and add them to the database
	app.Put("/artist/get-featured-songs/:id", artist.GetFeaturedSongs)

	app.Get("/utils/artist/get-id/:name", utils.GetIdByName)
	app.Get("/utils/artist/get-name/:id", utils.GetNameById)
}
