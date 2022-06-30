package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danial-riazati/http-monitoring-server/database"
	"github.com/danial-riazati/http-monitoring-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func RequestHTTP(userId string, url models.URL) {
	for true {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		var isURLExists bool = false
		var index int = 0
		for i, x := range user.Urls {
			if x.URL == url.URL {
				isURLExists = true
				index = i
				break
			}
		}
		if !isURLExists {
			break
		}
		resp, err := http.Get(url.URL)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("url: %s  statuscode: %d\n", url.URL, resp.StatusCode)
		var history models.History
		history.URL = url
		history.StatusCode = resp.StatusCode
		history.Requested_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if resp.StatusCode < 200 && resp.StatusCode > 299 {
			user.Urls[index].Failed++
			if user.Urls[index].Failed == user.Urls[index].Threshold {
				user.Alerts = append(user.Alerts, history)
				user.Urls[index].Failed = 0
			}

		} else {
			user.Urls[index].Succeed++
		}
		user.History = append(user.History, history)
		filter := bson.M{"user_id": userId}
		userCollection.ReplaceOne(ctx, filter, user)
		time.Sleep(10 * time.Second)
	}
}
