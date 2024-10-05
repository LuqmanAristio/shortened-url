package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
    "os"
	"github.com/joho/godotenv"
)

var Client *mongo.Client

func Connect() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
    
	mongoUri := os.Getenv("DATABASE_URI")
	clientOptions := options.Client().ApplyURI(mongoUri)

	var err error
    Client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = Client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")
}