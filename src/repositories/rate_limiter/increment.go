package rate_limiter

import (
	"context"
	"time"
)

func (r *rateLimiterRepository) Increment(ctx context.Context, key string, expiration time.Duration) (int64, error) {
	value, err := r.redis.Incr(ctx, key, expiration)
	if err != nil {
		return 0, err
	}

	return value, nil
}
