package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(server *gin.Engine) {
	// public roup
	public_router := server.Group("/api/public")
	{
		public_router.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Welcome to the public endpoint")
		})
		public_router.GET("/info", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "public info")
		})
	}

	// private group
	private := server.Group("/api/private")
	{
		private.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Welcome to the private endpoint")
		})
		private.GET("/info", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "private info")
		})
	}
}
