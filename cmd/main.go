package main

import (
	v1 "go-gin-api/api/v1"
	"go-gin-api/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", v1.GetAlbums)

	db.LoadDatabase("db/database.json")

	router.Run("localhost:8080")
}
