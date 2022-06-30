package controllers

import (
	"context"
	"net/http"

	"github.com/danial-riazati/http-monitoring-server/configs"
	"github.com/danial-riazati/http-monitoring-server/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAlerts(cnx *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), configs.Cfg.DataBase.Timeout)
	userId, _ := cnx.Get("user_id")
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
	defer cancel()
	if err != nil {
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cnx.JSON(http.StatusOK, user.Alerts)
}
