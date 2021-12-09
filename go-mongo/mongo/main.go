package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	//playing around create collection and db in mongo
	quickstartDatabase := client.Database("quickstart")
	podcastsCollection := quickstartDatabase.Collection("podcasts")
	episodesCollection := quickstartDatabase.Collection("episodes")

	// inserts a data into podcasts collection here bson.D is document and bson.A is array
	// podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{
	// 	{"title", "The Polyglot Developer Podcast"},
	// 	{"author", "Nic Raboy"},
	// 	{"tags", bson.A{"development", "programming", "coding"}},
	// })

	// // //insert a data into eposodes collection
	// episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
	// 	bson.D{
	// 		{"podcast", podcastResult.InsertedID},
	// 		{"title", "GraphQL for API Development"},
	// 		{"description", "Learn about GraphQL from the co-creator of GraphQL, Lee Byron."},
	// 		{"duration", 25},
	// 	},
	// 	bson.D{
	// 		{"podcast", podcastResult.InsertedID},
	// 		{"title", "Progressive Web Application Development"},
	// 		{"description", "Learn about PWA development with Tara Manicsic."},
	// 		{"duration", 32},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Inserted %v documents into episode collection!\n", len(episodeResult.InsertedIDs))

	//finding single
	var podcast bson.M
	if err = podcastsCollection.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
		log.Fatal(err)
	}
	fmt.Println("podcast", podcast)

	//finding all
	cursor, err := episodesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println("episodes", episodes)

	// delete documets
	result, err := podcastsCollection.DeleteOne(ctx, bson.M{"title": "The Polyglot Developer Podcast"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)

	result, err = episodesCollection.DeleteMany(ctx, bson.M{"duration": 25})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteMany removed %v document(s)\n", result.DeletedCount)

	//list the number of database available in mongodb cluster
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

}
