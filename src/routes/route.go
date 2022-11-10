package routes

import (
	artistCONTR "github.com/RubenPari/feat_eminem/src/controllers/artist"
	authCONTR "github.com/RubenPari/feat_eminem/src/controllers/auth"
	utilsCONTR "github.com/RubenPari/feat_eminem/src/controllers/utils"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	artist := app.Group("/artist")

	artist.Post("/add/:id", artistCONTR.Add)
	artist.Get("/check-if-saved/:id", artistCONTR.CheckIfSaved)
	artist.Delete("/delete/:id", artistCONTR.Delete)
	artist.Get("/get-all-songs/:id", artistCONTR.GetAllSongs)
	artist.Put("/get-featured-songs/:id", artistCONTR.GetFeaturedSongs)

	utils := app.Group("/utils")

	utils.Get("/artist/get-id/:name", utilsCONTR.GetIdByName)
	utils.Get("/artist/get-name/:id", utilsCONTR.GetNameById)

	auth := app.Group("/auth")

	auth.Get("/login", authCONTR.Login)
	auth.Get("/callback", authCONTR.Callback)
}
