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
		public_router.GET("/blog", HandleGetAllBlogs(dbHandler))               // done - tested
		public_router.GET("/blog/:blog_url", handleGetOneBlogByURL(dbHandler)) // done - tested

		// tutorial api's
		public_router.GET("/tutorial", HandleGetAllTutorial(dbHandler))                                      // done - tested
		public_router.GET("/tutorial/:tutorial_url", HandleGetOneTutorialByURL(dbHandler))                   // done - tested
		public_router.GET("/tutorial/:tutorial_url/:sub_tutorial_url", HandleGetSubTutorialByURL(dbHandler)) // done - tested

		// user api's
		// TODO: update jwt logic, make it more secure same like shown in the udemy course
		// public_router.GET("/users/:email", HandlerGetOneUserByEmail(dbHandler))
		public_router.POST("/user/signup", Signup(dbHandler))                                                     // done - tested
		public_router.POST("/user/login", Login(dbHandler))                                                       // done - tested
		public_router.POST("/user/logout", Logout(dbHandler))                                                     // done - tested
		public_router.POST("/user/email_or_username_unique", HandleIsUsernameOrEmailPresent(dbHandler))           // done - tested
		public_router.PUT("/user/update/:username_or_email", HandlerUpdateOneUserByUsernameOrEmail(dbHandler))    // done - tested
		public_router.DELETE("/user/delete/:username_or_email", HandlerDeleteOneUserByUsernameOrEmail(dbHandler)) // done - tested

		// search api's
		public_router.GET("/search/blog_title/:query", HandleSearchBlogTitle(dbHandler))         // done - tested
		public_router.GET("/search/tutorial_title/:query", HandleSearchTutorialTitle(dbHandler)) // done - tested
		public_router.GET("/search/:title/:sub_title", HandlerSearchSubTutorialTitle(dbHandler)) // done - tested
		public_router.GET("/search/all/:query", HandlerSearch(dbHandler))                        // done - tested
	}

	// private group
	private_router := server.Group("/api/private/admin")
	{
		private_router.POST("/", CreateOneAdmin(dbHandler))                                                         // done - tested
		private_router.DELETE("/:username", DeleteAdminByUsername(dbHandler))                                       // done - tested
		private_router.POST("/blog", HandleCreateOneBlog(dbHandler))                                                // done - tested
		private_router.PUT("/blog/:blog_url", HandleUpdateOneBlogByURL(dbHandler))                                  // done - tested
		private_router.DELETE("/blog/:blog_url", HandleDeleteOneBlogByURL(dbHandler))                               // done - tested
		private_router.POST("/tutorial", HandleCreateOneTutorial(dbHandler))                                        // done - tested
		private_router.POST("/tutorial/:tutorial_url", HandleCreateSubTutorial(dbHandler))                          // done - tested
		private_router.PUT("/tutorial/:tutorial_url", HandleUpdateOneTutorialByURL(dbHandler))                      // done - tested
		private_router.PUT("/tutorial/:tutorial_url/:sub_tutorial_url", HandleUpdateSubTutorialByURL(dbHandler))    // done - tested
		private_router.DELETE("/tutorial/:tutorial_url", HandleDeleteOneTutorialByURL(dbHandler))                   // done - tested
		private_router.DELETE("/tutorial/:tutorial_url/:sub_tutorial_url", HandleDeleteSubTutorialByURL(dbHandler)) // done - tested
		private_router.GET("/user/getAllUser", HandlerGetAllUsers(dbHandler))                                       // done - tested
	}
}
