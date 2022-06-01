package entity_models

import (
	"encoding/json"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	Request    string                 `json:"request" bson:"request"`
	Requests   []*Request             `json:"requests,omitempty" bson:"requests,omitempty"`
	Response   map[string]interface{} `json:"response,omitempty" bson:"response,omitempty"`
	LastUpdate time.Time              `json:"lastUpdate,omitempty" bson:"lastUpdate,omitmpty"`
}

type APIData struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id"`
	Title        string             `json:"title" bson:"title"`
	Description  string             `json:"description" bson:"description"`
	ExternalURL  string             `json:"externalURL" bson:"externalURL"`
	DailyCount   int                `json:"dailyCount" bson:"dailyCount"`
	WeeklyCount  int                `json:"weeklyCount" bson:"weeklyCount"`
	MonthlyCount int                `json:"monthlyCount" bson:"monthlyCount"`
	YearlyCount  int                `json:"yearlyCount" bson:"yearlyCount"`
	TotalCount   int                `json:"totalCount" bson:"totalCount"`
	Categories   []string           `json:"categories" bson:"categories"`
	Base         string             `json:"base" bson:"base"`
	Requests     []*Request         `json:"requests" bson:"requests"`
}

// Given a bson.M primitive from the apiData collected in MongoDB, make a new APIData struct
func NewAPIDataStruct(bsonData bson.M) (APIData, error) {
	var data APIData
	jsonData, err := bson.Marshal(bsonData)
	if err != nil {
		return APIData{}, err
	}
	err = bson.Unmarshal(jsonData, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (a APIData) GetURL(requestIndexes []int) (string, *Request) {
	url := a.Base
	currentRequests := a.Requests
	currentRequest := currentRequests[requestIndexes[0]]
	url += currentRequest.Request
	for r := 1; r < len(requestIndexes); r++ {
		currentRequests = currentRequest.Requests
		currentRequest = currentRequests[requestIndexes[r]]
		url += currentRequest.Request
	}
	return url, currentRequest
}

func (r *Request) UpdateResponse(body []byte) error {
	if time.Since(r.LastUpdate).Hours() > 24 {
		var res any
		err := json.Unmarshal(body, &res)
		if err != nil {
			return err
		}
		if reflect.TypeOf(res).Kind() != reflect.Map {
			r.Response = bson.M{"data": res}
		} else {
			r.Response = res.(map[string]interface{})
		}
		r.LastUpdate = time.Now()
	}
	return nil
}
