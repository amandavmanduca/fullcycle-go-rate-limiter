package env

import (
	"os"
	"strconv"

	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/structs"
	"github.com/joho/godotenv"
)

func GetConfigs(path string) (structs.Configs, error) {
	err := godotenv.Load(path)
	if err != nil {
		return structs.Configs{}, err
	}
	rateLimitInterval, err := strconv.Atoi(os.Getenv("IP_RATE_LIMIT_PER_SECOND"))
	if err != nil {
		return structs.Configs{}, err
	}
	rateLimitApiKey, err := strconv.Atoi(os.Getenv("API_KEY_RATE_LIMIT_PER_SECOND"))
	if err != nil {
		return structs.Configs{}, err
	}
	configs := structs.Configs{
		RedisHost:         os.Getenv("REDIS_HOST"),
		RedisPort:         os.Getenv("REDIS_PORT"),
		RedisPassword:     os.Getenv("REDIS_PASSWORD"),
		RedisDB:           0,
		ApiPort:           os.Getenv("API_PORT"),
		ApiKey:            os.Getenv("API_KEY"),
		RateLimitInterval: rateLimitInterval,
		RateLimitApiKey:   rateLimitApiKey,
	}

	return configs, nil
}
