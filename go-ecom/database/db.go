package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
	Mongodb := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(Mongodb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect to mongodb")
		return nil
	}

	fmt.Println("Successfully Connected to the mongodb")
	return client

}

var ClientMongo *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {

	var usercollection *mongo.Collection = client.Database("e-com").Collection(collectionName)
	return usercollection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {

	var productcollection *mongo.Collection = client.Database("e-com").Collection(collectionName)
	return productcollection
}
