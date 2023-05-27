package interceptor

import (
	"context"

	metric "github.com/satanaroom/auth/internal/metrics"
	"google.golang.org/grpc"
)

func MetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	res, err := handler(ctx, req)
	if err != nil {
		metric.IncRequestTotal(info.FullMethod, "error")
	} else {
		metric.IncRequestTotal(info.FullMethod, "success")
	}

	return res, err
}
