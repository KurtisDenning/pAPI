package services

import (
	"fmt"
	"io/ioutil"
	"net/http"

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
		return bson.M{"Message": err}
	}
	apiToUpdate := data.GetAPIData(id)
	apiModel, err := entity_models.NewAPIDataStruct(apiToUpdate)
	if err != nil {
		return bson.M{"Message": err}
	}
	url, req := apiModel.GetURL(requestIndexes)
	resp, err := http.Get(url)
	if err != nil {
		return bson.M{"Message": err}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bson.M{"Message": err}
	}
	err = req.UpdateResponse(body)
	if err != nil {
		return bson.M{"Message": err}
	}
	binData, err := bson.Marshal(bson.M{"requests": apiModel.Requests})
	if err != nil {
		return bson.M{"Message": err}
	}
	var bsonData bson.M
	bson.Unmarshal(binData, &bsonData)
	bsonResponse := data.UpdateAPIResponse(id, bsonData)
	return bsonResponse

}
func DeleteAPIData() {
}
