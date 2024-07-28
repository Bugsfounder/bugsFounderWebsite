package handler

import (
	"net/http"

	"github.com/bugsfounder/bugsfounderweb/db"
	"github.com/gin-gonic/gin"
)

type HandlerForDBHandlers struct {
	Client db.Client
}

func (h_DB *HandlerForDBHandlers) DemoFuncHandler() {
	h_DB.Client.DemoFunc()
}

func (h_DB *HandlerForDBHandlers) CreateOneBlog() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetAllBlogs(ctx *gin.Context) {
	// access db functions ex: h.Client.DemoFunc() // in db
	allBlogs, err := h_DB.Client.GetAllBlogs()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, allBlogs)
}
func (h_DB *HandlerForDBHandlers) GetOneBlogByURL() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) UpdateOneBlogByURL() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) DeleteOneBlogByURL() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) CreateOneTutorial() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetAllTutorial(ctx *gin.Context) {
	allTutorial, err := h_DB.Client.GetAllTutorial()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, allTutorial)
}
func (h_DB *HandlerForDBHandlers) GetOneTutorialByURL() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) UpdateOneTutorialByURL() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) DeleteOneTutorialByURL() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetAllUsers(ctx *gin.Context) {
	allUser, err := h_DB.Client.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, allUser)
}
func (h_DB *HandlerForDBHandlers) HandlerGetOneUserByUsernameOrEmail() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetOneUserByUsername() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetOneUserByEmail() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) CreateOneUser() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) UpdateOneUserByUsernameOrEmail() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) DeleteOneUserByUsernameOrEmail() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
