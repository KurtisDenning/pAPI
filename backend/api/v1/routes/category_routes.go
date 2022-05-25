package routes

import (
	"github.com/oliverschweikert/pAPI/backend/api/v1/controllers"
)

func enableCategoryRoutes() {
	categories := router.Group("/api/v1")
	{
		categories.POST("/categories", controllers.CreateCategory)
		categories.GET("/categories", controllers.GetCategories)
		categories.GET("/categories/:id", controllers.GetCategory)
		categories.PATCH("/categories/:id", controllers.UpdateCategory)
		categories.DELETE("/categories/:id", controllers.DeleteCategory)
	}
}
