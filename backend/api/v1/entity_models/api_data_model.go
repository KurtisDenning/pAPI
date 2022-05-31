package entity_models

import (
	"time"

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
