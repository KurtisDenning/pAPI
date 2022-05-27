package services

import (
	"fmt"

	"github.com/oliverschweikert/pAPI/backend/api/v1/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAPIData() {
}
func GetAPIDatas() []bson.Raw {
	return data.GetAPIDatas()
}
func GetAPIData(oid string) bson.Raw {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		fmt.Println("The provided hex value is not a valid object ID")
		return bson.Raw{}
	}
	return data.GetAPIData(id)
}
func UpdateAPIData() {
}
func DeleteAPIData() {
}
