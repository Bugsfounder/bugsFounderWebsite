package router

import (
	"github.com/bugsfounder/bugsfounderweb/router/router_handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/", router_handler.Home)

		// User routes
		api.POST("/signup", router_handler.Signup)
		api.POST("/login", router_handler.Login)
		api.GET("/logout", router_handler.Logout)
		api.POST("/create_user", router_handler.CreateUser)
		api.GET("/user/:username", router_handler.GetUserByUsername)
		api.GET("/user_email/:email", router_handler.GetUserByEmail)
		api.POST("/update_user/:id", router_handler.UpdateUser)
		api.POST("/delete_user/:id", router_handler.DeleteUser)

		// Blog routes
		api.POST("/blog", router_handler.CreateBlog)
		api.GET("/blogs", router_handler.GetAllBlogs)
		api.GET("/blogs/:blog_id", router_handler.GetBlogById)
		api.POST("/update_blog/:blog_id", router_handler.UpdateBlog)
		api.POST("/delete_blog/:blog_id", router_handler.DeleteBlog)
		api.POST("/hide_blog/:blog_id", router_handler.HideBlog)
		api.POST("/unhide_blog/:blog_id", router_handler.UnHideBlog)

		// Tutorial routes
		api.POST("/tutorial", router_handler.CreateTutorial)
		api.GET("/tutorials", router_handler.GetAllTutorials)
		api.GET("/tutorials/:tutorial_id", router_handler.GetTutorialById)
		api.POST("/update_tutorial/:tutorial_id", router_handler.UpdateTutorial)
		api.POST("/delete_tutorial/:tutorial_id", router_handler.DeleteTutorial)
		api.POST("/hide_tutorial/:tutorial_id", router_handler.HideTutorial)
		api.POST("/unhide_tutorial/:tutorial_id", router_handler.UnHideTutorial)
	}
}
