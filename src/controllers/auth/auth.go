package auth

import (
	"context"
	spotifyAPI "github.com/zmb3/spotify/v2"
	spotifyAUTH "github.com/zmb3/spotify/v2/auth"
	"log"
	"net/http"

	authMO "github.com/RubenPari/feat_eminem/src/modules/auth"
	"github.com/gofiber/fiber/v2"
)

var (
	oauthConf, stateGlobal = authMO.GetOAuthConfig()
)

func Login(c *fiber.Ctx) error {
	url := oauthConf.AuthCodeURL(stateGlobal)
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

func Callback(c *fiber.Ctx) error {
	state := c.Query("state")

	if state != stateGlobal {
		log.Printf("invalid oauth state, expected '%s', got '%s'\n", stateGlobal, state)
		_ = c.SendStatus(http.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status": "error",
			"error":  "invalid oauth state",
		})
	}

	code := c.Query("code")

	token, err := oauthConf.Exchange(context.Background(), code)

	if err != nil {
		log.Printf("Code exchange failed with '%s'\n", err)
		_ = c.SendStatus(http.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status": "error",
			"error":  "Code exchange failed",
		})
	}

	// create http client
	httpClient := spotifyAUTH.New().Client(context.Background(), token)

	// create spotify http client
	spotifyClient := spotifyAPI.New(httpClient)

	// export spotify client to modules
	authMO.SpotifyClient = spotifyClient

	resp, err := spotifyClient.CurrentUser(context.Background())

	if err != nil {
		_ = c.SendStatus(http.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Failed authenticating",
		})
	}

	_ = c.SendStatus(http.StatusOK)
	return c.JSON(fiber.Map{
		"status":       "success",
		"message":      "login successful",
		"current_user": resp,
	})
}
