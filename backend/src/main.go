package main

import (
	"github.com/bugsfounder/bugsfounderweb/db"
	"github.com/bugsfounder/bugsfounderweb/handler"
	"github.com/bugsfounder/bugsfounderweb/logger"
	"github.com/bugsfounder/bugsfounderweb/router"
	"github.com/gin-gonic/gin"
)

var LOG = logger.Logging()

func main() {
	LOG.Debug("Starting server")
	server := gin.Default() // gin server

	// connecting to db
	client, ctx, cancel, err := db.ConnectToDBAndGetClientCtxCancelErr()
	if err != nil {
		LOG.Error("Failed to connect to the database: %v", err)
		LOG.Fatal(err)
	}
	// closing db
	defer db.Close(client, ctx, cancel)

	// ping
	if err := db.Ping(client, ctx); err != nil {
		LOG.Error("Failed to ping the database: %v", err)
		LOG.Fatal(err)
	}

	// creating handler which is present in handler to access db methods
	dbHandler := handler.HandlerForDBHandlers{
		Client: db.Client{
			Client_Obj: client,
			Ctx:        ctx,
		},
	}

	// passing handler to api routes
	router.ApiRoutes(server, &dbHandler)
	server.Run(":8080")
}
