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

func (r *Redis) Get(ctx context.Context, key string) (interface{}, error) {
	client, err := r.GetClient(ctx)
	if err != nil {
		return "", err
	}
	return client.Get(ctx, key).Result()
}

func (r *Redis) Incr(ctx context.Context, key string, expiration time.Duration) (int64, error) {
	client, err := r.GetClient(ctx)
	if err != nil {
		return 0, err
	}
	err = client.Incr(ctx, key).Err()
	if err != nil {
		return 0, err
	}
	val, err := client.Get(ctx, key).Int64()
	if err != nil {
		return 0, err
	}
	if val == 1 {
		err = client.Expire(ctx, key, expiration).Err()
		if err != nil {
			return 0, err
		}
	}
	return val, err
}
