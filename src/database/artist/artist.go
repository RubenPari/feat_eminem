package artist

import (
	"log"

	"github.com/RubenPari/feat_eminem/src/database"
	"github.com/RubenPari/feat_eminem/src/models"
	"github.com/zmb3/spotify/v2"
)

func Add(artist models.Artist) bool {
	db := database.GetDB()

	exist := db.QueryRow("SELECT * FROM artists WHERE id = ?", artist.Id.String())

	var artistFounded models.Artist

	_ = exist.Scan(&artistFounded.Id, &artistFounded.Name, &artistFounded.Uri)

	if artistFounded.Id != spotify.ID("") {
		log.Default().Println("Artist already exists")
		log.Default().Println(artistFounded)
		return false
	}

	result, errInsert := db.Exec("INSERT INTO artists (id, name, uri) VALUES (?, ?, ?)", artist.Id, artist.Name, artist.Uri)

	if errInsert != nil {
		log.Default().Println("Error inserting artist")
		log.Default().Println(errInsert)
		return false
	}

	rowsAffected, _ := result.RowsAffected()

	return rowsAffected == int64(1)
}
