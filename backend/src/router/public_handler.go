package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlePublic(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Welcome to the public endpoint")
}
func handlePublicInfo(ctx *gin.Context) {
	ctx.String(http.StatusOK, "public info")
}
