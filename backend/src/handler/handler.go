package handler

import (
	"net/http"

	"github.com/bugsfounder/bugsfounderweb/db"
	"github.com/bugsfounder/bugsfounderweb/logger"
	"github.com/gin-gonic/gin"
)

var LOG = logger.Logging()

type HandlerForDBHandlers struct {
	Client db.Client
}

func (h_DB *HandlerForDBHandlers) DemoFuncHandler() {
	LOG.Debug("")
	h_DB.Client.DemoFunc()
}

func (h_DB *HandlerForDBHandlers) CreateOneBlog() {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetAllBlogs(ctx *gin.Context) {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
	allBlogs, err := h_DB.Client.GetAllBlogs()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, allBlogs)
}
func (h_DB *HandlerForDBHandlers) GetOneBlogByURL(ctx *gin.Context) {
	LOG.Debug("")

	blogURL := ctx.Param("blog_url")

	blog, err := h_DB.Client.GetOneBlogByURL(blogURL)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, blog)

}
func (h_DB *HandlerForDBHandlers) UpdateOneBlogByURL() {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) DeleteOneBlogByURL() {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) CreateOneTutorial() {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetAllTutorial(ctx *gin.Context) {
	LOG.Debug("")
	allTutorial, err := h_DB.Client.GetAllTutorial()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, allTutorial)
}
func (h_DB *HandlerForDBHandlers) GetOneTutorialByURL(ctx *gin.Context) {
	LOG.Debug("")

	tutorial_url := ctx.Param("tutorial_url")

	tutorial, err := h_DB.Client.GetOneTutorialByURL(tutorial_url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, tutorial)
}
func (h_DB *HandlerForDBHandlers) UpdateOneTutorialByURL() {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) DeleteOneTutorialByURL() {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetAllUsers(ctx *gin.Context) {
	LOG.Debug("")
	allUser, err := h_DB.Client.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, allUser)
}
func (h_DB *HandlerForDBHandlers) HandlerGetOneUserByUsernameOrEmail() {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetOneUserByUsername() {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetOneUserByEmail() {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) CreateOneUser() {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) UpdateOneUserByUsernameOrEmail() {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) DeleteOneUserByUsernameOrEmail() {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
}
