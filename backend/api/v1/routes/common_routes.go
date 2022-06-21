package routes

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
	"github.com/oliverschweikert/pAPI/backend/api/v1/controllers"
	docs "github.com/oliverschweikert/pAPI/backend/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var router *gin.Engine

func EnableRoutes(r *gin.Engine) {
	router = r
	docs.SwaggerInfo.BasePath = "/api/v1"
	limiter := tollbooth.NewLimiter(1/10.0, nil)
	router.GET("/", tollbooth_gin.LimitHandler(limiter), controllers.Index)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	enableCategoryRoutes()
	enableApiDataRoutes()
}
