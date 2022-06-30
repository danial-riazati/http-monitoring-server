package routes

import (
	controllers "github.com/danial-riazati/http-monitoring-server/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {

	users := r.Group("/users")

	users.POST("signup", controllers.SignUp)
	users.POST("login", controllers.Login)
	// users.Use(middlewares.Auth)
	// users.GET("history", controllers.GetHistory)
	// users.GET("alerts", controllers.GetAlerts)
}
