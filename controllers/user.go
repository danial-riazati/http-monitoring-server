package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/danial-riazati/http-monitoring-server/configs"
	"github.com/danial-riazati/http-monitoring-server/handlers"
	"github.com/danial-riazati/http-monitoring-server/models"
	"github.com/danial-riazati/http-monitoring-server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(cnx *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), configs.Cfg.DataBase.Timeout)
	var user models.User
	var foundUser models.User

	if err := cnx.BindJSON(&user); err != nil {
		cnx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	err := userCollection.FindOne(ctx, bson.M{"name": user.Name}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": "Username or Password is incorrect"})
		return
	}

	passwordIsValid, msg := utils.VerifyPassword(*user.Password, *foundUser.Password)
	defer cancel()
	if !passwordIsValid {
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	token, refreshToken, _ := utils.GenerateAllTokens(*foundUser.Name, foundUser.User_id)

	utils.UpdateAllTokens(token, refreshToken, foundUser.User_id)
	handlers.MonitorAllRequests(foundUser)
	cnx.JSON(http.StatusOK, foundUser.Token)

}

func SignUp(cnx *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User

	if err := cnx.BindJSON(&user); err != nil {
		cnx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		cnx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		defer cancel()
		return
	}

	count, err := userCollection.CountDocuments(ctx, bson.M{"name": user.Name})
	defer cancel()
	if err != nil {
		log.Panic(err)
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the username"})
		return
	}
	if count > 0 {
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": "this username already exists"})
		return
	}

	password := utils.HashPassword(*user.Password)
	user.Password = &password

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	token, refreshToken, _ := utils.GenerateAllTokens(*user.Name, user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken

	resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		msg := "User item was not created"
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()

	cnx.JSON(http.StatusOK, resultInsertionNumber)

}
func GetTodayHistory(cnx *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), configs.Cfg.DataBase.Timeout)
	userId, _ := cnx.Get("user_id")
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
	defer cancel()
	if err != nil {
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var filtered []models.History
	for _, x := range user.History {
		if x.Requested_at.Day() == time.Now().Day() {
			filtered = append(filtered, x)
		}
	}
	cnx.JSON(http.StatusOK, filtered)
}
