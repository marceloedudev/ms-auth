package grpc

import (
	"ms-auth/config"
	clientRepository "ms-auth/internal/client/repository"
	clientUsecase "ms-auth/internal/client/usecase"
	userRepository "ms-auth/internal/user/repository"
	userUsecase "ms-auth/internal/user/usecase"
	"net"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

func StartGrpcServer(conf *config.Config, postgresDB *sqlx.DB, redisDB *redis.Client) (*grpc.Server, net.Listener, error) {

	userPgRepository := userRepository.NewUserPgRepository(postgresDB)
	userRedisRepository := userRepository.NewUserRedisRepository(redisDB)
	userService := userUsecase.NewService(conf, userPgRepository, userRedisRepository)

	clientPgRepository := clientRepository.NewClientPgRepository(postgresDB)
	clientService := clientUsecase.NewService(conf, clientPgRepository)

	server, lis, err := CreateGrpcServer(conf, userService, clientService)

	if err != nil {
		return nil, nil, err
	}

	return server, lis, nil

}
