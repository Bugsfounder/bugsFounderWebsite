package router

import (
	"net/http"

	"github.com/bugsfounder/bugsfounderweb/handler"
	"github.com/gin-gonic/gin"
)

func handlePublic(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dbHandler.DemoFuncHandler()
		ctx.String(http.StatusOK, "Welcome to the public endpoint")
	}
}
func handlePublicInfo(ctx *gin.Context) {
	ctx.String(http.StatusOK, "public info")
}
