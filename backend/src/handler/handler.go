package handler

import (
	"net/http"

	"github.com/bugsfounder/bugsfounderweb/db"
	"github.com/bugsfounder/bugsfounderweb/logger"
	"github.com/bugsfounder/bugsfounderweb/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var LOG = logger.Logging()

type HandlerForDBHandlers struct {
	Client db.Client
}

func (h_DB *HandlerForDBHandlers) DemoFuncHandler() {
	LOG.Debug("")
	h_DB.Client.DemoFunc()
}

func (h_DB *HandlerForDBHandlers) CreateOneBlog(ctx *gin.Context) {
	LOG.Debug("")

	var blog models.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	result, err := h_DB.Client.CreateOneBlog(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "no document found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"inserted_id": result.InsertedID})

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
func (h_DB *HandlerForDBHandlers) CreateOneTutorial(ctx *gin.Context) {
	LOG.Debug("")

	var tutorial models.Tutorial
	if err := ctx.ShouldBindJSON(&tutorial); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	result, err := h_DB.Client.CreateOneTutorial(&tutorial)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "no document found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"inserted_id": result.InsertedID})
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
func (h_DB *HandlerForDBHandlers) GetSubTutorialByURL(ctx *gin.Context) {
	LOG.Debug("")

	tutorial_url := ctx.Param("tutorial_url")
	sub_tutorial_url := ctx.Param("sub_tutorial_url")

	subTutorial, err := h_DB.Client.GetSubTutorialByURL(tutorial_url, sub_tutorial_url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, subTutorial)
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
func (h_DB *HandlerForDBHandlers) CreateOneUser(ctx *gin.Context) {
	LOG.Debug("")
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	result, err := h_DB.Client.CreateOneUser(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "no document found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"inserted_id": result.InsertedID})
}
func (h_DB *HandlerForDBHandlers) UpdateOneUserByUsernameOrEmail() {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) DeleteOneUserByUsernameOrEmail() {
	LOG.Debug("")
	// access db functions ex: h.Client.DemoFunc() // in db
}
