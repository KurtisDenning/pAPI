package data

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAPIDatas() []bson.M {
	client, ctx, cancel, err := connect("mongodb+srv://db_admin:Disobey4-Reflux-Crying@apidata.ino8ejr.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		fmt.Println("Can't connect to MongoDB")
		return []bson.M{}
	}
	defer close(client, ctx, cancel)

	opts := options.Find().SetProjection(bson.D{{Key: "title", Value: 1}, {Key: "description", Value: 1}})
	apiData, err := client.Database("papi_db").Collection("apiData").Find(ctx, bson.D{}, opts)
	if err != nil {
		fmt.Println("Can't get API data from database")
		return []bson.M{}
	}

	var docs []bson.M
	for apiData.Next(ctx) {
		var doc bson.M
		err = apiData.Decode(&doc)
		if err != nil {
			fmt.Println("Can't convert to bson.M")
			continue
		}
		docs = append(docs, doc)
	}
	return docs
}
func GetAllAPIDatas() []bson.M {
	client, ctx, cancel, err := connect("mongodb+srv://db_admin:Disobey4-Reflux-Crying@apidata.ino8ejr.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		fmt.Println("Can't connect to MongoDB")
		return []bson.M{}
	}
	defer close(client, ctx, cancel)

	apiData, err := client.Database("papi_db").Collection("apiData").Find(ctx, bson.D{})
	if err != nil {
		fmt.Println("Can't get API data from database")
		return []bson.M{}
	}
	var docs []bson.M
	for apiData.Next(ctx) {
		var doc bson.M
		err = apiData.Decode(&doc)
		if err != nil {
			fmt.Println("Can't convert to bson.M")
			continue
		}
		docs = append(docs, doc)
	}
	return docs
}
func GetAPIData(oid primitive.ObjectID) bson.M {
	client, ctx, cancel, err := connect("mongodb+srv://db_admin:Disobey4-Reflux-Crying@apidata.ino8ejr.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		fmt.Println("Can't connect to MongoDB")
		return bson.M{}
	}
	defer close(client, ctx, cancel)

	category := client.Database("papi_db").Collection("apiData").FindOne(ctx, bson.M{"_id": oid})
	var doc bson.M
	category.Decode(&doc)
	fmt.Println(doc)
	return doc
}
