package main

import (
	"fmt"
	api "gin_ws/api/album"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/albums", api.GetAlbums)
	router.POST("/albums", api.PostAlbums)
	router.GET("/", func(c *gin.Context) {
		// If the client is 192.168.1.2, use the X-Forwarded-For
		// header to deduce the original client IP from the trust-
		// worthy parts of that header.
		// Otherwise, simply return the direct client IP
		fmt.Printf("ClientIPAdress: %s\n", c.ClientIP())
	})
	router.Run("0.0.0.0:8080")
}
