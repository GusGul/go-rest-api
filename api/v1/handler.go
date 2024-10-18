package v1

import (
    "net/http"
    "go-gin-api/internal/db"
    "github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, db.Albums)
}