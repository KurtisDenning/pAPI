package routes

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
	"github.com/oliverschweikert/pAPI/backend/api/v1/controllers"
)

var router *gin.Engine

func EnableRoutes(r *gin.Engine) {
	router = r
	limiter := tollbooth.NewLimiter(1/10.0, nil)
	router.GET("/", tollbooth_gin.LimitHandler(limiter), controllers.Index)
	enableCategoryRoutes()
	enableApiDataRoutes()
}
