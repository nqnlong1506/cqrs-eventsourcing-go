package repository

import (
	"context"

	mongopkg "github.com/nqnlong1506/cqrs-event-sourcing-go/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection

func InitUserCollection() {
	UserCollection = mongopkg.DB.Collection("user")
	UserCollection.InsertOne(context.Background(), bson.M{
		"info": "User Collection",
	})
}
