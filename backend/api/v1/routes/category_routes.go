package routes

import (
	"github.com/oliverschweikert/pAPI/backend/api/v1/controllers"
)

func enableCategoryRoutes() {
	categories := router.Group("/api/v1")
	{
		categories.POST("/categories", controllers.CreateCategory)
		categories.GET("/categories", controllers.GetCategories)
		categories.GET("/categories/:oid", controllers.GetCategory)
		categories.PATCH("/categories/:oid", controllers.UpdateCategory)
		categories.DELETE("/categories/:oid", controllers.DeleteCategory)
	}
}
