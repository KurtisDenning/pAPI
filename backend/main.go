package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oliverschweikert/pAPI/backend/api/v1/routes"
)

// @title           pAPI
// @version         1.0
// @description     A public API that brings together the data from many other API's and makes them publicly accessible.
// @host      		papi-project.herokuapp.com
// @BasePath  		/api/v1
func main() {
	r := gin.Default()
	r.Use(cors.Default())
	routes.EnableRoutes(r)
	r.Run()
}
