package main

import (
	"os"

	"github.com/danial-riazati/http-monitoring-server/configs"
	"github.com/danial-riazati/http-monitoring-server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.Create("server.log")
	app := gin.New()
	cfg := configs.New()
	app.Use(gin.LoggerWithWriter(f))
	routes.UserRoutes(app)
	app.Run("localhost" + cfg.Listen)

}
