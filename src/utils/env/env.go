package env

import (
	"os"
	"strconv"
	"time"

	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/structs"
	"github.com/joho/godotenv"
)

func GetConfigs(path string) (structs.Configs, error) {
	err := godotenv.Load(path)
	if err != nil {
		return structs.Configs{}, err
	}
	rateLimitInterval, err := strconv.Atoi(os.Getenv("IP_RATE_LIMIT"))
	if err != nil {
		return structs.Configs{}, err
	}
	rateLimitApiKey, err := strconv.Atoi(os.Getenv("API_KEY_RATE_LIMIT"))
	if err != nil {
		return structs.Configs{}, err
	}
	rateLimitIpInterval, err := strconv.Atoi(os.Getenv("RATE_LIMIT_IP_INTERVAL_IN_SECONDS"))
	if err != nil {
		return structs.Configs{}, err
	}
	rateLimitApiKeyInterval, err := strconv.Atoi(os.Getenv("RATE_LIMIT_KEY_INTERVAL_IN_SECONDS"))
	if err != nil {
		return structs.Configs{}, err
	}
	configs := structs.Configs{
		RedisHost:               os.Getenv("REDIS_HOST"),
		RedisPort:               os.Getenv("REDIS_PORT"),
		RedisPassword:           os.Getenv("REDIS_PASSWORD"),
		RedisDB:                 0,
		ApiPort:                 os.Getenv("API_PORT"),
		ApiKey:                  os.Getenv("API_KEY"),
		RateLimitIp:             rateLimitInterval,
		RateLimitApiKey:         rateLimitApiKey,
		RateLimitIpInterval:     time.Duration(rateLimitIpInterval) * time.Second,
		RateLimitApiKeyInterval: time.Duration(rateLimitApiKeyInterval) * time.Second,
	}

	return configs, nil
}
