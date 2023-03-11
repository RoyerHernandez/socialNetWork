package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoCN = connectionDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://royhernandez:Israfel2112@clustersocialnetwork.i0l2kku.mongodb.net/?retryWrites=true&w=majority")

func connectionDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Fatal("connection successful to data base")
	return client
}

func ConnectionCheck() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
