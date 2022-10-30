package track

import (
	"github.com/RubenPari/feat_eminem/src/database"
	"github.com/RubenPari/feat_eminem/src/models"
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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

	result, errInsert := db.Exec("INSERT INTO tracks (id, name, uri, album, artist, featuring) VALUES (?, ?, ?, ?, ?, ?)", track.Id, track.Name, track.Uri, track.Album, track.Artist, track.Featuring)

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

func AddFeatured(track models.Track) bool {
	db := database.GetDB()

	result, errInsert := db.Exec("INSERT INTO tracks_feat (id, name, uri, album, artist, featuring) VALUES (?, ?, ?, ?, ?, ?)", track.Id, track.Name, track.Uri, track.Album, track.Artist, track.Featuring)

	if errInsert != nil {
		log.Default().Println("Error inserting featured track")
		log.Default().Println(errInsert)
		return false
	}

	rowsAffected, _ := result.RowsAffected()

	return rowsAffected == int64(1)
}

func AddsFeatured(tracks []models.Track) bool {
	success := true

	for i := 0; i < len(tracks); i++ {
		added := AddFeatured(tracks[i])
		if !added {
			success = false
			log.Default().Println("Error inserting featured track n. ", i+1, "of ", len(tracks))
		}
	}

	return success
}

func GetAllByArtist(id string) []models.Track {
	db := database.GetDB()

	_ = godotenv.Load()
	port := os.Getenv("PORT")

	// get name of artist by id with endpoint
	responseName, errGetName := http.Get("http://localhost:" + port + "/utils/artist/get-name/" + id)
	if errGetName != nil {
		log.Default().Println("Error getting name of artist")
		log.Default().Println(errGetName)
	}

	// TODO: change this to an appropriate way that is more short
	// extract name of artist from response of type json
	responseNameBytes, _ := io.ReadAll(responseName.Body)
	_ = responseName.Body.Close()
	responseNameString := string(responseNameBytes)
	_, name, _ := strings.Cut(responseNameString, "{\"name\":\"")
	name = name[:len(name)-2]

	rows, err := db.Query("SELECT * FROM tracks WHERE artist = ?", name)

	if err != nil {
		log.Default().Println("Error getting all tracks")
		log.Default().Println(err)
	}

	var tracks []models.Track

	for rows.Next() {
		var track models.Track
		_ = rows.Scan(&track.Id, &track.Name, &track.Uri, &track.Album, &track.Artist, &track.Featuring)
		tracks = append(tracks, track)
	}

	return tracks
}

func IsFeaturing(track models.Track) bool {
	artistsFeaturing := track.Featuring

	if strings.Contains(artistsFeaturing, "Eminem") {
		return true
	} else {
		return false
	}
}

func FilterByFeaturing(tracks []models.Track) []models.Track {
	var tracksFiltered []models.Track

	for i := 0; i < len(tracks); i++ {
		if IsFeaturing(tracks[i]) {
			tracksFiltered = append(tracksFiltered, tracks[i])
		}
	}

	return tracksFiltered
}
