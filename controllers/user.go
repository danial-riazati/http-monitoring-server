package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(cnx *gin.Context) {

	cnx.IndentedJSON(http.StatusOK, "hi")
}
