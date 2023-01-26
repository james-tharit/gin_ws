package main

import (
	api "gin_ws/api/album"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", api.GetAlbums)
	router.POST("/albums", api.PostAlbums)
	router.Run("0.0.0.0:8080")
}
