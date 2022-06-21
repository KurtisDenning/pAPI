package routes

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/oliverschweikert/pAPI/backend/api/v1/controllers"
)

func enableApiDataRoutes() {
	limiter := tollbooth.NewLimiter(1/10.0, nil)
	apiData := router.Group("/api/v1")
	{
		apiData.GET("/apidata", tollbooth_gin.LimitHandler(limiter), controllers.GetAPIDatas)
		apiData.GET("/apidata/:oid", tollbooth_gin.LimitHandler(limiter), controllers.GetAPIData)
		apiData.GET("/apidata/:oid/:request", tollbooth_gin.LimitHandler(limiter), controllers.GetAPIDataRequest)
	}
}
