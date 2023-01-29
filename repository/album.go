package repository

import "go.mongodb.org/mongo-driver/mongo"

type MongoDatabase struct {
	Client *mongo.Client
}
