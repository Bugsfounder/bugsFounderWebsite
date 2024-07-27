package router

import (
	"github.com/bugsfounder/bugsfounderweb/handler"
	"github.com/bugsfounder/bugsfounderweb/logger"
	"github.com/gin-gonic/gin"
)

var LOG = logger.Logging()

func ApiRoutes(server *gin.Engine, dbHandler *handler.HandlerForDBHandlers) {
	LOG.Debug("Creating public and private routes")

	// public roup
	public_router := server.Group("/api/public")
	{
		// blog api's
		public_router.GET("/", handlePublic(dbHandler))
		public_router.POST("/blog", HandleCreateOneBlog(dbHandler))
		public_router.GET("/blogs", HandleGetAllBlogs(dbHandler))
		public_router.GET("/blogs/:blog_id", handleGetOneBlogByID(dbHandler))
		public_router.POST("/blogs/:blog_id", HandleUpdateOneBlogById(dbHandler))
		public_router.DELETE("/blogs/:blog_id", HandleDeleteOneBlogById(dbHandler))

		// tutorial api's
		public_router.POST("/tutorial", HandleCreateOneTutorail(dbHandler))
		public_router.GET("/tutorials", HandleGetAllTutorial(dbHandler))
		public_router.GET("/tutorials/:tutorial_id", handleGetOneTutorialByID(dbHandler))
		public_router.POST("/tutorials/:tutorial_id", HandleUpdateOneTutorialById(dbHandler))
		public_router.DELETE("/tutorials/:tutorial_id", HandleDeleteOneTutorialById(dbHandler))

		// user api's
		public_router.GET("/users", HandlerGetAllUsers(dbHandler))
		public_router.GET("/users/:username_or_email", HandlerGetOneUserByUsernameOrEmail(dbHandler))
		// public_router.GET("/users/:email", HandlerGetOneUserByEmail(dbHandler))
		public_router.POST("/user", HandlerCreateOneUser(dbHandler))
		public_router.POST("/user/:username_or_email", HandlerUpdateOneUserByUsernameOrEmail(dbHandler))
		public_router.DELETE("/user/:username_or_email", HandlerDeleteOneUserByUsernameOrEmail(dbHandler))
	}

	// private group
	private := server.Group("/api/private")
	{
		private.GET("/", handlePrivate)
		private.GET("/info", handlePrivateInfo)
	}
}
