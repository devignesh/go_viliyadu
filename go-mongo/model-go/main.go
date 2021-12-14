package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("quickstart")
	podcastsCollection := quickstartDatabase.Collection("podcasts")
	episodesCollection := quickstartDatabase.Collection("episodes")

	//dummy data
	mongopodcat := Podcast{
		Title:  "golang mongodb aggregation",
		Author: "vicky kumar",
		Tags:   []string{"mongo", "nosql"},
	}

	insert, err := podcastsCollection.InsertOne(ctx, mongopodcat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted %v!", insert.InsertedID)

	var podcast []Podcast
	podcastCursor, err := podcastsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = podcastCursor.All(ctx, &podcast); err != nil {
		log.Fatal(err)
	}
	log.Println(podcast)

	//episode

	var episode []Episode

	episodeCursor, err := episodesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = episodeCursor.All(ctx, &episode); err != nil {
		log.Fatal(err)
	}
	log.Println(episode[0].Title)

}
