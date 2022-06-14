package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oliverschweikert/pAPI/backend/api/v1/services"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAPIDatas godoc
// @Summary Get all API data
// @Description Get a list of all API's that are currently stored in the database, and an overview of their information
// @tags APIData
// @Produce json
// @Success 200 {array} entity_models.APIDataOverview "Successful request."
// @Failure 429 {string} string "Too many requests - please only test once every 10 seconds"
// @Router /apidata [get]
func GetAPIDatas(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, services.GetAPIDatas())
}

// GetAPIData godoc
// @Summary Get single API data
// @Description Get complete information about a single API stored in the database
// @tags APIData
// @Produce json
// @Param oid path string true "The Object ID (_id) of the API data"
// @Param requestID query int false "The request ID to check for - is recursive in practice"
// @Success 200 {object} entity_models.APIData "Successful request. Please note that requests are recursive, and can have many request layers"
// @Failure 400 {object} entity_models.Message "Problem with user request - read message for more details"
// @Failure 500 {object} entity_models.Message "Problem with server processing - read message for more details"
// @Failure 502 {object} entity_models.Message "Problem connecting to external API - read message for more details"
// @Failure 429 {string} string "Too many requests - please only test once every 10 seconds"
// @Router /apidata/{oid} [get]
func GetAPIData(ctx *gin.Context) {
	code, response := services.GetAPIData(ctx.Param("oid"))
	ctx.JSON(code, response)

}

func GetAPIDataRequest(ctx *gin.Context) {
	request, err := strconv.Atoi(ctx.Param("request"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, bson.M{"Error": err.Error()})
	} else {
		code, response := services.RefreshAPIData(ctx.Param("oid"), request)
		ctx.JSON(code, response)
	}
}
