package routes

import (
	artistContr "github.com/RubenPari/feat_eminem/src/controllers/artist"
	utilsContr "github.com/RubenPari/feat_eminem/src/controllers/utils"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	artist := app.Group("/artist")

	artist.Post("/add/:id", artistContr.Add)
	// get all songs of a specific artist
	// and add them to the database
	artist.Get("/get-all-songs/:id", artistContr.GetAllSongs)
	// filters all songs of a specific artist
	// where Eminem is featured
	// and add them to the database
	artist.Put("/get-featured-songs/:id", artistContr.GetFeaturedSongs)

	utils := app.Group("/utils")

	utils.Get("/artist/get-id/:name", utilsContr.GetIdByName)
	utils.Get("/artist/get-name/:id", utilsContr.GetNameById)
}
