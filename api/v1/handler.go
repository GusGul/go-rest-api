package v1

import (
	"encoding/json"
	"go-gin-api/internal/db"
	"net/http"
)

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.Albums)
}
