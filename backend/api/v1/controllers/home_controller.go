package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
}
