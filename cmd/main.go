package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jluo926/GoProjects/web-service/internal/album"
)

// albums slice to seed record album data.
func main() {
	router := gin.Default()
	router.GET("/albums", album.GetAlbums)
	router.POST("/albums", album.PostAlbums)
	router.GET("/albums/:id", album.GetAlbumByID)
	router.Run("localhost:6001")
}
