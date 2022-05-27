package data

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAPIDatas() []bson.Raw {
	client, ctx, cancel, err := connect("mongodb+srv://db_admin:Disobey4-Reflux-Crying@apidata.ino8ejr.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
		data, _ := bson.Marshal(bson.M{"msg": "Cannot connect to MongoDB", "err": fmt.Sprintf("%v", err)})
		return []bson.Raw{bson.Raw(data)}
	}
	defer close(client, ctx, cancel)

	apiData, err := client.Database("papi_db").Collection("apiData").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		data, _ := bson.Marshal(bson.M{"msg": "Cannot find APIData collection in pAPI database", "err": fmt.Sprintf("%v", err)})
		return []bson.Raw{bson.Raw(data)}
	}
	var output []bson.Raw
	for apiData.Next(ctx) {
		output = append(output, apiData.Current)
	}
	return output
}
func GetAPIData(oid primitive.ObjectID) bson.Raw {
	client, ctx, cancel, err := connect("mongodb+srv://db_admin:Disobey4-Reflux-Crying@apidata.ino8ejr.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
		data, _ := bson.Marshal(bson.M{"msg": "Cannot connect to MongoDB", "err": fmt.Sprintf("%v", err)})
		return bson.Raw(data)
	}
	defer close(client, ctx, cancel)

	category := client.Database("papi_db").Collection("apiData").FindOne(ctx, bson.M{"_id": oid})
	var out bson.Raw
	category.Decode(&out)
	fmt.Println(out)
	return out
}
