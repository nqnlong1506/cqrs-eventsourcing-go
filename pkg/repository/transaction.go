package repository

import (
	"context"

	mongopkg "github.com/nqnlong1506/cqrs-event-sourcing-go/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var TransactionCollection *mongo.Collection

func InitTransactionCollection() {
	TransactionCollection = mongopkg.DB.Collection("transaction")
	TransactionCollection.InsertOne(context.Background(), bson.M{
		"info": "Transaction Collection",
	})
}
