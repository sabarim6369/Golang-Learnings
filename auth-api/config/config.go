package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port         string
	MongoURI     string
	DatabaseName string
	JWTSecret    string
)

func LoadConfig() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Port = os.Getenv("PORT")
	MongoURI = os.Getenv("MONGODB_URI")
	DatabaseName = os.Getenv("DATABASE_NAME")
	JWTSecret = os.Getenv("JWT_SECRET")
}