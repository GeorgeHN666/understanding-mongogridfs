package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBConnect = connection()

var userOpt = options.Client().ApplyURI("MONGO URI HERE")

func connection() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), userOpt)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection success with mongo DataBase")

	return client

}

func CheckConnection() int {
	err := DBConnect.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
