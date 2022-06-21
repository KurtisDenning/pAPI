package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func GetAllCollections() []string {
	client, ctx, cancel, err := connect("mongodb+srv://db_admin:Disobey4-Reflux-Crying@apidata.ino8ejr.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
		return []string{"Cannot connect to MongoDB", fmt.Sprintf("%v", err)}
	}
	defer close(client, ctx, cancel)

	collections, err := client.Database("papi_db").ListCollectionNames(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return []string{"Cannot connect to pAPI database", fmt.Sprintf("%v", err)}

	}
	return collections
}
