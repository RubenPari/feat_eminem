package artist

import (
	"log"
	"net/http"

	trackDB "github.com/RubenPari/feat_eminem/src/database/track"

	artistDB "github.com/RubenPari/feat_eminem/src/database/artist"
	"github.com/RubenPari/feat_eminem/src/models"
	spotifyMO "github.com/RubenPari/feat_eminem/src/modules/spotify"
	"github.com/gofiber/fiber/v2"
	spotifyAPI "github.com/zmb3/spotify/v2"
)

func Add(c *fiber.Ctx) error {
	id := c.Params("id")

	spotifyClient, ctx := spotifyMO.GetClient()

	artistApi, err := spotifyClient.GetArtist(ctx, spotifyAPI.ID(id))

	if err != nil || artistApi == nil {
		log.Fatalf("couldn't get artist information: %v", err)
	}

	artistObj := models.Artist{
		Id:   artistApi.ID,
		Name: artistApi.Name,
		Uri:  artistApi.URI,
	}

	success := artistDB.Add(artistObj)

	if success {
		_ = c.SendStatus(http.StatusCreated)
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "artist added",
		})
	} else {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "error adding artist",
		})
	}
}

func GetAllSongs(c *fiber.Ctx) error {
	id := c.Params("id")

	spotifyClient, ctx := spotifyMO.GetClient()

	artistObj := artistDB.Get(id)

	if artistObj.Id == spotifyAPI.ID("") {
		_ = c.SendStatus(http.StatusNotFound)
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
		artistObj.Id,
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

		for _, song := range tracks.Tracks {
			featuringArtists := ""
			for i := 1; i < len(song.Artists); i++ {
				featuringArtists += song.Artists[i].Name + ", "
			}

			trackObj := models.Track{
				Id:        song.ID,
				Name:      song.Name,
				Uri:       song.URI,
				Album:     album.Name,
				Artist:    song.Artists[0].Name,
				Featuring: featuringArtists,
			}

			tracksObj = append(tracksObj, trackObj)
		}
	}

	// save songs to db
	success := trackDB.Adds(tracksObj)

	if success {
		_ = c.SendStatus(http.StatusCreated)
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "songs added",
		})
	} else {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "error adding songs",
		})
	}
}

func GetFeaturedSongs(c *fiber.Ctx) error {
	id := c.Params("id")

	tracksObj := trackDB.GetAllByArtist(id)

	tracksFiltersObj := trackDB.FilterByFeaturing(tracksObj)

	success := trackDB.AddsFeatured(tracksFiltersObj)

	if success {
		_ = c.SendStatus(http.StatusCreated)
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "songs added",
		})
	} else {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "error adding songs",
		})
	}
}
