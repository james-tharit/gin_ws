package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (app *application) initiateMongo() (*mongo.Client, error) {
	// Get Client, Context, CancelFunc and
	// err from connect method.
	client, ctx, _, err := connect("mongodb://root:example@mongo:27017")
	if err != nil {
		panic(err)
	}

	if err != nil {
		print("connect error", err)
	}

	// Check mongoDB connection with Ping method
	err = connectionStatus(client, ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (app *application) disposeMongo(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func connectionStatus(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("Mongo connected successfully")
	return nil
}
