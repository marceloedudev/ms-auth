package usecase_test

import (
	"log"
	"ms-auth/config"
	"ms-auth/internal/user/repository/fakes"
	"ms-auth/internal/user/usecase"
	"ms-auth/pkg/postgres/dbfake"
)

func getConfig() *config.Config {

	config, err := config.GetConfig("../../../config/config-test.yaml")

	if err != nil {
		log.Fatalf("Config failed %s", err)
	}

	err = dbfake.SetEnvDBFakePath("../../..")
	if err != nil {
		log.Fatalf("dbfake_path failed: %v", err)
	}

	return config

}

func getModel() *usecase.Service {

	config := getConfig()

	pgRepo := fakes.NewUserPgRepositoryFake()
	redisRepo := fakes.NewUserRedisRepositoryFake()
	model := usecase.NewService(config, pgRepo, redisRepo)

	return model

}
