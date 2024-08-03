package main

import (
	"fmt"
	"os"

	"github.com/bugsfounder/bugsfounderweb/db"
	"github.com/bugsfounder/bugsfounderweb/handler"
	"github.com/bugsfounder/bugsfounderweb/logger"
	"github.com/bugsfounder/bugsfounderweb/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var LOG = logger.Logging()

func CustomPanic(message string) {
	const (
		RedBG = "\033[41m"
		White = "\033[97m"
		Reset = "\033[0m"
	)
	fmt.Printf("%s%sPanic: %s%s\n", RedBG, White, message, Reset)
	os.Exit(1)
}

func main() {

	LOG.Debug("Starting server")
	server := gin.Default() // gin server
	server.SetTrustedProxies([]string{"localhost", "0.0.0.0", "127.0.0.1"})
	// CORS configuration
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Change this to your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Set trusted proxies for production (replace with actual proxy IPs)
	// r.SetTrustedProxies([]string{"192.168.1.1", "192.168.1.2"})

	// connecting to db
	client, ctx, cancel, err := db.ConnectToDBAndGetClientCtxCancelErr()
	if err != nil {
		LOG.Error("Failed to connect to the database: %v", err)
		LOG.Fatal("%v", err)
	}
	// closing db
	defer db.Close(client, ctx, cancel)

	// ping
	if err := db.Ping(client, ctx); err != nil {
		LOG.Error("Failed to ping the database: %v", err)
		LOG.Fatal("%v", err)
	}

	// creating handler which is present in handler to access db methods
	dbHandler := handler.HandlerForDBHandlers{
		Client: db.Client{
			Client_Obj: client,
		},
	}

	// passing handler to api routes
	router.ApiRoutes(server, &dbHandler)
	server.Run(":8080")
}
