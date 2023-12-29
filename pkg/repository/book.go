package repository

import (
	"context"

	database "github.com/nqnlong1506/cqrs-event-sourcing-go/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var BookCollection *mongo.Collection

func InitBookCollection() {
	BookCollection = database.DB.Collection("book")
	BookCollection.InsertOne(context.Background(), bson.M{
		"info": "book collection",
	})
}
