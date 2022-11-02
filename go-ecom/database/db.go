package database

import "go.mongodb.org/mongo-driver/mongo"

func DBSet() *mongo.Client {

}

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {}
