package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	ShortDesc string             `json:"shortDesc" bson:"shortDesc"`
	LongDesc  string             `json:"longDesc" bson:"longDesc"`
}
