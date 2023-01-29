package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct{}

func main() {
	// set application config
	var app application

	// Get Client, Context, CancelFunc and
	// err from connect method.
	client, ctx, cancel, err := app.databaseConnect("mongodb://mongo:27017")
	if err != nil {
		panic(err)
	}

	if err != nil {
		print("connect error", err)
	}
	// Release resource when the main
	// function is returned.
	defer app.databaseDisconnect(client, ctx, cancel)

	// Check mongoDB connection with Ping method
	app.checkDatabaseConnection(client, ctx)

	// start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
