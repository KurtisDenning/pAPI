package entity_models

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	Request    string                 `json:"request" bson:"request"`
	Requests   []Request              `json:"requests,omitempty" bson:"requests,omitempty"`
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
	Requests     []Request          `json:"requests" bson:"requests"`
}

// Given a bson.M primitive from the apiData collected in MongoDB, make a new APIData struct
func NewAPIDataStruct(bsonData bson.M) (APIData, error) {
	var data APIData
	jsonData, err := bson.MarshalExtJSON(bsonData, false, false)
	if err != nil {
		return APIData{}, err
	}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (a APIData) GetURL(requestIndexes []int) string {
	url := a.Base
	currentRequests := a.Requests
	for _, r := range requestIndexes {
		currentRequest := currentRequests[r]
		url += currentRequest.Request
		currentRequests = currentRequests[r].Requests
	}
	return url
}
