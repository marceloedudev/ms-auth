package grpc

import (
	"ms-auth/cmd/server/grpc/interceptors"
	"ms-auth/config"
	"ms-auth/internal/user/infra/grpc_user"

	"net"
	"time"

	userProto "ms-auth/internal/user/infra/grpc_user/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	clientUsecase "ms-auth/internal/client/usecase"
	userUsecase "ms-auth/internal/user/usecase"
)

func CreateGrpcServer(conf *config.Config, userService userUsecase.UseCaseUser, clientService clientUsecase.UseCaseClient) (*grpc.Server, net.Listener, error) {

	lis, err := net.Listen("tcp", conf.Grpc.Port)

	if err != nil {
		return nil, nil, err
	}

	grpcInteceptors := interceptors.NewInterceptors()

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: conf.Grpc.MaxConnectionIdle * time.Minute,
		Timeout:           conf.Grpc.Timeout * time.Second,
		MaxConnectionAge:  conf.Grpc.MaxConnectionAge * time.Minute,
		Time:              conf.Grpc.Timeout * time.Minute,
	}),
		grpc.UnaryInterceptor(grpcInteceptors.Logger),
	)

	userGrpcService := grpc_user.NewUserGRpcService(userService, clientService)
	userProto.RegisterAuthUserServiceServer(server, userGrpcService)

	if conf.Server.Mode != "production" {
		reflection.Register(server)
	}

	return server, lis, nil

}

func CreateGrpcClient(conf *config.Config) (*grpc.ClientConn, error) {

	opts := grpc.WithInsecure()

	clientConn, err := grpc.Dial(conf.Grpc.Port, opts)

	if err != nil {
		return nil, err
	}

	return clientConn, nil
}
