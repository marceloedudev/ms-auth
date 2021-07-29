package main

import (
	"log"
	"ms-auth/cmd/server/grpc"
	"ms-auth/cmd/server/http/routes"
	"ms-auth/config"
	_ "ms-auth/docs"
	"ms-auth/pkg/postgres"
	"ms-auth/pkg/redisdb"
	"os"
)

// @title Auth API
// @version 1.0
// @description This is a basic API Auth using Gin.

// @host localhost:3000
// @BasePath /

func main() {
	configName := config.GetConfigByName(os.Getenv("config"))

	config, err := config.GetConfig(configName)

	if err != nil {
		log.Fatalf("Config failed %s", err)
	}

	postgresDB, err := postgres.InitPostgres(config)
	if err != nil {
		log.Fatalf("Postgres failed %s", err)
	}
	defer postgresDB.Close()

	log.Printf("Postgres connected: %#v", postgresDB.Stats())

	redisDB, err := redisdb.InitRedis(config)
	if err != nil {
		log.Fatalf("Redis failed %s", err)
	}

	log.Println("Redis connected", redisDB)

	gRpcServer, lis, err := grpc.StartGrpcServer(config, postgresDB, redisDB)

	if err != nil {
		log.Fatal("GRPC failed ", err)
	}

	defer lis.Close()

	app := routes.MakeRouters(config, postgresDB, redisDB)

	go func() {
		log.Println("GRPC Server listening")

		if err := gRpcServer.Serve(lis); err != nil {
			log.Fatalln("GRPC failed ", err)
		}
	}()

	app.Run(config.Server.Port)

	log.Println("Http Server listening")

	gRpcServer.GracefulStop()
}
