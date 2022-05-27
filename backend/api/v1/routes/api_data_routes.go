package routes

import (
	"github.com/oliverschweikert/pAPI/backend/api/v1/controllers"
)

func enableApiDataRoutes() {
	apiData := router.Group("/api/v1")
	{
		apiData.POST("/apidata", controllers.CreateAPIData)
		apiData.GET("/apidata", controllers.GetAPIDatas)
		apiData.GET("/apidata/:oid", controllers.GetAPIData)
		apiData.PATCH("/apidata/:oid", controllers.UpdateAPIData)
		apiData.DELETE("/apidata/:oid", controllers.DeleteAPIData)
	}
}
