package http_client

import (
	clientUsecase "ms-auth/internal/client/usecase"
	userUsecase "ms-auth/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

func MakeClientRouters(route *gin.RouterGroup, clientService clientUsecase.UseCaseClient, userService userUsecase.UseCaseUser) {
	api := route.Group("client", TokenAuthMiddleware(userService, clientService))
	{
		api.POST("/register", registerClient(clientService))
	}
}
