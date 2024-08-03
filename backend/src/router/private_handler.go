package router

import (
	"net/http"

	"github.com/bugsfounder/bugsfounderweb/handler"
	"github.com/gin-gonic/gin"
)

func handlePrivate(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Welcome to the private endpoint")
}
func handlePrivateInfo(ctx *gin.Context) {
	ctx.String(http.StatusOK, "private info")
}

func CreateOneAdmin(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	return dbHandler.CreateOneAdmin

}

func DeleteAdminByUsername(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	return dbHandler.DeleteAdminByUsername

}
