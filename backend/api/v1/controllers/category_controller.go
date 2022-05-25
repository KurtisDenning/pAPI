package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(ctx *gin.Context) {
	ctx.String(http.StatusOK, "POST : Create Category")
	// createdCategory := services.CreateCategory()
}
func GetCategories(ctx *gin.Context) {
	ctx.String(http.StatusOK, "GET : All Categories")
	// allCategories := services.GetCategories()
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
