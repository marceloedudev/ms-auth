package grpc_user

import (
	clientUsecase "ms-auth/internal/client/usecase"
	userUsecase "ms-auth/internal/user/usecase"
)

type UserGrpcService struct {
	UserService userUsecase.UseCaseUser
	ClientService clientUsecase.UseCaseClient
}

func NewUserGRpcService(userService userUsecase.UseCaseUser, clientService clientUsecase.UseCaseClient) *UserGrpcService {
	return &UserGrpcService{
		UserService: userService,
		ClientService: clientService,
	}
}
