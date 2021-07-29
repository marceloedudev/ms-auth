package grpc_user_test

import (
	"context"
	"ms-auth/internal/user/infra/grpc_user"
	userProto "ms-auth/internal/user/infra/grpc_user/pb"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_UserAuthorized(t *testing.T) {

	userService := getUserService()

	service := grpc_user.NewUserGRpcService(userService, nil)

	user, err := userService.RegisterUser(context.Background(), faker.Username(), faker.FirstName(), faker.LastName(), faker.Password(), faker.Email())
	assert.NotNil(t, user)
	assert.Nil(t, err)

	request := &userProto.UserAuthorizedRequest{
		ClientId:     "id",
		ClientSecret: "secret",
		AccessToken:  user.AccessToken,
	}

	response, err := service.UserAuthorized(context.Background(), request)

	require.Nil(t, err)
	require.NotNil(t, response)

}

// func Test_UserAuthorizedConn(t *testing.T) {

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	u := userMock.NewMockUseCaseUser(ctrl)
// 	c := clientMock.NewMockUseCaseClient(ctrl)

// 	user := &domain.UserDataToken{
// 		ID:         1,
// 		Username:   faker.Username(),
// 		FirstName:  faker.FirstName(),
// 		LastName:   faker.LastName(),
// 		Email:      faker.Email(),
// 		ManagingID: "master",
// 	}

// 	token := faker.UUIDDigit()

// 	u.EXPECT().InfoAccessToken(token).Return(user, nil)

// 	gRpcServer, lis, err := serverGrpc.CreateGrpcServer(getConfig(), u, c)

// 	if err != nil {
// 		log.Fatalf("ServerGrpc failed: %v", err)
// 	}

// 	go func() {
// 		log.Println("gRPC Server listening")

// 		if err := gRpcServer.Serve(lis); err != nil {
// 			log.Fatalln("gRPC Could not connect ", err)
// 		}
// 	}()

// 	conn := ConnectGrpc()

// 	defer conn.Close()

// 	for {

// 		client := userProto.NewAuthUserServiceClient(conn)
// 		request := &userProto.UserAuthorizedRequest{
// 			ClientId:     "id",
// 			ClientSecret: "secret",
// 			AccessToken:  token,
// 		}

// 		_, err := client.UserAuthorized(context.Background(), request)

// 		assert.NotNil(t, err)

// 		return

// 	}
// }
