package repository

import (
	"context"
	"gin_ws/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const timeout = time.Second * 5

type MongoDatabase struct {
	Client *mongo.Client
}

func (m *MongoDatabase) Sample() (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var album = models.Album{ID: "1", Title: "OMG", Artist: "NewJeans", Released: 2023}

	collection := m.Client.Database("db").Collection("Album")
	res, err := collection.InsertOne(ctx, album)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

func (m *MongoDatabase) AllAlbums() ([]models.Album, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	collection := m.Client.Database("db").Collection("Album")
	cur, err := collection.Find(ctx, bson.D{{}})

	var results []models.Album
	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem models.Album
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)

	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return results, nil
}
