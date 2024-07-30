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
func (h_DB *HandlerForDBHandlers) UpdateOneBlogByURL(ctx *gin.Context) {
	LOG.Debug("")
	blogURL := ctx.Param("blog_url")

	LOG.Debug("blog_url: %v", blogURL)

	// define a variable to hold the updated blog data
	var updatedBlog models.Blog

	// Bind the JSON data from the request body to the updatedBlog variable
	if err := ctx.BindJSON(&updatedBlog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request data"})
		return
	}

	// call the db handler method to update the blog
	result, err := h_DB.Client.UpdateOneBlogByURL(blogURL, &updatedBlog)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Return the result of the update operation
	ctx.JSON(http.StatusOK, result)

}
func (h_DB *HandlerForDBHandlers) DeleteOneBlogByURL(ctx *gin.Context) {
	LOG.Debug("")
	blogURL := ctx.Param("blog_url")
	LOG.Debug("blog_url: %v", blogURL)

	// call the db handler method to delete the blog
	result, err := h_DB.Client.DeleteOneBlogByURL(blogURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "blog not found"})
		return
	}

	// return success message
	ctx.JSON(http.StatusOK, gin.H{"message": "blog deleted successfully"})
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

func (h_DB *HandlerForDBHandlers) CreateSubTutorial(ctx *gin.Context) {
	LOG.Debug("")
	tutorialURL := ctx.Param("tutorial_url")

	var subTutorial models.Sub_Tutorial
	if err := ctx.BindJSON(&subTutorial); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	result, err := h_DB.Client.CreateSubTutorial(tutorialURL, &subTutorial)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
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
func (h_DB *HandlerForDBHandlers) UpdateOneTutorialByURL(ctx *gin.Context) {
	LOG.Debug("")
	tutorialURL := ctx.Param("tutorial_url")

	// define a variable to hold the updated blog data
	var updatedTutorial models.Tutorial

	// Bind the JSON data from the request body to the updatedBlog variable
	if err := ctx.BindJSON(&updatedTutorial); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request data"})
		return
	}

	// call the db handler method to update the blog
	result, err := h_DB.Client.UpdateOneTutorialByURL(tutorialURL, &updatedTutorial)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Return the result of the update operation
	ctx.JSON(http.StatusOK, result)

}
func (h_DB *HandlerForDBHandlers) UpdateSubTutorialByURL(ctx *gin.Context) {
	LOG.Debug("")
	tutorialURL := ctx.Param("tutorial_url")
	subTutorialURL := ctx.Param("sub_tutorial_url")

	var updatedSubTutorial models.Sub_Tutorial
	if err := ctx.ShouldBindJSON(&updatedSubTutorial); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	result, err := h_DB.Client.UpdateSubTutorialByURL(tutorialURL, subTutorialURL, &updatedSubTutorial)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Sub-Tutorial not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Sub-Tutorial updated successfully"})
}
func (h_DB *HandlerForDBHandlers) DeleteOneTutorialByURL(ctx *gin.Context) {
	LOG.Debug("")
	tutorialURL := ctx.Param("tutorial_url")

	result, err := h_DB.Client.DeleteOneTutorialByURL(tutorialURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Tutorial not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tutorial deleted successfully"})
}

func (h_DB *HandlerForDBHandlers) DeleteSubTutorialByURL(ctx *gin.Context) {
	LOG.Debug("")

	tutorialURL := ctx.Param("tutorial_url")
	subTutorialURL := ctx.Param("sub_tutorial_url")

	result, err := h_DB.Client.DeleteSubTutorialByURL(tutorialURL, subTutorialURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Sub-Tutorial not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Sub-Tutorial deleted successfully"})
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
func (h_DB *HandlerForDBHandlers) GetOneUserByUsernameOrEmail(ctx *gin.Context) {
	LOG.Debug("")
	usernameOrEmail := ctx.Param("username_or_email")

	user, err := h_DB.Client.GetOneUserByUsernameOrEmail(usernameOrEmail)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)

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
func (h_DB *HandlerForDBHandlers) UpdateOneUserByUsernameOrEmail(ctx *gin.Context) {
	LOG.Debug("")
	// Get the email or username from the URL parameters
	username_or_email := ctx.Param("username_or_email")
	LOG.Debug("username_or_email: %v", username_or_email)

	// Bind the JSON body to the User model
	var updatedUser models.User
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	// Call the DB handler method to update the user
	result, err := h_DB.Client.UpdateOneUserByUsernameOrEmail(username_or_email, &updatedUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// Return success message
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
func (h_DB *HandlerForDBHandlers) DeleteOneUserByUsernameOrEmail(ctx *gin.Context) {
	LOG.Debug("")
	usernameOrEmail := ctx.Param("username_or_email")

	result, err := h_DB.Client.DeleteOneUserByUsernameOrEmail(usernameOrEmail)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
