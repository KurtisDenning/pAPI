package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// func ping(client *mongo.Client, ctx context.Context) error {
// 	if err := client.Ping(ctx, readpref.Primary()); err != nil {
// 		return err
// 	}
// 	fmt.Println("connected successfully")
// 	return nil
// }

func GetAllDatabases() []mongo.DatabaseSpecification {
	client, ctx, cancel, err := connect("mongodb+srv://db_admin:Disobey4-Reflux-Crying@apidata.ino8ejr.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	databases, err := client.ListDatabases(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	return databases.Databases
}

func GetAllCollections() []string {
	client, ctx, cancel, err := connect("mongodb+srv://db_admin:Disobey4-Reflux-Crying@apidata.ino8ejr.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	collections, err := client.Database("papi_db").ListCollectionNames(ctx, bson.D{})
	if err != nil {
		panic(err)
	} else {
		return collections
	}
}

func GetCategories() []bson.Raw {
	client, ctx, cancel, err := connect("mongodb+srv://db_admin:Disobey4-Reflux-Crying@apidata.ino8ejr.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	categories, err := client.Database("papi_db").Collection("categories").Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	var output []bson.Raw
	for categories.Next(ctx) {
		output = append(output, categories.Current)
	}
	return output
}

func GetAPIData() []bson.Raw {
	client, ctx, cancel, err := connect("mongodb+srv://db_admin:Disobey4-Reflux-Crying@apidata.ino8ejr.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	apiData, err := client.Database("papi_db").Collection("apiData").Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	var output []bson.Raw
	for apiData.Next(ctx) {
		output = append(output, apiData.Current)
	}
	return output
}
