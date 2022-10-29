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

func GetAllSongs(c *fiber.Ctx) error {
	id := c.Params("id")

	spotifyClient, ctx := spotify.GetClient()

	artist := artist.Get(id)

	if artist.Id == spotifyAPI.ID("") {
		c.SendStatus(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "artist not found",
		})
	}

	albumsObj := make([]models.Album, 0)
	tracksObj := make([]models.Track, 0)

	// get all albums
	albums, err := spotifyClient.GetArtistAlbums(
		ctx,
		artist.Id,
		[]spotifyAPI.AlbumType{
			spotifyAPI.AlbumTypeAlbum,
			spotifyAPI.AlbumTypeSingle,
			spotifyAPI.AlbumTypeCompilation,
		})

	if err != nil {
		log.Fatalf("couldn't get artist albums: %v", err)
	}

	for _, album := range albums.Albums {
		albumObj := models.Album{
			Id:     album.ID,
			Name:   album.Name,
			Uri:    album.URI,
			Artist: album.Artists[0].Name,
		}

		albumsObj = append(albumsObj, albumObj)
	}

	// get all songs
	for _, album := range albumsObj {
		tracks, err := spotifyClient.GetAlbumTracks(ctx, album.Id)

		if err != nil {
			log.Fatalf("couldn't get album tracks: %v", err)
		}

		for _, track := range tracks.Tracks {
			featuringArtists := make([]string, 0)

			for i := 1; i < len(track.Artists); i++ {
				featuringArtists = append(featuringArtists, track.Artists[i].Name)
			}

			trackObj := models.Track{
				Id:        track.ID,
				Name:      track.Name,
				Uri:       track.URI,
				Album:     album.Name,
				Artist:    track.Artists[0].Name,
				Featuring: featuringArtists,
			}

			tracksObj = append(tracksObj, trackObj)
		}
	}

	c.SendStatus(http.StatusOK)
	return c.JSON(tracksObj)
}

func GetFeaturedSongs(c *fiber.Ctx) error {
	return nil
}
