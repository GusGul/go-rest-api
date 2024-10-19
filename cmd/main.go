package main

import (
	v1 "go-rest-api/api/v1"
	"go-rest-api/internal/db"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/albums", v1.GetAlbums)
	router.Get("/albums/{id}", v1.GetAlbumById)
	router.Post("/albums", v1.CreateAlbum)

	db.LoadDatabase("db/database.json")

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		return
	}
}
