package main

import (
	"context"
	"log"

	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/containers"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/handlers"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/redis"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/utils/env"
)

func main() {
	configs, err := env.GetConfigs(".env")
	if err != nil {
		log.Fatalf("Failed to get configs: %v", err)
		return
	}

	ctx := context.Background()
	redis, err := redis.NewRedis(ctx, redis.Config{
		Host:     configs.RedisHost,
		Port:     configs.RedisPort,
		Password: configs.RedisPassword,
		DB:       configs.RedisDB,
	})
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
		return
	}

	repositories := containers.NewRepositoryContainer(redis)
	services := containers.NewServiceContainer(configs, repositories)

	handlers.Start(configs, services)
}
