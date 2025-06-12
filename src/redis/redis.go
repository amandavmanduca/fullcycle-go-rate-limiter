package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
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

func NewRedis(ctx context.Context, config Config) (*Redis, error) {
	client := &Redis{
		client: redis.NewClient(&redis.Options{
			Addr:     config.Host + ":" + config.Port,
			Password: config.Password,
			DB:       config.DB,
		}),
	}
	if err := client.client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return client, nil
}

func (r *Redis) GetClient(ctx context.Context) (*redis.Client, error) {
	if r.client == nil {
		reconnect, err := NewRedis(ctx, r.config)
		if err != nil {
			return nil, err
		}
		r.client = reconnect.client
	}
	return r.client, nil
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	client, err := r.GetClient(ctx)
	if err != nil {
		return err
	}
	return client.Set(ctx, key, value, expiration).Err()
}

func (r *Redis) Get(ctx context.Context, key string) (interface{}, error) {
	client, err := r.GetClient(ctx)
	if err != nil {
		return "", err
	}
	return client.Get(ctx, key).Result()
}
