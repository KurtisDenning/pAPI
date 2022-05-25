package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func enableCategoryRoutes() {
	postCategory()
	getCategories()
	getCategory()
	updateCategory()
	deleteCategory()
}

func postCategory() {
	router.POST(
		"/api/v1/categories",
		func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "POST : Create Category")
		})
}

func getCategories() {
	router.GET("/api/v1/categories", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "GET : All Categories")
	})
}

func getCategory() {
	router.GET("/api/v1/categories/:id", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "GET : Category %v", ctx.Param("id"))
	})
}

func updateCategory() {
	router.PATCH(
		"/api/v1/categories/:id",
		func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "PATCH : Update Category %v", ctx.Param("id"))
		})
}

func deleteCategory() {
	router.DELETE(
		"/api/v1/categories/:id",
		func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "DELETE : Delete Category %v", ctx.Param("id"))
		})
}
