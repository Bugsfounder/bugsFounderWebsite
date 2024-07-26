package main

import (
	"github.com/bugsfounder/bugsfounderweb/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.RegisterRoutes(r)

	r.Run()
}
