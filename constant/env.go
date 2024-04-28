package constant

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type IEnv struct {
	DBHost        string
	DBPort        string
	DBName        string
	DBUsername    string
	DBPassword    string
	AdminUsername string
	AdminPassword string
	RedisHost     string
	RedisPort     string
	RedisPassword string
	SmtpUsername  string
	SmtpPassword  string
}

var env *IEnv

func GetEnv() *IEnv {
	if env == nil {
		rootDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		errEnv := godotenv.Load(filepath.Join(rootDir, ".env"))
		if errEnv != nil {
			log.Fatal("Error loading .env file")
		}

		log.Println("env init")
		env = &IEnv{
			DBHost:        os.Getenv("DB_HOST"),
			DBPort:        os.Getenv("DB_PORT"),
			DBName:        os.Getenv("DB_NAME"),
			DBUsername:    os.Getenv("DB_USERNAME"),
			DBPassword:    os.Getenv("DB_PASSWORD"),
			AdminUsername: os.Getenv("ADMIN_USERNAME"),
			AdminPassword: os.Getenv("ADMIN_PASSWORD"),
			RedisHost:     os.Getenv("REDIS_HOST"),
			RedisPort:     os.Getenv("REDIS_HOST"),
			RedisPassword: os.Getenv("REDIS_PASSWORD"),
			SmtpUsername:  os.Getenv("SMTP_USERNAME"),
			SmtpPassword:  os.Getenv("SMTP_PASSWORD"),
		}
	}

	return env
}
