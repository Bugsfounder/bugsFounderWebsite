package router

import (
	"github.com/bugsfounder/bugsfounderweb/logger"
	"github.com/gin-gonic/gin"
)

var LOG = logger.Logging()

func ApiRoutes(server *gin.Engine) {
	LOG.Debug("Creating public and private routes")

	// public roup
	public_router := server.Group("/api/public")
	{
		public_router.GET("/", handlePublic)
		public_router.GET("/info", handlePublicInfo)
	}

	// private group
	private := server.Group("/api/private")
	{
		private.GET("/", handlePrivate)
		private.GET("/info", handlePrivateInfo)
	}
}
