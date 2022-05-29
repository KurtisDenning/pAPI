package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oliverschweikert/pAPI/backend/api/v1/services"
)

func CreateCategory(ctx *gin.Context) {
	ctx.String(http.StatusOK, "POST : Create Category")
	// createdCategory := services.CreateCategory()
}
func GetCategories(ctx *gin.Context) {
	allCategories := services.GetCategories()
	ctx.JSON(http.StatusOK, allCategories)
}
func GetCategory(ctx *gin.Context) {
	category := services.GetCategory(ctx.Param("oid"))
	ctx.String(http.StatusOK, category.String())
}
func UpdateCategory(ctx *gin.Context) {
	ctx.String(http.StatusOK, "PATCH : Update Category %v", ctx.Param("id"))
	// updatedCategory := services.UpdateCategory()
}
func DeleteCategory(ctx *gin.Context) {
	ctx.String(http.StatusOK, "DELETE : Delete Category %v", ctx.Param("id"))
	// deletedCategory := services.DeleteCategory()
}
