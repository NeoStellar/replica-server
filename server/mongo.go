package server

import (
	"context"
	"log"

	"iharacee/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Mongo          *mongo.Database
	UserCollection *mongo.Collection
)

func init() {
	mongoURI := config.MONGO_URI
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	log.Println("Connecting to MongoDB...")
	if err != nil {
		log.Println(err.Error())
		log.Println("Error connecting to MongoDB")
	} else {
		var result bson.M
		Mongo = client.Database(config.MONGO_DBNAME)
		if err := Mongo.RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
			log.Fatalln("Error pinging MongoDB deployment")
		} else {
			log.Println("Successfully connected to MongoDB!")
		}
	}
	UserCollection = Mongo.Collection("Users")
}
