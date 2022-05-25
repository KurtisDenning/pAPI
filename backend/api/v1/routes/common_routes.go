package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverschweikert/pAPI/backend/api/v1/controllers"
)

var router *gin.Engine

func EnableRoutes(r *gin.Engine) {
	router = r
	router.GET("/", controllers.Index)
	enableCategoryRoutes()
	enableApiDataRoutes()
}
