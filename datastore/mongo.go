package datastore

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func NewMongo() *mongo.Database{
	clientOptions := options.Client().ApplyURI(viper.GetString("database.url"))
	client, err := mongo.NewClient(clientOptions)

	timeout := time.Duration(viper.GetInt("context.timeout")) * time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	err = client.Connect(ctx)
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	db := client.Database("clapdb")
	return db
}