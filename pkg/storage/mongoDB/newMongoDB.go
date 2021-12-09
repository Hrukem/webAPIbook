package mongoDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewMongoDB() (*mongo.Collection, error) {
	var collection *mongo.Collection
	var ctx = context.TODO()

	clientOption := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Println("error create MongoDB in mongoDB.NewMongoDB()  ", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("error ping to MongoDB in mongoDB.NewMongoDB() ", err)
		return nil, err
	}
	fmt.Println("Ping to MongoDB ok!")

	collection = client.Database("webapibook").Collection("logging_actions")
	return collection, nil
}
