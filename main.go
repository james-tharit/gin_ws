package main

import (
	"fmt"
	api "gin_ws/api/album"
	db "gin_ws/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	// Get Client, Context, CancelFunc and
	// err from connect method.
	client, ctx, cancel, err := db.Connect("mongodb://mongo:27017")
	if err != nil {
		panic(err)
	}

	if err != nil {
		print("connect error", err)
	}
	// Release resource when the main
	// function is returned.
	defer db.Close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	db.Ping(client, ctx)
	router.GET("/albums", api.GetAlbums)
	router.POST("/albums", api.PostAlbums)
	router.GET("/clientIP", func(c *gin.Context) {
		// If the client is 192.168.1.2, use the X-Forwarded-For
		// header to deduce the original client IP from the trust-
		// worthy parts of that header.
		// Otherwise, simply return the direct client IP
		fmt.Printf("ClientIPAdress: %s\n", c.ClientIP())
	})
	router.Run("0.0.0.0:8080")
}
