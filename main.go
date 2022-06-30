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
	app.Use(gin.LoggerWithWriter(f))
	routes.InitialRoutes(app)
	app.Run("localhost" + configs.Cfg.Listen)

}
