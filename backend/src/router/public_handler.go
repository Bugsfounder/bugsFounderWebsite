package router

import (
	"net/http"

	"github.com/bugsfounder/bugsfounderweb/handler"
	"github.com/gin-gonic/gin"
)

func handlePublic(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return func(ctx *gin.Context) {
		dbHandler.DemoFuncHandler()
		ctx.String(http.StatusOK, "Welcome to the public endpoint")
	}
}

func HandleCreateOneBlog(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleCreateOneBlog")
	}
}
func HandleGetAllBlogs(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return func(ctx *gin.Context) {
		allBlogs, err := dbHandler.Client.GetAllBlogs()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, allBlogs)
	}
}

func handleGetOneBlogByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /handleGetOneBlogByURL")
	}
}
func HandleUpdateOneBlogByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleUpdateOneBlogByURL")
	}
}
func HandleDeleteOneBlogByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleDeleteOneBlogByURL")
	}
}
func HandleCreateOneTutorail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleCreateOneTutorail")
	}
}
func HandleGetAllTutorial(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		allBlogs, err := dbHandler.Client.GetAllTutorial()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, allBlogs)
	}
}
func handleGetOneTutorialByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /handleGetOneTutorialByURL")
	}
}
func handleGetOneTutorialBySubURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /handleGetOneTutorialBySubURL")
	}
}
func HandleUpdateOneTutorialByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleUpdateOneTutorialByURL")
	}
}
func HandleDeleteOneTutorialByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleDeleteOneTutorialByURL")
	}
}
func HandlerGetAllUsers(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		allUser, err := dbHandler.Client.GetAllUsers()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, allUser)
	}
}
func HandlerGetOneUserByUsernameOrEmail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerGetOneUserByUsernameOrEmail")
	}
}
func HandlerGetOneUserByUsername(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerGetOneUserByUsername")
	}
}
func HandlerGetOneUserByEmail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerGetOneUserByEmail")
	}
}
func HandlerCreateOneUser(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerCreateOneUser")
	}
}
func HandlerUpdateOneUserByUsernameOrEmail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerUpdateOneUserByUsernameOrEmail")
	}
}
func HandlerDeleteOneUserByUsernameOrEmail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerDeleteOneUserByUsernameOrEmail")
	}
}
