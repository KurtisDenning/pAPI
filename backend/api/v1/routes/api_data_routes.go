package routes

import (
	"github.com/oliverschweikert/pAPI/backend/api/v1/controllers"
)

func enableApiDataRoutes() {
	apiData := router.Group("/api/v1")
	{
		apiData.POST("/apidata", controllers.CreateAPIData)
		apiData.GET("/apidata", controllers.GetAPIDatas)
		apiData.GET("/apidata/:id", controllers.GetAPIData)
		apiData.PATCH("/apidata/:id", controllers.UpdateAPIData)
		apiData.DELETE("/apidata/:id", controllers.DeleteAPIData)
	}
}
