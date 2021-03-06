package keys

import (
	"go-cron-job-microservice/models"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var instance *models.Keys
var once sync.Once

func loadEnv(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
		return ""
	}

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return ""
}

var envVariables = models.Keys{
	MONGO_URI:     loadEnv("MONGO_URI"),
	MONGO_DB_NAME: loadEnv("MONGO_DB_NAME"),
	PORT:          loadEnv("PORT"),
}

// GetKeys returns all environment variables.
func GetKeys() *models.Keys {
	once.Do(func() {
		instance = &envVariables
	})

	return instance
}
