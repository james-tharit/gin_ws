package api

import (
	"net/http"

	"gin_ws/models"

	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
var albums = []models.Album{
	{ID: "1", Title: "OMG", Artist: "NewJeans", Released: 2023},
	{ID: "2", Title: "Ditto", Artist: "NewJeans", Released: 2022},
	{ID: "3", Title: "Attention", Artist: "NewJeans", Released: 2022},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
