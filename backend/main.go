package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oliverschweikert/pAPI/backend/api/v1/routes"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	routes.EnableRoutes(r)
	r.Run()
}
