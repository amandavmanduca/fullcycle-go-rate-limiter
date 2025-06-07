package main

import (
	"log"
	"os"

	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/containers"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/handlers"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/redis"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	redis, err := redis.NewRedis(redis.Config{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
		return
	}

	repositories := containers.NewRepositoryContainer(redis)
	services := containers.NewServiceContainer(repositories)

	handlers.Start(services)
}
