package testutils

import (
	"context"
	"testing"

	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/redis"
)

func NewRedisDB() (*redis.Redis, error) {
	return redis.NewRedis(context.Background(), redis.Config{
		Host:     "localhost",
		Port:     "6380",
		Password: "redis-test",
		DB:       0,
	})
}

func WithDB(ts *testing.T, name string, fn func(t *testing.T, redis *redis.Redis)) {
	db, err := NewRedisDB()
	if err != nil {
		ts.Fatalf("Failed to create Redis DB: %v", err)
	}

	fn(ts, db)
}
