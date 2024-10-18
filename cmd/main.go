package main

import (
	v1 "go-gin-api/api/v1"
	"go-gin-api/internal/db"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/albums", v1.GetAlbums)

	db.LoadDatabase("db/database.json")

	http.ListenAndServe("localhost:8080", router)
}
