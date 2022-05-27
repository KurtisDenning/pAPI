package services

import (
	"fmt"

	"github.com/oliverschweikert/pAPI/backend/api/v1/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCategory() {
}
func GetCategories() []bson.Raw {
	return data.GetCategories()
}
func GetCategory(oid string) bson.Raw {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		fmt.Println("The provided hex value is not a valid object ID")
		return bson.Raw{}
	}
	return data.GetCategory(id)
}
func UpdateCategory() {
}
func DeleteCategory() {
}
