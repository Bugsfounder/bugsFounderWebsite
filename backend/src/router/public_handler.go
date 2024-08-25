package router

import (
	"net/http"

	"github.com/bugsfounder/bugsfounderweb/handler"
	"github.com/gin-gonic/gin"
)

func handlePublic(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return func(ctx *gin.Context) {
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

func Login(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.Login
}

func Signup(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.Signup
}

func HandleIsUsernameOrEmailPresent(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.IsUsernameOrEmailPresent
}
func Logout(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.Logout
}

func HandlerUpdateOneUserByUsernameOrEmail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.UpdateOneUserByUsernameOrEmail
}

func HandlerDeleteOneUserByUsernameOrEmail(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.DeleteOneUserByUsernameOrEmail
}

func HandlerSearch(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.Search

}
func HandleSearchBlogTitle(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.SearchBlogURL

}
func HandleSearchTutorialTitle(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.SearchTutorialURL
}
func HandlerSearchSubTutorialTitle(dbHandler *handler.HandlerForDBHandlers) gin.HandlerFunc {
	LOG.Debug("")
	return dbHandler.SearchSubTutorialURL

}
