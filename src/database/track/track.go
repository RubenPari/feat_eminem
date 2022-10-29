package track

import (
	"github.com/RubenPari/feat_eminem/src/database"
	"github.com/RubenPari/feat_eminem/src/models"
	"github.com/zmb3/spotify/v2"
	"log"
)

func Add(track models.Track) bool {
	db := database.GetDB()

	exist := db.QueryRow("SELECT * FROM tracks WHERE id = ?", track.Id.String())

	var trackFounded models.Track

	_ = exist.Scan(&trackFounded.Id, &trackFounded.Name, &trackFounded.Uri)

	if trackFounded.Id != spotify.ID("") {
		log.Default().Println("Track already exists")
		log.Default().Println(trackFounded)
		return false
	}

	result, errInsert := db.Exec("INSERT INTO tracks (id, name, uri, album, artist) VALUES (?, ?, ?, ?, ?)", track.Id, track.Name, track.Uri, track.Album, track.Artist)

	if errInsert != nil {
		log.Default().Println("Error inserting track")
		log.Default().Println(errInsert)
		return false
	}

	rowsAffected, _ := result.RowsAffected()

	return rowsAffected == int64(1)
}

func Adds(tracks []models.Track) bool {
	success := true

	for i := 0; i < len(tracks); i++ {
		added := Add(tracks[i])
		if !added {
			success = false
			log.Default().Println("Error inserting track n. ", i+1, "of ", len(tracks))
		}
	}

	return success
}
