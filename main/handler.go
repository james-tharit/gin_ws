package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) GetAlbums(c *gin.Context) {
	albums, err := app.DB.AllAlbums()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func (app *application) PostAlbums(c *gin.Context) {
	result, err := app.DB.Sample()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusCreated, result)

}
