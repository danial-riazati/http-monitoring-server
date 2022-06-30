package routes

import (
	"github.com/danial-riazati/http-monitoring-server/controllers"
	"github.com/danial-riazati/http-monitoring-server/middlewares"
	"github.com/gin-gonic/gin"
)

func UrlRoutes(r *gin.Engine) {

	urls := r.Group("/urls")

	urls.Use(middlewares.Auth)
	urls.POST("create", controllers.CreateUrl)
	urls.POST("delete", controllers.DeleteUrl)
	urls.GET("get", controllers.GetUrl)
}
