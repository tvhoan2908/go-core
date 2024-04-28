package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

var RedisClient *redis.Client

func ConnectRedis() *redis.Client {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: 0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	RedisClient = client
	return client
}