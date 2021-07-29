package grpc_user

import (
	"context"
	userProto "ms-auth/internal/user/infra/grpc_user/pb"
	"ms-auth/pkg/grpc_exception"
)

func (u *UserGrpcService) UserAuthorized(ctx context.Context, r *userProto.UserAuthorizedRequest) (*userProto.UserAuthorizedResponse, error) {

	user, err := u.UserService.InfoAccessToken(r.AccessToken)

	if err != nil {
		return nil, grpc_exception.InternalServerExceptionGRPC("UserAuthorized.UserService.InfoAccessToken", err)
	}

	return &userProto.UserAuthorizedResponse{
		Id:        uint64(user.ID),
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil

}
