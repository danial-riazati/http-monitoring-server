package main

import (
	"os"

	"github.com/danial-riazati/http-monitoring-server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.Create("gin.log")
	app := gin.New()
	app.Use(gin.LoggerWithWriter(f))
	routes.UserRoutes(app)
	app.Run("localhost:1234")

}
