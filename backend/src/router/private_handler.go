package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlePrivate(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Welcome to the private endpoint")
}
func handlePrivateInfo(ctx *gin.Context) {
	ctx.String(http.StatusOK, "private info")
}
