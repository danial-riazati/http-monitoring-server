package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/danial-riazati/http-monitoring-server/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DBinstance func
func DBinstance() *mongo.Client {

	MongoDb := configs.Cfg.DataBase.ConnectionString
	fmt.Println(MongoDb)
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to MongoDB!")

	return client
}

//Client Database instance
var Client *mongo.Client = DBinstance()

//OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)

	return collection
}
