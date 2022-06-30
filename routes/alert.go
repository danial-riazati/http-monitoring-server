package routes

import (
	"github.com/danial-riazati/http-monitoring-server/controllers"
	"github.com/danial-riazati/http-monitoring-server/middlewares"
	"github.com/gin-gonic/gin"
)

func AlertRoutes(r *gin.Engine) {

	alerts := r.Group("/alerts")
	alerts.Use(middlewares.Auth)
	alerts.GET("alerts", controllers.GetAlerts)
}
