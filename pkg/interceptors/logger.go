package interceptors

import (
	"context"
	"time"

	grpclogging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// ZerologUnaryServerInterceptor returns a new unary server interceptors that adds zerolog.Logger to the context.
func ZerologUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		startTime := time.Now()
		resp, err := handler(ctx, req)
		code := grpclogging.DefaultErrorToCode(err)
		duration := time.Since(startTime)

		log.Info().Msgf(
			"METHOD [%s] - STATUS [%s] - DUR %s",
			info.FullMethod, code.String(), duration.String(),
		)

		return resp, err
	}
}
