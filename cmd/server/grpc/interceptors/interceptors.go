package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Interceptor struct {
}

func NewInterceptors() *Interceptor {
	return &Interceptor{}
}

func (i Interceptor) Logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	
	start := time.Now()

	h, err := handler(ctx, req)

	md, _ := metadata.FromIncomingContext(ctx)

	log.Printf("GRPC Request: Method: %s\tDuration: %v\tMetadata: %v\tError: %v\n", info.FullMethod, time.Since(start), md, err)

	return h, err

}
