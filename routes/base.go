package routes

import "github.com/gin-gonic/gin"

func InitialRoutes(app *gin.Engine) {
	UserRoutes(app)
	UrlRoutes(app)
	AlertRoutes(app)
}
