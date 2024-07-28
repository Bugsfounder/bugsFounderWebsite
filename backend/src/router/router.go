package router

import (
	"github.com/bugsfounder/bugsfounderweb/handler"
	"github.com/bugsfounder/bugsfounderweb/logger"
	"github.com/gin-gonic/gin"
)

var LOG = logger.Logging()

func ApiRoutes(server *gin.Engine, dbHandler *handler.HandlerForDBHandlers) {
	LOG.Debug("Creating public and private routes")

	// public group
	public_router := server.Group("/api/public")
	{
		// blog api's
		public_router.GET("/", handlePublic(dbHandler))
		public_router.POST("/blog", HandleCreateOneBlog(dbHandler))
		public_router.GET("/blogs", HandleGetAllBlogs(dbHandler))               // done
		public_router.GET("/blogs/:blog_url", handleGetOneBlogByURL(dbHandler)) // done
		public_router.POST("/blogs/:blog_url", HandleUpdateOneBlogByURL(dbHandler))
		public_router.DELETE("/blogs/:blog_url", HandleDeleteOneBlogByURL(dbHandler))

		// tutorial api's
		public_router.POST("/tutorial", HandleCreateOneTutorial(dbHandler))
		public_router.GET("/tutorials", HandleGetAllTutorial(dbHandler))                                      // done
		public_router.GET("/tutorials/:tutorial_url", HandleGetOneTutorialByURL(dbHandler))                   // done
		public_router.GET("/tutorials/:tutorial_url/:sub_tutorial_url", HandleGetSubTutorialByURL(dbHandler)) // done
		public_router.POST("/tutorials/:tutorial_url", HandleUpdateOneTutorialByURL(dbHandler))
		public_router.DELETE("/tutorials/:tutorial_url", HandleDeleteOneTutorialByURL(dbHandler))

		// user api's
		public_router.GET("/users", HandlerGetAllUsers(dbHandler)) // done
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
