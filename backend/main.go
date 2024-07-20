package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const PORT = 8080

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost", "0.0.0.0", "127.0.0.1"})
	router.Use(cors.Default())

	// define your routes here
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to BugsFounder.com"})
	})

	// start the server in a goroutine
	go func() {
		if err := router.Run(fmt.Sprintf(":%d", PORT)); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
	fmt.Println("Server is Successfully Listening at", PORT)

	// Block the main goroutine to keep the application running
	select {}
}
