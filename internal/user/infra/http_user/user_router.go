package http_user

import (
	clientUsecase "ms-auth/internal/client/usecase"
	userUsecase "ms-auth/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

func MakeUserRouters(route *gin.RouterGroup, userService userUsecase.UseCaseUser, clientService clientUsecase.UseCaseClient) {
	api := route.Group("user", RequestLimit(userService))
	{
		api.POST("/revoke", revokeUser(userService))

		api.POST("/tokens", tokenUser(userService))

		api.POST("/authorize", authorizeUser(userService, clientService))

		api.POST("/register", registerUser(userService))

		api.GET("/me", infoUser(userService))
	}
}
