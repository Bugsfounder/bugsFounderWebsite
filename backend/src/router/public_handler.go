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
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleCreateOneBlog")
	}
}
func HandleGetAllBlogs(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return func(ctx *gin.Context) {
		allBlogs := dbHandler.Client.GetAllBlogs()
		ctx.JSON(http.StatusOK, allBlogs)
	}
}

func handleGetOneBlogByID(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /handleGetOneBlogByID")
	}
}
func HandleUpdateOneBlogById(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleUpdateOneBlogById")
	}
}
func HandleDeleteOneBlogById(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleDeleteOneBlogById")
	}
}
func HandleCreateOneTutorail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleCreateOneTutorail")
	}
}
func HandleGetAllTutorial(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		allBlogs := dbHandler.Client.GetAllTutorial()
		ctx.JSON(http.StatusOK, allBlogs)
	}
}
func handleGetOneTutorialByID(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /handleGetOneTutorialByID")
	}
}
func HandleUpdateOneTutorialById(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleUpdateOneTutorialById")
	}
}
func HandleDeleteOneTutorialById(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandleDeleteOneTutorialById")
	}
}
func HandlerGetAllUsers(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		allUser := dbHandler.Client.GetAllUsers()
		ctx.JSON(http.StatusOK, allUser)
	}
}
func HandlerGetOneUserByUsernameOrEmail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerGetOneUserByUsernameOrEmail")
	}
}
func HandlerGetOneUserByUsername(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerGetOneUserByUsername")
	}
}
func HandlerGetOneUserByEmail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerGetOneUserByEmail")
	}
}
func HandlerCreateOneUser(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerCreateOneUser")
	}
}
func HandlerUpdateOneUserByUsernameOrEmail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerUpdateOneUserByUsernameOrEmail")
	}
}
func HandlerDeleteOneUserByUsernameOrEmail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	// we can access handler functions here, ex: dbHandler.DemoFuncHandler()
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "endpoint: /HandlerDeleteOneUserByUsernameOrEmail")
	}
}
