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
	LOG.Debug("Main of server")
	server := gin.Default() // gin server

	// connecting to db
	client, ctx, cancel, err := db.ConnectToDBAndGetClientCtxCancelErr()
	LOG.Error("Error: %v", err)
	if err != nil {
		LOG.Error("Error: %v", err)
		panic(err)
	}
	// closing db
	defer db.Close(client, ctx, cancel)
	// ping
	if err := db.Ping(client, ctx); err != nil {
		LOG.Error("%v", err)
		LOG.Fatal(err)
	}

	c := db.Client{Client_Obj: client, Ctx: ctx}

	// c.DemoFunc()
	dbHandler := handler.DB_Handler{Client: c}

	router.ApiRoutes(server, &dbHandler)
	server.Run(":8080")
}
