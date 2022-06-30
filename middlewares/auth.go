package middlewares

import (
	"net/http"

	"github.com/danial-riazati/http-monitoring-server/utils"
	"github.com/gin-gonic/gin"
)

// Authz validates token and authorizes users
func Auth(cnx *gin.Context) {
	clientToken := cnx.Request.Header.Get("bearer_token")
	if clientToken == "" {
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization header provided"})
		cnx.Abort()
		return
	}

	claims, err := utils.ValidateToken(clientToken)
	if err != "" {
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		cnx.Abort()
		return
	}

	cnx.Set("name", claims.Name)
	cnx.Set("user_id", claims.Uid)
	cnx.Next()

}
