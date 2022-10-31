package utils

import (
	"net/http"

	spotifyMO "github.com/RubenPari/feat_eminem/src/modules/spotify"
	"github.com/gofiber/fiber/v2"
	spotifyAPI "github.com/zmb3/spotify/v2"
)

func GetIdByName(c *fiber.Ctx) error {
	name := c.Params("name")

	spotifyClient, ctx := spotifyMO.GetClient()

	artistApi, err := spotifyClient.Search(ctx, name, spotifyAPI.SearchTypeArtist)

	if err != nil || artistApi == nil {
		_ = c.SendStatus(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "artist with given name not found",
		})
	} else {
		id := artistApi.Artists.Artists[0].ID.String()

		_ = c.SendStatus(http.StatusOK)
		return c.JSON(fiber.Map{
			"id": id,
		})
	}
}

func GetNameById(c *fiber.Ctx) error {
	id := c.Params("id")

	spotifyClient, ctx := spotifyMO.GetClient()

	artistApi, err := spotifyClient.GetArtist(ctx, spotifyAPI.ID(id))

	if err != nil || artistApi == nil {
		_ = c.SendStatus(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "artist with given id not found",
		})
	} else {
		name := artistApi.Name

		_ = c.SendStatus(http.StatusOK)
		return c.JSON(fiber.Map{
			"name": name,
		})
	}
}
