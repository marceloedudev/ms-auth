package http_user_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"ms-auth/cmd/server/http/middleware"
	clientUsecase "ms-auth/internal/client/usecase"
	clientMock "ms-auth/internal/client/usecase/mock"
	"ms-auth/internal/user/domain"
	"ms-auth/internal/user/infra/http_user"
	userUsecase "ms-auth/internal/user/usecase"
	userMock "ms-auth/internal/user/usecase/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func AddRouters(userService userUsecase.UseCaseUser, clientService clientUsecase.UseCaseClient) *gin.Engine {
	gin.SetMode(gin.TestMode)

	app := gin.New()

	app.Use(middleware.HandleErrors())

	api := app.Group("api")
	{
		http_user.MakeUserRouters(api, userService, clientService)
	}

	return app
}

func Test_RegisterUser(t *testing.T) {

	t.Run("should be able register user", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		u := userMock.NewMockUseCaseUser(ctrl)
		c := clientMock.NewMockUseCaseClient(ctrl)

		user := &domain.TokenResponse{
			UserID:       1,
			AccessToken:  "3f0dcdfe-7117-461d-b54e-bfc6c039a173",
			ExpiresIn:    3600,
			TokenType:    "Bearer",
			RefreshToken: "28591e38-6f52-4fbc-97ce-747d76edfa2f",
		}

		InputUser := &http_user.UserRegisterRequest{
			Username:  "_example_",
			FirstName: "Example",
			LastName:  "Admin",
			Password:  "123456",
			Email:     "example@gmail.com",
		}

		u.EXPECT().RegisterUser(gomock.Any(), InputUser.Username, InputUser.FirstName, InputUser.LastName, InputUser.Password, InputUser.Email).Return(user, nil)
		u.EXPECT().RequestLimit(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		testRouter := AddRouters(u, c)

		data, _ := json.Marshal(InputUser)

		req, err := http.NewRequest("POST", "/api/user/register", bytes.NewBufferString(string(data)))

		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(err)
		}

		resp := httptest.NewRecorder()

		testRouter.ServeHTTP(resp, req)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		res := domain.TokenResponse{}

		json.Unmarshal(body, &res)

		assert.Equal(t, resp.Code, http.StatusOK)

		assert.Equal(t, res.TokenType, "Bearer")

	})

}
