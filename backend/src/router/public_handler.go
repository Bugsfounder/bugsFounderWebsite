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

func HandleCreateOneBlog(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleCreateOneBlog")
	}
}
func HandleGetAllBlogs(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleGetAllBlogs")
	}
}
func handleGetOneBlogByID(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /handleGetOneBlogByID")
	}
}
func HandleUpdateOneBlogById(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleUpdateOneBlogById")
	}
}
func HandleDeleteOneBlogById(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleDeleteOneBlogById")
	}
}
func HandleCreateOneTutorail(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleCreateOneTutorail")
	}
}
func HandleGetAllTutorial(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleGetAllTutorial")
	}
}
func handleGetOneTutorialByID(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /handleGetOneTutorialByID")
	}
}
func HandleUpdateOneTutorialById(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleUpdateOneTutorialById")
	}
}
func HandleDeleteOneTutorialById(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleDeleteOneTutorialById")
	}
}
func HandlerGetAllUsers(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerGetAllUsers")
	}
}
func HandlerGetOneUserByUsername(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerGetOneUserByUsername")
	}
}
func HandlerGetOneUserByEmail(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerGetOneUserByEmail")
	}
}
func HandlerCreateOneUser(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerCreateOneUser")
	}
}
func HandlerUpdateOneUserByUsernameOrEmail(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerUpdateOneUserByUsernameOrEmail")
	}
}
func HandlerDeleteOneUserByUsernameOrEmail(dbHandler *handler.DB_Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerDeleteOneUserByUsernameOrEmail")
	}
}
