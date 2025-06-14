package rate_limiter

import (
	"context"
)

func (s *rateLimiterService) checkIPLimit(ctx context.Context, ip string) (bool, error) {
	if s.limits.ipDuration == 0 {
		return true, nil
	}
	ipVal, err := s.rateLimiterRepository.Increment(ctx, ip, s.limits.ipDuration)
	if err != nil {
		return false, err
	}
	if ipVal > int64(s.limits.ipRateLimit) {
		return false, nil
	}
	return true, nil
}
