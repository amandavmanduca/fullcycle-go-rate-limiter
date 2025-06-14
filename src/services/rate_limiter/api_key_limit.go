package rate_limiter

import "context"

func (s *rateLimiterService) checkApiKeyLimit(ctx context.Context, apiKey string) (bool, error) {
	if s.limits.apiKeyDuration == 0 {
		return true, nil
	}
	if apiKey == "" {
		return false, nil
	}
	apiKeyVal, err := s.rateLimiterRepository.Increment(ctx, apiKey, s.limits.apiKeyDuration)
	if err != nil {
		return false, err
	}
	if apiKeyVal > int64(s.limits.apiKeyRateLimit) {
		return false, nil
	}
	return true, nil
}
