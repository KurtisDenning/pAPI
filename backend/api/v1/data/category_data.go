package data

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCategories() []bson.M {
	client, ctx, cancel, err := connect("mongodb+srv://db_admin:Disobey4-Reflux-Crying@apidata.ino8ejr.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		fmt.Printf("Error connecting to MongoDB\n%v\n", err)
		return []bson.M{}
	}
	defer close(client, ctx, cancel)

	categories, err := client.Database("papi_db").Collection("categories").Find(ctx, bson.D{})
	if err != nil {
		fmt.Printf("Error getting collection from database\n%v\n", err)
		return []bson.M{}
	}

	var docs []bson.M
	for categories.Next(ctx) {
		var doc bson.M
		err := categories.Decode(&doc)
		if err != nil {
			fmt.Println(err)
			continue
		}
		docs = append(docs, doc)
	}
	return docs
}

func GetCategory(oid primitive.ObjectID) bson.M {
	client, ctx, cancel, err := connect("mongodb+srv://db_admin:Disobey4-Reflux-Crying@apidata.ino8ejr.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}
	defer close(client, ctx, cancel)

	category := client.Database("papi_db").Collection("categories").FindOne(ctx, bson.M{"_id": oid})
	var doc bson.M
	err = category.Decode(&doc)
	if err != nil {
		fmt.Println(err)
	}
	return doc
}
