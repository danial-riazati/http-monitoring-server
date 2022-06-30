package handlers

import (
	"fmt"

	"github.com/danial-riazati/http-monitoring-server/models"
)

func MonitorAllRequests(user models.User) {
	for _, x := range user.Urls {
		fmt.Printf("url: %s  userId: %s\n", x.URL, user.User_id)
		go RequestHTTP(user.User_id, x)
	}
}
