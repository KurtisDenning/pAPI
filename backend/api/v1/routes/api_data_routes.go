package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func enableApiDataRoutes() {
	postAPIData()
	getAPIDatas()
	getAPIData()
	updateAPIData()
	deleteAPIData()
}

func postAPIData() {
	router.POST(
		"/api/v1/apidata",
		func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "POST : Create API Data Value")
		})
}

func getAPIDatas() {
	router.GET("/api/v1/apidata", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "GET : All API Data Values")
	})
}

func getAPIData() {
	router.GET("/api/v1/apidata/:id", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "GET : API Data Value %v", ctx.Param("id"))
	})
}

func updateAPIData() {
	router.PATCH(
		"/api/v1/apidata/:id",
		func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "PATCH : Update API Data Value %v", ctx.Param("id"))
		})
}

func deleteAPIData() {
	router.DELETE(
		"/api/v1/apidata/:id",
		func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "DELETE : Delete API Data Value %v", ctx.Param("id"))
		})
}
