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
		log.Fatal(err)
	}
	defer close(client, ctx, cancel)

	categories, err := client.Database("papi_db").Collection("categories").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var docs []bson.M
	for categories.Next(ctx) {
		var doc bson.M
		err := categories.Decode(&doc)
		if err != nil {
			fmt.Println(err)
		}
		docs = append(docs, doc)
	}
	return docs
}

func GetCategory(oid primitive.ObjectID) bson.Raw {
	client, ctx, cancel, err := connect("mongodb+srv://db_admin:Disobey4-Reflux-Crying@apidata.ino8ejr.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
		data, _ := bson.Marshal(bson.M{"msg": "Cannot connect to MongoDB", "err": fmt.Sprintf("%v", err)})
		return bson.Raw(data)
	}
	defer close(client, ctx, cancel)

	category := client.Database("papi_db").Collection("categories").FindOne(ctx, bson.M{"_id": oid})
	var out bson.Raw
	category.Decode(&out)
	return out
}
