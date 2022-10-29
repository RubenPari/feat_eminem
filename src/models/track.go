package models

import (
	"github.com/zmb3/spotify/v2"
)

type Track struct {
	Id        spotify.ID  `json:"id"`
	Name      string      `json:"name"`
	Uri       spotify.URI `json:"uri"`
	Album     string      `json:"album"`
	Artist    string      `json:"artist"`
	Featuring []string    `json:"featuring"`
}
