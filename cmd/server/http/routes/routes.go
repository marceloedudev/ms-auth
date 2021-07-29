package routes

import (
	"errors"
	"fmt"
	"ms-auth/cmd/server/http/middleware"
	"ms-auth/config"
	"ms-auth/internal/client/infra/http_client"
	clientRepository "ms-auth/internal/client/repository"
	clientUsecase "ms-auth/internal/client/usecase"
	"ms-auth/internal/user/infra/http_user"
	userRepository "ms-auth/internal/user/repository"
	userUsecase "ms-auth/internal/user/usecase"
	"ms-auth/pkg/http_exception"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func getUserService(config *config.Config, postgresDB *sqlx.DB, redisDB *redis.Client) *userUsecase.Service {
	userPgRepository := userRepository.NewUserPgRepository(postgresDB)
	userRedisRepository := userRepository.NewUserRedisRepository(redisDB)
	return userUsecase.NewService(config, userPgRepository, userRedisRepository)
}

func getClientService(config *config.Config, postgresDB *sqlx.DB) *clientUsecase.Service {
	clientPgRepository := clientRepository.NewClientPgRepository(postgresDB)
	return clientUsecase.NewService(config, clientPgRepository)
}

func MakeRouters(config *config.Config, postgresDB *sqlx.DB, redisDB *redis.Client) *gin.Engine {

	app := gin.Default()

	app.Use(middleware.RequestID())
	app.Use(middleware.CORS())
	app.Use(middleware.HandleErrors())

	userService := getUserService(config, postgresDB, redisDB)
	clientService := getClientService(config, postgresDB)

	api := app.Group("auth/v1")
	{
		http_user.MakeUserRouters(api, userService, clientService)
		http_client.MakeClientRouters(api, clientService, userService)
	}

	if config.Server.Mode == "development" {
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	app.NoRoute(func(c *gin.Context) {
		panic(http_exception.NotFoundExeception(errors.New(fmt.Sprintf("Route '%s' was not found", c.Request.URL.Path))))
	})

	return app

}
