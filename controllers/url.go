package controllers

import (
	"context"

	"log"
	"net/http"

	"github.com/danial-riazati/http-monitoring-server/configs"
	"github.com/danial-riazati/http-monitoring-server/handlers"
	"github.com/danial-riazati/http-monitoring-server/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUrl(cnx *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), configs.Cfg.DataBase.Timeout)
	userId, _ := cnx.Get("user_id")
	log.Println(userId)
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
	defer cancel()
	if err != nil {
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var url models.URL
	if err := cnx.BindJSON(&url); err != nil {
		cnx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(url.URL)
	if len(user.Urls) >= int(configs.Cfg.User.NoOfUrls) {
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": "You can Just add 20 Urls"})
		return
	}
	var result bool = false
	for _, x := range user.Urls {
		if x.URL == url.URL {
			result = true
			break
		}
	}
	if result {
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": "this url already exists"})
		return
	}
	url.Failed = 0
	user.Urls = append(user.Urls, url)
	filter := bson.M{"user_id": userId}
	userCollection.ReplaceOne(ctx, filter, user)

	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}
	go handlers.HTTPCaller(user.User_id, url)
	cnx.JSON(http.StatusOK, gin.H{"msg": "added sucessfully"})
}

func GetUrl(cnx *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), configs.Cfg.DataBase.Timeout)
	userId, _ := cnx.Get("user_id")
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
	defer cancel()
	if err != nil {
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cnx.JSON(http.StatusOK, user.Urls)
}

func DeleteUrl(cnx *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), configs.Cfg.DataBase.Timeout)
	userId, _ := cnx.Get("user_id")
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
	defer cancel()
	if err != nil {
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var url models.URL
	if err := cnx.BindJSON(&url); err != nil {
		cnx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var result bool = false
	var index int = 0
	for i, x := range user.Urls {
		if x.URL == url.URL {
			result = true
			index = i
			break
		}
	}
	if !result {
		cnx.JSON(http.StatusInternalServerError, gin.H{"error": "this url not exists"})
		return
	}
	user.Urls = append(user.Urls[:index], user.Urls[index+1:]...)
	filter := bson.M{"user_id": userId}
	userCollection.ReplaceOne(ctx, filter, user)

	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}
	cnx.JSON(http.StatusOK, gin.H{"msg": "deleted sucessfully"})
}
