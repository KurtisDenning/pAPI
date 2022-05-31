package services

import (
	"fmt"

	"github.com/oliverschweikert/pAPI/backend/api/v1/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCategory() {
}
func GetCategories() []bson.M {
	bsonData := data.GetCategories()
	fmt.Println(bsonData)
	return bsonData
}
func GetCategory(oid string) bson.M {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		fmt.Println("The provided hex value is not a valid object ID")
		return bson.M{"Message": "The provided hex value is not a valid object ID"}
	}
	return data.GetCategory(id)
}
func UpdateCategory() {
}
func DeleteCategory() {
}
