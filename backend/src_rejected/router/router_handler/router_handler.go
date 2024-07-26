package router_handler

import (
	"net/http"

	"github.com/bugsfounder/bugsfounderweb/db/models"
	"github.com/bugsfounder/bugsfounderweb/handlers"
	"github.com/bugsfounder/bugsfounderweb/logger"
	"github.com/gin-gonic/gin"
)

var LOG = logger.Logging()

func Home(c *gin.Context) {
	LOG.Info("Home endpoint hit")
	c.JSON(http.StatusOK, gin.H{"api": "home"})
}

func Signup(c *gin.Context) {
	var userJSON []byte
	if err := c.BindJSON(&userJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	response, err := handlers.Signup(userJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func Login(c *gin.Context) {
	var credentialsJSON []byte
	if err := c.BindJSON(&credentialsJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	response, err := handlers.Login(credentialsJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func Logout(c *gin.Context) {
	response, err := handlers.Logout()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

// Existing functions ...

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	err := handlers.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "blog created successfully"})

}

func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	response, err := handlers.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	response, err := handlers.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updateJSON []byte
	if err := c.BindJSON(&updateJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	response, err := handlers.UpdateUser(id, updateJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	response, err := handlers.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	err := handlers.CreateBlog(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "blog created successfully"})
}

func GetAllBlogs(c *gin.Context) {
	response, err := handlers.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func GetBlogById(c *gin.Context) {
	id := c.Param("id")
	response, err := handlers.GetBlogById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	var updateJSON models.Blog
	if err := c.BindJSON(&updateJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	response, err := handlers.UpdateBlog(id, updateJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	response, err := handlers.DeleteBlog(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func HideBlog(c *gin.Context) {
	id := c.Param("id")
	response, err := handlers.HideBlog(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func UnHideBlog(c *gin.Context) {
	id := c.Param("id")
	response, err := handlers.UnHideBlog(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func CreateTutorial(c *gin.Context) {
	var tutorial models.Tutorial
	if err := c.BindJSON(&tutorial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	err := handlers.CreateTutorial(tutorial)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tutorial Created Successfully"})
}

func GetAllTutorials(c *gin.Context) {
	response, err := handlers.GetAllTutorials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func GetTutorialById(c *gin.Context) {
	id := c.Param("id")
	response, err := handlers.GetTutorialById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func UpdateTutorial(c *gin.Context) {
	id := c.Param("id")
	var update models.Tutorial
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	response, err := handlers.UpdateTutorial(id, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func DeleteTutorial(c *gin.Context) {
	id := c.Param("id")
	response, err := handlers.DeleteTutorial(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func HideTutorial(c *gin.Context) {
	id := c.Param("id")
	response, err := handlers.HideTutorial(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}

func UnHideTutorial(c *gin.Context) {
	id := c.Param("id")
	response, err := handlers.UnHideTutorial(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", response)
}
