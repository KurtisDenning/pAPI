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
	for _, category := range allCategories {
		ctx.String(http.StatusOK, "%v\n\n", category)
	}
}
func GetCategory(ctx *gin.Context) {
	ctx.String(http.StatusOK, "GET : Category %v", ctx.Param("id"))
	// category := services.GetCategory()
}
func UpdateCategory(ctx *gin.Context) {
	ctx.String(http.StatusOK, "PATCH : Update Category %v", ctx.Param("id"))
	// updatedCategory := services.UpdateCategory()
}
func DeleteCategory(ctx *gin.Context) {
	ctx.String(http.StatusOK, "DELETE : Delete Category %v", ctx.Param("id"))
	// deletedCategory := services.DeleteCategory()
}
