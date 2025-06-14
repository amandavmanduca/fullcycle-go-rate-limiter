package rate_limiter

import "context"

func (r *rateLimiterRepository) Get(ctx context.Context, key string) (interface{}, error) {
	value, err := r.redis.Get(ctx, key)
	if err != nil {
		return "", err
	}

	return value, nil
}
