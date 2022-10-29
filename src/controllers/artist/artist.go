package artist

import (
	"log"
	"net/http"

	"github.com/RubenPari/feat_eminem/src/database/artist"
	"github.com/RubenPari/feat_eminem/src/models"
	"github.com/RubenPari/feat_eminem/src/modules/spotify"
	"github.com/gofiber/fiber/v2"
	spotifyAPI "github.com/zmb3/spotify/v2"
)

func Add(c *fiber.Ctx) error {
	id := c.Params("id")

	spotifyClient, ctx := spotify.GetClient()

	artistApi, err := spotifyClient.GetArtist(ctx, spotifyAPI.ID(id))

	if err != nil || artistApi == nil {
		log.Fatalf("couldn't get artist information: %v", err)
	}

	artistObj := models.Artist{
		Id:   artistApi.ID,
		Name: artistApi.Name,
		Uri:  artistApi.URI,
	}

	success := artist.Add(artistObj)

	if success {
		c.SendStatus(http.StatusCreated)
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "artist added",
		})
	} else {
		c.SendStatus(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "error adding artist",
		})
	}
}
