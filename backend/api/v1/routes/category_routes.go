package routes

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/oliverschweikert/pAPI/backend/api/v1/controllers"
)

func enableCategoryRoutes() {
	limiter := tollbooth.NewLimiter(1/10.0, nil)
	categories := router.Group("/api/v1")
	{
		categories.POST("/categories", tollbooth_gin.LimitHandler(limiter), controllers.CreateCategory)
		categories.GET("/categories", tollbooth_gin.LimitHandler(limiter), controllers.GetCategories)
		categories.GET("/categories/:oid", tollbooth_gin.LimitHandler(limiter), controllers.GetCategory)
		categories.PATCH("/categories/:oid", tollbooth_gin.LimitHandler(limiter), controllers.UpdateCategory)
		categories.DELETE("/categories/:oid", tollbooth_gin.LimitHandler(limiter), controllers.DeleteCategory)
	}
}
