package services

import (
	"io/ioutil"
	"net/http"

	"github.com/oliverschweikert/pAPI/backend/api/v1/data"
	"github.com/oliverschweikert/pAPI/backend/api/v1/entity_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAPIDatas() []bson.M {
	return data.GetAPIDatas()
}
func GetAPIData(oid string) (int, bson.M) {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return http.StatusBadRequest, bson.M{"Message": "The provided hex value is not a valid object ID"}
	}
	return http.StatusOK, data.GetAPIData(id)
}
func RefreshAPIData(oid string, requestIndexes []int) (int, bson.M) {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return http.StatusBadRequest, bson.M{"Message": "The provided hex value is not a valid object ID"}
	}
	apiToUpdate := data.GetAPIData(id)
	apiModel, err := entity_models.NewAPIDataStruct(apiToUpdate)
	if err != nil {
		return http.StatusInternalServerError, bson.M{"Message": err.Error()}
	}
	url, req, err := apiModel.GetURL(requestIndexes)
	if err != nil {
		return http.StatusBadRequest, bson.M{"Message": err.Error()}
	}
	resp, err := http.Get(url)
	if err != nil {
		return http.StatusBadGateway, bson.M{"Message": err.Error()}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusBadGateway, bson.M{"Message": err.Error()}
	}
	err = req.UpdateResponse(body)
	if err != nil {
		return http.StatusInternalServerError, bson.M{"Message": err.Error()}
	}
	binData, err := bson.Marshal(bson.M{"requests": apiModel.Requests})
	if err != nil {
		return http.StatusInternalServerError, bson.M{"Message": err.Error()}
	}
	var bsonData bson.M
	bson.Unmarshal(binData, &bsonData)
	data.UpdateAPIResponse(id, bsonData)
	reqData, err := bson.Marshal(req)
	if err != nil {
		return http.StatusInternalServerError, bson.M{"Message": err.Error()}
	}
	var reqResponse bson.M
	bson.Unmarshal(reqData, &reqResponse)
	return http.StatusOK, reqResponse
}
