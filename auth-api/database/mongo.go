package database

import (
	"auth-api/config"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {

	client, err := mongo.Connect(
		options.Client().ApplyURI(config.MongoURI),
	)

	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("MongoDB Ping Failed:", err)
	}

	DB = client.Database(config.DatabaseName)

	log.Println("✅ MongoDB Connected Successfully")
}