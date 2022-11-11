package utils

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	spotifyMO "github.com/RubenPari/feat_eminem/src/modules/spotify"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	spotifyAPI "github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
	spotifyAUTH "golang.org/x/oauth2/spotify"
)

// LoadEnv loads environment
// variables from .env file
// in the root directory
// upDir: number of directories
// to move up from the current path
func LoadEnv(upDir int) bool {
	// get current path
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	var numUpDir = ""

	for i := 0; i < upDir; i++ {
		numUpDir += "../"
	}

	// move up of n directories
	rootPath := filepath.Join(basePath, numUpDir)

	// load env variables
	err := godotenv.Load(rootPath + "/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return true
}

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

func GenerateStateString() string {
	// TODO: generate random string
	return "state"
}

func GetOAuthConfig() (*oauth2.Config, string) {
	LoadEnv(3)

	return &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes:       []string{"user-read-private", "user-read-email"},
		Endpoint:     spotifyAUTH.Endpoint,
	}, GenerateStateString()
}
