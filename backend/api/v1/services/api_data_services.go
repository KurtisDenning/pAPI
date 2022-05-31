package services

import (
	"fmt"

	"github.com/oliverschweikert/pAPI/backend/api/v1/data"
	"github.com/oliverschweikert/pAPI/backend/api/v1/entity_models"
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
func UpdateAPIData(oid string) {
}
func RefreshAPIData(oid string, requestIndexes []int) bson.M {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return bson.M{"Message": "Invalid Object ID"}
	} else {
		apiToUpdate := data.GetAPIData(id)
		apiModel, err := entity_models.NewAPIDataStruct(apiToUpdate)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(apiModel.GetURL(requestIndexes))
		}
		return apiToUpdate
	}
}
func DeleteAPIData() {
}
