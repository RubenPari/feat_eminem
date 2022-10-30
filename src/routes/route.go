package routes

import (
	artistCONTR "github.com/RubenPari/feat_eminem/src/controllers/artist"
	utilsCONTR "github.com/RubenPari/feat_eminem/src/controllers/utils"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	artist := app.Group("/artist")

	artist.Post("/add/:id", artistCONTR.Add)
	// get all songs of a specific artist
	// and add them to the database
	artist.Get("/get-all-songs/:id", artistCONTR.GetAllSongs)
	// filters all songs of a specific artist
	// where Eminem is featured
	// and add them to the database
	artist.Put("/get-featured-songs/:id", artistCONTR.GetFeaturedSongs)

	utils := app.Group("/utils")

	utils.Get("/artist/get-id/:name", utilsCONTR.GetIdByName)
	utils.Get("/artist/get-name/:id", utilsCONTR.GetNameById)
}
