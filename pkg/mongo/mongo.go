package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://nqnlong:iz7gRgckiizWPDpg@cluster0.ujb303l.mongodb.net/"

var DB *mongo.Database

func ConnectDB() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}
	DB = client.Database("librabry")
}
