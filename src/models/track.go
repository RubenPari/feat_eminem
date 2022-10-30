package models

import (
	"github.com/zmb3/spotify/v2"
)

// Track NOTE: featuring attribute is a string concatenation
// of all name of the artist featuring separated by ', '
type Track struct {
	Id        spotify.ID  `json:"id"`
	Name      string      `json:"name"`
	Uri       spotify.URI `json:"uri"`
	Album     string      `json:"album"`
	Artist    string      `json:"artist"`
	Featuring string      `json:"featuring"`
}
