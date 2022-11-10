package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/RubenPari/feat_eminem/src/controllers/utils"
	"github.com/gofiber/fiber/v2"
)

var (
	oauthConf, stateString = utils.GetOAuthConfig()
)

func Login(c *fiber.Ctx) error {
	url := oauthConf.AuthCodeURL(stateString)
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

func Callback(c *fiber.Ctx) error {
	state := c.Query("state")

	if state != stateString {
		log.Printf("invalid oauth state, expected '%s', got '%s'\n", stateString, state)
		return c.SendStatus(http.StatusUnauthorized)
	}

	code := c.Query("code")

	token, err := oauthConf.Exchange(context.Background(), code)

	if err != nil {
		log.Printf("Code exchange failed with '%s'\n", err)
		return c.SendStatus(http.StatusUnauthorized)
	}

	client := oauthConf.Client(context.Background(), token)

	resp, err := client.Get("https://api.spotify.com/v1/me")

	if err != nil {
		c.SendStatus(http.StatusUnauthorized)
		c.JSON(fiber.Map{
			"status":  "error",
			"message": "Failed authenticating",
		})
	}

	defer resp.Body.Close()

	c.SendStatus(http.StatusOK)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "login successful",
	})
}
