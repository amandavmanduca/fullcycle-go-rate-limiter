package rate_limiter

import (
	"context"
	"errors"
)

var errLimitExceeded = errors.New("you have reached the maximum number of requests or actions allowed within a certain time frame")

func (s *rateLimiterService) CheckRateLimit(ctx context.Context, ip string, apiKey string) error {
	validIP, err := s.checkIPLimit(ctx, ip)
	if err != nil {
		return err
	}

	validApiKey, err := s.checkApiKeyLimit(ctx, apiKey)
	if err != nil {
		return err
	}

	if validApiKey {
		return nil
	}

	if validIP && apiKey == "" {
		return nil
	}

	return errLimitExceeded
}
