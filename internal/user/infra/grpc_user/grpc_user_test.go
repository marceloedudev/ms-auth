package grpc_user_test

import (
	"log"
	serverGrpc "ms-auth/cmd/server/grpc"
	"ms-auth/config"
	clientFakes "ms-auth/internal/client/repository/fakes"
	clientUsecase "ms-auth/internal/client/usecase"
	userFakes "ms-auth/internal/user/repository/fakes"
	userUsecase "ms-auth/internal/user/usecase"
	"ms-auth/pkg/postgres/dbfake"

	"google.golang.org/grpc"
)

func getConfig() *config.Config {

	config, err := config.GetConfig("../../../../config/config-test.yaml")

	if err != nil {
		log.Fatalf("Config failed %s", err)
	}

	err = dbfake.SetEnvDBFakePath("../../../..")
	if err != nil {
		log.Fatalf("dbfake_path failed: %v", err)
	}

	return config

}

func getUserService() *userUsecase.Service {

	config := getConfig()

	pgRepo := userFakes.NewUserPgRepositoryFake()
	redisRepo := userFakes.NewUserRedisRepositoryFake()
	service := userUsecase.NewService(config, pgRepo, redisRepo)

	return service

}

func getClientService() *clientUsecase.Service {

	config := getConfig()

	pgRepo := clientFakes.NewClientPgRepositoryFake()
	service := clientUsecase.NewService(config, pgRepo)

	return service

}

func ConnectGrpc() *grpc.ClientConn {

	conn, err := serverGrpc.CreateGrpcClient(getConfig())

	if err != nil {
		log.Fatalf("Connect grpc failed %s", err)
	}

	return conn

}
