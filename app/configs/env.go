package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DBUSER := os.Getenv("DB_USER")
	DBPASSWORD := os.Getenv("DB_PASSWORD")
	DBHOST := os.Getenv("DB_HOST")
	DBPORT := os.Getenv("MONGODB_DOCKER_PORT")
	DBNAME := os.Getenv("DB_NAME")

	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", DBUSER, DBPASSWORD, DBHOST, DBPORT, DBNAME)
}
