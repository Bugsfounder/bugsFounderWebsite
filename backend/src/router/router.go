package router

import (
	"github.com/bugsfounder/bugsfounderweb/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/", handlers.Home)
		api.GET("/blogs", handlers.Blogs)
		api.GET("/blogs/:blog_id", handlers.Blog)
		api.GET("/tutorials", handlers.Tutorials)
		api.GET("/tutorials/:tutorial_id", handlers.Tutorial)
		api.POST("/login", handlers.Login)
		api.POST("/signup", handlers.Signup)
		api.GET("/logout", handlers.Signup)
	}
}
