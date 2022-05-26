package services

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverschweikert/pAPI/backend/api/v1/data"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateAPIData(ctx *gin.Context) {
}
func GetAPIDatas() []bson.Raw {
	return data.GetAPIData()
}
func GetAPIData(ctx *gin.Context) {
}
func UpdateAPIData(ctx *gin.Context) {
}
func DeleteAPIData(ctx *gin.Context) {
}
