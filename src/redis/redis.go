package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type Config struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type Redis struct {
	client *redis.Client
	config Config
}

func NewRedis(config Config) (*Redis, error) {
	client := &Redis{
		client: redis.NewClient(&redis.Options{
			Addr:     config.Host + ":" + config.Port,
			Password: config.Password,
			DB:       config.DB,
		}),
	}
	if err := client.client.Ping().Err(); err != nil {
		return nil, err
	}
	return client, nil
}

func (r *Redis) GetClient() (*redis.Client, error) {
	if r.client == nil {
		reconnect, err := NewRedis(r.config)
		if err != nil {
			return nil, err
		}
		r.client = reconnect.client
	}
	return r.client, nil
}

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	client, err := r.GetClient()
	if err != nil {
		return err
	}
	return client.Set(key, value, expiration).Err()
}

func (r *Redis) Get(key string) (interface{}, error) {
	client, err := r.GetClient()
	if err != nil {
		return "", err
	}
	return client.Get(key).Result()
}
