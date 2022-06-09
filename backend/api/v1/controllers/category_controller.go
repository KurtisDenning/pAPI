package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oliverschweikert/pAPI/backend/api/v1/services"
)

// GetCategories godoc
// @Summary Get all categories
// @Description Get a list of all categories that are currently stored in the database
// @tags Category
// @Produce json
// @Success 200 {array} entity_models.Category "Successful request"
// @Router /categories [get]
func GetCategories(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, services.GetCategories())
}

// GetCategory godoc
// @Summary Get a category
// @Description Get a single category from the database
// @tags Category
// @Produce json
// @Param oid path string true "The Object ID (_id) of the category"
// @Success 200 {object} entity_models.Category "Successful request"
// @Failure 400 {object} entity_models.Message "Invalid object ID"
// @Router /categories/{oid} [get]
func GetCategory(ctx *gin.Context) {
	code, response := services.GetCategoryResponse(ctx.Param("oid"))
	ctx.JSON(code, response)
}
