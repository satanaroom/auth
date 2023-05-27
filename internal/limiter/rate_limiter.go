package limiter

import (
	"context"
	"time"
)

type TokenBucketLimiter struct {
	tokenBucketCh chan struct{}
}

func NewTokenBucketLimiter(ctx context.Context, limit int, period time.Duration) *TokenBucketLimiter {
	limiter := &TokenBucketLimiter{
		tokenBucketCh: make(chan struct{}, limit),
	}

	for i := 0; i < limit; i++ {
		limiter.tokenBucketCh <- struct{}{}
	}

	replenishmentInterval := period.Nanoseconds() / int64(limit)
	go limiter.startPeriodReplenishment(ctx, time.Duration(replenishmentInterval))

	return limiter
}

func (l *TokenBucketLimiter) startPeriodReplenishment(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			l.tokenBucketCh <- struct{}{}
		}
	}
}

func (l *TokenBucketLimiter) Allow() bool {
	select {
	case <-l.tokenBucketCh:
		return true
	default:
		return false
	}
}

//func clientOpts() {
//	var opts []grpc.DialOption
//	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
//		grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
//		grpc_retry.WithMax(5),
//		grpc_retry.WithBackoff(grpc_retry.BackoffLinear())
//		)))
//}
