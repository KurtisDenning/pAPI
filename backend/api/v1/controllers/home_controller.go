package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Currently two API endpoints:\n/api/v1/categories\n/api/v1/apidata")
}
