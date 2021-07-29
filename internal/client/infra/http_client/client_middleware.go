package http_client

import (
	"errors"
	clientUsecase "ms-auth/internal/client/usecase"
	userUsecase "ms-auth/internal/user/usecase"
	"ms-auth/pkg/http_exception"
	"ms-auth/pkg/utils"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(us userUsecase.UseCaseUser, cvs clientUsecase.UseCaseClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := utils.ExtractTokenFromRequest(c.Request)

		user, err := us.InfoAccessToken(tokenString)

		if err != nil {
			panic(err)
		}

		existsManaging, err := cvs.VerifyManaging(c, user.ManagingID)

		if err != nil {
			panic(err)
		}

		if existsManaging == nil {
			panic(http_exception.BadRequestException(errors.New("Access not allowed")))
		}

		c.Next()

	}
}
