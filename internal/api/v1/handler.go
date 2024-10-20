package v1

import (
	"encoding/json"
	"go-rest-api/internal/db"
	"go-rest-api/pkg/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func GetAlbumsAverageByGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	genreParam := chi.URLParam(r, "genre")
	if genreParam == "" {
		http.Error(w, "Invalid genre", http.StatusBadRequest)
		return
	}

	var total, average, count float64
	for _, album := range db.Albums {
		if strings.ToLower(album.Genre) != strings.ToLower(genreParam) {
			continue
		}
		total += album.Price
		count++
	}

	if count == 0 {
		http.Error(w, "No albums found for the specified genre", http.StatusNotFound)
		return
	}

	average = total / count

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]float64{"average": average})
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func GetAlbums(w http.ResponseWriter, r *http.Request) {
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

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid album ID", http.StatusBadRequest)
		return
	}

	var updatedAlbum models.Album
	if err := json.NewDecoder(r.Body).Decode(&updatedAlbum); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, album := range db.Albums {
		if album.ID == id {
			db.Albums[i] = updatedAlbum
			db.SaveDatabase("db/database.json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(updatedAlbum)
			if err != nil {
				return
			}
			return
		}
	}

	http.Error(w, "Album not found", http.StatusNotFound)
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid album ID", http.StatusBadRequest)
		return
	}

	for i, album := range db.Albums {
		if album.ID == id {
			db.Albums = append(db.Albums[:i], db.Albums[i+1:]...)
			db.SaveDatabase("db/database.json")
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "Album not found", http.StatusNotFound)
}
