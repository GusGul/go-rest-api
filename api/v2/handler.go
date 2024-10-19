package v2

import (
	"encoding/json"
	"go-rest-api/internal/db"
	"go-rest-api/pkg/models"
	"net/http"
)

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.DB.Query("SELECT id, title, artist, price FROM albums")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Next()

	var albums []models.Album
	for rows.Next() {
		var album models.Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		albums = append(albums, album)
	}

	if err := json.NewEncoder(w).Encode(albums); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
