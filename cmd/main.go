package main

import (
	v1 "go-rest-api/api/v1"
	v2 "go-rest-api/api/v2"
	"go-rest-api/internal/db"
	"log"
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
	router.Put("/albums/{id}", v1.UpdateAlbum)
	router.Delete("/albums/{id}", v1.DeleteAlbum)

	router.Get("/albums", v2.GetAlbums)
	router.Get("/albums/{id}", v2.GetAlbumById)

	db.LoadDatabase("db/database.json")
	err := db.InitDatabase("gopher:Gopher@tcp(localhost:3306)/golang")
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	err = http.ListenAndServe("localhost:8080", router)
	if err != nil {
		return
	}
}
