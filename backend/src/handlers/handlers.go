package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"api": "home"})
}

func About(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"api": "about"})
}

func Blogs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"api": "blogs"})
}

func Blog(c *gin.Context) {
	blogID := c.Param("blog_id")
	c.JSON(http.StatusOK, gin.H{"api": "blog", "blog_id": blogID})
}

func Tutorials(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"api": "tutorials"})
}

func Tutorial(c *gin.Context) {
	tutorialID := c.Param("tutorial_id")
	c.JSON(http.StatusOK, gin.H{"api": "tutorial", "tutorial_id": tutorialID})
}
