package serverinterceptors

import (
	"context"

	"google.golang.org/grpc"

	"github.com/gzorm/commons/zrpc/internal/auth"
)

// StreamAuthorizeInterceptor returns a func that uses given authenticator in processing stream requests.
func StreamAuthorizeInterceptor(authenticator *auth.Authenticator) grpc.StreamServerInterceptor {
	return func(svr any, stream grpc.ServerStream, info *grpc.StreamServerInfo,
		handler grpc.StreamHandler) error {
		if err := authenticator.Authenticate(stream.Context()); err != nil {
			return err
		}

		return handler(svr, stream)
	}
}

// UnaryAuthorizeInterceptor returns a func that uses given authenticator in processing unary requests.
func UnaryAuthorizeInterceptor(authenticator *auth.Authenticator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (any, error) {
		if err := authenticator.Authenticate(ctx); err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}
