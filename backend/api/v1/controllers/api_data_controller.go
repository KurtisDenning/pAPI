package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oliverschweikert/pAPI/backend/api/v1/services"
)

func CreateAPIData(ctx *gin.Context) {
	ctx.String(http.StatusOK, "POST : Create API Data Value")
	// createdAPIData := services.CreateAPIData()
}
func GetAPIDatas(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, services.GetAPIDatas())
}
func GetAPIData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, services.GetAPIData(ctx.Param("oid")))
}
func UpdateAPIData(ctx *gin.Context) {
	ctx.String(http.StatusOK, "PATCH : Update API Data Value %v", ctx.Param("id"))
	// updatedAPIData := services.UpdateAPIData()
}
func DeleteAPIData(ctx *gin.Context) {
	ctx.String(http.StatusOK, "DELETE : Delete API Data Value %v", ctx.Param("id"))
	// deletedAPIData := services.DeleteAPIData()
}
