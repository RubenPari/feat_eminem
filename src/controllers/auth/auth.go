package auth

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/RubenPari/feat_eminem/src/controllers/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

var (
	oauthConf, stateGlobal = utils.GetOAuthConfig()
	AccessToken            = &oauth2.Token{}
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

	// save token to session
	AccessToken = token

	client := oauthConf.Client(context.Background(), token)

	resp, err := client.Get("https://api.spotify.com/v1/me")

	if err != nil {
		_ = c.SendStatus(http.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Failed authenticating",
		})
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	_ = c.SendStatus(http.StatusOK)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "login successful",
	})
}
