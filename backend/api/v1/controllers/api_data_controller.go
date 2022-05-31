package controllers

import (
	"net/http"
	"strconv"

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
func RefreshAPIData(ctx *gin.Context) {
	requestIds, hasIds := ctx.GetQueryArray("requestID")
	if !hasIds {
		ctx.String(http.StatusOK, "There were no request ID's supplied for object %v", ctx.Param("oid"))
	} else {
		requestIndexes := []int{}
		for _, s := range requestIds {
			i, err := strconv.Atoi(s)
			if err != nil {
				ctx.String(http.StatusOK, "Invalid request ID - please enter integers only")
				break
			} else {
				requestIndexes = append(requestIndexes, i)
			}
		}
		ctx.JSON(http.StatusOK, services.RefreshAPIData(ctx.Param("oid"), requestIndexes))
	}

	// updatedAPIData := services.UpdateAPIData()
}
func DeleteAPIData(ctx *gin.Context) {
	ctx.String(http.StatusOK, "DELETE : Delete API Data Value %v", ctx.Param("id"))
	// deletedAPIData := services.DeleteAPIData()
}
