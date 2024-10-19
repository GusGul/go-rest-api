package v2

import (
	"database/sql"
	"encoding/json"
	"go-rest-api/internal/db"
	"go-rest-api/pkg/models"
	"net/http"
	"strings"
)

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	artist := r.URL.Query().Get("artist")
	title := r.URL.Query().Get("title")

	var rows *sql.Rows
	var err error

	query := "SELECT al.id, al.title, ar.name AS artistName, al.price FROM albums al JOIN artists ar ON al.artistId = ar.id"
	var queryParams []interface{}

	if artist != "" || title != "" {
		query += " WHERE"
		conditions := []string{}

		if artist != "" {
			conditions = append(conditions, "ar.name = ?")
			queryParams = append(queryParams, artist)
		}

		if title != "" {
			conditions = append(conditions, "al.title = ?")
			queryParams = append(queryParams, title)
		}

		query += " " + strings.Join(conditions, " AND ")
	}

	rows, err = db.DB.Query(query, queryParams...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

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
