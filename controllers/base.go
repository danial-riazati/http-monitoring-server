package controllers

import (
	"github.com/danial-riazati/http-monitoring-server/database"
	validator "github.com/go-playground/validator/v10"
)

var userCollection = database.OpenCollection(database.Client, "user")
var validate = validator.New()
