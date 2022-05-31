package services

import (
	"fmt"

	"github.com/oliverschweikert/pAPI/backend/api/v1/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAPIData() {
}
func GetAPIDatas() []bson.M {
	return data.GetAPIDatas()
}
func GetAllAPIDatas() []bson.M {
	return data.GetAllAPIDatas()
}
func GetAPIData(oid string) bson.M {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		fmt.Println("The provided hex value is not a valid object ID")
		return bson.M{"Message": "The provided hex value is not a valid object ID"}
	}
	return data.GetAPIData(id)
}
func UpdateAPIData() {
}
func DeleteAPIData() {
}
