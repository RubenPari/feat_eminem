package auth

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	spotifyOAuth "golang.org/x/oauth2/spotify"
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
		Endpoint:     spotifyOAuth.Endpoint,
	}, GenerateStateString()
}
