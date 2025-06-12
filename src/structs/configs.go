package structs

type Configs struct {
	RedisHost         string
	RedisPort         string
	RedisPassword     string
	RedisDB           int
	ApiPort           string
	ApiKey            string
	RateLimitInterval int
	RateLimitApiKey   int
}
