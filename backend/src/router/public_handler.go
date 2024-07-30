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
	return dbHandler.CreateOneBlog
}

func HandleGetAllBlogs(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.GetAllBlogs
}

func handleGetOneBlogByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.GetOneBlogByURL
}

func HandleUpdateOneBlogByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.UpdateOneBlogByURL
}

func HandleDeleteOneBlogByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.DeleteOneBlogByURL
}

func HandleCreateOneTutorial(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return dbHandler.CreateOneTutorial
}

func HandleCreateSubTutorial(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")

	return dbHandler.CreateSubTutorial
}

func HandleGetAllTutorial(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.GetAllTutorial
}

func HandleGetOneTutorialByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.GetOneTutorialByURL
}

func HandleGetSubTutorialByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.GetSubTutorialByURL
}

func HandleUpdateOneTutorialByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.UpdateOneTutorialByURL
}

func HandleUpdateSubTutorialByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.UpdateSubTutorialByURL
}

func HandleDeleteOneTutorialByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.DeleteOneTutorialByURL
}

func HandleDeleteSubTutorialByURL(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.DeleteSubTutorialByURL
}

func HandlerGetAllUsers(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.GetAllUsers
}

func HandlerGetOneUserByUsernameOrEmail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.GetOneUserByUsernameOrEmail
}

func HandlerCreateOneUser(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.CreateOneUser
}

func HandlerUpdateOneUserByUsernameOrEmail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.UpdateOneUserByUsernameOrEmail
}

func HandlerDeleteOneUserByUsernameOrEmail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.DeleteOneUserByUsernameOrEmail
}
