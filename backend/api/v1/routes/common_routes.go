package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func EnableAllRoutes(r *gin.Engine) {
	router = r
	enableCategoryRoutes()
	enableApiDataRoutes()
	enableHomeRoute()
}

func enableHomeRoute() {
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Currently two API endpoints:\n/api/v1/categories\n/api/v1/apidata")
	})
}
