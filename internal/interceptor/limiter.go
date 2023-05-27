package interceptor

import (
	"context"

	"github.com/satanaroom/auth/internal/limiter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RateLimiterInterceptor struct {
	limiter *limiter.TokenBucketLimiter
}

func NewRateLimiterInterceptor(limiter *limiter.TokenBucketLimiter) *RateLimiterInterceptor {
	return &RateLimiterInterceptor{
		limiter: limiter,
	}
}

func (r *RateLimiterInterceptor) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if !r.limiter.Allow() {
		return nil, status.Errorf(codes.ResourceExhausted, "rate limit exceeded")
	}

	return handler(ctx, req)
}
