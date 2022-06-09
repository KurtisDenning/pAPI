package services

import (
	"fmt"
	"net/http"

	"github.com/oliverschweikert/pAPI/backend/api/v1/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCategory() {
}
func GetCategories() []bson.M {
	bsonData := data.GetCategories()
	fmt.Println(bsonData)
	return bsonData
}
func GetCategoryResponse(oid string) (int, bson.M) {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		fmt.Println("The provided hex value is not a valid object ID")
		return http.StatusBadRequest, bson.M{"message": fmt.Sprintf("%v is not a valid object ID", oid)}
	}
	return http.StatusOK, data.GetCategory(id)
}
func UpdateCategory() {
}
func DeleteCategory() {
}
