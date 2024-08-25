package handler

import (
	"net/http"
	"strconv"

	"github.com/bugsfounder/bugsfounderweb/db"
	"github.com/bugsfounder/bugsfounderweb/logger"
	"github.com/bugsfounder/bugsfounderweb/models"
	"github.com/bugsfounder/bugsfounderweb/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var LOG = logger.Logging()

type HandlerForDBHandlers struct {
	Client db.Client
}

// CreateOneBlog
func (h_DB *HandlerForDBHandlers) CreateOneBlog(ctx *gin.Context) {
	LOG.Debug("")

	var blog models.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "err": err.Error()})
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

// GetAllBlogs
func (h_DB *HandlerForDBHandlers) GetAllBlogs(ctx *gin.Context) {
	LOG.Debug("")

	// Parse pagination parameters
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// Fetch blogs with offset and limit
	allBlogs, err := h_DB.Client.GetAllBlogs(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"blogs": allBlogs})
}

// GetOneBlogByURL
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

// UpdateOneBlogByURL
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
	result, updated_blog, err := h_DB.Client.UpdateOneBlogByURL(blogURL, &updatedBlog)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Return the result of the update operation
	ctx.JSON(http.StatusOK, gin.H{"result": result, "updated_fields": updated_blog})

}

// DeleteOneBlogByURL
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

// CreateOneTutorial
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

// CreateSubTutorial
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

// GetAllTutorial
func (h_DB *HandlerForDBHandlers) GetAllTutorial(ctx *gin.Context) {
	LOG.Debug("")

	// Parse Pagination parameters
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset := (page + 1) + limit

	// fetch tutorials with offset and limit
	allTutorial, err := h_DB.Client.GetAllTutorial(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tutorials": allTutorial})
}

// GetOneTutorialByURL
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

// GetSubTutorialByURL
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

// UpdateOneTutorialByURL
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
	result, updated_fields, err := h_DB.Client.UpdateOneTutorialByURL(tutorialURL, &updatedTutorial)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Return the result of the update operation
	ctx.JSON(http.StatusOK, gin.H{"result": result, "updated_fields": updated_fields})

}

// UpdateSubTutorialByURL
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

// DeleteOneTutorialByURL
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

// DeleteSubTutorialByURL
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

// GetAllUsers
func (h_DB *HandlerForDBHandlers) GetAllUsers(ctx *gin.Context) {
	LOG.Debug("")
	allUser, err := h_DB.Client.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, allUser)
}

// Login
func (h_DB *HandlerForDBHandlers) Login(ctx *gin.Context) {
	LOG.Debug("")

	var loginRequest struct {
		UsernameOrEmail string `json:"username_or_email" binding:"required"`
		Password        string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := h_DB.Client.GetOneUserByUsernameOrEmail(loginRequest.UsernameOrEmail, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// generating jwt token
	token, err := utils.GenerateJWT(user.Username, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{"token": token, "user": user})
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// Signup
func (h_DB *HandlerForDBHandlers) Signup(ctx *gin.Context) {
	LOG.Debug("")

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	_, err := h_DB.Client.CreateOneUser(&user)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	// generating jwt token
	token, err := utils.GenerateJWT(user.Username, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (h_DB *HandlerForDBHandlers) IsUsernameOrEmailPresent(ctx *gin.Context) {
	LOG.Debug("")

	var usernameOrEmail map[string]string
	if err := ctx.ShouldBindJSON(&usernameOrEmail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	isPresent, err := h_DB.Client.IsUsernameOrEmailPresent(usernameOrEmail["usernameOrEmail"])
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"isPresent": isPresent})
}

func (h_DB *HandlerForDBHandlers) Logout(ctx *gin.Context) {
	LOG.Debug("Received a logout request")

	// Clear JWT token by setting an empty value and expired time in the cookie
	ctx.SetCookie("token", "", -1, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

// UpdateOneUserByUsernameOrEmail
func (h_DB *HandlerForDBHandlers) UpdateOneUserByUsernameOrEmail(ctx *gin.Context) {
	LOG.Debug("")
	// Get the email or username from the URL parameters
	username_or_email := ctx.Param("username_or_email")
	LOG.Debug("username_or_email: %v", username_or_email)

	// Bind the JSON body to the User model
	var updatedUser models.UpdateUser
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

// DeleteOneUserByUsernameOrEmail
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

func (h_DB *HandlerForDBHandlers) TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		isBlacklisted, err := h_DB.Client.IsTokenBlacklisted(tokenString)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify token"})
			ctx.Abort()
			return
		}

		if isBlacklisted {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is blacklisted"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

// search
func (h_DB *HandlerForDBHandlers) SearchBlogURL(ctx *gin.Context) {
	query := ctx.Param("query")

	resultCount, err := h_DB.Client.SearchBlogURL(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": resultCount})

}
func (h_DB *HandlerForDBHandlers) SearchTutorialURL(ctx *gin.Context) {
	query := ctx.Param("query")

	resultCount, err := h_DB.Client.SearchTutorialURL(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": resultCount})
}
func (h_DB *HandlerForDBHandlers) SearchSubTutorialURL(ctx *gin.Context) {
	title := ctx.Param("title")
	sub_title := ctx.Param("sub_title")

	resultCount, err := h_DB.Client.SearchSubTutorialURL(title, sub_title)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": resultCount})
}

func (h_DB *HandlerForDBHandlers) Search(ctx *gin.Context) {
	query := ctx.Param("query")

	result, err := h_DB.Client.Search(query)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": result})
}

// admin

func (h_DB *HandlerForDBHandlers) CreateOneAdmin(ctx *gin.Context) {
	LOG.Debug("")

	var adminModel models.Admin
	if err := ctx.ShouldBindJSON(&adminModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	// Create the admin
	result, err := h_DB.Client.CreateOneAdmin(&adminModel)
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

func (h_DB *HandlerForDBHandlers) DeleteAdminByUsername(ctx *gin.Context) {
	LOG.Debug("")

	username := ctx.Param("username")
	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "username is required"})
		return
	}

	// Delete the admin
	result, err := h_DB.Client.DeleteAdminByUsername(username)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "no document found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"deleted_count": result.DeletedCount})
}
