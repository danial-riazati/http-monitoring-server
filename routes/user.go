package routes

import (
	controllers "github.com/danial-riazati/http-monitoring-server/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/login", controllers.Login)
}
