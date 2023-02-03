package main

import (
	"fmt"
	"gin_ws/repository"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DB repository.MongoDatabase
}

func main() {
	// set application config
	var app application

	client, err := app.initiateMongo()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = repository.MongoDatabase{Client: client}
	// Release resource when the main
	// function is returned.
	defer app.disposeMongo(client)
	// start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
