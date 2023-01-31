package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/albums", app.GetAlbums)
	router.POST("/albums", app.PostAlbums)
	// router.GET("/clientIP", func(c *gin.Context) {
	// 	// If the client is 192.168.1.2, use the X-Forwarded-For
	// 	// header to deduce the original client IP from the trust-
	// 	// worthy parts of that header.
	// 	// Otherwise, simply return the direct client IP
	// 	fmt.Printf("ClientIPAdress: %s\n", c.ClientIP())
	// })

	return router
}
