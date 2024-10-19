package v1

import (
	"encoding/json"
	"go-rest-api/internal/db"
	"go-rest-api/pkg/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetAlbums(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(db.Albums)
	if err != nil {
		return
	}
}

func GetAlbumById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid album ID", http.StatusBadRequest)
		return
	}

	for _, album := range db.Albums {
		if album.ID == id {
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(album)
			if err != nil {
				return
			}
			return
		}
	}

	http.Error(w, "Album not found", http.StatusNotFound)
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var newAlbum models.Album
	if err := json.NewDecoder(r.Body).Decode(&newAlbum); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db.Albums = append(db.Albums, newAlbum)
	db.SaveDatabase("db/database.json")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(newAlbum)
	if err != nil {
		return
	}
}
