package spotify

import (
	"context"

	authCONTR "github.com/RubenPari/feat_eminem/src/controllers/auth"
	spotifyAPI "github.com/zmb3/spotify/v2"
	spotifyAUTH "github.com/zmb3/spotify/v2/auth"
)

func GetClient() (*spotifyAPI.Client, context.Context) {
	ctx := context.Background()
	token := authCONTR.AccessToken

	// create http client
	httpClient := spotifyAUTH.New().Client(ctx, token)

	// return spotify http client
	return spotifyAPI.New(httpClient), ctx
}
