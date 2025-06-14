package structs

import "time"

type Configs struct {
	RedisHost               string
	RedisPort               string
	RedisPassword           string
	RedisDB                 int
	ApiPort                 string
	ApiKey                  string
	RateLimitIp             int
	RateLimitIpInterval     time.Duration
	RateLimitApiKey         int
	RateLimitApiKeyInterval time.Duration
}
