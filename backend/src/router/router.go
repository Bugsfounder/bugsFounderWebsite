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
		public_router.GET("/", handlePublic(dbHandler))

		// blog api's
		public_router.POST("/blogs", HandleCreateOneBlog(dbHandler))                  // done
		public_router.GET("/blogs", HandleGetAllBlogs(dbHandler))                     // done
		public_router.GET("/blogs/:blog_url", handleGetOneBlogByURL(dbHandler))       // done
		public_router.PUT("/blogs/:blog_url", HandleUpdateOneBlogByURL(dbHandler))    // done
		public_router.DELETE("/blogs/:blog_url", HandleDeleteOneBlogByURL(dbHandler)) // done

		// tutorial api's
		public_router.POST("/tutorials", HandleCreateOneTutorial(dbHandler))                                        // done
		public_router.POST("/tutorials/:tutorial_url", HandleCreateSubTutorial(dbHandler))                          // done
		public_router.GET("/tutorials", HandleGetAllTutorial(dbHandler))                                            // done
		public_router.GET("/tutorials/:tutorial_url", HandleGetOneTutorialByURL(dbHandler))                         // done
		public_router.GET("/tutorials/:tutorial_url/:sub_tutorial_url", HandleGetSubTutorialByURL(dbHandler))       // done
		public_router.PUT("/tutorials/:tutorial_url", HandleUpdateOneTutorialByURL(dbHandler))                      // done
		public_router.PUT("/tutorials/:tutorial_url/:sub_tutorial_url", HandleUpdateSubTutorialByURL(dbHandler))    // done
		public_router.DELETE("/tutorials/:tutorial_url", HandleDeleteOneTutorialByURL(dbHandler))                   // done
		public_router.DELETE("/tutorials/:tutorial_url/:sub_tutorial_url", HandleDeleteSubTutorialByURL(dbHandler)) // done

		// user api's
		// public_router.GET("/users/:email", HandlerGetOneUserByEmail(dbHandler))
		public_router.POST("/users", HandlerCreateOneUser(dbHandler))                                       // done
		public_router.GET("/users", HandlerGetAllUsers(dbHandler))                                          // done
		public_router.PUT("/users/:username_or_email", HandlerUpdateOneUserByUsernameOrEmail(dbHandler))    // done
		public_router.GET("/users/:username_or_email", HandlerGetOneUserByUsernameOrEmail(dbHandler))       // done
		public_router.DELETE("/users/:username_or_email", HandlerDeleteOneUserByUsernameOrEmail(dbHandler)) // done
	}

	// private group
	private := server.Group("/api/private")
	{
		private.GET("/", handlePrivate)
		private.GET("/info", handlePrivateInfo)
	}
}
