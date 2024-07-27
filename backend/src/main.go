package main

import (
	"github.com/bugsfounder/bugsfounderweb/logger"
	"github.com/bugsfounder/bugsfounderweb/router"
	"github.com/gin-gonic/gin"
)

var LOG = logger.Logging()

func main() {

	LOG.Debug("Main of server")
	server := gin.Default()

	router.ApiRoutes(server)
	server.Run(":8080")
}
